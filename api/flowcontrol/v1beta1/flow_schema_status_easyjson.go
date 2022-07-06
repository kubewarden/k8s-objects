// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

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

func easyjson3312364bDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta1(in *jlexer.Lexer, out *FlowSchemaStatus) {
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
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*FlowSchemaCondition, 0, 8)
					} else {
						out.Conditions = []*FlowSchemaCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *FlowSchemaCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(FlowSchemaCondition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson3312364bEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta1(out *jwriter.Writer, in FlowSchemaStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		first = false
		out.RawString(prefix[1:])
		{
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FlowSchemaStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3312364bEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FlowSchemaStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3312364bEncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FlowSchemaStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3312364bDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FlowSchemaStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3312364bDecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1beta1(l, v)
}
