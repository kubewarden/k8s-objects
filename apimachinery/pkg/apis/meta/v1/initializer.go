// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Initializer Initializer is information about an initializer that has not yet completed.
//
// swagger:model Initializer
type Initializer struct {

	// name of the process that is responsible for initializing this object.
	// Required: true
	Name *string `json:"name"`
}
