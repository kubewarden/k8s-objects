// Code generated by go-swagger; DO NOT EDIT.

package v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// ExternalMetricSource ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). Exactly one "target" type should be set.
//
// swagger:model ExternalMetricSource
type ExternalMetricSource struct {

	// metricName is the name of the metric in question.
	// Required: true
	MetricName *string `json:"metricName"`

	// metricSelector is used to identify a specific time series within a given metric.
	MetricSelector apimachinery_pkg_apis_meta_v1.LabelSelector `json:"metricSelector,omitempty"`

	// targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.
	TargetAverageValue apimachinery_pkg_api_resource.Quantity `json:"targetAverageValue,omitempty"`

	// targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.
	TargetValue apimachinery_pkg_api_resource.Quantity `json:"targetValue,omitempty"`
}
