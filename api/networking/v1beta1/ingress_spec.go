// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// IngressSpec IngressSpec describes the Ingress the user wishes to exist.
//
// swagger:model IngressSpec
type IngressSpec struct {

	// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
	Backend *IngressBackend `json:"backend,omitempty"`

	// IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.
	IngressClassName string `json:"ingressClassName,omitempty"`

	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
	Rules []*IngressRule `json:"rules"`

	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
	TLS []*IngressTLS `json:"tls"`
}
