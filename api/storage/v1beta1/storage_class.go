// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// StorageClass StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.
//
// StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
//
// swagger:model StorageClass
type StorageClass struct {

	// AllowVolumeExpansion shows whether the storage class allow volume expand
	AllowVolumeExpansion bool `json:"allowVolumeExpansion,omitempty"`

	// Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
	AllowedTopologies []api_core_v1.TopologySelectorTerm `json:"allowedTopologies"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	Metadata apimachinery_pkg_apis_meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
	MountOptions []string `json:"mountOptions"`

	// Parameters holds the parameters for the provisioner that should create volumes of this storage class.
	Parameters map[string]string `json:"parameters,omitempty"`

	// Provisioner indicates the type of the provisioner.
	// Required: true
	Provisioner *string `json:"provisioner"`

	// Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
	ReclaimPolicy string `json:"reclaimPolicy,omitempty"`

	// VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
	VolumeBindingMode string `json:"volumeBindingMode,omitempty"`
}
