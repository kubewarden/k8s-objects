// Code generated by go-swagger; DO NOT EDIT.

package v1beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
	apimachinery_pkg_runtime "github.com/kubewarden/k8s-objects/apimachinery/pkg/runtime"
)

// ControllerRevision DEPRECATED - This group version of ControllerRevision is deprecated by apps/v1/ControllerRevision. See the release notes for more information. ControllerRevision implements an immutable snapshot of state data. Clients are responsible for serializing and deserializing the objects that contain their internal state. Once a ControllerRevision has been successfully created, it can not be updated. The API Server will fail validation of all requests that attempt to mutate the Data field. ControllerRevisions may, however, be deleted. Note that, due to its use by both the DaemonSet and StatefulSet controllers for update and rollback, this object is beta. However, it may be subject to name and representation changes in future releases, and clients should not depend on its stability. It is primarily for internal use by controllers.
//
// swagger:model ControllerRevision
type ControllerRevision struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// Data is the serialized representation of the state.
	Data apimachinery_pkg_runtime.RawExtension `json:"data,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata apimachinery_pkg_apis_meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Revision indicates the revision of the state represented by Data.
	// Required: true
	Revision *int64 `json:"revision"`
}
