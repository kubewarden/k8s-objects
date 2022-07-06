// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v11 "github.com/kubewarden/k8s-objects/api/batch/v1"
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

func easyjson761e00a6DecodeGithubComKubewardenK8sObjectsApiBatchV1beta1(in *jlexer.Lexer, out *JobTemplateSpec) {
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
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(_v11.JobSpec)
				}
				(*out.Spec).UnmarshalEasyJSON(in)
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
func easyjson761e00a6EncodeGithubComKubewardenK8sObjectsApiBatchV1beta1(out *jwriter.Writer, in JobTemplateSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Spec).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JobTemplateSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson761e00a6EncodeGithubComKubewardenK8sObjectsApiBatchV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JobTemplateSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson761e00a6EncodeGithubComKubewardenK8sObjectsApiBatchV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JobTemplateSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson761e00a6DecodeGithubComKubewardenK8sObjectsApiBatchV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JobTemplateSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson761e00a6DecodeGithubComKubewardenK8sObjectsApiBatchV1beta1(l, v)
}
