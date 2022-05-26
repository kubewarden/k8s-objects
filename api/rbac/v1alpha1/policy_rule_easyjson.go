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

func easyjsonD6e22eadDecodeGithubComKubewardenK8sObjectsApiRbacV1alpha1(in *jlexer.Lexer, out *PolicyRule) {
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
		case "apiGroups":
			if in.IsNull() {
				in.Skip()
				out.APIGroups = nil
			} else {
				in.Delim('[')
				if out.APIGroups == nil {
					if !in.IsDelim(']') {
						out.APIGroups = make([]string, 0, 4)
					} else {
						out.APIGroups = []string{}
					}
				} else {
					out.APIGroups = (out.APIGroups)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.APIGroups = append(out.APIGroups, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "nonResourceURLs":
			if in.IsNull() {
				in.Skip()
				out.NonResourceURLs = nil
			} else {
				in.Delim('[')
				if out.NonResourceURLs == nil {
					if !in.IsDelim(']') {
						out.NonResourceURLs = make([]string, 0, 4)
					} else {
						out.NonResourceURLs = []string{}
					}
				} else {
					out.NonResourceURLs = (out.NonResourceURLs)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.NonResourceURLs = append(out.NonResourceURLs, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "resourceNames":
			if in.IsNull() {
				in.Skip()
				out.ResourceNames = nil
			} else {
				in.Delim('[')
				if out.ResourceNames == nil {
					if !in.IsDelim(']') {
						out.ResourceNames = make([]string, 0, 4)
					} else {
						out.ResourceNames = []string{}
					}
				} else {
					out.ResourceNames = (out.ResourceNames)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.ResourceNames = append(out.ResourceNames, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "resources":
			if in.IsNull() {
				in.Skip()
				out.Resources = nil
			} else {
				in.Delim('[')
				if out.Resources == nil {
					if !in.IsDelim(']') {
						out.Resources = make([]string, 0, 4)
					} else {
						out.Resources = []string{}
					}
				} else {
					out.Resources = (out.Resources)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Resources = append(out.Resources, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "verbs":
			if in.IsNull() {
				in.Skip()
				out.Verbs = nil
			} else {
				in.Delim('[')
				if out.Verbs == nil {
					if !in.IsDelim(']') {
						out.Verbs = make([]string, 0, 4)
					} else {
						out.Verbs = []string{}
					}
				} else {
					out.Verbs = (out.Verbs)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.Verbs = append(out.Verbs, v5)
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
func easyjsonD6e22eadEncodeGithubComKubewardenK8sObjectsApiRbacV1alpha1(out *jwriter.Writer, in PolicyRule) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"apiGroups\":"
		out.RawString(prefix[1:])
		if in.APIGroups == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.APIGroups {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"nonResourceURLs\":"
		out.RawString(prefix)
		if in.NonResourceURLs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.NonResourceURLs {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"resourceNames\":"
		out.RawString(prefix)
		if in.ResourceNames == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.ResourceNames {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"resources\":"
		out.RawString(prefix)
		if in.Resources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Resources {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"verbs\":"
		out.RawString(prefix)
		if in.Verbs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Verbs {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PolicyRule) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD6e22eadEncodeGithubComKubewardenK8sObjectsApiRbacV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PolicyRule) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD6e22eadEncodeGithubComKubewardenK8sObjectsApiRbacV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PolicyRule) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD6e22eadDecodeGithubComKubewardenK8sObjectsApiRbacV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PolicyRule) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD6e22eadDecodeGithubComKubewardenK8sObjectsApiRbacV1alpha1(l, v)
}