// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// APIServiceStatus APIServiceStatus contains derived information about an API server
//
// swagger:model APIServiceStatus
type APIServiceStatus struct {

	// Current service state of apiService.
	Conditions []*APIServiceCondition `json:"conditions"`
}
