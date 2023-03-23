// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// VsphereVirtualDiskVolumeSource Represents a vSphere volume resource.
//
// swagger:model VsphereVirtualDiskVolumeSource
type VsphereVirtualDiskVolumeSource struct {

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FSType string `json:"fsType,omitempty"`

	// Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	StoragePolicyID string `json:"storagePolicyID,omitempty"`

	// Storage Policy Based Management (SPBM) profile name.
	StoragePolicyName string `json:"storagePolicyName,omitempty"`

	// Path that identifies vSphere volume vmdk
	// Required: true
	VolumePath *string `json:"volumePath"`
}
