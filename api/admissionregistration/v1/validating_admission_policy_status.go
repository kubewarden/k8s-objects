// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// ValidatingAdmissionPolicyStatus ValidatingAdmissionPolicyStatus represents the status of an admission validation policy.
//
// swagger:model ValidatingAdmissionPolicyStatus
type ValidatingAdmissionPolicyStatus struct {

	// The conditions represent the latest available observations of a policy's current state.
	Conditions []*apimachinery_pkg_apis_meta_v1.Condition `json:"conditions,omitempty"`

	// The generation observed by the controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// The results of type checking for each expression. Presence of this field indicates the completion of the type checking.
	TypeChecking *TypeChecking `json:"typeChecking,omitempty"`
}