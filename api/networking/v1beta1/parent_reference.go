// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ParentReference ParentReference describes a reference to a parent object.
//
// swagger:model ParentReference
type ParentReference struct {

	// Group is the group of the object being referenced.
	Group string `json:"group,omitempty"`

	// Name is the name of the object being referenced.
	// Required: true
	Name *string `json:"name"`

	// Namespace is the namespace of the object being referenced.
	Namespace string `json:"namespace,omitempty"`

	// Resource is the resource of the object being referenced.
	// Required: true
	Resource *string `json:"resource"`
}