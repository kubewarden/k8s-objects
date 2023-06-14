// Code generated by go-swagger; DO NOT EDIT.

package v1beta3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Subject Subject matches the originator of a request, as identified by the request authentication system. There are three ways of matching an originator; by user, group, or service account.
//
// swagger:model Subject
type Subject struct {

	// `group` matches based on user group name.
	Group *GroupSubject `json:"group,omitempty"`

	// `kind` indicates which one of the other fields is non-empty. Required
	// Required: true
	Kind *string `json:"kind"`

	// `serviceAccount` matches ServiceAccounts.
	ServiceAccount *ServiceAccountSubject `json:"serviceAccount,omitempty"`

	// `user` matches based on username.
	User *UserSubject `json:"user,omitempty"`
}