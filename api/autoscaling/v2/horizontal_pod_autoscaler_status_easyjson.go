// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2

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

func easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(in *jlexer.Lexer, out *HorizontalPodAutoscalerStatus) {
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
						(*v1).UnmarshalEasyJSON(in)
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
						easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV21(in, v2)
					}
					out.CurrentMetrics = append(out.CurrentMetrics, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "currentReplicas":
			out.CurrentReplicas = int32(in.Int32())
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
func easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(out *jwriter.Writer, in HorizontalPodAutoscalerStatus) {
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
					(*v4).MarshalEasyJSON(out)
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
					easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV21(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	if in.CurrentReplicas != 0 {
		const prefix string = ",\"currentReplicas\":"
		out.RawString(prefix)
		out.Int32(int32(in.CurrentReplicas))
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

// MarshalJSON supports json.Marshaler interface
func (v HorizontalPodAutoscalerStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HorizontalPodAutoscalerStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HorizontalPodAutoscalerStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HorizontalPodAutoscalerStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(l, v)
}
func easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV21(in *jlexer.Lexer, out *MetricStatus) {
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
				easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV22(in, out.Object)
			}
		case "pods":
			if in.IsNull() {
				in.Skip()
				out.Pods = nil
			} else {
				if out.Pods == nil {
					out.Pods = new(PodsMetricStatus)
				}
				easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV23(in, out.Pods)
			}
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(ResourceMetricStatus)
				}
				easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV24(in, out.Resource)
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
func easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV21(out *jwriter.Writer, in MetricStatus) {
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
		easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV22(out, *in.Object)
	}
	if in.Pods != nil {
		const prefix string = ",\"pods\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV23(out, *in.Pods)
	}
	if in.Resource != nil {
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV24(out, *in.Resource)
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
func easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV24(in *jlexer.Lexer, out *ResourceMetricStatus) {
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
				easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV25(in, out.Current)
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
func easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV24(out *jwriter.Writer, in ResourceMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"current\":"
		out.RawString(prefix[1:])
		if in.Current == nil {
			out.RawString("null")
		} else {
			easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV25(out, *in.Current)
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
func easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV25(in *jlexer.Lexer, out *MetricValueStatus) {
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
func easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV25(out *jwriter.Writer, in MetricValueStatus) {
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
func easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV23(in *jlexer.Lexer, out *PodsMetricStatus) {
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
				easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV25(in, out.Current)
			}
		case "metric":
			if in.IsNull() {
				in.Skip()
				out.Metric = nil
			} else {
				if out.Metric == nil {
					out.Metric = new(MetricIdentifier)
				}
				easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV26(in, out.Metric)
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
func easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV23(out *jwriter.Writer, in PodsMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"current\":"
		out.RawString(prefix[1:])
		if in.Current == nil {
			out.RawString("null")
		} else {
			easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV25(out, *in.Current)
		}
	}
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix)
		if in.Metric == nil {
			out.RawString("null")
		} else {
			easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV26(out, *in.Metric)
		}
	}
	out.RawByte('}')
}
func easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV26(in *jlexer.Lexer, out *MetricIdentifier) {
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
func easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV26(out *jwriter.Writer, in MetricIdentifier) {
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
func easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV22(in *jlexer.Lexer, out *ObjectMetricStatus) {
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
				easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV25(in, out.Current)
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
				easyjson4e8edfcfDecodeGithubComKubewardenK8sObjectsApiAutoscalingV26(in, out.Metric)
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
func easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV22(out *jwriter.Writer, in ObjectMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"current\":"
		out.RawString(prefix[1:])
		if in.Current == nil {
			out.RawString("null")
		} else {
			easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV25(out, *in.Current)
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
			easyjson4e8edfcfEncodeGithubComKubewardenK8sObjectsApiAutoscalingV26(out, *in.Metric)
		}
	}
	out.RawByte('}')
}
