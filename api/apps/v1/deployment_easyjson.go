// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	_v11 "github.com/kubewarden/k8s-objects/api/core/v1"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
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

func easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV1(in *jlexer.Lexer, out *Deployment) {
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
					out.Spec = new(DeploymentSpec)
				}
				easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(DeploymentStatus)
				}
				easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV12(in, out.Status)
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
func easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV1(out *jwriter.Writer, in Deployment) {
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
		easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Deployment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Deployment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Deployment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Deployment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV1(l, v)
}
func easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV12(in *jlexer.Lexer, out *DeploymentStatus) {
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
						out.Conditions = make([]*DeploymentCondition, 0, 8)
					} else {
						out.Conditions = []*DeploymentCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *DeploymentCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DeploymentCondition)
						}
						easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV13(in, v1)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
		case "readyReplicas":
			out.ReadyReplicas = int32(in.Int32())
		case "replicas":
			out.Replicas = int32(in.Int32())
		case "unavailableReplicas":
			out.UnavailableReplicas = int32(in.Int32())
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
func easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV12(out *jwriter.Writer, in DeploymentStatus) {
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
					easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV13(out, *v3)
				}
			}
			out.RawByte(']')
		}
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
	if in.UnavailableReplicas != 0 {
		const prefix string = ",\"unavailableReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.UnavailableReplicas))
	}
	if in.UpdatedReplicas != 0 {
		const prefix string = ",\"updatedReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.UpdatedReplicas))
	}
	out.RawByte('}')
}
func easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV13(in *jlexer.Lexer, out *DeploymentCondition) {
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
		case "lastUpdateTime":
			if in.IsNull() {
				in.Skip()
				out.LastUpdateTime = nil
			} else {
				if out.LastUpdateTime == nil {
					out.LastUpdateTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastUpdateTime).UnmarshalJSON(data))
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
func easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV13(out *jwriter.Writer, in DeploymentCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastTransitionTime != nil {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.LastTransitionTime).MarshalJSON())
	}
	if in.LastUpdateTime != nil {
		const prefix string = ",\"lastUpdateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.LastUpdateTime).MarshalJSON())
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
func easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV11(in *jlexer.Lexer, out *DeploymentSpec) {
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
		case "paused":
			out.Paused = bool(in.Bool())
		case "progressDeadlineSeconds":
			out.ProgressDeadlineSeconds = int32(in.Int32())
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
		case "strategy":
			if in.IsNull() {
				in.Skip()
				out.Strategy = nil
			} else {
				if out.Strategy == nil {
					out.Strategy = new(DeploymentStrategy)
				}
				easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV14(in, out.Strategy)
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
func easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV11(out *jwriter.Writer, in DeploymentSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MinReadySeconds != 0 {
		const prefix string = ",\"minReadySeconds\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.MinReadySeconds))
	}
	if in.Paused {
		const prefix string = ",\"paused\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Paused))
	}
	if in.ProgressDeadlineSeconds != 0 {
		const prefix string = ",\"progressDeadlineSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ProgressDeadlineSeconds))
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
	if in.Strategy != nil {
		const prefix string = ",\"strategy\":"
		out.RawString(prefix)
		easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV14(out, *in.Strategy)
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
	out.RawByte('}')
}
func easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV14(in *jlexer.Lexer, out *DeploymentStrategy) {
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
					out.RollingUpdate = new(RollingUpdateDeployment)
				}
				easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV15(in, out.RollingUpdate)
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
func easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV14(out *jwriter.Writer, in DeploymentStrategy) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RollingUpdate != nil {
		const prefix string = ",\"rollingUpdate\":"
		first = false
		out.RawString(prefix[1:])
		easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV15(out, *in.RollingUpdate)
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
func easyjsonC7c6e959DecodeGithubComKubewardenK8sObjectsApiAppsV15(in *jlexer.Lexer, out *RollingUpdateDeployment) {
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
		case "maxSurge":
			if in.IsNull() {
				in.Skip()
				out.MaxSurge = nil
			} else {
				if out.MaxSurge == nil {
					out.MaxSurge = new(intstr.IntOrString)
				}
				*out.MaxSurge = intstr.IntOrString(in.String())
			}
		case "maxUnavailable":
			if in.IsNull() {
				in.Skip()
				out.MaxUnavailable = nil
			} else {
				if out.MaxUnavailable == nil {
					out.MaxUnavailable = new(intstr.IntOrString)
				}
				*out.MaxUnavailable = intstr.IntOrString(in.String())
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
func easyjsonC7c6e959EncodeGithubComKubewardenK8sObjectsApiAppsV15(out *jwriter.Writer, in RollingUpdateDeployment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MaxSurge != nil {
		const prefix string = ",\"maxSurge\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.MaxSurge))
	}
	if in.MaxUnavailable != nil {
		const prefix string = ",\"maxUnavailable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.MaxUnavailable))
	}
	out.RawByte('}')
}
