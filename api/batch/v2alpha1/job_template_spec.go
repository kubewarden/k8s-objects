// Code generated by go-swagger; DO NOT EDIT.

package v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_batch_v1 "github.com/kubewarden/k8s-objects/api/batch/v1"
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// JobTemplateSpec JobTemplateSpec describes the data a Job should have when created from a template
//
// swagger:model JobTemplateSpec
type JobTemplateSpec struct {

	// Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata apimachinery_pkg_apis_meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec api_batch_v1.JobSpec `json:"spec,omitempty"`
}