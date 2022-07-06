// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
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

func easyjson3a9ed663DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta1(in *jlexer.Lexer, out *Endpoint) {
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
		case "addresses":
			if in.IsNull() {
				in.Skip()
				out.Addresses = nil
			} else {
				in.Delim('[')
				if out.Addresses == nil {
					if !in.IsDelim(']') {
						out.Addresses = make([]string, 0, 4)
					} else {
						out.Addresses = []string{}
					}
				} else {
					out.Addresses = (out.Addresses)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Addresses = append(out.Addresses, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				if out.Conditions == nil {
					out.Conditions = new(EndpointConditions)
				}
				easyjson3a9ed663DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta11(in, out.Conditions)
			}
		case "hostname":
			out.Hostname = string(in.String())
		case "nodeName":
			out.NodeName = string(in.String())
		case "targetRef":
			if in.IsNull() {
				in.Skip()
				out.TargetRef = nil
			} else {
				if out.TargetRef == nil {
					out.TargetRef = new(_v1.ObjectReference)
				}
				(*out.TargetRef).UnmarshalEasyJSON(in)
			}
		case "topology":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Topology = make(map[string]string)
				} else {
					out.Topology = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 string
					v2 = string(in.String())
					(out.Topology)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
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
func easyjson3a9ed663EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta1(out *jwriter.Writer, in Endpoint) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"addresses\":"
		out.RawString(prefix[1:])
		if in.Addresses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Addresses {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if in.Conditions != nil {
		const prefix string = ",\"conditions\":"
		out.RawString(prefix)
		easyjson3a9ed663EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta11(out, *in.Conditions)
	}
	if in.Hostname != "" {
		const prefix string = ",\"hostname\":"
		out.RawString(prefix)
		out.String(string(in.Hostname))
	}
	if in.NodeName != "" {
		const prefix string = ",\"nodeName\":"
		out.RawString(prefix)
		out.String(string(in.NodeName))
	}
	if in.TargetRef != nil {
		const prefix string = ",\"targetRef\":"
		out.RawString(prefix)
		(*in.TargetRef).MarshalEasyJSON(out)
	}
	if len(in.Topology) != 0 {
		const prefix string = ",\"topology\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Topology {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.String(string(v5Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Endpoint) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3a9ed663EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Endpoint) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3a9ed663EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Endpoint) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3a9ed663DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Endpoint) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3a9ed663DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta1(l, v)
}
func easyjson3a9ed663DecodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta11(in *jlexer.Lexer, out *EndpointConditions) {
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
		case "ready":
			out.Ready = bool(in.Bool())
		case "serving":
			out.Serving = bool(in.Bool())
		case "terminating":
			out.Terminating = bool(in.Bool())
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
func easyjson3a9ed663EncodeGithubComKubewardenK8sObjectsApiDiscoveryV1beta11(out *jwriter.Writer, in EndpointConditions) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ready {
		const prefix string = ",\"ready\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.Ready))
	}
	if in.Serving {
		const prefix string = ",\"serving\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Serving))
	}
	if in.Terminating {
		const prefix string = ",\"terminating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Terminating))
	}
	out.RawByte('}')
}
