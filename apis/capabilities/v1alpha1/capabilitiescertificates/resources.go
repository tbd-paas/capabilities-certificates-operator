/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package capabilitiescertificates

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	capabilitiesv1alpha1 "github.com/tbd-paas/capabilities-certificates-operator/apis/capabilities/v1alpha1"
)

// sampleCertificatesCapability is a sample containing all fields
const sampleCertificatesCapability = `apiVersion: capabilities.platform.tbd.io/v1alpha1
kind: CertificatesCapability
metadata:
  name: certificatescapability-sample
spec:
  namespace: "tbd-certificates-system"
  certManager:
    injector:
      replicas: 2
      image: "quay.io/jetstack/cert-manager-cainjector:v1.14.4"
      resources:
        requests:
          cpu: "50m"
          memory: "64Mi"
        limits:
          memory: "128Mi"
    controller:
      replicas: 2
      image: "quay.io/jetstack/cert-manager-controller:v1.14.4"
      resources:
        requests:
          cpu: "25m"
          memory: "32Mi"
        limits:
          memory: "64Mi"
    webhook:
      replicas: 2
      image: "quay.io/jetstack/cert-manager-webhook:v1.14.4"
      resources:
        requests:
          cpu: "25m"
          memory: "32Mi"
        limits:
          memory: "64Mi"
`

// sampleCertificatesCapabilityRequired is a sample containing only required fields
const sampleCertificatesCapabilityRequired = `apiVersion: capabilities.platform.tbd.io/v1alpha1
kind: CertificatesCapability
metadata:
  name: certificatescapability-sample
spec:
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return sampleCertificatesCapabilityRequired
	}

	return sampleCertificatesCapability
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(
	workloadObj capabilitiesv1alpha1.CertificatesCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	resourceObjects := []client.Object{}

	for _, f := range CreateFuncs {
		resources, err := f(&workloadObj, reconciler, req)

		if err != nil {
			return nil, err
		}

		resourceObjects = append(resourceObjects, resources...)
	}

	return resourceObjects, nil
}

// GenerateForCLI returns the child resources that are associated with this workload given
// appropriate YAML manifest files.
func GenerateForCLI(workloadFile []byte) ([]client.Object, error) {
	var workloadObj capabilitiesv1alpha1.CertificatesCapability
	if err := yaml.Unmarshal(workloadFile, &workloadObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into workload, %w", err)
	}

	if err := workload.Validate(&workloadObj); err != nil {
		return nil, fmt.Errorf("error validating workload yaml, %w", err)
	}

	return Generate(workloadObj, nil, nil)
}

// CreateFuncs is an array of functions that are called to create the child resources for the controller
// in memory during the reconciliation loop prior to persisting the changes or updates to the Kubernetes
// database.
var CreateFuncs = []func(
	*capabilitiesv1alpha1.CertificatesCapability,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateNamespaceNamespace,
	CreateCRDCertificaterequestsCertManagerIo,
	CreateCRDCertificatesCertManagerIo,
	CreateCRDChallengesAcmeCertManagerIo,
	CreateCRDClusterissuersCertManagerIo,
	CreateCRDIssuersCertManagerIo,
	CreateCRDOrdersAcmeCertManagerIo,
	CreateDeploymentNamespaceCertManagerCainjector,
	CreateDeploymentNamespaceCertManager,
	CreateDeploymentNamespaceCertManagerWebhook,
	CreateServiceAccountNamespaceCertManagerCainjector,
	CreateServiceAccountNamespaceCertManager,
	CreateServiceAccountNamespaceCertManagerWebhook,
	CreateClusterRoleCertManagerCainjector,
	CreateClusterRoleCertManagerControllerIssuers,
	CreateClusterRoleCertManagerControllerClusterissuers,
	CreateClusterRoleCertManagerControllerCertificates,
	CreateClusterRoleCertManagerControllerOrders,
	CreateClusterRoleCertManagerControllerChallenges,
	CreateClusterRoleCertManagerControllerIngressShim,
	CreateClusterRoleCertManagerClusterView,
	CreateClusterRoleCertManagerView,
	CreateClusterRoleCertManagerEdit,
	CreateClusterRoleCertManagerControllerApproveCertManagerIo,
	CreateClusterRoleCertManagerControllerCertificatesigningrequests,
	CreateClusterRoleCertManagerWebhookSubjectaccessreviews,
	CreateClusterRoleBindingCertManagerCainjector,
	CreateClusterRoleBindingCertManagerControllerIssuers,
	CreateClusterRoleBindingCertManagerControllerClusterissuers,
	CreateClusterRoleBindingCertManagerControllerCertificates,
	CreateClusterRoleBindingCertManagerControllerOrders,
	CreateClusterRoleBindingCertManagerControllerChallenges,
	CreateClusterRoleBindingCertManagerControllerIngressShim,
	CreateClusterRoleBindingCertManagerControllerApproveCertManagerIo,
	CreateClusterRoleBindingCertManagerControllerCertificatesigningrequests,
	CreateClusterRoleBindingCertManagerWebhookSubjectaccessreviews,
	CreateRoleNamespaceCertManagerCainjectorLeaderelection,
	CreateRoleNamespaceCertManagerLeaderelection,
	CreateRoleNamespaceCertManagerWebhookDynamicServing,
	CreateRoleBindingNamespaceCertManagerCainjectorLeaderelection,
	CreateRoleBindingNamespaceCertManagerLeaderelection,
	CreateRoleBindingNamespaceCertManagerWebhookDynamicServing,
	CreateServiceNamespaceCertManager,
	CreateServiceNamespaceCertManagerWebhook,
	CreateMutatingWebhookCertManagerWebhook,
	CreateValidatingWebhookCertManagerWebhook,
	CreateClusterIssuerInternal,
}

// InitFuncs is an array of functions that are called prior to starting the controller manager.  This is
// necessary in instances which the controller needs to "own" objects which depend on resources to
// pre-exist in the cluster. A common use case for this is the need to own a custom resource.
// If the controller needs to own a custom resource type, the CRD that defines it must
// first exist. In this case, the InitFunc will create the CRD so that the controller
// can own custom resources of that type.  Without the InitFunc the controller will
// crash loop because when it tries to own a non-existent resource type during manager
// setup, it will fail.
var InitFuncs = []func(
	*capabilitiesv1alpha1.CertificatesCapability,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateCRDCertificaterequestsCertManagerIo,
	CreateCRDCertificatesCertManagerIo,
	CreateCRDChallengesAcmeCertManagerIo,
	CreateCRDClusterissuersCertManagerIo,
	CreateCRDIssuersCertManagerIo,
	CreateCRDOrdersAcmeCertManagerIo,
}

func ConvertWorkload(component workload.Workload) (*capabilitiesv1alpha1.CertificatesCapability, error) {
	p, ok := component.(*capabilitiesv1alpha1.CertificatesCapability)
	if !ok {
		return nil, capabilitiesv1alpha1.ErrUnableToConvertCertificatesCapability
	}

	return p, nil
}
