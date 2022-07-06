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

func easyjsonD8856ac0DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1(in *jlexer.Lexer, out *EndpointSlice) {
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
		case "addressType":
			if in.IsNull() {
				in.Skip()
				out.AddressType = nil
			} else {
				if out.AddressType == nil {
					out.AddressType = new(string)
				}
				*out.AddressType = string(in.String())
			}
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "endpoints":
			if in.IsNull() {
				in.Skip()
				out.Endpoints = nil
			} else {
				in.Delim('[')
				if out.Endpoints == nil {
					if !in.IsDelim(']') {
						out.Endpoints = make([]*Endpoint, 0, 8)
					} else {
						out.Endpoints = []*Endpoint{}
					}
				} else {
					out.Endpoints = (out.Endpoints)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Endpoint
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Endpoint)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Endpoints = append(out.Endpoints, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "kind":
			out.Kind = string(in.String())
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
		case "ports":
			if in.IsNull() {
				in.Skip()
				out.Ports = nil
			} else {
				in.Delim('[')
				if out.Ports == nil {
					if !in.IsDelim(']') {
						out.Ports = make([]*EndpointPort, 0, 8)
					} else {
						out.Ports = []*EndpointPort{}
					}
				} else {
					out.Ports = (out.Ports)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *EndpointPort
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(EndpointPort)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.Ports = append(out.Ports, v2)
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
func easyjsonD8856ac0EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1(out *jwriter.Writer, in EndpointSlice) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"addressType\":"
		out.RawString(prefix[1:])
		if in.AddressType == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.AddressType))
		}
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		out.RawString(prefix)
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"endpoints\":"
		out.RawString(prefix)
		if in.Endpoints == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Endpoints {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
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
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		(*in.Metadata).MarshalEasyJSON(out)
	}
	if len(in.Ports) != 0 {
		const prefix string = ",\"ports\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Ports {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EndpointSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD8856ac0EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EndpointSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD8856ac0EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EndpointSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD8856ac0DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EndpointSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD8856ac0DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1(l, v)
}
