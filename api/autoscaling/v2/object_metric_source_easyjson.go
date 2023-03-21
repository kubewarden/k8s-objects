// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2

import (
	json "encoding/json"
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

func easyjsonE86a02c4DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(in *jlexer.Lexer, out *ObjectMetricSource) {
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
		case "describedObject":
			if in.IsNull() {
				in.Skip()
				out.DescribedObject = nil
			} else {
				if out.DescribedObject == nil {
					out.DescribedObject = new(CrossVersionObjectReference)
				}
				(*out.DescribedObject).UnmarshalEasyJSON(in)
			}
		case "metric":
			if in.IsNull() {
				in.Skip()
				out.Metric = nil
			} else {
				if out.Metric == nil {
					out.Metric = new(MetricIdentifier)
				}
				(*out.Metric).UnmarshalEasyJSON(in)
			}
		case "target":
			if in.IsNull() {
				in.Skip()
				out.Target = nil
			} else {
				if out.Target == nil {
					out.Target = new(MetricTarget)
				}
				(*out.Target).UnmarshalEasyJSON(in)
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
func easyjsonE86a02c4EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(out *jwriter.Writer, in ObjectMetricSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"describedObject\":"
		out.RawString(prefix[1:])
		if in.DescribedObject == nil {
			out.RawString("null")
		} else {
			(*in.DescribedObject).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix)
		if in.Metric == nil {
			out.RawString("null")
		} else {
			(*in.Metric).MarshalEasyJSON(out)
		}
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ObjectMetricSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE86a02c4EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ObjectMetricSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE86a02c4EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ObjectMetricSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE86a02c4DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ObjectMetricSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE86a02c4DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(l, v)
}
