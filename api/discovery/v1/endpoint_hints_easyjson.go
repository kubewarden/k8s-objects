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

func easyjson471dc258DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1(in *jlexer.Lexer, out *EndpointHints) {
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
		case "forZones":
			if in.IsNull() {
				in.Skip()
				out.ForZones = nil
			} else {
				in.Delim('[')
				if out.ForZones == nil {
					if !in.IsDelim(']') {
						out.ForZones = make([]*ForZone, 0, 8)
					} else {
						out.ForZones = []*ForZone{}
					}
				} else {
					out.ForZones = (out.ForZones)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *ForZone
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(ForZone)
						}
						easyjson471dc258DecodeGithubComKubewardenK8sObjectsApiDiscoveryV11(in, v1)
					}
					out.ForZones = append(out.ForZones, v1)
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
func easyjson471dc258EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1(out *jwriter.Writer, in EndpointHints) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"forZones\":"
		out.RawString(prefix[1:])
		if in.ForZones == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.ForZones {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson471dc258EncodeGithubComKubewardenK8sObjectsApiDiscoveryV11(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EndpointHints) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson471dc258EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EndpointHints) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson471dc258EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EndpointHints) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson471dc258DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EndpointHints) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson471dc258DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1(l, v)
}
func easyjson471dc258DecodeGithubComKubewardenK8sObjectsApiDiscoveryV11(in *jlexer.Lexer, out *ForZone) {
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
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
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
func easyjson471dc258EncodeGithubComKubewardenK8sObjectsApiDiscoveryV11(out *jwriter.Writer, in ForZone) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}
