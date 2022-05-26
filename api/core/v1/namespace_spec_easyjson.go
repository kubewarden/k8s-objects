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

func easyjsonC7539851DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *NamespaceSpec) {
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
		case "finalizers":
			if in.IsNull() {
				in.Skip()
				out.Finalizers = nil
			} else {
				in.Delim('[')
				if out.Finalizers == nil {
					if !in.IsDelim(']') {
						out.Finalizers = make([]string, 0, 4)
					} else {
						out.Finalizers = []string{}
					}
				} else {
					out.Finalizers = (out.Finalizers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Finalizers = append(out.Finalizers, v1)
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
func easyjsonC7539851EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in NamespaceSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"finalizers\":"
		out.RawString(prefix[1:])
		if in.Finalizers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Finalizers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NamespaceSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC7539851EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NamespaceSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC7539851EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NamespaceSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC7539851DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NamespaceSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC7539851DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
