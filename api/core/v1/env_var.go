// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// EnvVar EnvVar represents an environment variable present in a Container.
//
// swagger:model EnvVar
type EnvVar struct {

	// Name of the environment variable. Must be a C_IDENTIFIER.
	// Required: true
	Name *string `json:"name"`

	// Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	Value string `json:"value,omitempty"`

	// Source for the environment variable's value. Cannot be used if value is not empty.
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty"`
}
