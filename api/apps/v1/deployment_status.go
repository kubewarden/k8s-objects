// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// DeploymentStatus DeploymentStatus is the most recently observed status of the Deployment.
//
// swagger:model DeploymentStatus
type DeploymentStatus struct {

	// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`

	// Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	CollisionCount int32 `json:"collisionCount,omitempty"`

	// Represents the latest available observations of a deployment's current state.
	Conditions []*DeploymentCondition `json:"conditions"`

	// The generation observed by the deployment controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Total number of ready pods targeted by this deployment.
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Replicas int32 `json:"replicas,omitempty"`

	// Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty"`

	// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`
}
