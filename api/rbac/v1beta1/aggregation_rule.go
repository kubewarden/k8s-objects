// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// AggregationRule AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
//
// swagger:model AggregationRule
type AggregationRule struct {

	// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
	ClusterRoleSelectors []apimachinery_pkg_apis_meta_v1.LabelSelector `json:"clusterRoleSelectors"`
}
