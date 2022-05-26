// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	intstr "github.com/kubewarden/k8s-objects/apimachinery/pkg/util/intstr"
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

func easyjson2b5a1f69DecodeGithubComKubewardenK8sObjectsApiAppsV1(in *jlexer.Lexer, out *DaemonSetUpdateStrategy) {
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
		case "rollingUpdate":
			if in.IsNull() {
				in.Skip()
				out.RollingUpdate = nil
			} else {
				if out.RollingUpdate == nil {
					out.RollingUpdate = new(RollingUpdateDaemonSet)
				}
				easyjson2b5a1f69DecodeGithubComKubewardenK8sObjectsApiAppsV11(in, out.RollingUpdate)
			}
		case "type":
			out.Type = string(in.String())
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
func easyjson2b5a1f69EncodeGithubComKubewardenK8sObjectsApiAppsV1(out *jwriter.Writer, in DaemonSetUpdateStrategy) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RollingUpdate != nil {
		const prefix string = ",\"rollingUpdate\":"
		first = false
		out.RawString(prefix[1:])
		easyjson2b5a1f69EncodeGithubComKubewardenK8sObjectsApiAppsV11(out, *in.RollingUpdate)
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DaemonSetUpdateStrategy) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2b5a1f69EncodeGithubComKubewardenK8sObjectsApiAppsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DaemonSetUpdateStrategy) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2b5a1f69EncodeGithubComKubewardenK8sObjectsApiAppsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DaemonSetUpdateStrategy) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2b5a1f69DecodeGithubComKubewardenK8sObjectsApiAppsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DaemonSetUpdateStrategy) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2b5a1f69DecodeGithubComKubewardenK8sObjectsApiAppsV1(l, v)
}
func easyjson2b5a1f69DecodeGithubComKubewardenK8sObjectsApiAppsV11(in *jlexer.Lexer, out *RollingUpdateDaemonSet) {
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
		case "maxUnavailable":
			out.MaxUnavailable = intstr.IntOrString(in.String())
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
func easyjson2b5a1f69EncodeGithubComKubewardenK8sObjectsApiAppsV11(out *jwriter.Writer, in RollingUpdateDaemonSet) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MaxUnavailable != "" {
		const prefix string = ",\"maxUnavailable\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.MaxUnavailable))
	}
	out.RawByte('}')
}
