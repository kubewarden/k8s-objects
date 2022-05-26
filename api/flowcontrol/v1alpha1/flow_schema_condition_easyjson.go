// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjson8bb5296eDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(in *jlexer.Lexer, out *FlowSchemaCondition) {
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
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastTransitionTime).UnmarshalJSON(data))
			}
		case "message":
			out.Message = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "status":
			out.Status = string(in.String())
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
func easyjson8bb5296eEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(out *jwriter.Writer, in FlowSchemaCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.LastTransitionTime).MarshalJSON())
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
	if in.Status != "" {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FlowSchemaCondition) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8bb5296eEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FlowSchemaCondition) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8bb5296eEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FlowSchemaCondition) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8bb5296eDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FlowSchemaCondition) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8bb5296eDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(l, v)
}
