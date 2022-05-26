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

func easyjson15425908DecodeGithubComKubewardenK8sObjectsApiEventsV1beta1(in *jlexer.Lexer, out *EventSeries) {
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
		case "count":
			if in.IsNull() {
				in.Skip()
				out.Count = nil
			} else {
				if out.Count == nil {
					out.Count = new(int32)
				}
				*out.Count = int32(in.Int32())
			}
		case "lastObservedTime":
			if in.IsNull() {
				in.Skip()
				out.LastObservedTime = nil
			} else {
				if out.LastObservedTime == nil {
					out.LastObservedTime = new(_v1.MicroTime)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastObservedTime).UnmarshalJSON(data))
				}
			}
		case "state":
			if in.IsNull() {
				in.Skip()
				out.State = nil
			} else {
				if out.State == nil {
					out.State = new(string)
				}
				*out.State = string(in.String())
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
func easyjson15425908EncodeGithubComKubewardenK8sObjectsApiEventsV1beta1(out *jwriter.Writer, in EventSeries) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix[1:])
		if in.Count == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Count))
		}
	}
	{
		const prefix string = ",\"lastObservedTime\":"
		out.RawString(prefix)
		if in.LastObservedTime == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.LastObservedTime).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix)
		if in.State == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.State))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventSeries) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15425908EncodeGithubComKubewardenK8sObjectsApiEventsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventSeries) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15425908EncodeGithubComKubewardenK8sObjectsApiEventsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventSeries) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15425908DecodeGithubComKubewardenK8sObjectsApiEventsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventSeries) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15425908DecodeGithubComKubewardenK8sObjectsApiEventsV1beta1(l, v)
}
