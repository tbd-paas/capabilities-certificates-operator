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

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDBundlesTrustCertManagerIo creates the CustomResourceDefinition resource with name bundles.trust.cert-manager.io.
func CreateCRDBundlesTrustCertManagerIo(
	parent *certificatesv1alpha1.TrustManager,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name":        "bundles.trust.cert-manager.io",
				"annotations": map[string]interface{}{},
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
				"group": "trust.cert-manager.io",
				"names": map[string]interface{}{
					"kind":     "Bundle",
					"listKind": "BundleList",
					"plural":   "bundles",
					"singular": "bundle",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Bundle Target Key",
								"jsonPath":    ".status.target.configMap.key",
								"name":        "Target",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Bundle has been synced",
								"jsonPath":    ".status.conditions[?(@.type == \"Synced\")].status",
								"name":        "Synced",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Reason Bundle has Synced status",
								"jsonPath":    ".status.conditions[?(@.type == \"Synced\")].reason",
								"name":        "Reason",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Timestamp Bundle was created",
								"jsonPath":    ".metadata.creationTimestamp",
								"name":        "Age",
								"type":        "date",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"type": "object",
								"required": []interface{}{
									"spec",
								},
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "Desired state of the Bundle resource.",
										"type":        "object",
										"required": []interface{}{
											"sources",
											"target",
										},
										"properties": map[string]interface{}{
											"sources": map[string]interface{}{
												"description": "Sources is a set of references to data whose data will sync to the target.",
												"type":        "array",
												"items": map[string]interface{}{
													"description": `BundleSource is the set of sources whose data will be appended and synced to
the BundleTarget in all Namespaces.`,
													"type": "object",
													"properties": map[string]interface{}{
														"configMap": map[string]interface{}{
															"description": `ConfigMap is a reference (by name) to a ConfigMap's ` + "`" + `data` + "`" + ` key, or to a
list of ConfigMap's ` + "`" + `data` + "`" + ` key using label selector, in the trust Namespace.`,
															"type": "object",
															"required": []interface{}{
																"key",
															},
															"properties": map[string]interface{}{
																"key": map[string]interface{}{
																	"description": "Key is the key of the entry in the object's `data` field to be used.",
																	"type":        "string",
																},
																"name": map[string]interface{}{
																	"description": `Name is the name of the source object in the trust Namespace.
This field must be left empty when ` + "`" + `selector` + "`" + ` is set`,
																	"type": "string",
																},
																"selector": map[string]interface{}{
																	"description": `Selector is the label selector to use to fetch a list of objects. Must not be set
when ` + "`" + `Name` + "`" + ` is set.`,
																	"type": "object",
																	"properties": map[string]interface{}{
																		"matchExpressions": map[string]interface{}{
																			"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			"type":        "array",
																			"items": map[string]interface{}{
																				"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																				"type": "object",
																				"required": []interface{}{
																					"key",
																					"operator",
																				},
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "key is the label key that the selector applies to.",
																						"type":        "string",
																					},
																					"operator": map[string]interface{}{
																						"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																						"type": "string",
																					},
																					"values": map[string]interface{}{
																						"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																						"type": "array",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																					},
																				},
																			},
																		},
																		"matchLabels": map[string]interface{}{
																			"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																			"type": "object",
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																		},
																	},
																	"x-kubernetes-map-type": "atomic",
																},
															},
														},
														"inLine": map[string]interface{}{
															"description": "InLine is a simple string to append as the source data.",
															"type":        "string",
														},
														"secret": map[string]interface{}{
															"description": `Secret is a reference (by name) to a Secret's ` + "`" + `data` + "`" + ` key, or to a
list of Secret's ` + "`" + `data` + "`" + ` key using label selector, in the trust Namespace.`,
															"type": "object",
															"required": []interface{}{
																"key",
															},
															"properties": map[string]interface{}{
																"key": map[string]interface{}{
																	"description": "Key is the key of the entry in the object's `data` field to be used.",
																	"type":        "string",
																},
																"name": map[string]interface{}{
																	"description": `Name is the name of the source object in the trust Namespace.
This field must be left empty when ` + "`" + `selector` + "`" + ` is set`,
																	"type": "string",
																},
																"selector": map[string]interface{}{
																	"description": `Selector is the label selector to use to fetch a list of objects. Must not be set
when ` + "`" + `Name` + "`" + ` is set.`,
																	"type": "object",
																	"properties": map[string]interface{}{
																		"matchExpressions": map[string]interface{}{
																			"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			"type":        "array",
																			"items": map[string]interface{}{
																				"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																				"type": "object",
																				"required": []interface{}{
																					"key",
																					"operator",
																				},
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "key is the label key that the selector applies to.",
																						"type":        "string",
																					},
																					"operator": map[string]interface{}{
																						"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																						"type": "string",
																					},
																					"values": map[string]interface{}{
																						"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																						"type": "array",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																					},
																				},
																			},
																		},
																		"matchLabels": map[string]interface{}{
																			"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																			"type": "object",
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																		},
																	},
																	"x-kubernetes-map-type": "atomic",
																},
															},
														},
														"useDefaultCAs": map[string]interface{}{
															"description": `UseDefaultCAs, when true, requests the default CA bundle to be used as a source.
Default CAs are available if trust-manager was installed via Helm
or was otherwise set up to include a package-injecting init container by using the
"--default-package-location" flag when starting the trust-manager controller.
If default CAs were not configured at start-up, any request to use the default
CAs will fail.
The version of the default CA package which is used for a Bundle is stored in the
defaultCAPackageVersion field of the Bundle's status field.`,
															"type": "boolean",
														},
													},
												},
											},
											"target": map[string]interface{}{
												"description": "Target is the target location in all namespaces to sync source data to.",
												"type":        "object",
												"properties": map[string]interface{}{
													"additionalFormats": map[string]interface{}{
														"description": "AdditionalFormats specifies any additional formats to write to the target",
														"type":        "object",
														"properties": map[string]interface{}{
															"jks": map[string]interface{}{
																"description": `JKS requests a JKS-formatted binary trust bundle to be written to the target.
The bundle has "changeit" as the default password.
For more information refer to this link https://cert-manager.io/docs/faq/#keystore-passwords`,
																"type": "object",
																"required": []interface{}{
																	"key",
																},
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "Key is the key of the entry in the object's `data` field to be used.",
																		"type":        "string",
																	},
																	"password": map[string]interface{}{
																		"description": "Password for JKS trust store",
																		"type":        "string",
																		"default":     "changeit",
																		"maxLength":   128,
																		"minLength":   1,
																	},
																},
															},
															"pkcs12": map[string]interface{}{
																"description": `PKCS12 requests a PKCS12-formatted binary trust bundle to be written to the target.
The bundle is by default created without a password.`,
																"type": "object",
																"required": []interface{}{
																	"key",
																},
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "Key is the key of the entry in the object's `data` field to be used.",
																		"type":        "string",
																	},
																	"password": map[string]interface{}{
																		"description": "Password for PKCS12 trust store",
																		"type":        "string",
																		"default":     "",
																		"maxLength":   128,
																	},
																},
															},
														},
													},
													"configMap": map[string]interface{}{
														"description": `ConfigMap is the target ConfigMap in Namespaces that all Bundle source
data will be synced to.`,
														"type": "object",
														"required": []interface{}{
															"key",
														},
														"properties": map[string]interface{}{
															"key": map[string]interface{}{
																"description": "Key is the key of the entry in the object's `data` field to be used.",
																"type":        "string",
															},
														},
													},
													"namespaceSelector": map[string]interface{}{
														"description": `NamespaceSelector will, if set, only sync the target resource in
Namespaces which match the selector.`,
														"type": "object",
														"properties": map[string]interface{}{
															"matchLabels": map[string]interface{}{
																"description": `MatchLabels matches on the set of labels that must be present on a
Namespace for the Bundle target to be synced there.`,
																"type": "object",
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
															},
														},
													},
													"secret": map[string]interface{}{
														"description": `Secret is the target Secret that all Bundle source data will be synced to.
Using Secrets as targets is only supported if enabled at trust-manager startup.
By default, trust-manager has no permissions for writing to secrets and can only read secrets in the trust namespace.`,
														"type": "object",
														"required": []interface{}{
															"key",
														},
														"properties": map[string]interface{}{
															"key": map[string]interface{}{
																"description": "Key is the key of the entry in the object's `data` field to be used.",
																"type":        "string",
															},
														},
													},
												},
											},
										},
									},
									"status": map[string]interface{}{
										"description": "Status of the Bundle. This is set and managed automatically.",
										"type":        "object",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"description": `List of status conditions to indicate the status of the Bundle.
Known condition types are ` + "`" + `Bundle` + "`" + `.`,
												"type": "array",
												"items": map[string]interface{}{
													"description": "BundleCondition contains condition information for a Bundle.",
													"type":        "object",
													"required": []interface{}{
														"lastTransitionTime",
														"reason",
														"status",
														"type",
													},
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"description": `LastTransitionTime is the timestamp corresponding to the last status
change of this condition.`,
															"type":   "string",
															"format": "date-time",
														},
														"message": map[string]interface{}{
															"description": `Message is a human-readable description of the details of the last
transition, complementing reason.`,
															"type":      "string",
															"maxLength": 32768,
														},
														"observedGeneration": map[string]interface{}{
															"description": `If set, this represents the .metadata.generation that the condition was
set based upon.
For instance, if .metadata.generation is currently 12, but the
.status.condition[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the Bundle.`,
															"type":    "integer",
															"format":  "int64",
															"minimum": 0,
														},
														"reason": map[string]interface{}{
															"description": `Reason is a brief machine-readable explanation for the condition's last
transition.
The value should be a CamelCase string.
This field may not be empty.`,
															"type":      "string",
															"maxLength": 1024,
															"minLength": 1,
															"pattern":   "^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$",
														},
														"status": map[string]interface{}{
															"description": "Status of the condition, one of True, False, Unknown.",
															"type":        "string",
															"enum": []interface{}{
																"True",
																"False",
																"Unknown",
															},
														},
														"type": map[string]interface{}{
															"description": "Type of the condition, known values are (`Synced`).",
															"type":        "string",
															"maxLength":   316,
															"pattern":     `^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`,
														},
													},
												},
												"x-kubernetes-list-map-keys": []interface{}{
													"type",
												},
												"x-kubernetes-list-type": "map",
											},
											"defaultCAVersion": map[string]interface{}{
												"description": `DefaultCAPackageVersion, if set and non-empty, indicates the version information
which was retrieved when the set of default CAs was requested in the bundle
source. This should only be set if useDefaultCAs was set to "true" on a source,
and will be the same for the same version of a bundle with identical certificates.`,
												"type": "string",
											},
										},
									},
								},
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDBundlesTrustCertManagerIo(resourceObj, parent, reconciler, req)
}
