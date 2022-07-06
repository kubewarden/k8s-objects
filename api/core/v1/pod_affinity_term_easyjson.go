// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjson18999177DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *PodAffinityTerm) {
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
		case "labelSelector":
			if in.IsNull() {
				in.Skip()
				out.LabelSelector = nil
			} else {
				if out.LabelSelector == nil {
					out.LabelSelector = new(_v1.LabelSelector)
				}
				(*out.LabelSelector).UnmarshalEasyJSON(in)
			}
		case "namespaces":
			if in.IsNull() {
				in.Skip()
				out.Namespaces = nil
			} else {
				in.Delim('[')
				if out.Namespaces == nil {
					if !in.IsDelim(']') {
						out.Namespaces = make([]string, 0, 4)
					} else {
						out.Namespaces = []string{}
					}
				} else {
					out.Namespaces = (out.Namespaces)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Namespaces = append(out.Namespaces, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "topologyKey":
			if in.IsNull() {
				in.Skip()
				out.TopologyKey = nil
			} else {
				if out.TopologyKey == nil {
					out.TopologyKey = new(string)
				}
				*out.TopologyKey = string(in.String())
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
func easyjson18999177EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in PodAffinityTerm) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LabelSelector != nil {
		const prefix string = ",\"labelSelector\":"
		first = false
		out.RawString(prefix[1:])
		(*in.LabelSelector).MarshalEasyJSON(out)
	}
	if len(in.Namespaces) != 0 {
		const prefix string = ",\"namespaces\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Namespaces {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"topologyKey\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.TopologyKey == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TopologyKey))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodAffinityTerm) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18999177EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodAffinityTerm) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18999177EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodAffinityTerm) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18999177DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodAffinityTerm) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18999177DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
