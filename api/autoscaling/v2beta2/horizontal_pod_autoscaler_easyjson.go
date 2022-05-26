// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2beta2

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
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

func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta2(in *jlexer.Lexer, out *HorizontalPodAutoscaler) {
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
			(out.Metadata).UnmarshalEasyJSON(in)
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(HorizontalPodAutoscalerSpec)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta21(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(HorizontalPodAutoscalerStatus)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta22(in, out.Status)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta2(out *jwriter.Writer, in HorizontalPodAutoscaler) {
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
	if true {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta21(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta22(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HorizontalPodAutoscaler) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HorizontalPodAutoscaler) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HorizontalPodAutoscaler) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HorizontalPodAutoscaler) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta2(l, v)
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta22(in *jlexer.Lexer, out *HorizontalPodAutoscalerStatus) {
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
						out.Conditions = make([]*HorizontalPodAutoscalerCondition, 0, 8)
					} else {
						out.Conditions = []*HorizontalPodAutoscalerCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *HorizontalPodAutoscalerCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(HorizontalPodAutoscalerCondition)
						}
						easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta23(in, v1)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "currentMetrics":
			if in.IsNull() {
				in.Skip()
				out.CurrentMetrics = nil
			} else {
				in.Delim('[')
				if out.CurrentMetrics == nil {
					if !in.IsDelim(']') {
						out.CurrentMetrics = make([]*MetricStatus, 0, 8)
					} else {
						out.CurrentMetrics = []*MetricStatus{}
					}
				} else {
					out.CurrentMetrics = (out.CurrentMetrics)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *MetricStatus
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(MetricStatus)
						}
						easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta24(in, v2)
					}
					out.CurrentMetrics = append(out.CurrentMetrics, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "currentReplicas":
			if in.IsNull() {
				in.Skip()
				out.CurrentReplicas = nil
			} else {
				if out.CurrentReplicas == nil {
					out.CurrentReplicas = new(int32)
				}
				*out.CurrentReplicas = int32(in.Int32())
			}
		case "desiredReplicas":
			if in.IsNull() {
				in.Skip()
				out.DesiredReplicas = nil
			} else {
				if out.DesiredReplicas == nil {
					out.DesiredReplicas = new(int32)
				}
				*out.DesiredReplicas = int32(in.Int32())
			}
		case "lastScaleTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastScaleTime).UnmarshalJSON(data))
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta22(out *jwriter.Writer, in HorizontalPodAutoscalerStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"conditions\":"
		out.RawString(prefix[1:])
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Conditions {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta23(out, *v4)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"currentMetrics\":"
		out.RawString(prefix)
		if in.CurrentMetrics == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.CurrentMetrics {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta24(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"currentReplicas\":"
		out.RawString(prefix)
		if in.CurrentReplicas == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.CurrentReplicas))
		}
	}
	{
		const prefix string = ",\"desiredReplicas\":"
		out.RawString(prefix)
		if in.DesiredReplicas == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.DesiredReplicas))
		}
	}
	if true {
		const prefix string = ",\"lastScaleTime\":"
		out.RawString(prefix)
		out.Raw((in.LastScaleTime).MarshalJSON())
	}
	if in.ObservedGeneration != 0 {
		const prefix string = ",\"observedGeneration\":"
		out.RawString(prefix)
		out.Int64(int64(in.ObservedGeneration))
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta24(in *jlexer.Lexer, out *MetricStatus) {
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
		case "containerResource":
			if in.IsNull() {
				in.Skip()
				out.ContainerResource = nil
			} else {
				if out.ContainerResource == nil {
					out.ContainerResource = new(ContainerResourceMetricStatus)
				}
				(*out.ContainerResource).UnmarshalEasyJSON(in)
			}
		case "external":
			if in.IsNull() {
				in.Skip()
				out.External = nil
			} else {
				if out.External == nil {
					out.External = new(ExternalMetricStatus)
				}
				(*out.External).UnmarshalEasyJSON(in)
			}
		case "object":
			if in.IsNull() {
				in.Skip()
				out.Object = nil
			} else {
				if out.Object == nil {
					out.Object = new(ObjectMetricStatus)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta25(in, out.Object)
			}
		case "pods":
			if in.IsNull() {
				in.Skip()
				out.Pods = nil
			} else {
				if out.Pods == nil {
					out.Pods = new(PodsMetricStatus)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta26(in, out.Pods)
			}
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(ResourceMetricStatus)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta27(in, out.Resource)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta24(out *jwriter.Writer, in MetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ContainerResource != nil {
		const prefix string = ",\"containerResource\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ContainerResource).MarshalEasyJSON(out)
	}
	if in.External != nil {
		const prefix string = ",\"external\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.External).MarshalEasyJSON(out)
	}
	if in.Object != nil {
		const prefix string = ",\"object\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta25(out, *in.Object)
	}
	if in.Pods != nil {
		const prefix string = ",\"pods\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta26(out, *in.Pods)
	}
	if in.Resource != nil {
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta27(out, *in.Resource)
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta27(in *jlexer.Lexer, out *ResourceMetricStatus) {
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
		case "current":
			if in.IsNull() {
				in.Skip()
				out.Current = nil
			} else {
				if out.Current == nil {
					out.Current = new(MetricValueStatus)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta28(in, out.Current)
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta27(out *jwriter.Writer, in ResourceMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"current\":"
		out.RawString(prefix[1:])
		if in.Current == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta28(out, *in.Current)
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta28(in *jlexer.Lexer, out *MetricValueStatus) {
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
		case "averageUtilization":
			out.AverageUtilization = int32(in.Int32())
		case "averageValue":
			out.AverageValue = resource.Quantity(in.String())
		case "value":
			out.Value = resource.Quantity(in.String())
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta28(out *jwriter.Writer, in MetricValueStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AverageUtilization != 0 {
		const prefix string = ",\"averageUtilization\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AverageUtilization))
	}
	if in.AverageValue != "" {
		const prefix string = ",\"averageValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AverageValue))
	}
	if in.Value != "" {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta26(in *jlexer.Lexer, out *PodsMetricStatus) {
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
		case "current":
			if in.IsNull() {
				in.Skip()
				out.Current = nil
			} else {
				if out.Current == nil {
					out.Current = new(MetricValueStatus)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta28(in, out.Current)
			}
		case "metric":
			if in.IsNull() {
				in.Skip()
				out.Metric = nil
			} else {
				if out.Metric == nil {
					out.Metric = new(MetricIdentifier)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(in, out.Metric)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta26(out *jwriter.Writer, in PodsMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"current\":"
		out.RawString(prefix[1:])
		if in.Current == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta28(out, *in.Current)
		}
	}
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix)
		if in.Metric == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(out, *in.Metric)
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(in *jlexer.Lexer, out *MetricIdentifier) {
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
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "selector":
			(out.Selector).UnmarshalEasyJSON(in)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(out *jwriter.Writer, in MetricIdentifier) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	if true {
		const prefix string = ",\"selector\":"
		out.RawString(prefix)
		(in.Selector).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta25(in *jlexer.Lexer, out *ObjectMetricStatus) {
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
		case "current":
			if in.IsNull() {
				in.Skip()
				out.Current = nil
			} else {
				if out.Current == nil {
					out.Current = new(MetricValueStatus)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta28(in, out.Current)
			}
		case "describedObject":
			if in.IsNull() {
				in.Skip()
				out.DescribedObject = nil
			} else {
				if out.DescribedObject == nil {
					out.DescribedObject = new(CrossVersionObjectReference)
				}
				(*out.DescribedObject).UnmarshalEasyJSON(in)
			}
		case "metric":
			if in.IsNull() {
				in.Skip()
				out.Metric = nil
			} else {
				if out.Metric == nil {
					out.Metric = new(MetricIdentifier)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(in, out.Metric)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta25(out *jwriter.Writer, in ObjectMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"current\":"
		out.RawString(prefix[1:])
		if in.Current == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta28(out, *in.Current)
		}
	}
	{
		const prefix string = ",\"describedObject\":"
		out.RawString(prefix)
		if in.DescribedObject == nil {
			out.RawString("null")
		} else {
			(*in.DescribedObject).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix)
		if in.Metric == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(out, *in.Metric)
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta23(in *jlexer.Lexer, out *HorizontalPodAutoscalerCondition) {
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
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastTransitionTime).UnmarshalJSON(data))
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta23(out *jwriter.Writer, in HorizontalPodAutoscalerCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.LastTransitionTime).MarshalJSON())
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
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta21(in *jlexer.Lexer, out *HorizontalPodAutoscalerSpec) {
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
		case "behavior":
			if in.IsNull() {
				in.Skip()
				out.Behavior = nil
			} else {
				if out.Behavior == nil {
					out.Behavior = new(HorizontalPodAutoscalerBehavior)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta210(in, out.Behavior)
			}
		case "maxReplicas":
			if in.IsNull() {
				in.Skip()
				out.MaxReplicas = nil
			} else {
				if out.MaxReplicas == nil {
					out.MaxReplicas = new(int32)
				}
				*out.MaxReplicas = int32(in.Int32())
			}
		case "metrics":
			if in.IsNull() {
				in.Skip()
				out.Metrics = nil
			} else {
				in.Delim('[')
				if out.Metrics == nil {
					if !in.IsDelim(']') {
						out.Metrics = make([]*MetricSpec, 0, 8)
					} else {
						out.Metrics = []*MetricSpec{}
					}
				} else {
					out.Metrics = (out.Metrics)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *MetricSpec
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(MetricSpec)
						}
						easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta211(in, v7)
					}
					out.Metrics = append(out.Metrics, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "minReplicas":
			out.MinReplicas = int32(in.Int32())
		case "scaleTargetRef":
			if in.IsNull() {
				in.Skip()
				out.ScaleTargetRef = nil
			} else {
				if out.ScaleTargetRef == nil {
					out.ScaleTargetRef = new(CrossVersionObjectReference)
				}
				(*out.ScaleTargetRef).UnmarshalEasyJSON(in)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta21(out *jwriter.Writer, in HorizontalPodAutoscalerSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Behavior != nil {
		const prefix string = ",\"behavior\":"
		first = false
		out.RawString(prefix[1:])
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta210(out, *in.Behavior)
	}
	{
		const prefix string = ",\"maxReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MaxReplicas == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.MaxReplicas))
		}
	}
	{
		const prefix string = ",\"metrics\":"
		out.RawString(prefix)
		if in.Metrics == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Metrics {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta211(out, *v9)
				}
			}
			out.RawByte(']')
		}
	}
	if in.MinReplicas != 0 {
		const prefix string = ",\"minReplicas\":"
		out.RawString(prefix)
		out.Int32(int32(in.MinReplicas))
	}
	{
		const prefix string = ",\"scaleTargetRef\":"
		out.RawString(prefix)
		if in.ScaleTargetRef == nil {
			out.RawString("null")
		} else {
			(*in.ScaleTargetRef).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta211(in *jlexer.Lexer, out *MetricSpec) {
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
		case "containerResource":
			if in.IsNull() {
				in.Skip()
				out.ContainerResource = nil
			} else {
				if out.ContainerResource == nil {
					out.ContainerResource = new(ContainerResourceMetricSource)
				}
				(*out.ContainerResource).UnmarshalEasyJSON(in)
			}
		case "external":
			if in.IsNull() {
				in.Skip()
				out.External = nil
			} else {
				if out.External == nil {
					out.External = new(ExternalMetricSource)
				}
				(*out.External).UnmarshalEasyJSON(in)
			}
		case "object":
			if in.IsNull() {
				in.Skip()
				out.Object = nil
			} else {
				if out.Object == nil {
					out.Object = new(ObjectMetricSource)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta212(in, out.Object)
			}
		case "pods":
			if in.IsNull() {
				in.Skip()
				out.Pods = nil
			} else {
				if out.Pods == nil {
					out.Pods = new(PodsMetricSource)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta213(in, out.Pods)
			}
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(ResourceMetricSource)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta214(in, out.Resource)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta211(out *jwriter.Writer, in MetricSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ContainerResource != nil {
		const prefix string = ",\"containerResource\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ContainerResource).MarshalEasyJSON(out)
	}
	if in.External != nil {
		const prefix string = ",\"external\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.External).MarshalEasyJSON(out)
	}
	if in.Object != nil {
		const prefix string = ",\"object\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta212(out, *in.Object)
	}
	if in.Pods != nil {
		const prefix string = ",\"pods\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta213(out, *in.Pods)
	}
	if in.Resource != nil {
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta214(out, *in.Resource)
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta214(in *jlexer.Lexer, out *ResourceMetricSource) {
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
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "target":
			if in.IsNull() {
				in.Skip()
				out.Target = nil
			} else {
				if out.Target == nil {
					out.Target = new(MetricTarget)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta215(in, out.Target)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta214(out *jwriter.Writer, in ResourceMetricSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"target\":"
		out.RawString(prefix)
		if in.Target == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta215(out, *in.Target)
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta215(in *jlexer.Lexer, out *MetricTarget) {
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
		case "averageUtilization":
			out.AverageUtilization = int32(in.Int32())
		case "averageValue":
			out.AverageValue = resource.Quantity(in.String())
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
		case "value":
			out.Value = resource.Quantity(in.String())
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta215(out *jwriter.Writer, in MetricTarget) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AverageUtilization != 0 {
		const prefix string = ",\"averageUtilization\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AverageUtilization))
	}
	if in.AverageValue != "" {
		const prefix string = ",\"averageValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AverageValue))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	if in.Value != "" {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta213(in *jlexer.Lexer, out *PodsMetricSource) {
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
		case "metric":
			if in.IsNull() {
				in.Skip()
				out.Metric = nil
			} else {
				if out.Metric == nil {
					out.Metric = new(MetricIdentifier)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(in, out.Metric)
			}
		case "target":
			if in.IsNull() {
				in.Skip()
				out.Target = nil
			} else {
				if out.Target == nil {
					out.Target = new(MetricTarget)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta215(in, out.Target)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta213(out *jwriter.Writer, in PodsMetricSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix[1:])
		if in.Metric == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(out, *in.Metric)
		}
	}
	{
		const prefix string = ",\"target\":"
		out.RawString(prefix)
		if in.Target == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta215(out, *in.Target)
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta212(in *jlexer.Lexer, out *ObjectMetricSource) {
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
		case "describedObject":
			if in.IsNull() {
				in.Skip()
				out.DescribedObject = nil
			} else {
				if out.DescribedObject == nil {
					out.DescribedObject = new(CrossVersionObjectReference)
				}
				(*out.DescribedObject).UnmarshalEasyJSON(in)
			}
		case "metric":
			if in.IsNull() {
				in.Skip()
				out.Metric = nil
			} else {
				if out.Metric == nil {
					out.Metric = new(MetricIdentifier)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(in, out.Metric)
			}
		case "target":
			if in.IsNull() {
				in.Skip()
				out.Target = nil
			} else {
				if out.Target == nil {
					out.Target = new(MetricTarget)
				}
				easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta215(in, out.Target)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta212(out *jwriter.Writer, in ObjectMetricSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"describedObject\":"
		out.RawString(prefix[1:])
		if in.DescribedObject == nil {
			out.RawString("null")
		} else {
			(*in.DescribedObject).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix)
		if in.Metric == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta29(out, *in.Metric)
		}
	}
	{
		const prefix string = ",\"target\":"
		out.RawString(prefix)
		if in.Target == nil {
			out.RawString("null")
		} else {
			easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta215(out, *in.Target)
		}
	}
	out.RawByte('}')
}
func easyjson23e3e9dcDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta210(in *jlexer.Lexer, out *HorizontalPodAutoscalerBehavior) {
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
		case "scaleDown":
			if in.IsNull() {
				in.Skip()
				out.ScaleDown = nil
			} else {
				if out.ScaleDown == nil {
					out.ScaleDown = new(HPAScalingRules)
				}
				(*out.ScaleDown).UnmarshalEasyJSON(in)
			}
		case "scaleUp":
			if in.IsNull() {
				in.Skip()
				out.ScaleUp = nil
			} else {
				if out.ScaleUp == nil {
					out.ScaleUp = new(HPAScalingRules)
				}
				(*out.ScaleUp).UnmarshalEasyJSON(in)
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
func easyjson23e3e9dcEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta210(out *jwriter.Writer, in HorizontalPodAutoscalerBehavior) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ScaleDown != nil {
		const prefix string = ",\"scaleDown\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ScaleDown).MarshalEasyJSON(out)
	}
	if in.ScaleUp != nil {
		const prefix string = ",\"scaleUp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ScaleUp).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
