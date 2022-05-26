// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// EventSeries EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in "k8s.io/client-go/tools/events/event_broadcaster.go" shows how this struct is updated on heartbeats and can guide customized reporter implementations.
//
// swagger:model EventSeries
type EventSeries struct {

	// count is the number of occurrences in this series up to the last heartbeat time.
	// Required: true
	Count *int32 `json:"count"`

	// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
	// Required: true
	LastObservedTime *apimachinery_pkg_apis_meta_v1.MicroTime `json:"lastObservedTime"`
}
