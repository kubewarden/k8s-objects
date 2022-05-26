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

func easyjson57e69005DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(in *jlexer.Lexer, out *CustomResourceDefinitionNames) {
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
		case "categories":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]string, 0, 4)
					} else {
						out.Categories = []string{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Categories = append(out.Categories, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
			}
		case "listKind":
			out.ListKind = string(in.String())
		case "plural":
			if in.IsNull() {
				in.Skip()
				out.Plural = nil
			} else {
				if out.Plural == nil {
					out.Plural = new(string)
				}
				*out.Plural = string(in.String())
			}
		case "shortNames":
			if in.IsNull() {
				in.Skip()
				out.ShortNames = nil
			} else {
				in.Delim('[')
				if out.ShortNames == nil {
					if !in.IsDelim(']') {
						out.ShortNames = make([]string, 0, 4)
					} else {
						out.ShortNames = []string{}
					}
				} else {
					out.ShortNames = (out.ShortNames)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.ShortNames = append(out.ShortNames, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "singular":
			out.Singular = string(in.String())
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
func easyjson57e69005EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(out *jwriter.Writer, in CustomResourceDefinitionNames) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"categories\":"
		out.RawString(prefix[1:])
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Categories {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Kind))
		}
	}
	if in.ListKind != "" {
		const prefix string = ",\"listKind\":"
		out.RawString(prefix)
		out.String(string(in.ListKind))
	}
	{
		const prefix string = ",\"plural\":"
		out.RawString(prefix)
		if in.Plural == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Plural))
		}
	}
	{
		const prefix string = ",\"shortNames\":"
		out.RawString(prefix)
		if in.ShortNames == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.ShortNames {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if in.Singular != "" {
		const prefix string = ",\"singular\":"
		out.RawString(prefix)
		out.String(string(in.Singular))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomResourceDefinitionNames) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson57e69005EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CustomResourceDefinitionNames) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson57e69005EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomResourceDefinitionNames) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson57e69005DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CustomResourceDefinitionNames) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson57e69005DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(l, v)
}