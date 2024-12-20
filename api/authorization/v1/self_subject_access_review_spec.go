// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// SelfSubjectAccessReviewSpec SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
//
// swagger:model SelfSubjectAccessReviewSpec
type SelfSubjectAccessReviewSpec struct {

	// NonResourceAttributes describes information for a non-resource access request
	NonResourceAttributes *NonResourceAttributes `json:"nonResourceAttributes,omitempty"`

	// ResourceAuthorizationAttributes describes information for a resource access request
	ResourceAttributes *ResourceAttributes `json:"resourceAttributes,omitempty"`
}