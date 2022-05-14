// Code generated by go-swagger; DO NOT EDIT.

package v1beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// FlowSchemaSpec FlowSchemaSpec describes how the FlowSchema's specification looks like.
//
// swagger:model FlowSchemaSpec
type FlowSchemaSpec struct {

	// `distinguisherMethod` defines how to compute the flow distinguisher for requests that match this schema. `nil` specifies that the distinguisher is disabled and thus will always be the empty string.
	DistinguisherMethod *FlowDistinguisherMethod `json:"distinguisherMethod,omitempty"`

	// `matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.
	MatchingPrecedence int32 `json:"matchingPrecedence,omitempty"`

	// `priorityLevelConfiguration` should reference a PriorityLevelConfiguration in the cluster. If the reference cannot be resolved, the FlowSchema will be ignored and marked as invalid in its status. Required.
	// Required: true
	PriorityLevelConfiguration *PriorityLevelConfigurationReference `json:"priorityLevelConfiguration"`

	// `rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.
	Rules []*PolicyRulesWithSubjects `json:"rules"`
}
