// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// Overhead Overhead structure represents the resource overhead associated with running a pod.
//
// swagger:model Overhead
type Overhead struct {

	// PodFixed represents the fixed resource overhead associated with running a pod.
	PodFixed map[string]*apimachinery_pkg_api_resource.Quantity `json:"podFixed,omitempty"`
}
