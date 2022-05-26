// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// RuntimeClassStrategyOptions RuntimeClassStrategyOptions define the strategy that will dictate the allowable RuntimeClasses for a pod.
//
// swagger:model RuntimeClassStrategyOptions
type RuntimeClassStrategyOptions struct {

	// allowedRuntimeClassNames is an allowlist of RuntimeClass names that may be specified on a pod. A value of "*" means that any RuntimeClass name is allowed, and must be the only item in the list. An empty list requires the RuntimeClassName field to be unset.
	// Required: true
	AllowedRuntimeClassNames []string `json:"allowedRuntimeClassNames"`

	// defaultRuntimeClassName is the default RuntimeClassName to set on the pod. The default MUST be allowed by the allowedRuntimeClassNames list. A value of nil does not mutate the Pod.
	DefaultRuntimeClassName string `json:"defaultRuntimeClassName,omitempty"`
}
