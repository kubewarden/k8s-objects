// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// TokenRequest TokenRequest contains parameters of a service account token.
//
// swagger:model TokenRequest
type TokenRequest struct {

	// audience is the intended audience of the token in "TokenRequestSpec". It will default to the audiences of kube apiserver.
	// Required: true
	Audience *string `json:"audience"`

	// expirationSeconds is the duration of validity of the token in "TokenRequestSpec". It has the same default value of "ExpirationSeconds" in "TokenRequestSpec".
	ExpirationSeconds int64 `json:"expirationSeconds,omitempty"`
}
