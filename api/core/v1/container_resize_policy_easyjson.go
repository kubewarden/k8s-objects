// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjson8e24b043DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ContainerResizePolicy) {
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
		case "resourceName":
			if in.IsNull() {
				in.Skip()
				out.ResourceName = nil
			} else {
				if out.ResourceName == nil {
					out.ResourceName = new(string)
				}
				*out.ResourceName = string(in.String())
			}
		case "restartPolicy":
			if in.IsNull() {
				in.Skip()
				out.RestartPolicy = nil
			} else {
				if out.RestartPolicy == nil {
					out.RestartPolicy = new(string)
				}
				*out.RestartPolicy = string(in.String())
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
func easyjson8e24b043EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ContainerResizePolicy) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"resourceName\":"
		out.RawString(prefix[1:])
		if in.ResourceName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ResourceName))
		}
	}
	{
		const prefix string = ",\"restartPolicy\":"
		out.RawString(prefix)
		if in.RestartPolicy == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.RestartPolicy))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContainerResizePolicy) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8e24b043EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContainerResizePolicy) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8e24b043EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContainerResizePolicy) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8e24b043DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContainerResizePolicy) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8e24b043DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}