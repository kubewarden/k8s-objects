// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CustomResourceDefinitionSpec CustomResourceDefinitionSpec describes how a user wants their resource to appear
//
// swagger:model CustomResourceDefinitionSpec
type CustomResourceDefinitionSpec struct {

	// AdditionalPrinterColumns are additional columns shown e.g. in kubectl next to the name. Defaults to a created-at column. Optional, the global columns for all versions. Top-level and per-version columns are mutually exclusive.
	AdditionalPrinterColumns []*CustomResourceColumnDefinition `json:"additionalPrinterColumns,omitempty"`

	// `conversion` defines conversion settings for the CRD.
	Conversion *CustomResourceConversion `json:"conversion,omitempty"`

	// Group is the group this resource belongs in
	// Required: true
	Group *string `json:"group"`

	// Names are the names used to describe this custom resource
	// Required: true
	Names *CustomResourceDefinitionNames `json:"names"`

	// preserveUnknownFields disables pruning of object fields which are not specified in the OpenAPI schema. apiVersion, kind, metadata and known fields inside metadata are always preserved. Defaults to true in v1beta and will default to false in v1.
	PreserveUnknownFields bool `json:"preserveUnknownFields,omitempty"`

	// Scope indicates whether this resource is cluster or namespace scoped.  Default is namespaced
	// Required: true
	Scope *string `json:"scope"`

	// Subresources describes the subresources for CustomResource Optional, the global subresources for all versions. Top-level and per-version subresources are mutually exclusive.
	Subresources *CustomResourceSubresources `json:"subresources,omitempty"`

	// Validation describes the validation methods for CustomResources Optional, the global validation schema for all versions. Top-level and per-version schemas are mutually exclusive.
	Validation *CustomResourceValidation `json:"validation,omitempty"`

	// Version is the version this resource belongs in Should be always first item in Versions field if provided. Optional, but at least one of Version or Versions must be set. Deprecated: Please use `Versions`.
	Version string `json:"version,omitempty"`

	// Versions is the list of all supported versions for this resource. If Version field is provided, this field is optional. Validation: All versions must use the same validation schema for now. i.e., top level Validation field is applied to all of these versions. Order: The version name will be used to compute the order. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
	Versions []*CustomResourceDefinitionVersion `json:"versions,omitempty"`
}
