// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1.KubeFed":       schema_pkg_apis_operator_v1alpha1_KubeFed(ref),
		"github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1.KubeFedSpec":   schema_pkg_apis_operator_v1alpha1_KubeFedSpec(ref),
		"github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1.KubeFedStatus": schema_pkg_apis_operator_v1alpha1_KubeFedStatus(ref),
	}
}

func schema_pkg_apis_operator_v1alpha1_KubeFed(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KubeFed is the Schema for the installs API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1.KubeFedSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1.KubeFedStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1.KubeFedSpec", "github.com/openshift/kubefed-operator/pkg/apis/operator/v1alpha1.KubeFedStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1alpha1_KubeFedSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KubeFedSpec defines the desired state of KubeFed",
				Properties: map[string]spec.Schema{
					"scope": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"scope"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_operator_v1alpha1_KubeFedStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KubeFedStatus defines the observed state of KubeFed",
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "The version of the installed release",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
