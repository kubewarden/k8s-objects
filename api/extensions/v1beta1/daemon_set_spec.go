// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// DaemonSetSpec DaemonSetSpec is the specification of a daemon set.
//
// swagger:model DaemonSetSpec
type DaemonSetSpec struct {

	// The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`

	// The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	RevisionHistoryLimit int32 `json:"revisionHistoryLimit,omitempty"`

	// A label query over pods that are managed by the daemon set. Must match in order to be controlled. If empty, defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *apimachinery_pkg_apis_meta_v1.LabelSelector `json:"selector,omitempty"`

	// An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	// Required: true
	Template *api_core_v1.PodTemplateSpec `json:"template"`

	// DEPRECATED. A sequence number representing a specific generation of the template. Populated by the system. It can be set only during the creation.
	TemplateGeneration int64 `json:"templateGeneration,omitempty"`

	// An update strategy to replace existing DaemonSet pods with new pods.
	UpdateStrategy *DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"`
}
