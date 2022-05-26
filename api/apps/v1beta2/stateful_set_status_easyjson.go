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

func easyjson30c5d56cDecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(in *jlexer.Lexer, out *StatefulSetStatus) {
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
						out.Conditions = make([]*StatefulSetCondition, 0, 8)
					} else {
						out.Conditions = []*StatefulSetCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *StatefulSetCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(StatefulSetCondition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "currentReplicas":
			out.CurrentReplicas = int32(in.Int32())
		case "currentRevision":
			out.CurrentRevision = string(in.String())
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
		case "readyReplicas":
			out.ReadyReplicas = int32(in.Int32())
		case "replicas":
			if in.IsNull() {
				in.Skip()
				out.Replicas = nil
			} else {
				if out.Replicas == nil {
					out.Replicas = new(int32)
				}
				*out.Replicas = int32(in.Int32())
			}
		case "updateRevision":
			out.UpdateRevision = string(in.String())
		case "updatedReplicas":
			out.UpdatedReplicas = int32(in.Int32())
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
func easyjson30c5d56cEncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(out *jwriter.Writer, in StatefulSetStatus) {
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
	if in.CurrentReplicas != 0 {
		const prefix string = ",\"currentReplicas\":"
		out.RawString(prefix)
		out.Int32(int32(in.CurrentReplicas))
	}
	if in.CurrentRevision != "" {
		const prefix string = ",\"currentRevision\":"
		out.RawString(prefix)
		out.String(string(in.CurrentRevision))
	}
	if in.ObservedGeneration != 0 {
		const prefix string = ",\"observedGeneration\":"
		out.RawString(prefix)
		out.Int64(int64(in.ObservedGeneration))
	}
	if in.ReadyReplicas != 0 {
		const prefix string = ",\"readyReplicas\":"
		out.RawString(prefix)
		out.Int32(int32(in.ReadyReplicas))
	}
	{
		const prefix string = ",\"replicas\":"
		out.RawString(prefix)
		if in.Replicas == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Replicas))
		}
	}
	if in.UpdateRevision != "" {
		const prefix string = ",\"updateRevision\":"
		out.RawString(prefix)
		out.String(string(in.UpdateRevision))
	}
	if in.UpdatedReplicas != 0 {
		const prefix string = ",\"updatedReplicas\":"
		out.RawString(prefix)
		out.Int32(int32(in.UpdatedReplicas))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StatefulSetStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30c5d56cEncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StatefulSetStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30c5d56cEncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StatefulSetStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30c5d56cDecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StatefulSetStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30c5d56cDecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(l, v)
}