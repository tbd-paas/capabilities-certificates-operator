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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	capabilitiesv1alpha1 "github.com/tbd-paas/capabilities-certificates-operator/apis/capabilities/v1alpha1"
	"github.com/tbd-paas/capabilities-certificates-operator/apis/capabilities/v1alpha1/capabilitiescertificates/mutate"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookTrustManager creates the ValidatingWebhookConfiguration resource with name trust-manager.
func CreateValidatingWebhookTrustManager(
	parent *capabilitiesv1alpha1.CertificatesCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "trust-manager",
				"labels": map[string]interface{}{
					"app":                                  "trust-manager",
					"app.kubernetes.io/name":               "trust-manager",
					"app.kubernetes.io/instance":           "trust-manager",
					"app.kubernetes.io/version":            "v0.9.2",
					"app.kubernetes.io/component":          "trust-manager",
					"capabilities.tbd.io/capability":       "certificates",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "certificates-operator",
				},
				"annotations": map[string]interface{}{
					// controlled by field: namespace
					"cert-manager.io/inject-ca-from": "" + parent.Spec.Namespace + "/trust-manager",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "trust.cert-manager.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"trust.cert-manager.io",
							},
							"apiVersions": []interface{}{
								"*",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"*/*",
							},
						},
					},
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"timeoutSeconds": 5,
					"failurePolicy":  "Fail",
					"sideEffects":    "None",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "trust-manager",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-trust-cert-manager-io-v1alpha1-bundle",
						},
					},
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookTrustManager(resourceObj, parent, reconciler, req)
}
