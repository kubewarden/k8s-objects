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

func easyjson545ee795DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *IPBlock) {
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
		case "cidr":
			if in.IsNull() {
				in.Skip()
				out.Cidr = nil
			} else {
				if out.Cidr == nil {
					out.Cidr = new(string)
				}
				*out.Cidr = string(in.String())
			}
		case "except":
			if in.IsNull() {
				in.Skip()
				out.Except = nil
			} else {
				in.Delim('[')
				if out.Except == nil {
					if !in.IsDelim(']') {
						out.Except = make([]string, 0, 4)
					} else {
						out.Except = []string{}
					}
				} else {
					out.Except = (out.Except)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Except = append(out.Except, v1)
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
func easyjson545ee795EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in IPBlock) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cidr\":"
		out.RawString(prefix[1:])
		if in.Cidr == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Cidr))
		}
	}
	{
		const prefix string = ",\"except\":"
		out.RawString(prefix)
		if in.Except == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Except {
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
func (v IPBlock) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson545ee795EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IPBlock) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson545ee795EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IPBlock) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson545ee795DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IPBlock) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson545ee795DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}