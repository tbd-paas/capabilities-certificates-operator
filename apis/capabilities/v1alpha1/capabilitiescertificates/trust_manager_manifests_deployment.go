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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceTrustManager creates the Deployment resource with name trust-manager.
func CreateDeploymentNamespaceTrustManager(
	parent *capabilitiesv1alpha1.CertificatesCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
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
				// controlled by field: trustManager.controller.replicas
				//  Number of replicas to use for the trust-manager controller deployment.
				"replicas": parent.Spec.TrustManager.Controller.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":      "trust-manager",
						"app.kubernetes.io/instance":  "trust-manager",
						"app.kubernetes.io/component": "trust-manager",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
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
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "trust-manager",
						"initContainers": []interface{}{
							map[string]interface{}{
								"name":            "cert-manager-package-debian",
								"image":           "quay.io/jetstack/cert-manager-package-debian:20210119.0",
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"/copyandmaybepause",
									"/debian-package",
									"/packages",
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/packages",
										"name":      "packages",
										"readOnly":  false,
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
									"runAsNonRoot":           true,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
								},
							},
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name":            "trust-manager",
								"image":           "quay.io/jetstack/trust-manager:v0.9.2",
								"imagePullPolicy": "IfNotPresent",
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 6443,
									},
									map[string]interface{}{
										"containerPort": 9402,
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"port": 6060,
										"path": "/readyz",
									},
									"initialDelaySeconds": 3,
									"periodSeconds":       7,
								},
								"command": []interface{}{
									"trust-manager",
								},
								"args": []interface{}{
									"--log-level=1",
									"--metrics-port=9402",
									"--readiness-probe-port=6060",
									"--readiness-probe-path=/readyz",
									"--trust-namespace=nukleros-certs-system",
									"--webhook-host=0.0.0.0",
									"--webhook-port=6443",
									"--webhook-certificate-dir=/tls",
									"--default-package-location=/packages/cert-manager-package-debian.json",
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/tls",
										"name":      "tls",
										"readOnly":  true,
									},
									map[string]interface{}{
										"mountPath": "/packages",
										"name":      "packages",
										"readOnly":  true,
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "25m",
										"memory": "32Mi",
									},
									"limits": map[string]interface{}{
										"memory": "64Mi",
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   true,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":   "linux",
							"tbd.io/node-type":   "platform",
							"kubernetes.io/arch": "arm64",
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "packages",
								"emptyDir": map[string]interface{}{
									"sizeLimit": "50M",
								},
							},
							map[string]interface{}{
								"name": "tls",
								"secret": map[string]interface{}{
									"defaultMode": 420,
									"secretName":  "trust-manager-tls",
								},
							},
						},
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"weight": 100,
										"podAffinityTerm": map[string]interface{}{
											"topologyKey": "kubernetes.io/hostname",
											"labelSelector": map[string]interface{}{
												"matchExpressions": []interface{}{
													map[string]interface{}{
														"key":      "app.kubernetes.io/name",
														"operator": "In",
														"values": []interface{}{
															"trust-manager",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceTrustManager(resourceObj, parent, reconciler, req)
}
