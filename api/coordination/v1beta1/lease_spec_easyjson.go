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

func easyjsonDdb58c80DecodeGithubComKubewardenK8sObjectsApiCoordinationV1beta1(in *jlexer.Lexer, out *LeaseSpec) {
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
		case "acquireTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.AcquireTime).UnmarshalJSON(data))
			}
		case "holderIdentity":
			out.HolderIdentity = string(in.String())
		case "leaseDurationSeconds":
			out.LeaseDurationSeconds = int32(in.Int32())
		case "leaseTransitions":
			out.LeaseTransitions = int32(in.Int32())
		case "renewTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.RenewTime).UnmarshalJSON(data))
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
func easyjsonDdb58c80EncodeGithubComKubewardenK8sObjectsApiCoordinationV1beta1(out *jwriter.Writer, in LeaseSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"acquireTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.AcquireTime).MarshalJSON())
	}
	if in.HolderIdentity != "" {
		const prefix string = ",\"holderIdentity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HolderIdentity))
	}
	if in.LeaseDurationSeconds != 0 {
		const prefix string = ",\"leaseDurationSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.LeaseDurationSeconds))
	}
	if in.LeaseTransitions != 0 {
		const prefix string = ",\"leaseTransitions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.LeaseTransitions))
	}
	if true {
		const prefix string = ",\"renewTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.RenewTime).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LeaseSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDdb58c80EncodeGithubComKubewardenK8sObjectsApiCoordinationV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LeaseSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDdb58c80EncodeGithubComKubewardenK8sObjectsApiCoordinationV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LeaseSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDdb58c80DecodeGithubComKubewardenK8sObjectsApiCoordinationV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LeaseSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDdb58c80DecodeGithubComKubewardenK8sObjectsApiCoordinationV1beta1(l, v)
}
