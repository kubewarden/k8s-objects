// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *PersistentVolumeClaim) {
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
					out.Spec = new(PersistentVolumeClaimSpec)
				}
				easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(PersistentVolumeClaimStatus)
				}
				easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.Status)
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
func easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in PersistentVolumeClaim) {
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
		easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PersistentVolumeClaim) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PersistentVolumeClaim) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PersistentVolumeClaim) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PersistentVolumeClaim) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *PersistentVolumeClaimStatus) {
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
		case "accessModes":
			if in.IsNull() {
				in.Skip()
				out.AccessModes = nil
			} else {
				in.Delim('[')
				if out.AccessModes == nil {
					if !in.IsDelim(']') {
						out.AccessModes = make([]string, 0, 4)
					} else {
						out.AccessModes = []string{}
					}
				} else {
					out.AccessModes = (out.AccessModes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.AccessModes = append(out.AccessModes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "allocatedResources":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.AllocatedResources = make(map[string]resource.Quantity)
				} else {
					out.AllocatedResources = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 resource.Quantity
					v2 = resource.Quantity(in.String())
					(out.AllocatedResources)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "capacity":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Capacity = make(map[string]resource.Quantity)
				} else {
					out.Capacity = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 resource.Quantity
					v3 = resource.Quantity(in.String())
					(out.Capacity)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*PersistentVolumeClaimCondition, 0, 8)
					} else {
						out.Conditions = []*PersistentVolumeClaimCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *PersistentVolumeClaimCondition
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(PersistentVolumeClaimCondition)
						}
						easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV13(in, v4)
					}
					out.Conditions = append(out.Conditions, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "phase":
			out.Phase = string(in.String())
		case "resizeStatus":
			out.ResizeStatus = string(in.String())
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
func easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in PersistentVolumeClaimStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"accessModes\":"
		out.RawString(prefix[1:])
		if in.AccessModes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.AccessModes {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if len(in.AllocatedResources) != 0 {
		const prefix string = ",\"allocatedResources\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v7First := true
			for v7Name, v7Value := range in.AllocatedResources {
				if v7First {
					v7First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v7Name))
				out.RawByte(':')
				out.String(string(v7Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Capacity) != 0 {
		const prefix string = ",\"capacity\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v8First := true
			for v8Name, v8Value := range in.Capacity {
				if v8First {
					v8First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v8Name))
				out.RawByte(':')
				out.String(string(v8Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"conditions\":"
		out.RawString(prefix)
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Conditions {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *v10)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Phase != "" {
		const prefix string = ",\"phase\":"
		out.RawString(prefix)
		out.String(string(in.Phase))
	}
	if in.ResizeStatus != "" {
		const prefix string = ",\"resizeStatus\":"
		out.RawString(prefix)
		out.String(string(in.ResizeStatus))
	}
	out.RawByte('}')
}
func easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV13(in *jlexer.Lexer, out *PersistentVolumeClaimCondition) {
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
		case "lastProbeTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastProbeTime).UnmarshalJSON(data))
			}
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
func easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV13(out *jwriter.Writer, in PersistentVolumeClaimCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"lastProbeTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.LastProbeTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"lastTransitionTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
func easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *PersistentVolumeClaimSpec) {
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
		case "accessModes":
			if in.IsNull() {
				in.Skip()
				out.AccessModes = nil
			} else {
				in.Delim('[')
				if out.AccessModes == nil {
					if !in.IsDelim(']') {
						out.AccessModes = make([]string, 0, 4)
					} else {
						out.AccessModes = []string{}
					}
				} else {
					out.AccessModes = (out.AccessModes)[:0]
				}
				for !in.IsDelim(']') {
					var v11 string
					v11 = string(in.String())
					out.AccessModes = append(out.AccessModes, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "dataSource":
			if in.IsNull() {
				in.Skip()
				out.DataSource = nil
			} else {
				if out.DataSource == nil {
					out.DataSource = new(TypedLocalObjectReference)
				}
				easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV14(in, out.DataSource)
			}
		case "dataSourceRef":
			if in.IsNull() {
				in.Skip()
				out.DataSourceRef = nil
			} else {
				if out.DataSourceRef == nil {
					out.DataSourceRef = new(TypedLocalObjectReference)
				}
				easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV14(in, out.DataSourceRef)
			}
		case "resources":
			if in.IsNull() {
				in.Skip()
				out.Resources = nil
			} else {
				if out.Resources == nil {
					out.Resources = new(ResourceRequirements)
				}
				easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV15(in, out.Resources)
			}
		case "selector":
			(out.Selector).UnmarshalEasyJSON(in)
		case "storageClassName":
			out.StorageClassName = string(in.String())
		case "volumeMode":
			out.VolumeMode = string(in.String())
		case "volumeName":
			out.VolumeName = string(in.String())
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
func easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in PersistentVolumeClaimSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"accessModes\":"
		out.RawString(prefix[1:])
		if in.AccessModes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.AccessModes {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	if in.DataSource != nil {
		const prefix string = ",\"dataSource\":"
		out.RawString(prefix)
		easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV14(out, *in.DataSource)
	}
	if in.DataSourceRef != nil {
		const prefix string = ",\"dataSourceRef\":"
		out.RawString(prefix)
		easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV14(out, *in.DataSourceRef)
	}
	if in.Resources != nil {
		const prefix string = ",\"resources\":"
		out.RawString(prefix)
		easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV15(out, *in.Resources)
	}
	if true {
		const prefix string = ",\"selector\":"
		out.RawString(prefix)
		(in.Selector).MarshalEasyJSON(out)
	}
	if in.StorageClassName != "" {
		const prefix string = ",\"storageClassName\":"
		out.RawString(prefix)
		out.String(string(in.StorageClassName))
	}
	if in.VolumeMode != "" {
		const prefix string = ",\"volumeMode\":"
		out.RawString(prefix)
		out.String(string(in.VolumeMode))
	}
	if in.VolumeName != "" {
		const prefix string = ",\"volumeName\":"
		out.RawString(prefix)
		out.String(string(in.VolumeName))
	}
	out.RawByte('}')
}
func easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV15(in *jlexer.Lexer, out *ResourceRequirements) {
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
		case "limits":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Limits = make(map[string]resource.Quantity)
				} else {
					out.Limits = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v14 resource.Quantity
					v14 = resource.Quantity(in.String())
					(out.Limits)[key] = v14
					in.WantComma()
				}
				in.Delim('}')
			}
		case "requests":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Requests = make(map[string]resource.Quantity)
				} else {
					out.Requests = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v15 resource.Quantity
					v15 = resource.Quantity(in.String())
					(out.Requests)[key] = v15
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV15(out *jwriter.Writer, in ResourceRequirements) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Limits) != 0 {
		const prefix string = ",\"limits\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('{')
			v16First := true
			for v16Name, v16Value := range in.Limits {
				if v16First {
					v16First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v16Name))
				out.RawByte(':')
				out.String(string(v16Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Requests) != 0 {
		const prefix string = ",\"requests\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v17First := true
			for v17Name, v17Value := range in.Requests {
				if v17First {
					v17First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v17Name))
				out.RawByte(':')
				out.String(string(v17Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}
func easyjsonF52ef087DecodeGithubComKubewardenK8sObjectsApiCoreV14(in *jlexer.Lexer, out *TypedLocalObjectReference) {
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
		case "apiGroup":
			out.APIGroup = string(in.String())
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
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
func easyjsonF52ef087EncodeGithubComKubewardenK8sObjectsApiCoreV14(out *jwriter.Writer, in TypedLocalObjectReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIGroup != "" {
		const prefix string = ",\"apiGroup\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIGroup))
	}
	{
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Kind))
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
