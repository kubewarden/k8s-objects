// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ContainerStatus ContainerStatus contains details for the current status of this container.
//
// swagger:model ContainerStatus
type ContainerStatus struct {

	// Container's ID in the format '<type>://<container_id>'.
	ContainerID string `json:"containerID,omitempty"`

	// The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images.
	// Required: true
	Image *string `json:"image"`

	// ImageID of the container's image.
	// Required: true
	ImageID *string `json:"imageID"`

	// Details about the container's last termination condition.
	LastState *ContainerState `json:"lastState,omitempty"`

	// This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated.
	// Required: true
	Name *string `json:"name"`

	// Specifies whether the container has passed its readiness probe.
	// Required: true
	Ready *bool `json:"ready"`

	// The number of times the container has been restarted.
	// Required: true
	RestartCount *int32 `json:"restartCount"`

	// Specifies whether the container has passed its startup probe. Initialized as false, becomes true after startupProbe is considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. Is always true when no startupProbe is defined.
	Started bool `json:"started,omitempty"`

	// Details about the container's current condition.
	State *ContainerState `json:"state,omitempty"`
}
