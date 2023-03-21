// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjson1a48f9c9DecodeGithubComKubewardenK8sObjectsApiNetworkingV1alpha1(in *jlexer.Lexer, out *ClusterCIDRSpec) {
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
		case "ipv4":
			out.IPV4 = string(in.String())
		case "ipv6":
			out.IPV6 = string(in.String())
		case "nodeSelector":
			if in.IsNull() {
				in.Skip()
				out.NodeSelector = nil
			} else {
				if out.NodeSelector == nil {
					out.NodeSelector = new(_v1.NodeSelector)
				}
				(*out.NodeSelector).UnmarshalEasyJSON(in)
			}
		case "perNodeHostBits":
			if in.IsNull() {
				in.Skip()
				out.PerNodeHostBits = nil
			} else {
				if out.PerNodeHostBits == nil {
					out.PerNodeHostBits = new(int32)
				}
				*out.PerNodeHostBits = int32(in.Int32())
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
func easyjson1a48f9c9EncodeGithubComKubewardenK8sObjectsApiNetworkingV1alpha1(out *jwriter.Writer, in ClusterCIDRSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IPV4 != "" {
		const prefix string = ",\"ipv4\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.IPV4))
	}
	if in.IPV6 != "" {
		const prefix string = ",\"ipv6\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IPV6))
	}
	if in.NodeSelector != nil {
		const prefix string = ",\"nodeSelector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.NodeSelector).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"perNodeHostBits\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.PerNodeHostBits == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.PerNodeHostBits))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClusterCIDRSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1a48f9c9EncodeGithubComKubewardenK8sObjectsApiNetworkingV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClusterCIDRSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1a48f9c9EncodeGithubComKubewardenK8sObjectsApiNetworkingV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClusterCIDRSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1a48f9c9DecodeGithubComKubewardenK8sObjectsApiNetworkingV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClusterCIDRSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1a48f9c9DecodeGithubComKubewardenK8sObjectsApiNetworkingV1alpha1(l, v)
}