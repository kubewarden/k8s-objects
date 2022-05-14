// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// NodeStatus NodeStatus is information about the current status of a node.
//
// swagger:model NodeStatus
type NodeStatus struct {

	// List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-replacement patch. See http://pr.k8s.io/79391 for an example.
	Addresses []*NodeAddress `json:"addresses"`

	// Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
	Allocatable map[string]apimachinery_pkg_api_resource.Quantity `json:"allocatable,omitempty"`

	// Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Capacity map[string]apimachinery_pkg_api_resource.Quantity `json:"capacity,omitempty"`

	// Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
	Conditions []*NodeCondition `json:"conditions"`

	// Status of the config assigned to the node via the dynamic Kubelet config feature.
	Config *NodeConfigStatus `json:"config,omitempty"`

	// Endpoints of daemons running on the Node.
	DaemonEndpoints *NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"`

	// List of container images on this node
	Images []*ContainerImage `json:"images"`

	// Set of ids/uuids to uniquely identify the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#info
	NodeInfo *NodeSystemInfo `json:"nodeInfo,omitempty"`

	// NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
	//
	// Possible enum values:
	//  - `"Pending"` means the node has been created/added by the system, but not configured.
	//  - `"Running"` means the node has been configured and has Kubernetes components running.
	//  - `"Terminated"` means the node has been removed from the cluster.
	// Enum: [Pending Running Terminated]
	Phase string `json:"phase,omitempty"`

	// List of volumes that are attached to the node.
	VolumesAttached []*AttachedVolume `json:"volumesAttached"`

	// List of attachable volumes in use (mounted) by the node.
	VolumesInUse []string `json:"volumesInUse"`
}
