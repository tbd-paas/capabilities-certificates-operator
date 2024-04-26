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

package v1alpha1

import (
	"errors"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrUnableToConvertCertificatesCapability = errors.New("unable to convert to CertificatesCapability")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CertificatesCapabilitySpec defines the desired state of CertificatesCapability.
type CertificatesCapabilitySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Required
	//
	//	Namespace to use where underlying certificates capability will be deployed.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Aws CertificatesCapabilitySpecAws `json:"aws,omitempty"`

	// +kubebuilder:validation:Optional
	CertManager CertificatesCapabilitySpecCertManager `json:"certManager,omitempty"`

	// +kubebuilder:validation:Optional
	TrustManager CertificatesCapabilitySpecTrustManager `json:"trustManager,omitempty"`
}

type CertificatesCapabilitySpecAws struct {
	// +kubebuilder:default=""
	// +kubebuilder:validation:Optional
	// (Default: "")
	//
	//	The AWS IAM Role ARN to use for validating public DNS records for issuing public certificates.
	RoleARN string `json:"roleARN,omitempty"`
}

type CertificatesCapabilitySpecCertManager struct {
	// +kubebuilder:validation:Optional
	Injector CertificatesCapabilitySpecCertManagerInjector `json:"injector,omitempty"`

	// +kubebuilder:validation:Optional
	Controller CertificatesCapabilitySpecCertManagerController `json:"controller,omitempty"`

	// +kubebuilder:validation:Optional
	Webhook CertificatesCapabilitySpecCertManagerWebhook `json:"webhook,omitempty"`
}

type CertificatesCapabilitySpecCertManagerInjector struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the cert-manager cainjector deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="quay.io/jetstack/cert-manager-cainjector:v1.14.4"
	// +kubebuilder:validation:Optional
	// (Default: "quay.io/jetstack/cert-manager-cainjector:v1.14.4")
	//
	//	Image to use for cert-manager CA injector deployment.
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Resources CertificatesCapabilitySpecCertManagerInjectorResources `json:"resources,omitempty"`
}

type CertificatesCapabilitySpecCertManagerInjectorResources struct {
	// +kubebuilder:validation:Optional
	Requests CertificatesCapabilitySpecCertManagerInjectorResourcesRequests `json:"requests,omitempty"`

	// +kubebuilder:validation:Optional
	Limits CertificatesCapabilitySpecCertManagerInjectorResourcesLimits `json:"limits,omitempty"`
}

type CertificatesCapabilitySpecCertManagerInjectorResourcesRequests struct {
	// +kubebuilder:default="50m"
	// +kubebuilder:validation:Optional
	// (Default: "50m")
	//
	//	CPU requests to use for cert-manager CA injector deployment.
	Cpu string `json:"cpu,omitempty"`

	// +kubebuilder:default="64Mi"
	// +kubebuilder:validation:Optional
	// (Default: "64Mi")
	//
	//	Memory requests to use for cert-manager CA injector deployment.
	Memory string `json:"memory,omitempty"`
}

type CertificatesCapabilitySpecCertManagerInjectorResourcesLimits struct {
	// +kubebuilder:default="128Mi"
	// +kubebuilder:validation:Optional
	// (Default: "128Mi")
	//
	//	Memory limits to use for cert-manager CA injector deployment.
	Memory string `json:"memory,omitempty"`
}

type CertificatesCapabilitySpecCertManagerController struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the cert-manager controller deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="quay.io/jetstack/cert-manager-controller:v1.14.4"
	// +kubebuilder:validation:Optional
	// (Default: "quay.io/jetstack/cert-manager-controller:v1.14.4")
	//
	//	Image to use for cert-manager controller deployment.
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Resources CertificatesCapabilitySpecCertManagerControllerResources `json:"resources,omitempty"`
}

type CertificatesCapabilitySpecCertManagerControllerResources struct {
	// +kubebuilder:validation:Optional
	Requests CertificatesCapabilitySpecCertManagerControllerResourcesRequests `json:"requests,omitempty"`

	// +kubebuilder:validation:Optional
	Limits CertificatesCapabilitySpecCertManagerControllerResourcesLimits `json:"limits,omitempty"`
}

type CertificatesCapabilitySpecCertManagerControllerResourcesRequests struct {
	// +kubebuilder:default="25m"
	// +kubebuilder:validation:Optional
	// (Default: "25m")
	//
	//	CPU requests to use for cert-manager controller deployment.
	Cpu string `json:"cpu,omitempty"`

	// +kubebuilder:default="32Mi"
	// +kubebuilder:validation:Optional
	// (Default: "32Mi")
	//
	//	Memory requests to use for cert-manager controller deployment.
	Memory string `json:"memory,omitempty"`
}

type CertificatesCapabilitySpecCertManagerControllerResourcesLimits struct {
	// +kubebuilder:default="64Mi"
	// +kubebuilder:validation:Optional
	// (Default: "64Mi")
	//
	//	Memory limits to use for cert-manager controller deployment.
	Memory string `json:"memory,omitempty"`
}

type CertificatesCapabilitySpecCertManagerWebhook struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the cert-manager webhook deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="quay.io/jetstack/cert-manager-webhook:v1.14.4"
	// +kubebuilder:validation:Optional
	// (Default: "quay.io/jetstack/cert-manager-webhook:v1.14.4")
	//
	//	Image to use for cert-manager webhook deployment.
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Resources CertificatesCapabilitySpecCertManagerWebhookResources `json:"resources,omitempty"`
}

type CertificatesCapabilitySpecCertManagerWebhookResources struct {
	// +kubebuilder:validation:Optional
	Requests CertificatesCapabilitySpecCertManagerWebhookResourcesRequests `json:"requests,omitempty"`

	// +kubebuilder:validation:Optional
	Limits CertificatesCapabilitySpecCertManagerWebhookResourcesLimits `json:"limits,omitempty"`
}

type CertificatesCapabilitySpecCertManagerWebhookResourcesRequests struct {
	// +kubebuilder:default="25m"
	// +kubebuilder:validation:Optional
	// (Default: "25m")
	//
	//	CPU requests to use for cert-manager webhook deployment.
	Cpu string `json:"cpu,omitempty"`

	// +kubebuilder:default="32Mi"
	// +kubebuilder:validation:Optional
	// (Default: "32Mi")
	//
	//	Memory requests to use for cert-manager webhook deployment.
	Memory string `json:"memory,omitempty"`
}

type CertificatesCapabilitySpecCertManagerWebhookResourcesLimits struct {
	// +kubebuilder:default="64Mi"
	// +kubebuilder:validation:Optional
	// (Default: "64Mi")
	//
	//	Memory limits to use for cert-manager webhook deployment.
	Memory string `json:"memory,omitempty"`
}

type CertificatesCapabilitySpecTrustManager struct {
	// +kubebuilder:validation:Optional
	Controller CertificatesCapabilitySpecTrustManagerController `json:"controller,omitempty"`
}

type CertificatesCapabilitySpecTrustManagerController struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the trust-manager controller deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="quay.io/jetstack/trust-manager:v0.9.2"
	// +kubebuilder:validation:Optional
	// (Default: "quay.io/jetstack/trust-manager:v0.9.2")
	//
	//	Image to use for trust-manager controller deployment.
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Resources CertificatesCapabilitySpecTrustManagerControllerResources `json:"resources,omitempty"`
}

type CertificatesCapabilitySpecTrustManagerControllerResources struct {
	// +kubebuilder:validation:Optional
	Requests CertificatesCapabilitySpecTrustManagerControllerResourcesRequests `json:"requests,omitempty"`

	// +kubebuilder:validation:Optional
	Limits CertificatesCapabilitySpecTrustManagerControllerResourcesLimits `json:"limits,omitempty"`
}

type CertificatesCapabilitySpecTrustManagerControllerResourcesRequests struct {
	// +kubebuilder:default="25m"
	// +kubebuilder:validation:Optional
	// (Default: "25m")
	//
	//	CPU requests to use for trust-manager controller deployment.
	Cpu string `json:"cpu,omitempty"`

	// +kubebuilder:default="32Mi"
	// +kubebuilder:validation:Optional
	// (Default: "32Mi")
	//
	//	Memory requests to use for trust-manager controller deployment.
	Memory string `json:"memory,omitempty"`
}

type CertificatesCapabilitySpecTrustManagerControllerResourcesLimits struct {
	// +kubebuilder:default="64Mi"
	// +kubebuilder:validation:Optional
	// (Default: "64Mi")
	//
	//	Memory limits to use for trust-manager controller deployment.
	Memory string `json:"memory,omitempty"`
}

// CertificatesCapabilityStatus defines the observed state of CertificatesCapability.
type CertificatesCapabilityStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Created               bool                     `json:"created,omitempty"`
	DependenciesSatisfied bool                     `json:"dependenciesSatisfied,omitempty"`
	Conditions            []*status.PhaseCondition `json:"conditions,omitempty"`
	Resources             []*status.ChildResource  `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// CertificatesCapability is the Schema for the certificatescapabilities API.
type CertificatesCapability struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificatesCapabilitySpec   `json:"spec,omitempty"`
	Status            CertificatesCapabilityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificatesCapabilityList contains a list of CertificatesCapability.
type CertificatesCapabilityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificatesCapability `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *CertificatesCapability) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *CertificatesCapability) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *CertificatesCapability) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *CertificatesCapability) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *CertificatesCapability) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *CertificatesCapability) SetPhaseCondition(condition *status.PhaseCondition) {
	for i, currentCondition := range component.GetPhaseConditions() {
		if currentCondition.Phase == condition.Phase {
			component.Status.Conditions[i] = condition

			return
		}
	}

	// phase not found, lets add it to the list.
	component.Status.Conditions = append(component.Status.Conditions, condition)
}

// GetResources returns the child resource status for a component.
func (component *CertificatesCapability) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *CertificatesCapability) SetChildResourceCondition(resource *status.ChildResource) {
	for i, currentResource := range component.GetChildResourceConditions() {
		if currentResource.Group == resource.Group && currentResource.Version == resource.Version && currentResource.Kind == resource.Kind {
			if currentResource.Name == resource.Name && currentResource.Namespace == resource.Namespace {
				component.Status.Resources[i] = resource

				return
			}
		}
	}

	// phase not found, lets add it to the collection
	component.Status.Resources = append(component.Status.Resources, resource)
}

// GetDependencies returns the dependencies for a component.
func (*CertificatesCapability) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*CertificatesCapability) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("CertificatesCapability")
}

func init() {
	SchemeBuilder.Register(&CertificatesCapability{}, &CertificatesCapabilityList{})
}
