// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

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

func easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *DeploymentSpec) {
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
		case "rollbackTo":
			if in.IsNull() {
				in.Skip()
				out.RollbackTo = nil
			} else {
				if out.RollbackTo == nil {
					out.RollbackTo = new(RollbackConfig)
				}
				easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in, out.RollbackTo)
			}
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
				easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(in, out.Strategy)
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
func easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in DeploymentSpec) {
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
	if in.RollbackTo != nil {
		const prefix string = ",\"rollbackTo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out, *in.RollbackTo)
	}
	if in.Selector != nil {
		const prefix string = ",\"selector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Selector).MarshalEasyJSON(out)
	}
	if in.Strategy != nil {
		const prefix string = ",\"strategy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(out, *in.Strategy)
	}
	{
		const prefix string = ",\"template\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Template == nil {
			out.RawString("null")
		} else {
			(*in.Template).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeploymentSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeploymentSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeploymentSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeploymentSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
func easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(in *jlexer.Lexer, out *DeploymentStrategy) {
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
				easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(in, out.RollingUpdate)
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
func easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(out *jwriter.Writer, in DeploymentStrategy) {
	out.RawByte('{')
	first := true
	_ = first
	if in.RollingUpdate != nil {
		const prefix string = ",\"rollingUpdate\":"
		first = false
		out.RawString(prefix[1:])
		easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(out, *in.RollingUpdate)
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
func easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(in *jlexer.Lexer, out *RollingUpdateDeployment) {
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
func easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(out *jwriter.Writer, in RollingUpdateDeployment) {
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
func easyjson9460807fDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in *jlexer.Lexer, out *RollbackConfig) {
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
		case "revision":
			out.Revision = int64(in.Int64())
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
func easyjson9460807fEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out *jwriter.Writer, in RollbackConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Revision != 0 {
		const prefix string = ",\"revision\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Revision))
	}
	out.RawByte('}')
}
