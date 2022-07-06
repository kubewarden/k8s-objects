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

func easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV1(in *jlexer.Lexer, out *StatefulSet) {
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
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(StatefulSetSpec)
				}
				easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(StatefulSetStatus)
				}
				easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV12(in, out.Status)
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
func easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV1(out *jwriter.Writer, in StatefulSet) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
	}
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StatefulSet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StatefulSet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StatefulSet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StatefulSet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV1(l, v)
}
func easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV12(in *jlexer.Lexer, out *StatefulSetStatus) {
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
		case "availableReplicas":
			out.AvailableReplicas = int32(in.Int32())
		case "collisionCount":
			out.CollisionCount = int32(in.Int32())
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*StatefulSetCondition, 0, 8)
					} else {
						out.Conditions = []*StatefulSetCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *StatefulSetCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(StatefulSetCondition)
						}
						easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV13(in, v1)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "currentReplicas":
			out.CurrentReplicas = int32(in.Int32())
		case "currentRevision":
			out.CurrentRevision = string(in.String())
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
		case "readyReplicas":
			out.ReadyReplicas = int32(in.Int32())
		case "replicas":
			if in.IsNull() {
				in.Skip()
				out.Replicas = nil
			} else {
				if out.Replicas == nil {
					out.Replicas = new(int32)
				}
				*out.Replicas = int32(in.Int32())
			}
		case "updateRevision":
			out.UpdateRevision = string(in.String())
		case "updatedReplicas":
			out.UpdatedReplicas = int32(in.Int32())
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
func easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV12(out *jwriter.Writer, in StatefulSetStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AvailableReplicas != 0 {
		const prefix string = ",\"availableReplicas\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AvailableReplicas))
	}
	if in.CollisionCount != 0 {
		const prefix string = ",\"collisionCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CollisionCount))
	}
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV13(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	if in.CurrentReplicas != 0 {
		const prefix string = ",\"currentReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CurrentReplicas))
	}
	if in.CurrentRevision != "" {
		const prefix string = ",\"currentRevision\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CurrentRevision))
	}
	if in.ObservedGeneration != 0 {
		const prefix string = ",\"observedGeneration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ObservedGeneration))
	}
	if in.ReadyReplicas != 0 {
		const prefix string = ",\"readyReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ReadyReplicas))
	}
	{
		const prefix string = ",\"replicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Replicas == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Replicas))
		}
	}
	if in.UpdateRevision != "" {
		const prefix string = ",\"updateRevision\":"
		out.RawString(prefix)
		out.String(string(in.UpdateRevision))
	}
	if in.UpdatedReplicas != 0 {
		const prefix string = ",\"updatedReplicas\":"
		out.RawString(prefix)
		out.Int32(int32(in.UpdatedReplicas))
	}
	out.RawByte('}')
}
func easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV13(in *jlexer.Lexer, out *StatefulSetCondition) {
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
		case "lastTransitionTime":
			if in.IsNull() {
				in.Skip()
				out.LastTransitionTime = nil
			} else {
				if out.LastTransitionTime == nil {
					out.LastTransitionTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastTransitionTime).UnmarshalJSON(data))
				}
			}
		case "message":
			out.Message = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(string)
				}
				*out.Status = string(in.String())
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
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
func easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV13(out *jwriter.Writer, in StatefulSetCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastTransitionTime != nil {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.LastTransitionTime).MarshalJSON())
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Reason))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Status == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Status))
		}
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}
func easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV11(in *jlexer.Lexer, out *StatefulSetSpec) {
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
				easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV14(in, out.UpdateStrategy)
			}
		case "volumeClaimTemplates":
			if in.IsNull() {
				in.Skip()
				out.VolumeClaimTemplates = nil
			} else {
				in.Delim('[')
				if out.VolumeClaimTemplates == nil {
					if !in.IsDelim(']') {
						out.VolumeClaimTemplates = make([]*_v11.PersistentVolumeClaim, 0, 8)
					} else {
						out.VolumeClaimTemplates = []*_v11.PersistentVolumeClaim{}
					}
				} else {
					out.VolumeClaimTemplates = (out.VolumeClaimTemplates)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *_v11.PersistentVolumeClaim
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(_v11.PersistentVolumeClaim)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.VolumeClaimTemplates = append(out.VolumeClaimTemplates, v4)
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
func easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV11(out *jwriter.Writer, in StatefulSetSpec) {
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
		easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV14(out, *in.UpdateStrategy)
	}
	if len(in.VolumeClaimTemplates) != 0 {
		const prefix string = ",\"volumeClaimTemplates\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.VolumeClaimTemplates {
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
func easyjsonB4d3c875DecodeGithubComKubewardenK8sObjectsApiAppsV14(in *jlexer.Lexer, out *StatefulSetUpdateStrategy) {
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
func easyjsonB4d3c875EncodeGithubComKubewardenK8sObjectsApiAppsV14(out *jwriter.Writer, in StatefulSetUpdateStrategy) {
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
