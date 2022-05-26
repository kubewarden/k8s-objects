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

func easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiEventsV1beta1(in *jlexer.Lexer, out *Event) {
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
		case "action":
			out.Action = string(in.String())
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "deprecatedCount":
			out.DeprecatedCount = int32(in.Int32())
		case "deprecatedFirstTimestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DeprecatedFirstTimestamp).UnmarshalJSON(data))
			}
		case "deprecatedLastTimestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DeprecatedLastTimestamp).UnmarshalJSON(data))
			}
		case "deprecatedSource":
			(out.DeprecatedSource).UnmarshalEasyJSON(in)
		case "eventTime":
			if in.IsNull() {
				in.Skip()
				out.EventTime = nil
			} else {
				if out.EventTime == nil {
					out.EventTime = new(_v1.MicroTime)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.EventTime).UnmarshalJSON(data))
				}
			}
		case "kind":
			out.Kind = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(_v1.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "note":
			out.Note = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "regarding":
			(out.Regarding).UnmarshalEasyJSON(in)
		case "related":
			(out.Related).UnmarshalEasyJSON(in)
		case "reportingController":
			out.ReportingController = string(in.String())
		case "reportingInstance":
			out.ReportingInstance = string(in.String())
		case "series":
			if in.IsNull() {
				in.Skip()
				out.Series = nil
			} else {
				if out.Series == nil {
					out.Series = new(EventSeries)
				}
				easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiEventsV1beta11(in, out.Series)
			}
		case "type":
			out.Type = string(in.String())
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
func easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiEventsV1beta1(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Action != "" {
		const prefix string = ",\"action\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Action))
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.APIVersion))
	}
	if in.DeprecatedCount != 0 {
		const prefix string = ",\"deprecatedCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.DeprecatedCount))
	}
	if true {
		const prefix string = ",\"deprecatedFirstTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.DeprecatedFirstTimestamp).MarshalJSON())
	}
	if true {
		const prefix string = ",\"deprecatedLastTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.DeprecatedLastTimestamp).MarshalJSON())
	}
	if true {
		const prefix string = ",\"deprecatedSource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.DeprecatedSource).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"eventTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.EventTime == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.EventTime).MarshalJSON())
		}
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil {
			out.RawString("null")
		} else {
			(*in.Metadata).MarshalEasyJSON(out)
		}
	}
	if in.Note != "" {
		const prefix string = ",\"note\":"
		out.RawString(prefix)
		out.String(string(in.Note))
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	if true {
		const prefix string = ",\"regarding\":"
		out.RawString(prefix)
		(in.Regarding).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"related\":"
		out.RawString(prefix)
		(in.Related).MarshalEasyJSON(out)
	}
	if in.ReportingController != "" {
		const prefix string = ",\"reportingController\":"
		out.RawString(prefix)
		out.String(string(in.ReportingController))
	}
	if in.ReportingInstance != "" {
		const prefix string = ",\"reportingInstance\":"
		out.RawString(prefix)
		out.String(string(in.ReportingInstance))
	}
	if in.Series != nil {
		const prefix string = ",\"series\":"
		out.RawString(prefix)
		easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiEventsV1beta11(out, *in.Series)
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiEventsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiEventsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiEventsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiEventsV1beta1(l, v)
}
func easyjsonF642ad3eDecodeGithubComKubewardenK8sObjectsApiEventsV1beta11(in *jlexer.Lexer, out *EventSeries) {
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
func easyjsonF642ad3eEncodeGithubComKubewardenK8sObjectsApiEventsV1beta11(out *jwriter.Writer, in EventSeries) {
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
	out.RawByte('}')
}
