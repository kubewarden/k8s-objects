// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// CertificateSigningRequestCondition CertificateSigningRequestCondition describes a condition of a CertificateSigningRequest object
//
// swagger:model CertificateSigningRequestCondition
type CertificateSigningRequestCondition struct {

	// lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
	LastTransitionTime apimachinery_pkg_apis_meta_v1.Time `json:"lastTransitionTime,omitempty"`

	// lastUpdateTime is the time of the last update to this condition
	LastUpdateTime apimachinery_pkg_apis_meta_v1.Time `json:"lastUpdateTime,omitempty"`

	// message contains a human readable message with details about the request state
	Message string `json:"message,omitempty"`

	// reason indicates a brief reason for the request state
	Reason string `json:"reason,omitempty"`

	// status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown".
	// Required: true
	Status *string `json:"status"`

	// type of the condition. Known conditions are "Approved", "Denied", and "Failed".
	//
	// An "Approved" condition is added via the /approval subresource, indicating the request was approved and should be issued by the signer.
	//
	// A "Denied" condition is added via the /approval subresource, indicating the request was denied and should not be issued by the signer.
	//
	// A "Failed" condition is added via the /status subresource, indicating the signer failed to issue the certificate.
	//
	// Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once added.
	//
	// Only one condition of a given type is allowed.
	//
	// Possible enum values:
	//  - `"Approved"` Approved indicates the request was approved and should be issued by the signer.
	//  - `"Denied"` Denied indicates the request was denied and should not be issued by the signer.
	//  - `"Failed"` Failed indicates the signer failed to issue the certificate.
	// Required: true
	// Enum: [Approved Denied Failed]
	Type *string `json:"type"`
}
