// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Subject Subject matches the originator of a request, as identified by the request authentication system. There are three ways of matching an originator; by user, group, or service account.
//
// swagger:model Subject
type Subject struct {

	// group
	Group *GroupSubject `json:"group,omitempty"`

	// Required
	// Required: true
	Kind *string `json:"kind"`

	// service account
	ServiceAccount *ServiceAccountSubject `json:"serviceAccount,omitempty"`

	// user
	User *UserSubject `json:"user,omitempty"`
}
