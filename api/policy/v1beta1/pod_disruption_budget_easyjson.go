// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
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

func easyjson817bc1ebDecodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(in *jlexer.Lexer, out *PodDisruptionBudget) {
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
					out.Spec = new(PodDisruptionBudgetSpec)
				}
				easyjson817bc1ebDecodeGithubComKubewardenK8sObjectsApiPolicyV1beta11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(PodDisruptionBudgetStatus)
				}
				easyjson817bc1ebDecodeGithubComKubewardenK8sObjectsApiPolicyV1beta12(in, out.Status)
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
func easyjson817bc1ebEncodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(out *jwriter.Writer, in PodDisruptionBudget) {
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
		easyjson817bc1ebEncodeGithubComKubewardenK8sObjectsApiPolicyV1beta11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson817bc1ebEncodeGithubComKubewardenK8sObjectsApiPolicyV1beta12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodDisruptionBudget) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson817bc1ebEncodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodDisruptionBudget) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson817bc1ebEncodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodDisruptionBudget) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson817bc1ebDecodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodDisruptionBudget) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson817bc1ebDecodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(l, v)
}
func easyjson817bc1ebDecodeGithubComKubewardenK8sObjectsApiPolicyV1beta12(in *jlexer.Lexer, out *PodDisruptionBudgetStatus) {
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
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*_v1.Condition, 0, 8)
					} else {
						out.Conditions = []*_v1.Condition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *_v1.Condition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(_v1.Condition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "currentHealthy":
			if in.IsNull() {
				in.Skip()
				out.CurrentHealthy = nil
			} else {
				if out.CurrentHealthy == nil {
					out.CurrentHealthy = new(int32)
				}
				*out.CurrentHealthy = int32(in.Int32())
			}
		case "desiredHealthy":
			if in.IsNull() {
				in.Skip()
				out.DesiredHealthy = nil
			} else {
				if out.DesiredHealthy == nil {
					out.DesiredHealthy = new(int32)
				}
				*out.DesiredHealthy = int32(in.Int32())
			}
		case "disruptedPods":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.DisruptedPods = make(map[string]*_v1.Time)
				} else {
					out.DisruptedPods = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 *_v1.Time
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(_v1.Time)
						}
						if data := in.Raw(); in.Ok() {
							in.AddError((*v2).UnmarshalJSON(data))
						}
					}
					(out.DisruptedPods)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "disruptionsAllowed":
			if in.IsNull() {
				in.Skip()
				out.DisruptionsAllowed = nil
			} else {
				if out.DisruptionsAllowed == nil {
					out.DisruptionsAllowed = new(int32)
				}
				*out.DisruptionsAllowed = int32(in.Int32())
			}
		case "expectedPods":
			if in.IsNull() {
				in.Skip()
				out.ExpectedPods = nil
			} else {
				if out.ExpectedPods == nil {
					out.ExpectedPods = new(int32)
				}
				*out.ExpectedPods = int32(in.Int32())
			}
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
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
func easyjson817bc1ebEncodeGithubComKubewardenK8sObjectsApiPolicyV1beta12(out *jwriter.Writer, in PodDisruptionBudgetStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v3, v4 := range in.Conditions {
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
	{
		const prefix string = ",\"currentHealthy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.CurrentHealthy == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.CurrentHealthy))
		}
	}
	{
		const prefix string = ",\"desiredHealthy\":"
		out.RawString(prefix)
		if in.DesiredHealthy == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.DesiredHealthy))
		}
	}
	if len(in.DisruptedPods) != 0 {
		const prefix string = ",\"disruptedPods\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.DisruptedPods {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				if v5Value == nil {
					out.RawString("null")
				} else {
					out.Raw((*v5Value).MarshalJSON())
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"disruptionsAllowed\":"
		out.RawString(prefix)
		if in.DisruptionsAllowed == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.DisruptionsAllowed))
		}
	}
	{
		const prefix string = ",\"expectedPods\":"
		out.RawString(prefix)
		if in.ExpectedPods == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.ExpectedPods))
		}
	}
	if in.ObservedGeneration != 0 {
		const prefix string = ",\"observedGeneration\":"
		out.RawString(prefix)
		out.Int64(int64(in.ObservedGeneration))
	}
	out.RawByte('}')
}
func easyjson817bc1ebDecodeGithubComKubewardenK8sObjectsApiPolicyV1beta11(in *jlexer.Lexer, out *PodDisruptionBudgetSpec) {
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
			if in.IsNull() {
				in.Skip()
				out.MaxUnavailable = nil
			} else {
				if out.MaxUnavailable == nil {
					out.MaxUnavailable = new(intstr.IntOrString)
				}
				*out.MaxUnavailable = intstr.IntOrString(in.String())
			}
		case "minAvailable":
			if in.IsNull() {
				in.Skip()
				out.MinAvailable = nil
			} else {
				if out.MinAvailable == nil {
					out.MinAvailable = new(intstr.IntOrString)
				}
				*out.MinAvailable = intstr.IntOrString(in.String())
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
func easyjson817bc1ebEncodeGithubComKubewardenK8sObjectsApiPolicyV1beta11(out *jwriter.Writer, in PodDisruptionBudgetSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MaxUnavailable != nil {
		const prefix string = ",\"maxUnavailable\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.MaxUnavailable))
	}
	if in.MinAvailable != nil {
		const prefix string = ",\"minAvailable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.MinAvailable))
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
	out.RawByte('}')
}
