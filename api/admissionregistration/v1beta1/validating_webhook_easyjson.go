// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson8684def1DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta1(in *jlexer.Lexer, out *ValidatingWebhook) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "admissionReviewVersions":
			if in.IsNull() {
				in.Skip()
				out.AdmissionReviewVersions = nil
			} else {
				in.Delim('[')
				if out.AdmissionReviewVersions == nil {
					if !in.IsDelim(']') {
						out.AdmissionReviewVersions = make([]string, 0, 4)
					} else {
						out.AdmissionReviewVersions = []string{}
					}
				} else {
					out.AdmissionReviewVersions = (out.AdmissionReviewVersions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.AdmissionReviewVersions = append(out.AdmissionReviewVersions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "clientConfig":
			if in.IsNull() {
				in.Skip()
				out.ClientConfig = nil
			} else {
				if out.ClientConfig == nil {
					out.ClientConfig = new(WebhookClientConfig)
				}
				easyjson8684def1DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta11(in, out.ClientConfig)
			}
		case "failurePolicy":
			out.FailurePolicy = string(in.String())
		case "matchPolicy":
			out.MatchPolicy = string(in.String())
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "namespaceSelector":
			if in.IsNull() {
				in.Skip()
				out.NamespaceSelector = nil
			} else {
				if out.NamespaceSelector == nil {
					out.NamespaceSelector = new(_v1.LabelSelector)
				}
				(*out.NamespaceSelector).UnmarshalEasyJSON(in)
			}
		case "objectSelector":
			if in.IsNull() {
				in.Skip()
				out.ObjectSelector = nil
			} else {
				if out.ObjectSelector == nil {
					out.ObjectSelector = new(_v1.LabelSelector)
				}
				(*out.ObjectSelector).UnmarshalEasyJSON(in)
			}
		case "rules":
			if in.IsNull() {
				in.Skip()
				out.Rules = nil
			} else {
				in.Delim('[')
				if out.Rules == nil {
					if !in.IsDelim(']') {
						out.Rules = make([]*RuleWithOperations, 0, 8)
					} else {
						out.Rules = []*RuleWithOperations{}
					}
				} else {
					out.Rules = (out.Rules)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *RuleWithOperations
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(RuleWithOperations)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.Rules = append(out.Rules, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sideEffects":
			out.SideEffects = string(in.String())
		case "timeoutSeconds":
			out.TimeoutSeconds = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8684def1EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta1(out *jwriter.Writer, in ValidatingWebhook) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.AdmissionReviewVersions) != 0 {
		const prefix string = ",\"admissionReviewVersions\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v3, v4 := range in.AdmissionReviewVersions {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"clientConfig\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.ClientConfig == nil {
			out.RawString("null")
		} else {
			easyjson8684def1EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta11(out, *in.ClientConfig)
		}
	}
	if in.FailurePolicy != "" {
		const prefix string = ",\"failurePolicy\":"
		out.RawString(prefix)
		out.String(string(in.FailurePolicy))
	}
	if in.MatchPolicy != "" {
		const prefix string = ",\"matchPolicy\":"
		out.RawString(prefix)
		out.String(string(in.MatchPolicy))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	if in.NamespaceSelector != nil {
		const prefix string = ",\"namespaceSelector\":"
		out.RawString(prefix)
		(*in.NamespaceSelector).MarshalEasyJSON(out)
	}
	if in.ObjectSelector != nil {
		const prefix string = ",\"objectSelector\":"
		out.RawString(prefix)
		(*in.ObjectSelector).MarshalEasyJSON(out)
	}
	if len(in.Rules) != 0 {
		const prefix string = ",\"rules\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Rules {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.SideEffects != "" {
		const prefix string = ",\"sideEffects\":"
		out.RawString(prefix)
		out.String(string(in.SideEffects))
	}
	if in.TimeoutSeconds != 0 {
		const prefix string = ",\"timeoutSeconds\":"
		out.RawString(prefix)
		out.Int32(int32(in.TimeoutSeconds))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ValidatingWebhook) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8684def1EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ValidatingWebhook) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8684def1EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ValidatingWebhook) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8684def1DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ValidatingWebhook) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8684def1DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta1(l, v)
}
func easyjson8684def1DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta11(in *jlexer.Lexer, out *WebhookClientConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "caBundle":
			if in.IsNull() {
				in.Skip()
				out.CaBundle = nil
			} else {
				out.CaBundle = in.Bytes()
			}
		case "service":
			if in.IsNull() {
				in.Skip()
				out.Service = nil
			} else {
				if out.Service == nil {
					out.Service = new(ServiceReference)
				}
				(*out.Service).UnmarshalEasyJSON(in)
			}
		case "url":
			out.URL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8684def1EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1beta11(out *jwriter.Writer, in WebhookClientConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.CaBundle) != 0 {
		const prefix string = ",\"caBundle\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(in.CaBundle)
	}
	if in.Service != nil {
		const prefix string = ",\"service\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Service).MarshalEasyJSON(out)
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	out.RawByte('}')
}