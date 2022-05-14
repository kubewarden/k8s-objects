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

func easyjson37ddbce8DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(in *jlexer.Lexer, out *HPAScalingPolicy) {
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
		case "periodSeconds":
			if in.IsNull() {
				in.Skip()
				out.PeriodSeconds = nil
			} else {
				if out.PeriodSeconds == nil {
					out.PeriodSeconds = new(int32)
				}
				*out.PeriodSeconds = int32(in.Int32())
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
			}
		case "value":
			if in.IsNull() {
				in.Skip()
				out.Value = nil
			} else {
				if out.Value == nil {
					out.Value = new(int32)
				}
				*out.Value = int32(in.Int32())
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
func easyjson37ddbce8EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(out *jwriter.Writer, in HPAScalingPolicy) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"periodSeconds\":"
		out.RawString(prefix[1:])
		if in.PeriodSeconds == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.PeriodSeconds))
		}
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		if in.Value == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Value))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HPAScalingPolicy) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson37ddbce8EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HPAScalingPolicy) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson37ddbce8EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HPAScalingPolicy) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson37ddbce8DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HPAScalingPolicy) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson37ddbce8DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(l, v)
}
