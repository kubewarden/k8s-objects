// Code generated by go-swagger; DO NOT EDIT.

package v1alpha3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// DeviceClassConfiguration DeviceClassConfiguration is used in DeviceClass.
//
// swagger:model DeviceClassConfiguration
type DeviceClassConfiguration struct {

	// Opaque provides driver-specific configuration parameters.
	Opaque *OpaqueDeviceConfiguration `json:"opaque,omitempty"`
}