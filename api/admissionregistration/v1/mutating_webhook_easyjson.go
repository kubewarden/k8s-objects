// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(in *jlexer.Lexer, out *MutatingWebhook) {
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
				easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV11(in, out.ClientConfig)
			}
		case "failurePolicy":
			out.FailurePolicy = string(in.String())
		case "matchConditions":
			if in.IsNull() {
				in.Skip()
				out.MatchConditions = nil
			} else {
				in.Delim('[')
				if out.MatchConditions == nil {
					if !in.IsDelim(']') {
						out.MatchConditions = make([]*MatchCondition, 0, 8)
					} else {
						out.MatchConditions = []*MatchCondition{}
					}
				} else {
					out.MatchConditions = (out.MatchConditions)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *MatchCondition
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(MatchCondition)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.MatchConditions = append(out.MatchConditions, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
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
		case "reinvocationPolicy":
			out.ReinvocationPolicy = string(in.String())
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
					var v3 *RuleWithOperations
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(RuleWithOperations)
						}
						easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV12(in, v3)
					}
					out.Rules = append(out.Rules, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sideEffects":
			if in.IsNull() {
				in.Skip()
				out.SideEffects = nil
			} else {
				if out.SideEffects == nil {
					out.SideEffects = new(string)
				}
				*out.SideEffects = string(in.String())
			}
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
func easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(out *jwriter.Writer, in MutatingWebhook) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"admissionReviewVersions\":"
		out.RawString(prefix[1:])
		if in.AdmissionReviewVersions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.AdmissionReviewVersions {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"clientConfig\":"
		out.RawString(prefix)
		if in.ClientConfig == nil {
			out.RawString("null")
		} else {
			easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV11(out, *in.ClientConfig)
		}
	}
	if in.FailurePolicy != "" {
		const prefix string = ",\"failurePolicy\":"
		out.RawString(prefix)
		out.String(string(in.FailurePolicy))
	}
	if len(in.MatchConditions) != 0 {
		const prefix string = ",\"matchConditions\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v6, v7 := range in.MatchConditions {
				if v6 > 0 {
					out.RawByte(',')
				}
				if v7 == nil {
					out.RawString("null")
				} else {
					(*v7).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
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
	if in.ReinvocationPolicy != "" {
		const prefix string = ",\"reinvocationPolicy\":"
		out.RawString(prefix)
		out.String(string(in.ReinvocationPolicy))
	}
	if len(in.Rules) != 0 {
		const prefix string = ",\"rules\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v8, v9 := range in.Rules {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV12(out, *v9)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"sideEffects\":"
		out.RawString(prefix)
		if in.SideEffects == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.SideEffects))
		}
	}
	if in.TimeoutSeconds != 0 {
		const prefix string = ",\"timeoutSeconds\":"
		out.RawString(prefix)
		out.Int32(int32(in.TimeoutSeconds))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MutatingWebhook) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MutatingWebhook) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MutatingWebhook) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MutatingWebhook) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV1(l, v)
}
func easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV12(in *jlexer.Lexer, out *RuleWithOperations) {
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
		case "apiGroups":
			if in.IsNull() {
				in.Skip()
				out.APIGroups = nil
			} else {
				in.Delim('[')
				if out.APIGroups == nil {
					if !in.IsDelim(']') {
						out.APIGroups = make([]string, 0, 4)
					} else {
						out.APIGroups = []string{}
					}
				} else {
					out.APIGroups = (out.APIGroups)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.APIGroups = append(out.APIGroups, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "apiVersions":
			if in.IsNull() {
				in.Skip()
				out.APIVersions = nil
			} else {
				in.Delim('[')
				if out.APIVersions == nil {
					if !in.IsDelim(']') {
						out.APIVersions = make([]string, 0, 4)
					} else {
						out.APIVersions = []string{}
					}
				} else {
					out.APIVersions = (out.APIVersions)[:0]
				}
				for !in.IsDelim(']') {
					var v11 string
					v11 = string(in.String())
					out.APIVersions = append(out.APIVersions, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "operations":
			if in.IsNull() {
				in.Skip()
				out.Operations = nil
			} else {
				in.Delim('[')
				if out.Operations == nil {
					if !in.IsDelim(']') {
						out.Operations = make([]string, 0, 4)
					} else {
						out.Operations = []string{}
					}
				} else {
					out.Operations = (out.Operations)[:0]
				}
				for !in.IsDelim(']') {
					var v12 string
					v12 = string(in.String())
					out.Operations = append(out.Operations, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "resources":
			if in.IsNull() {
				in.Skip()
				out.Resources = nil
			} else {
				in.Delim('[')
				if out.Resources == nil {
					if !in.IsDelim(']') {
						out.Resources = make([]string, 0, 4)
					} else {
						out.Resources = []string{}
					}
				} else {
					out.Resources = (out.Resources)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Resources = append(out.Resources, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "scope":
			out.Scope = string(in.String())
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
func easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV12(out *jwriter.Writer, in RuleWithOperations) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.APIGroups) != 0 {
		const prefix string = ",\"apiGroups\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v14, v15 := range in.APIGroups {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	if len(in.APIVersions) != 0 {
		const prefix string = ",\"apiVersions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.APIVersions {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	if len(in.Operations) != 0 {
		const prefix string = ",\"operations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.Operations {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.String(string(v19))
			}
			out.RawByte(']')
		}
	}
	if len(in.Resources) != 0 {
		const prefix string = ",\"resources\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.Resources {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	if in.Scope != "" {
		const prefix string = ",\"scope\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Scope))
	}
	out.RawByte('}')
}
func easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV11(in *jlexer.Lexer, out *WebhookClientConfig) {
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
				easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV13(in, out.Service)
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
func easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV11(out *jwriter.Writer, in WebhookClientConfig) {
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
		easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV13(out, *in.Service)
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
func easyjson28eb4f85DecodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV13(in *jlexer.Lexer, out *ServiceReference) {
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
		case "namespace":
			if in.IsNull() {
				in.Skip()
				out.Namespace = nil
			} else {
				if out.Namespace == nil {
					out.Namespace = new(string)
				}
				*out.Namespace = string(in.String())
			}
		case "path":
			out.Path = string(in.String())
		case "port":
			out.Port = int32(in.Int32())
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
func easyjson28eb4f85EncodeGithubComKubewardenK8sObjectsApiAdmissionregistrationV13(out *jwriter.Writer, in ServiceReference) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		if in.Namespace == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Namespace))
		}
	}
	if in.Path != "" {
		const prefix string = ",\"path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	if in.Port != 0 {
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.Int32(int32(in.Port))
	}
	out.RawByte('}')
}
