// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2beta1

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

func easyjson8d0d2808DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(in *jlexer.Lexer, out *ContainerResourceMetricStatus) {
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
		case "container":
			if in.IsNull() {
				in.Skip()
				out.Container = nil
			} else {
				if out.Container == nil {
					out.Container = new(string)
				}
				*out.Container = string(in.String())
			}
		case "currentAverageUtilization":
			out.CurrentAverageUtilization = int32(in.Int32())
		case "currentAverageValue":
			if in.IsNull() {
				in.Skip()
				out.CurrentAverageValue = nil
			} else {
				if out.CurrentAverageValue == nil {
					out.CurrentAverageValue = new(resource.Quantity)
				}
				*out.CurrentAverageValue = resource.Quantity(in.String())
			}
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
func easyjson8d0d2808EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(out *jwriter.Writer, in ContainerResourceMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"container\":"
		out.RawString(prefix[1:])
		if in.Container == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Container))
		}
	}
	if in.CurrentAverageUtilization != 0 {
		const prefix string = ",\"currentAverageUtilization\":"
		out.RawString(prefix)
		out.Int32(int32(in.CurrentAverageUtilization))
	}
	{
		const prefix string = ",\"currentAverageValue\":"
		out.RawString(prefix)
		if in.CurrentAverageValue == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.CurrentAverageValue))
		}
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContainerResourceMetricStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8d0d2808EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContainerResourceMetricStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8d0d2808EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContainerResourceMetricStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8d0d2808DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContainerResourceMetricStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8d0d2808DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(l, v)
}
