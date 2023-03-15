// Code generated by go-swagger; DO NOT EDIT.

package v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// AuditSinkSpec AuditSinkSpec holds the spec for the audit sink
//
// swagger:model AuditSinkSpec
type AuditSinkSpec struct {

	// Policy defines the policy for selecting which events should be sent to the webhook required
	// Required: true
	Policy *Policy `json:"policy"`

	// Webhook to send events required
	// Required: true
	Webhook *Webhook `json:"webhook"`
}