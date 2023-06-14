// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_authentication_v1 "github.com/kubewarden/k8s-objects/api/authentication/v1"
)

// SelfSubjectReviewStatus SelfSubjectReviewStatus is filled by the kube-apiserver and sent back to a user.
//
// swagger:model SelfSubjectReviewStatus
type SelfSubjectReviewStatus struct {

	// User attributes of the user making this request.
	UserInfo *api_authentication_v1.UserInfo `json:"userInfo,omitempty"`
}
