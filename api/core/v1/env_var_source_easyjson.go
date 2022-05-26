// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
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

func easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *EnvVarSource) {
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
		case "configMapKeyRef":
			if in.IsNull() {
				in.Skip()
				out.ConfigMapKeyRef = nil
			} else {
				if out.ConfigMapKeyRef == nil {
					out.ConfigMapKeyRef = new(ConfigMapKeySelector)
				}
				(*out.ConfigMapKeyRef).UnmarshalEasyJSON(in)
			}
		case "fieldRef":
			if in.IsNull() {
				in.Skip()
				out.FieldRef = nil
			} else {
				if out.FieldRef == nil {
					out.FieldRef = new(ObjectFieldSelector)
				}
				easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.FieldRef)
			}
		case "resourceFieldRef":
			if in.IsNull() {
				in.Skip()
				out.ResourceFieldRef = nil
			} else {
				if out.ResourceFieldRef == nil {
					out.ResourceFieldRef = new(ResourceFieldSelector)
				}
				easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.ResourceFieldRef)
			}
		case "secretKeyRef":
			if in.IsNull() {
				in.Skip()
				out.SecretKeyRef = nil
			} else {
				if out.SecretKeyRef == nil {
					out.SecretKeyRef = new(SecretKeySelector)
				}
				easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV13(in, out.SecretKeyRef)
			}
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
func easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in EnvVarSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConfigMapKeyRef != nil {
		const prefix string = ",\"configMapKeyRef\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ConfigMapKeyRef).MarshalEasyJSON(out)
	}
	if in.FieldRef != nil {
		const prefix string = ",\"fieldRef\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.FieldRef)
	}
	if in.ResourceFieldRef != nil {
		const prefix string = ",\"resourceFieldRef\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.ResourceFieldRef)
	}
	if in.SecretKeyRef != nil {
		const prefix string = ",\"secretKeyRef\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *in.SecretKeyRef)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnvVarSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnvVarSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnvVarSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnvVarSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV13(in *jlexer.Lexer, out *SecretKeySelector) {
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
		case "key":
			if in.IsNull() {
				in.Skip()
				out.Key = nil
			} else {
				if out.Key == nil {
					out.Key = new(string)
				}
				*out.Key = string(in.String())
			}
		case "name":
			out.Name = string(in.String())
		case "optional":
			out.Optional = bool(in.Bool())
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
func easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV13(out *jwriter.Writer, in SecretKeySelector) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix[1:])
		if in.Key == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Key))
		}
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.Optional {
		const prefix string = ",\"optional\":"
		out.RawString(prefix)
		out.Bool(bool(in.Optional))
	}
	out.RawByte('}')
}
func easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *ResourceFieldSelector) {
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
		case "containerName":
			out.ContainerName = string(in.String())
		case "divisor":
			out.Divisor = resource.Quantity(in.String())
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(string)
				}
				*out.Resource = string(in.String())
			}
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
func easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in ResourceFieldSelector) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ContainerName != "" {
		const prefix string = ",\"containerName\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ContainerName))
	}
	if in.Divisor != "" {
		const prefix string = ",\"divisor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Divisor))
	}
	{
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Resource == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Resource))
		}
	}
	out.RawByte('}')
}
func easyjson187983d5DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *ObjectFieldSelector) {
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
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "fieldPath":
			if in.IsNull() {
				in.Skip()
				out.FieldPath = nil
			} else {
				if out.FieldPath == nil {
					out.FieldPath = new(string)
				}
				*out.FieldPath = string(in.String())
			}
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
func easyjson187983d5EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in ObjectFieldSelector) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"fieldPath\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.FieldPath == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.FieldPath))
		}
	}
	out.RawByte('}')
}
