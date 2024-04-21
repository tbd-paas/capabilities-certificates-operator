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

// CreateDeploymentNamespaceCertManagerCainjector creates the Deployment resource with name cert-manager-cainjector.
func CreateDeploymentNamespaceCertManagerCainjector(
	parent *capabilitiesv1alpha1.CertificatesCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-cainjector",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                                  "cainjector",
					"app.kubernetes.io/name":               "cainjector",
					"app.kubernetes.io/instance":           "cert-manager",
					"app.kubernetes.io/component":          "cainjector",
					"app.kubernetes.io/version":            "v1.14.4",
					"capabilities.tbd.io/capability":       "certificates",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "certificates-operator",
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: certManager.injector.replicas
				//  Number of replicas to use for the cert-manager cainjector deployment.
				"replicas": parent.Spec.CertManager.Injector.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":      "cainjector",
						"app.kubernetes.io/instance":  "cert-manager",
						"app.kubernetes.io/component": "cainjector",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                                  "cainjector",
							"app.kubernetes.io/name":               "cainjector",
							"app.kubernetes.io/instance":           "cert-manager",
							"app.kubernetes.io/component":          "cainjector",
							"app.kubernetes.io/version":            "v1.14.4",
							"capabilities.tbd.io/capability":       "certificates",
							"capabilities.tbd.io/version":          "v0.0.1",
							"capabilities.tbd.io/platform-version": "unstable",
							"app.kubernetes.io/part-of":            "platform",
							"app.kubernetes.io/managed-by":         "certificates-operator",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "cert-manager-cainjector",
						"enableServiceLinks": false,
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "cert-manager-cainjector",
								// controlled by field: certManager.injector.image
								//  Image to use for cert-manager CA injector deployment.
								"image":           parent.Spec.CertManager.Injector.Image,
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"--v=2",
									"--leader-election-namespace=$(POD_NAMESPACE)",
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
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
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										// controlled by field: certManager.injector.resources.requests.cpu
										//  CPU requests to use for cert-manager CA injector deployment.
										"cpu": parent.Spec.CertManager.Injector.Resources.Requests.Cpu,
										// controlled by field: certManager.injector.resources.requests.memory
										//  Memory requests to use for cert-manager CA injector deployment.
										"memory": parent.Spec.CertManager.Injector.Resources.Requests.Memory,
									},
									"limits": map[string]interface{}{
										// controlled by field: certManager.injector.resources.limits.memory
										//  Memory limits to use for cert-manager CA injector deployment.
										"memory": parent.Spec.CertManager.Injector.Resources.Limits.Memory,
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":   "linux",
							"tbd.io/node-type":   "platform",
							"kubernetes.io/arch": "arm64",
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
															"cert-manager-cainjector",
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

	return mutate.MutateDeploymentNamespaceCertManagerCainjector(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceCertManager creates the Deployment resource with name cert-manager.
func CreateDeploymentNamespaceCertManager(
	parent *capabilitiesv1alpha1.CertificatesCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "cert-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                                  "cert-manager",
					"app.kubernetes.io/name":               "cert-manager",
					"app.kubernetes.io/instance":           "cert-manager",
					"app.kubernetes.io/component":          "controller",
					"app.kubernetes.io/version":            "v1.14.4",
					"capabilities.tbd.io/capability":       "certificates",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "certificates-operator",
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: certManager.controller.replicas
				//  Number of replicas to use for the cert-manager controller deployment.
				"replicas": parent.Spec.CertManager.Controller.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":      "cert-manager",
						"app.kubernetes.io/instance":  "cert-manager",
						"app.kubernetes.io/component": "controller",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                                  "cert-manager",
							"app.kubernetes.io/name":               "cert-manager",
							"app.kubernetes.io/instance":           "cert-manager",
							"app.kubernetes.io/component":          "controller",
							"app.kubernetes.io/version":            "v1.14.4",
							"capabilities.tbd.io/capability":       "certificates",
							"capabilities.tbd.io/version":          "v0.0.1",
							"capabilities.tbd.io/platform-version": "unstable",
							"app.kubernetes.io/part-of":            "platform",
							"app.kubernetes.io/managed-by":         "certificates-operator",
						},
						"annotations": map[string]interface{}{
							"prometheus.io/path":   "/metrics",
							"prometheus.io/scrape": "true",
							"prometheus.io/port":   "9402",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "cert-manager",
						"enableServiceLinks": false,
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "cert-manager-controller",
								// controlled by field: certManager.controller.image
								//  Image to use for cert-manager controller deployment.
								"image":           parent.Spec.CertManager.Controller.Image,
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"--v=2",
									"--cluster-resource-namespace=$(POD_NAMESPACE)",
									"--leader-election-namespace=$(POD_NAMESPACE)",
									"--acme-http01-solver-image=quay.io/jetstack/cert-manager-acmesolver:v1.14.4",
									"--max-concurrent-challenges=60",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 9402,
										"name":          "http-metrics",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 9403,
										"name":          "http-healthz",
										"protocol":      "TCP",
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
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
								},
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"port":   "http-healthz",
										"path":   "/livez",
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 10,
									"periodSeconds":       10,
									"timeoutSeconds":      15,
									"successThreshold":    1,
									"failureThreshold":    8,
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										// controlled by field: certManager.controller.resources.requests.cpu
										//  CPU requests to use for cert-manager controller deployment.
										"cpu": parent.Spec.CertManager.Controller.Resources.Requests.Cpu,
										// controlled by field: certManager.controller.resources.requests.memory
										//  Memory requests to use for cert-manager controller deployment.
										"memory": parent.Spec.CertManager.Controller.Resources.Requests.Memory,
									},
									"limits": map[string]interface{}{
										// controlled by field: certManager.controller.resources.limits.memory
										//  Memory limits to use for cert-manager controller deployment.
										"memory": parent.Spec.CertManager.Controller.Resources.Limits.Memory,
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":   "linux",
							"tbd.io/node-type":   "platform",
							"kubernetes.io/arch": "arm64",
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
															"cert-manager",
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

	return mutate.MutateDeploymentNamespaceCertManager(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceCertManagerWebhook creates the Deployment resource with name cert-manager-webhook.
func CreateDeploymentNamespaceCertManagerWebhook(
	parent *capabilitiesv1alpha1.CertificatesCapability,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                                  "webhook",
					"app.kubernetes.io/name":               "webhook",
					"app.kubernetes.io/instance":           "cert-manager",
					"app.kubernetes.io/component":          "webhook",
					"app.kubernetes.io/version":            "v1.14.4",
					"capabilities.tbd.io/capability":       "certificates",
					"capabilities.tbd.io/version":          "v0.0.1",
					"capabilities.tbd.io/platform-version": "unstable",
					"app.kubernetes.io/part-of":            "platform",
					"app.kubernetes.io/managed-by":         "certificates-operator",
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: certManager.webhook.replicas
				//  Number of replicas to use for the cert-manager webhook deployment.
				"replicas": parent.Spec.CertManager.Webhook.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":      "webhook",
						"app.kubernetes.io/instance":  "cert-manager",
						"app.kubernetes.io/component": "webhook",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                                  "webhook",
							"app.kubernetes.io/name":               "webhook",
							"app.kubernetes.io/instance":           "cert-manager",
							"app.kubernetes.io/component":          "webhook",
							"app.kubernetes.io/version":            "v1.14.4",
							"capabilities.tbd.io/capability":       "certificates",
							"capabilities.tbd.io/version":          "v0.0.1",
							"capabilities.tbd.io/platform-version": "unstable",
							"app.kubernetes.io/part-of":            "platform",
							"app.kubernetes.io/managed-by":         "certificates-operator",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "cert-manager-webhook",
						"enableServiceLinks": false,
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "cert-manager-webhook",
								// controlled by field: certManager.webhook.image
								//  Image to use for cert-manager webhook deployment.
								"image":           parent.Spec.CertManager.Webhook.Image,
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"--v=2",
									"--secure-port=10250",
									"--dynamic-serving-ca-secret-namespace=$(POD_NAMESPACE)",
									"--dynamic-serving-ca-secret-name=cert-manager-webhook-ca",
									"--dynamic-serving-dns-names=cert-manager-webhook",
									"--dynamic-serving-dns-names=cert-manager-webhook.$(POD_NAMESPACE)",
									"--dynamic-serving-dns-names=cert-manager-webhook.$(POD_NAMESPACE).svc",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"name":          "https",
										"protocol":      "TCP",
										"containerPort": 10250,
									},
									map[string]interface{}{
										"name":          "healthcheck",
										"protocol":      "TCP",
										"containerPort": 6080,
									},
								},
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/livez",
										"port":   6080,
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 60,
									"periodSeconds":       10,
									"timeoutSeconds":      1,
									"successThreshold":    1,
									"failureThreshold":    3,
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/healthz",
										"port":   6080,
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       5,
									"timeoutSeconds":      1,
									"successThreshold":    1,
									"failureThreshold":    3,
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
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										// controlled by field: certManager.webhook.resources.requests.cpu
										//  CPU requests to use for cert-manager webhook deployment.
										"cpu": parent.Spec.CertManager.Webhook.Resources.Requests.Cpu,
										// controlled by field: certManager.webhook.resources.requests.memory
										//  Memory requests to use for cert-manager webhook deployment.
										"memory": parent.Spec.CertManager.Webhook.Resources.Requests.Memory,
									},
									"limits": map[string]interface{}{
										// controlled by field: certManager.webhook.resources.limits.memory
										//  Memory limits to use for cert-manager webhook deployment.
										"memory": parent.Spec.CertManager.Webhook.Resources.Limits.Memory,
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os":   "linux",
							"tbd.io/node-type":   "platform",
							"kubernetes.io/arch": "arm64",
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
															"cert-manager-cainjector",
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

	return mutate.MutateDeploymentNamespaceCertManagerWebhook(resourceObj, parent, reconciler, req)
}
