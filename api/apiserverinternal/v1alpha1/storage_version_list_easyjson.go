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

func easyjson76341d05DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(in *jlexer.Lexer, out *StorageVersionList) {
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
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]*StorageVersion, 0, 8)
					} else {
						out.Items = []*StorageVersion{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *StorageVersion
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(StorageVersion)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Items = append(out.Items, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "kind":
			out.Kind = string(in.String())
		case "metadata":
			(out.Metadata).UnmarshalEasyJSON(in)
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
func easyjson76341d05EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(out *jwriter.Writer, in StorageVersionList) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"items\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Items {
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
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	if true {
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		(in.Metadata).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StorageVersionList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson76341d05EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StorageVersionList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson76341d05EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StorageVersionList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson76341d05DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StorageVersionList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson76341d05DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(l, v)
}
