// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta2

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

func easyjson1477c5c0DecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(in *jlexer.Lexer, out *DaemonSetStatus) {
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
		case "collisionCount":
			out.CollisionCount = int32(in.Int32())
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*DaemonSetCondition, 0, 8)
					} else {
						out.Conditions = []*DaemonSetCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *DaemonSetCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DaemonSetCondition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "currentNumberScheduled":
			if in.IsNull() {
				in.Skip()
				out.CurrentNumberScheduled = nil
			} else {
				if out.CurrentNumberScheduled == nil {
					out.CurrentNumberScheduled = new(int32)
				}
				*out.CurrentNumberScheduled = int32(in.Int32())
			}
		case "desiredNumberScheduled":
			if in.IsNull() {
				in.Skip()
				out.DesiredNumberScheduled = nil
			} else {
				if out.DesiredNumberScheduled == nil {
					out.DesiredNumberScheduled = new(int32)
				}
				*out.DesiredNumberScheduled = int32(in.Int32())
			}
		case "numberAvailable":
			out.NumberAvailable = int32(in.Int32())
		case "numberMisscheduled":
			if in.IsNull() {
				in.Skip()
				out.NumberMisscheduled = nil
			} else {
				if out.NumberMisscheduled == nil {
					out.NumberMisscheduled = new(int32)
				}
				*out.NumberMisscheduled = int32(in.Int32())
			}
		case "numberReady":
			if in.IsNull() {
				in.Skip()
				out.NumberReady = nil
			} else {
				if out.NumberReady == nil {
					out.NumberReady = new(int32)
				}
				*out.NumberReady = int32(in.Int32())
			}
		case "numberUnavailable":
			out.NumberUnavailable = int32(in.Int32())
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
		case "updatedNumberScheduled":
			out.UpdatedNumberScheduled = int32(in.Int32())
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
func easyjson1477c5c0EncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(out *jwriter.Writer, in DaemonSetStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CollisionCount != 0 {
		const prefix string = ",\"collisionCount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CollisionCount))
	}
	{
		const prefix string = ",\"conditions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"currentNumberScheduled\":"
		out.RawString(prefix)
		if in.CurrentNumberScheduled == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.CurrentNumberScheduled))
		}
	}
	{
		const prefix string = ",\"desiredNumberScheduled\":"
		out.RawString(prefix)
		if in.DesiredNumberScheduled == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.DesiredNumberScheduled))
		}
	}
	if in.NumberAvailable != 0 {
		const prefix string = ",\"numberAvailable\":"
		out.RawString(prefix)
		out.Int32(int32(in.NumberAvailable))
	}
	{
		const prefix string = ",\"numberMisscheduled\":"
		out.RawString(prefix)
		if in.NumberMisscheduled == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.NumberMisscheduled))
		}
	}
	{
		const prefix string = ",\"numberReady\":"
		out.RawString(prefix)
		if in.NumberReady == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.NumberReady))
		}
	}
	if in.NumberUnavailable != 0 {
		const prefix string = ",\"numberUnavailable\":"
		out.RawString(prefix)
		out.Int32(int32(in.NumberUnavailable))
	}
	if in.ObservedGeneration != 0 {
		const prefix string = ",\"observedGeneration\":"
		out.RawString(prefix)
		out.Int64(int64(in.ObservedGeneration))
	}
	if in.UpdatedNumberScheduled != 0 {
		const prefix string = ",\"updatedNumberScheduled\":"
		out.RawString(prefix)
		out.Int32(int32(in.UpdatedNumberScheduled))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DaemonSetStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1477c5c0EncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DaemonSetStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1477c5c0EncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DaemonSetStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1477c5c0DecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DaemonSetStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1477c5c0DecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(l, v)
}
