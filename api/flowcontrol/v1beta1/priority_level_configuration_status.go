// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PriorityLevelConfigurationStatus PriorityLevelConfigurationStatus represents the current state of a "request-priority".
//
// swagger:model PriorityLevelConfigurationStatus
type PriorityLevelConfigurationStatus struct {

	// `conditions` is the current state of "request-priority".
	Conditions []*PriorityLevelConfigurationCondition `json:"conditions"`
}
