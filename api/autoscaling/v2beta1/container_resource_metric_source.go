// Code generated by go-swagger; DO NOT EDIT.

package v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// ContainerResourceMetricSource ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
//
// swagger:model ContainerResourceMetricSource
type ContainerResourceMetricSource struct {

	// container is the name of the container in the pods of the scaling target
	// Required: true
	Container *string `json:"container"`

	// name is the name of the resource in question.
	// Required: true
	Name *string `json:"name"`

	// targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
	TargetAverageUtilization int32 `json:"targetAverageUtilization,omitempty"`

	// targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type.
	TargetAverageValue apimachinery_pkg_api_resource.Quantity `json:"targetAverageValue,omitempty"`
}
