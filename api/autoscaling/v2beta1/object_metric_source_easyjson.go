// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2beta1

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
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

func easyjsonE86a02c4DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(in *jlexer.Lexer, out *ObjectMetricSource) {
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
		case "averageValue":
			if in.IsNull() {
				in.Skip()
				out.AverageValue = nil
			} else {
				if out.AverageValue == nil {
					out.AverageValue = new(resource.Quantity)
				}
				*out.AverageValue = resource.Quantity(in.String())
			}
		case "metricName":
			if in.IsNull() {
				in.Skip()
				out.MetricName = nil
			} else {
				if out.MetricName == nil {
					out.MetricName = new(string)
				}
				*out.MetricName = string(in.String())
			}
		case "selector":
			if in.IsNull() {
				in.Skip()
				out.Selector = nil
			} else {
				if out.Selector == nil {
					out.Selector = new(_v1.LabelSelector)
				}
				(*out.Selector).UnmarshalEasyJSON(in)
			}
		case "target":
			if in.IsNull() {
				in.Skip()
				out.Target = nil
			} else {
				if out.Target == nil {
					out.Target = new(CrossVersionObjectReference)
				}
				(*out.Target).UnmarshalEasyJSON(in)
			}
		case "targetValue":
			if in.IsNull() {
				in.Skip()
				out.TargetValue = nil
			} else {
				if out.TargetValue == nil {
					out.TargetValue = new(resource.Quantity)
				}
				*out.TargetValue = resource.Quantity(in.String())
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
func easyjsonE86a02c4EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(out *jwriter.Writer, in ObjectMetricSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AverageValue != nil {
		const prefix string = ",\"averageValue\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.AverageValue))
	}
	{
		const prefix string = ",\"metricName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MetricName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.MetricName))
		}
	}
	if in.Selector != nil {
		const prefix string = ",\"selector\":"
		out.RawString(prefix)
		(*in.Selector).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"target\":"
		out.RawString(prefix)
		if in.Target == nil {
			out.RawString("null")
		} else {
			(*in.Target).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"targetValue\":"
		out.RawString(prefix)
		if in.TargetValue == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TargetValue))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ObjectMetricSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE86a02c4EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ObjectMetricSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE86a02c4EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ObjectMetricSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE86a02c4DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ObjectMetricSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE86a02c4DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(l, v)
}