// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	_v11 "github.com/kubewarden/k8s-objects/api/core/v1"
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

func easyjsonC8b5e123DecodeGithubComKubewardenK8sObjectsApiAppsV1(in *jlexer.Lexer, out *StatefulSetSpec) {
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
		case "minReadySeconds":
			out.MinReadySeconds = int32(in.Int32())
		case "podManagementPolicy":
			out.PodManagementPolicy = string(in.String())
		case "replicas":
			out.Replicas = int32(in.Int32())
		case "revisionHistoryLimit":
			out.RevisionHistoryLimit = int32(in.Int32())
		case "selector":
			if in.IsNull() {
				in.Skip()
				out.Selector = nil
			} else {
				if out.Selector == nil {
					out.Selector = new(_v1.LabelSelector)
				}
				(*out.Selector).UnmarshalEasyJSON(in)
			}
		case "serviceName":
			if in.IsNull() {
				in.Skip()
				out.ServiceName = nil
			} else {
				if out.ServiceName == nil {
					out.ServiceName = new(string)
				}
				*out.ServiceName = string(in.String())
			}
		case "template":
			if in.IsNull() {
				in.Skip()
				out.Template = nil
			} else {
				if out.Template == nil {
					out.Template = new(_v11.PodTemplateSpec)
				}
				(*out.Template).UnmarshalEasyJSON(in)
			}
		case "updateStrategy":
			if in.IsNull() {
				in.Skip()
				out.UpdateStrategy = nil
			} else {
				if out.UpdateStrategy == nil {
					out.UpdateStrategy = new(StatefulSetUpdateStrategy)
				}
				easyjsonC8b5e123DecodeGithubComKubewardenK8sObjectsApiAppsV11(in, out.UpdateStrategy)
			}
		case "volumeClaimTemplates":
			if in.IsNull() {
				in.Skip()
				out.VolumeClaimTemplates = nil
			} else {
				in.Delim('[')
				if out.VolumeClaimTemplates == nil {
					if !in.IsDelim(']') {
						out.VolumeClaimTemplates = make([]_v11.PersistentVolumeClaim, 0, 0)
					} else {
						out.VolumeClaimTemplates = []_v11.PersistentVolumeClaim{}
					}
				} else {
					out.VolumeClaimTemplates = (out.VolumeClaimTemplates)[:0]
				}
				for !in.IsDelim(']') {
					var v1 _v11.PersistentVolumeClaim
					(v1).UnmarshalEasyJSON(in)
					out.VolumeClaimTemplates = append(out.VolumeClaimTemplates, v1)
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
func easyjsonC8b5e123EncodeGithubComKubewardenK8sObjectsApiAppsV1(out *jwriter.Writer, in StatefulSetSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MinReadySeconds != 0 {
		const prefix string = ",\"minReadySeconds\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.MinReadySeconds))
	}
	if in.PodManagementPolicy != "" {
		const prefix string = ",\"podManagementPolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PodManagementPolicy))
	}
	if in.Replicas != 0 {
		const prefix string = ",\"replicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Replicas))
	}
	if in.RevisionHistoryLimit != 0 {
		const prefix string = ",\"revisionHistoryLimit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.RevisionHistoryLimit))
	}
	{
		const prefix string = ",\"selector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Selector == nil {
			out.RawString("null")
		} else {
			(*in.Selector).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"serviceName\":"
		out.RawString(prefix)
		if in.ServiceName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ServiceName))
		}
	}
	{
		const prefix string = ",\"template\":"
		out.RawString(prefix)
		if in.Template == nil {
			out.RawString("null")
		} else {
			(*in.Template).MarshalEasyJSON(out)
		}
	}
	if in.UpdateStrategy != nil {
		const prefix string = ",\"updateStrategy\":"
		out.RawString(prefix)
		easyjsonC8b5e123EncodeGithubComKubewardenK8sObjectsApiAppsV11(out, *in.UpdateStrategy)
	}
	{
		const prefix string = ",\"volumeClaimTemplates\":"
		out.RawString(prefix)
		if in.VolumeClaimTemplates == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.VolumeClaimTemplates {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StatefulSetSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC8b5e123EncodeGithubComKubewardenK8sObjectsApiAppsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StatefulSetSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC8b5e123EncodeGithubComKubewardenK8sObjectsApiAppsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StatefulSetSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC8b5e123DecodeGithubComKubewardenK8sObjectsApiAppsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StatefulSetSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC8b5e123DecodeGithubComKubewardenK8sObjectsApiAppsV1(l, v)
}
func easyjsonC8b5e123DecodeGithubComKubewardenK8sObjectsApiAppsV11(in *jlexer.Lexer, out *StatefulSetUpdateStrategy) {
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
					out.RollingUpdate = new(RollingUpdateStatefulSetStrategy)
				}
				(*out.RollingUpdate).UnmarshalEasyJSON(in)
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
func easyjsonC8b5e123EncodeGithubComKubewardenK8sObjectsApiAppsV11(out *jwriter.Writer, in StatefulSetUpdateStrategy) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RollingUpdate != nil {
		const prefix string = ",\"rollingUpdate\":"
		first = false
		out.RawString(prefix[1:])
		(*in.RollingUpdate).MarshalEasyJSON(out)
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
