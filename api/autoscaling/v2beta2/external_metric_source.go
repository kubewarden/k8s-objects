// Code generated by go-swagger; DO NOT EDIT.

package v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ExternalMetricSource ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
//
// swagger:model ExternalMetricSource
type ExternalMetricSource struct {

	// metric identifies the target metric by name and selector
	// Required: true
	Metric *MetricIdentifier `json:"metric"`

	// target specifies the target value for the given metric
	// Required: true
	Target *MetricTarget `json:"target"`
}