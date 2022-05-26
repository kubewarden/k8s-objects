// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ServiceReference ServiceReference holds a reference to Service.legacy.k8s.io
//
// swagger:model ServiceReference
type ServiceReference struct {

	// Name is the name of the service
	Name string `json:"name,omitempty"`

	// Namespace is the namespace of the service
	Namespace string `json:"namespace,omitempty"`

	// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	Port int32 `json:"port,omitempty"`
}
