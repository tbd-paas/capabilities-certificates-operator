package phases

import (
	"fmt"

	"github.com/nukleros/desired"
	"github.com/nukleros/operator-builder-tools/pkg/controller/phases"
	"github.com/nukleros/operator-builder-tools/pkg/controller/reconcile"
	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/resources"
	"github.com/nukleros/operator-builder-tools/pkg/status"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CreateResourcesPhase creates or updated the child resources of a workload during a reconciliation loop.
func CreateResourcesPhase(r workload.Reconciler, req *workload.Request) (bool, error) {
	// get the resources in memory
	desiredResources, err := r.GetResources(req)
	if err != nil {
		return false, fmt.Errorf("unable to retrieve resources, %w", err)
	}

	proceed := true

	for _, resource := range desiredResources {
		condition, created, err := phases.HandleResourcePhaseExit(
			persistResourcePhase(r, req, resource),
		)
		if err != nil {
			if !phases.IsOptimisticLockError(err) {
				req.Log.Error(err, "unable to create or update resource")
			}
		}

		if !created {
			return false, nil
		}

		resourceObject := status.ToCommonResource(resource)
		resourceObject.ChildResourceCondition = condition

		// update the status conditions and return any errors
		if err := phases.UpdateResourceConditions(r, req, resourceObject); err != nil {
			if !phases.IsOptimisticLockError(err) {
				r.GetLogger().Error(
					err, "failed to update resource conditions",
					"kind", resource.GetObjectKind().GroupVersionKind().Kind,
					"name", resource.GetName(),
					"namespace", resource.GetNamespace(),
				)

				created = false
			}
		}

		proceed = proceed && created
	}

	return proceed, err
}

// CreateOrUpdate creates a resource if it does not already exist or updates a resource
// if it does already exist.
func CreateOrUpdate(r workload.Reconciler, req *workload.Request, resource client.Object) error {
	// set ownership on the underlying resource being created or updated
	if err := ctrl.SetControllerReference(req.Workload, resource, r.Scheme()); err != nil {
		req.Log.Error(
			err, "unable to set owner reference on resource",
			"resourceName", resource.GetName(),
			"resourceNamespace", resource.GetNamespace(),
		)

		return fmt.Errorf("unable to set owner reference on %s, %w", resource.GetName(), err)
	}

	// get the resource from the cluster
	clusterResource, err := resources.Get(r, req, resource)
	if err != nil {
		return fmt.Errorf("unable to retrieve resource %s, %w", resource.GetName(), err)
	}

	// create the resource if we have a nil object, or update the resource if we have one
	// that exists in the cluster already
	if clusterResource == nil {
		if err := resources.Create(r, req, resource); err != nil {
			return fmt.Errorf("unable to create resource %s, %w", resource.GetName(), err)
		}

		// add the created event
		event := Created
		event.RegisterAction(r.GetEventRecorder(), resource, req.Workload)

		return reconcile.Watch(r, req, resource)
	}

	// return if the resource is already in a desired state (no update required)
	isDesired, err := desired.Desired(resource, clusterResource)
	if err != nil {
		r.GetLogger().Error(err, "found error attempting update")
	}

	if isDesired {
		return nil
	}

	r.GetLogger().Info("updating resource")

	if err := resources.Update(r, req, resource, clusterResource); err != nil {
		return fmt.Errorf("unable to update resource %s, %w", resource.GetName(), err)
	}

	// add the updated event
	event := Updated
	event.RegisterAction(r.GetEventRecorder(), resource, req.Workload)

	return nil
}

// persistResourcePhase executes persisting resources to the Kubernetes database.
func persistResourcePhase(
	r workload.Reconciler,
	req *workload.Request,
	resource client.Object,
) (bool, error) {
	ready := true

	if resource.GetNamespace() != "" {
		ready, err := resources.NamespaceForResourceIsReady(r, req, resource)
		if err != nil {
			return ready, fmt.Errorf("unable to determine if %s namespace is ready, %w", resource.GetNamespace(), err)
		}
	}

	// return the result if the object is not ready
	if !ready {
		r.GetLogger().Info(fmt.Sprintf("namespace '%s' is not ready", resource.GetNamespace()))

		return false, nil
	}

	// persist the resource
	if err := CreateOrUpdate(r, req, resource); err != nil {
		if phases.IsOptimisticLockError(err) {
			return true, nil
		}

		return false, fmt.Errorf("unable to create or update resource %s, %w", resource.GetName(), err)
	}

	// ensure the resource is ready before moving to the next phase
	current, err := resources.Get(r, req, resource)
	if err != nil {
		return false, fmt.Errorf("unable to get current resource status from cluster %s, %w", resource.GetName(), err)
	}

	ready, err = resources.IsReady(current)
	if err != nil {
		return false, fmt.Errorf("unable to get ready resource status from cluster %s, %w", resource.GetName(), err)
	}

	if !ready {
		r.GetLogger().Info(
			fmt.Sprintf(
				"resource '%s/%s' in namespace '%s' is not ready",
				resource.GetObjectKind().GroupVersionKind().Kind,
				resource.GetName(),
				resource.GetNamespace(),
			),
		)
	}

	return ready, nil
}
