// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ScaleIOPersistentVolumeSource ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume
//
// swagger:model ScaleIOPersistentVolumeSource
type ScaleIOPersistentVolumeSource struct {

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs"
	FSType string `json:"fsType,omitempty"`

	// The host address of the ScaleIO API Gateway.
	// Required: true
	Gateway *string `json:"gateway"`

	// The name of the ScaleIO Protection Domain for the configured storage.
	ProtectionDomain string `json:"protectionDomain,omitempty"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`

	// SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
	// Required: true
	SecretRef *SecretReference `json:"secretRef"`

	// Flag to enable/disable SSL communication with Gateway, default false
	SslEnabled bool `json:"sslEnabled,omitempty"`

	// Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
	StorageMode string `json:"storageMode,omitempty"`

	// The ScaleIO Storage Pool associated with the protection domain.
	StoragePool string `json:"storagePool,omitempty"`

	// The name of the storage system as configured in ScaleIO.
	// Required: true
	System *string `json:"system"`

	// The name of a volume already created in the ScaleIO system that is associated with this volume source.
	VolumeName string `json:"volumeName,omitempty"`
}
