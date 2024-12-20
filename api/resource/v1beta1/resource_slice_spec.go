// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
)

// ResourceSliceSpec ResourceSliceSpec contains the information published by the driver in one ResourceSlice.
//
// swagger:model ResourceSliceSpec
type ResourceSliceSpec struct {

	// AllNodes indicates that all nodes have access to the resources in the pool.
	//
	// Exactly one of NodeName, NodeSelector and AllNodes must be set.
	AllNodes bool `json:"allNodes,omitempty"`

	// Devices lists some or all of the devices in this pool.
	//
	// Must not have more than 128 entries.
	Devices []*Device `json:"devices,omitempty"`

	// Driver identifies the DRA driver providing the capacity information. A field selector can be used to list only ResourceSlice objects with a certain driver name.
	//
	// Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver. This field is immutable.
	// Required: true
	Driver *string `json:"driver"`

	// NodeName identifies the node which provides the resources in this pool. A field selector can be used to list only ResourceSlice objects belonging to a certain node.
	//
	// This field can be used to limit access from nodes to ResourceSlices with the same node name. It also indicates to autoscalers that adding new nodes of the same type as some old node might also make new resources available.
	//
	// Exactly one of NodeName, NodeSelector and AllNodes must be set. This field is immutable.
	NodeName string `json:"nodeName,omitempty"`

	// NodeSelector defines which nodes have access to the resources in the pool, when that pool is not limited to a single node.
	//
	// Must use exactly one term.
	//
	// Exactly one of NodeName, NodeSelector and AllNodes must be set.
	NodeSelector *api_core_v1.NodeSelector `json:"nodeSelector,omitempty"`

	// Pool describes the pool that this ResourceSlice belongs to.
	// Required: true
	Pool *ResourcePool `json:"pool"`
}