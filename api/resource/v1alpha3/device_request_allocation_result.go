// Code generated by go-swagger; DO NOT EDIT.

package v1alpha3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// DeviceRequestAllocationResult DeviceRequestAllocationResult contains the allocation result for one request.
//
// swagger:model DeviceRequestAllocationResult
type DeviceRequestAllocationResult struct {

	// AdminAccess indicates that this device was allocated for administrative access. See the corresponding request field for a definition of mode.
	//
	// This is an alpha field and requires enabling the DRAAdminAccess feature gate. Admin access is disabled if this field is unset or set to false, otherwise it is enabled.
	AdminAccess bool `json:"adminAccess,omitempty"`

	// Device references one device instance via its name in the driver's resource pool. It must be a DNS label.
	// Required: true
	Device *string `json:"device"`

	// Driver specifies the name of the DRA driver whose kubelet plugin should be invoked to process the allocation once the claim is needed on a node.
	//
	// Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	// Required: true
	Driver *string `json:"driver"`

	// This name together with the driver name and the device name field identify which device was allocated (`<driver name>/<pool name>/<device name>`).
	//
	// Must not be longer than 253 characters and may contain one or more DNS sub-domains separated by slashes.
	// Required: true
	Pool *string `json:"pool"`

	// Request is the name of the request in the claim which caused this device to be allocated. Multiple devices may have been allocated per request.
	// Required: true
	Request *string `json:"request"`
}