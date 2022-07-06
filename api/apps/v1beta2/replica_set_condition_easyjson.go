// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta2

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

func easyjson270ebef5DecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(in *jlexer.Lexer, out *ReplicaSetCondition) {
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
		case "lastTransitionTime":
			if in.IsNull() {
				in.Skip()
				out.LastTransitionTime = nil
			} else {
				if out.LastTransitionTime == nil {
					out.LastTransitionTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastTransitionTime).UnmarshalJSON(data))
				}
			}
		case "message":
			out.Message = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(string)
				}
				*out.Status = string(in.String())
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
func easyjson270ebef5EncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(out *jwriter.Writer, in ReplicaSetCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastTransitionTime != nil {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.LastTransitionTime).MarshalJSON())
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Reason))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Status == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Status))
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReplicaSetCondition) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson270ebef5EncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReplicaSetCondition) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson270ebef5EncodeGithubComKubewardenK8sObjectsApiAppsV1beta2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReplicaSetCondition) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson270ebef5DecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReplicaSetCondition) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson270ebef5DecodeGithubComKubewardenK8sObjectsApiAppsV1beta2(l, v)
}
