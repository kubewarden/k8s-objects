// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CustomResourceConversion CustomResourceConversion describes how to convert different versions of a CR.
//
// swagger:model CustomResourceConversion
type CustomResourceConversion struct {

	// ConversionReviewVersions is an ordered list of preferred `ConversionReview` versions the Webhook expects. API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, conversion will fail for this object. If a persisted Webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail. Default to `['v1beta1']`.
	ConversionReviewVersions []string `json:"conversionReviewVersions"`

	// `strategy` specifies the conversion strategy. Allowed values are: - `None`: The converter only change the apiVersion and would not touch any other field in the CR. - `Webhook`: API Server will call to an external webhook to do the conversion. Additional information
	//   is needed for this option. This requires spec.preserveUnknownFields to be false.
	// Required: true
	Strategy *string `json:"strategy"`

	// `webhookClientConfig` is the instructions for how to call the webhook if strategy is `Webhook`. This field is alpha-level and is only honored by servers that enable the CustomResourceWebhookConversion feature.
	WebhookClientConfig *WebhookClientConfig `json:"webhookClientConfig,omitempty"`
}
