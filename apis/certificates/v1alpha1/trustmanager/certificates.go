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

package trustmanager

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	certificatesv1alpha1 "github.com/tbd-paas/capabilities-certificates-operator/apis/certificates/v1alpha1"
	"github.com/tbd-paas/capabilities-certificates-operator/apis/certificates/v1alpha1/trustmanager/mutate"
)

// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete

// CreateCertNamespaceTrustManager creates the Certificate resource with name trust-manager.
func CreateCertNamespaceTrustManager(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Certificate",
			"metadata": map[string]interface{}{
				"name":      "trust-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":               "trust-manager",
					"app.kubernetes.io/instance":           "trust-manager",
					"app.kubernetes.io/version":            "v0.9.2",
					"app":                                  "trust-manager",
					"app.kubernetes.io/component":          "trust-manager",
					"capabilities.tbd.io/capability":       "certificates",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "certificates-operator",
				},
			},
			"spec": map[string]interface{}{
				"commonName": "trust-manager.tbd-certificates-system.svc",
				"dnsNames": []interface{}{
					"trust-manager",
					"trust-manager.tbd-certificates-system",
					"trust-manager.tbd-certificates-system.svc",
					"trust-manager.tbd-certificates-system.svc.local",
				},
				"secretName":           "trust-manager",
				"revisionHistoryLimit": 1,
				"issuerRef": map[string]interface{}{
					"name": "internal",
					"kind": "ClusterIssuer",
				},
			},
		},
	}

	return mutate.MutateCertNamespaceTrustManager(resourceObj, parent, reconciler, req)
}
