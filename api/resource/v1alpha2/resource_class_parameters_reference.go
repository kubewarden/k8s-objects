// Code generated by go-swagger; DO NOT EDIT.

package v1alpha2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ResourceClassParametersReference ResourceClassParametersReference contains enough information to let you locate the parameters for a ResourceClass.
//
// swagger:model ResourceClassParametersReference
type ResourceClassParametersReference struct {

	// APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
	APIGroup string `json:"apiGroup,omitempty"`

	// Kind is the type of resource being referenced. This is the same value as in the parameter object's metadata.
	// Required: true
	Kind *string `json:"kind"`

	// Name is the name of resource being referenced.
	// Required: true
	Name *string `json:"name"`

	// Namespace that contains the referenced resource. Must be empty for cluster-scoped resources and non-empty for namespaced resources.
	Namespace string `json:"namespace,omitempty"`
}