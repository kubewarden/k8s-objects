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

func easyjson1d7c9f89DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(in *jlexer.Lexer, out *ContainerResourceMetricSource) {
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
		case "targetAverageUtilization":
			out.TargetAverageUtilization = int32(in.Int32())
		case "targetAverageValue":
			out.TargetAverageValue = resource.Quantity(in.String())
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
func easyjson1d7c9f89EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(out *jwriter.Writer, in ContainerResourceMetricSource) {
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
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	if in.TargetAverageUtilization != 0 {
		const prefix string = ",\"targetAverageUtilization\":"
		out.RawString(prefix)
		out.Int32(int32(in.TargetAverageUtilization))
	}
	if in.TargetAverageValue != "" {
		const prefix string = ",\"targetAverageValue\":"
		out.RawString(prefix)
		out.String(string(in.TargetAverageValue))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContainerResourceMetricSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1d7c9f89EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContainerResourceMetricSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1d7c9f89EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContainerResourceMetricSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1d7c9f89DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContainerResourceMetricSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1d7c9f89DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(l, v)
}
