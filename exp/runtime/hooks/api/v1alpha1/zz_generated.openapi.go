//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.DiscoveryRequest":            schema_runtime_hooks_api_v1alpha1_DiscoveryRequest(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.DiscoveryResponse":           schema_runtime_hooks_api_v1alpha1_DiscoveryResponse(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ExtensionHandler":            schema_runtime_hooks_api_v1alpha1_ExtensionHandler(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GeneratePatchesRequest":      schema_runtime_hooks_api_v1alpha1_GeneratePatchesRequest(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GeneratePatchesRequestItem":  schema_runtime_hooks_api_v1alpha1_GeneratePatchesRequestItem(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GeneratePatchesResponse":     schema_runtime_hooks_api_v1alpha1_GeneratePatchesResponse(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GeneratePatchesResponseItem": schema_runtime_hooks_api_v1alpha1_GeneratePatchesResponseItem(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GroupVersionHook":            schema_runtime_hooks_api_v1alpha1_GroupVersionHook(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.HolderReference":             schema_runtime_hooks_api_v1alpha1_HolderReference(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ValidateTopologyRequest":     schema_runtime_hooks_api_v1alpha1_ValidateTopologyRequest(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ValidateTopologyRequestItem": schema_runtime_hooks_api_v1alpha1_ValidateTopologyRequestItem(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ValidateTopologyResponse":    schema_runtime_hooks_api_v1alpha1_ValidateTopologyResponse(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable":                    schema_runtime_hooks_api_v1alpha1_Variable(ref),
	}
}

func schema_runtime_hooks_api_v1alpha1_DiscoveryRequest(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DiscoveryRequest represents the object of a discovery request.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_runtime_hooks_api_v1alpha1_DiscoveryResponse(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DiscoveryResponse represents the object received as a discovery response.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the call. One of \"Success\" or \"Failure\".\n\nPossible enum values:\n - `\"Failure\"` represents a failure response.\n - `\"Success\"` represents the success response.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"Failure", "Success"}},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human-readable description of the status of the call.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"handlers": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Handlers defines the current ExtensionHandlers supported by an Extension.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ExtensionHandler"),
									},
								},
							},
						},
					},
				},
				Required: []string{"status", "message", "handlers"},
			},
		},
		Dependencies: []string{
			"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ExtensionHandler"},
	}
}

func schema_runtime_hooks_api_v1alpha1_ExtensionHandler(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExtensionHandler represents the discovery information of the extension which includes the hook it supports.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the ExtensionHandler.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"requestHook": {
						SchemaProps: spec.SchemaProps{
							Description: "RequestHook defines the versioned runtime hook which this ExtensionHandler serves.",
							Default:     map[string]interface{}{},
							Ref:         ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GroupVersionHook"),
						},
					},
					"timeoutSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "TimeoutSeconds defines the timeout duration for client calls to the ExtensionHandler.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"failurePolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "FailurePolicy defines how failures in calls to the ExtensionHandler should be handled by a client.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "requestHook"},
			},
		},
		Dependencies: []string{
			"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GroupVersionHook"},
	}
}

func schema_runtime_hooks_api_v1alpha1_GeneratePatchesRequest(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GeneratePatchesRequest is the request of the GeneratePatches hook.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"variables": {
						SchemaProps: spec.SchemaProps{
							Description: "Variables are global variables for all templates.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable"),
									},
								},
							},
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "Items is the list of templates to generate patches for.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GeneratePatchesRequestItem"),
									},
								},
							},
						},
					},
				},
				Required: []string{"variables", "items"},
			},
		},
		Dependencies: []string{
			"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GeneratePatchesRequestItem", "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable"},
	}
}

func schema_runtime_hooks_api_v1alpha1_GeneratePatchesRequestItem(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GeneratePatchesRequestItem represents a template to generate patches for.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"uid": {
						SchemaProps: spec.SchemaProps{
							Description: "UID is an identifier for this template. It allows us to correlate the template in the request with the corresponding generates patches in the response.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"holderReference": {
						SchemaProps: spec.SchemaProps{
							Description: "HolderReference is a reference to the object where the template is used.",
							Default:     map[string]interface{}{},
							Ref:         ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.HolderReference"),
						},
					},
					"object": {
						SchemaProps: spec.SchemaProps{
							Description: "Object contains the template as a raw object.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
						},
					},
					"variables": {
						SchemaProps: spec.SchemaProps{
							Description: "Variables are variables specific for the current template. For example some builtin variables like MachineDeployment replicas and version are context-sensitive and thus are only added to templates for MachineDeployments and with values which correspond to the current MachineDeployment.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable"),
									},
								},
							},
						},
					},
				},
				Required: []string{"uid", "holderReference", "object", "variables"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/runtime.RawExtension", "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.HolderReference", "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable"},
	}
}

func schema_runtime_hooks_api_v1alpha1_GeneratePatchesResponse(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GeneratePatchesResponse is the response of the GeneratePatches hook. NOTE: The patches in GeneratePatchesResponse will be applied in the order in which they are defined to the templates of the request. Thus applying changes consecutively when iterating through internal and external patches.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the call. One of: \"Success\" or \"Failure\".\n\nPossible enum values:\n - `\"Failure\"` represents a failure response.\n - `\"Success\"` represents the success response.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"Failure", "Success"}},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human-readable description of the status of the call.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "Items is the list of generated patches.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GeneratePatchesResponseItem"),
									},
								},
							},
						},
					},
				},
				Required: []string{"status", "items"},
			},
		},
		Dependencies: []string{
			"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GeneratePatchesResponseItem"},
	}
}

func schema_runtime_hooks_api_v1alpha1_GeneratePatchesResponseItem(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GeneratePatchesResponseItem is a generated patch.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"uid": {
						SchemaProps: spec.SchemaProps{
							Description: "UID identifies the corresponding template in the request on which the patch should be applied.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"patchType": {
						SchemaProps: spec.SchemaProps{
							Description: "PatchType defines the type of the patch. One of: \"JSONPatch\" or \"JSONMergePatch\".\n\nPossible enum values:\n - `\"JSONMergePatch\"` identifies a https://datatracker.ietf.org/doc/html/rfc7386 JSON merge patch.\n - `\"JSONPatch\"` identifies a https://datatracker.ietf.org/doc/html/rfc6902 JSON patch.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"JSONMergePatch", "JSONPatch"}},
					},
					"patch": {
						SchemaProps: spec.SchemaProps{
							Description: "Patch contains the patch which should be applied to the template. It must be of the corresponding PatchType.",
							Type:        []string{"string"},
							Format:      "byte",
						},
					},
				},
				Required: []string{"uid", "patchType", "patch"},
			},
		},
	}
}

func schema_runtime_hooks_api_v1alpha1_GroupVersionHook(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupVersionHook defines the runtime hook when the ExtensionHandler is called.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion is the Version of the Hook",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"hook": {
						SchemaProps: spec.SchemaProps{
							Description: "Hook is the name of the hook",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"apiVersion", "hook"},
			},
		},
	}
}

func schema_runtime_hooks_api_v1alpha1_HolderReference(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HolderReference represents a reference to an object which holds a template.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "API version of the referent.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"fieldPath": {
						SchemaProps: spec.SchemaProps{
							Description: "FieldPath is the path to the field of the object which references the template.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"apiVersion", "kind", "namespace", "name", "fieldPath"},
			},
		},
	}
}

func schema_runtime_hooks_api_v1alpha1_ValidateTopologyRequest(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ValidateTopologyRequest is the request of the ValidateTopology hook.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"variables": {
						SchemaProps: spec.SchemaProps{
							Description: "Variables are global variables for all templates.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable"),
									},
								},
							},
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "Items is the list of templates to validate.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ValidateTopologyRequestItem"),
									},
								},
							},
						},
					},
				},
				Required: []string{"variables", "items"},
			},
		},
		Dependencies: []string{
			"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ValidateTopologyRequestItem", "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable"},
	}
}

func schema_runtime_hooks_api_v1alpha1_ValidateTopologyRequestItem(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ValidateTopologyRequestItem represents a template to validate.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"holderReference": {
						SchemaProps: spec.SchemaProps{
							Description: "HolderReference is a reference to the object where the template is used.",
							Default:     map[string]interface{}{},
							Ref:         ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.HolderReference"),
						},
					},
					"object": {
						SchemaProps: spec.SchemaProps{
							Description: "Object contains the template as a raw object.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
						},
					},
					"variables": {
						SchemaProps: spec.SchemaProps{
							Description: "Variables are variables specific for the current template. For example some builtin variables like MachineDeployment replicas and version are context-sensitive and thus are only added to templates for MachineDeployments and with values which correspond to the current MachineDeployment.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable"),
									},
								},
							},
						},
					},
				},
				Required: []string{"holderReference", "object", "variables"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/runtime.RawExtension", "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.HolderReference", "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.Variable"},
	}
}

func schema_runtime_hooks_api_v1alpha1_ValidateTopologyResponse(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ValidateTopologyResponse is the response of the ValidateTopology hook.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the call. One of: \"Success\" or \"Failure\".\n\nPossible enum values:\n - `\"Failure\"` represents a failure response.\n - `\"Success\"` represents the success response.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"Failure", "Success"}},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human-readable description of the status of the call.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"status", "message"},
			},
		},
	}
}

func schema_runtime_hooks_api_v1alpha1_Variable(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Variable represents a variable value.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the variable.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Description: "Value of the variable.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON"),
						},
					},
				},
				Required: []string{"name", "value"},
			},
		},
		Dependencies: []string{
			"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON"},
	}
}
