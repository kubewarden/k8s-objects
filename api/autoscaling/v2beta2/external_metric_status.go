// Code generated by go-swagger; DO NOT EDIT.

package v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ExternalMetricStatus ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
//
// swagger:model ExternalMetricStatus
type ExternalMetricStatus struct {

	// current contains the current value for the given metric
	// Required: true
	Current *MetricValueStatus `json:"current"`

	// metric identifies the target metric by name and selector
	// Required: true
	Metric *MetricIdentifier `json:"metric"`
}
