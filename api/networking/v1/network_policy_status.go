// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// NetworkPolicyStatus NetworkPolicyStatus describes the current state of the NetworkPolicy.
//
// swagger:model NetworkPolicyStatus
type NetworkPolicyStatus struct {

	// conditions holds an array of metav1.Condition that describe the state of the NetworkPolicy. Current service state
	Conditions []*apimachinery_pkg_apis_meta_v1.Condition `json:"conditions,omitempty"`
}