// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// LeaseSpec LeaseSpec is a specification of a Lease.
//
// swagger:model LeaseSpec
type LeaseSpec struct {

	// acquireTime is a time when the current lease was acquired.
	AcquireTime apimachinery_pkg_apis_meta_v1.MicroTime `json:"acquireTime,omitempty"`

	// holderIdentity contains the identity of the holder of a current lease.
	HolderIdentity string `json:"holderIdentity,omitempty"`

	// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
	LeaseDurationSeconds int32 `json:"leaseDurationSeconds,omitempty"`

	// leaseTransitions is the number of transitions of a lease between holders.
	LeaseTransitions int32 `json:"leaseTransitions,omitempty"`

	// renewTime is a time when the current holder of a lease has last updated the lease.
	RenewTime apimachinery_pkg_apis_meta_v1.MicroTime `json:"renewTime,omitempty"`
}
