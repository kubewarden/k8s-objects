// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PersistentVolumeClaimVolumeSource PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).
//
// swagger:model PersistentVolumeClaimVolumeSource
type PersistentVolumeClaimVolumeSource struct {

	// ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// Required: true
	ClaimName *string `json:"claimName"`

	// Will force the ReadOnly setting in VolumeMounts. Default false.
	ReadOnly bool `json:"readOnly,omitempty"`
}
