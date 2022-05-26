// Code generated by go-swagger; DO NOT EDIT.

package v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CrossVersionObjectReference CrossVersionObjectReference contains enough information to let you identify the referred resource.
//
// swagger:model CrossVersionObjectReference
type CrossVersionObjectReference struct {

	// API version of the referent
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
	// Required: true
	Kind *string `json:"kind"`

	// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// Required: true
	Name *string `json:"name"`
}
