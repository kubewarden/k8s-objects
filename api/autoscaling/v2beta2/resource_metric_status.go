// Code generated by go-swagger; DO NOT EDIT.

package v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ResourceMetricStatus ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
//
// swagger:model ResourceMetricStatus
type ResourceMetricStatus struct {

	// current contains the current value for the given metric
	// Required: true
	Current *MetricValueStatus `json:"current"`

	// Name is the name of the resource in question.
	// Required: true
	Name *string `json:"name"`
}
