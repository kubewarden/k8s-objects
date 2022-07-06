// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// IPBlock DEPRECATED 1.9 - This group version of IPBlock is deprecated by networking/v1/IPBlock. IPBlock describes a particular CIDR (Ex. "192.168.1.1/24") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
//
// swagger:model IPBlock
type IPBlock struct {

	// CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24"
	// Required: true
	Cidr *string `json:"cidr"`

	// Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" Except values will be rejected if they are outside the CIDR range
	Except []string `json:"except,omitempty"`
}
