// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha2

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

func easyjson7f0e689bDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(in *jlexer.Lexer, out *ResourceClaimSchedulingStatus) {
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
		case "name":
			out.Name = string(in.String())
		case "unsuitableNodes":
			if in.IsNull() {
				in.Skip()
				out.UnsuitableNodes = nil
			} else {
				in.Delim('[')
				if out.UnsuitableNodes == nil {
					if !in.IsDelim(']') {
						out.UnsuitableNodes = make([]string, 0, 4)
					} else {
						out.UnsuitableNodes = []string{}
					}
				} else {
					out.UnsuitableNodes = (out.UnsuitableNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.UnsuitableNodes = append(out.UnsuitableNodes, v1)
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
func easyjson7f0e689bEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(out *jwriter.Writer, in ResourceClaimSchedulingStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if len(in.UnsuitableNodes) != 0 {
		const prefix string = ",\"unsuitableNodes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.UnsuitableNodes {
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
func (v ResourceClaimSchedulingStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7f0e689bEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResourceClaimSchedulingStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7f0e689bEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResourceClaimSchedulingStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7f0e689bDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResourceClaimSchedulingStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7f0e689bDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(l, v)
}