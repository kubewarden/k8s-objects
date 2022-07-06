// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
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

func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *Node) {
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
					out.Spec = new(NodeSpec)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(NodeStatus)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.Status)
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in Node) {
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
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Node) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Node) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Node) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Node) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *NodeStatus) {
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
		case "addresses":
			if in.IsNull() {
				in.Skip()
				out.Addresses = nil
			} else {
				in.Delim('[')
				if out.Addresses == nil {
					if !in.IsDelim(']') {
						out.Addresses = make([]*NodeAddress, 0, 8)
					} else {
						out.Addresses = []*NodeAddress{}
					}
				} else {
					out.Addresses = (out.Addresses)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *NodeAddress
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(NodeAddress)
						}
						easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV13(in, v1)
					}
					out.Addresses = append(out.Addresses, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "allocatable":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Allocatable = make(map[string]*resource.Quantity)
				} else {
					out.Allocatable = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 *resource.Quantity
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(resource.Quantity)
						}
						*v2 = resource.Quantity(in.String())
					}
					(out.Allocatable)[key] = v2
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
					out.Capacity = make(map[string]*resource.Quantity)
				} else {
					out.Capacity = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 *resource.Quantity
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(resource.Quantity)
						}
						*v3 = resource.Quantity(in.String())
					}
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
						out.Conditions = make([]*NodeCondition, 0, 8)
					} else {
						out.Conditions = []*NodeCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *NodeCondition
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(NodeCondition)
						}
						easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV14(in, v4)
					}
					out.Conditions = append(out.Conditions, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "config":
			if in.IsNull() {
				in.Skip()
				out.Config = nil
			} else {
				if out.Config == nil {
					out.Config = new(NodeConfigStatus)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV15(in, out.Config)
			}
		case "daemonEndpoints":
			if in.IsNull() {
				in.Skip()
				out.DaemonEndpoints = nil
			} else {
				if out.DaemonEndpoints == nil {
					out.DaemonEndpoints = new(NodeDaemonEndpoints)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV16(in, out.DaemonEndpoints)
			}
		case "images":
			if in.IsNull() {
				in.Skip()
				out.Images = nil
			} else {
				in.Delim('[')
				if out.Images == nil {
					if !in.IsDelim(']') {
						out.Images = make([]*ContainerImage, 0, 8)
					} else {
						out.Images = []*ContainerImage{}
					}
				} else {
					out.Images = (out.Images)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *ContainerImage
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(ContainerImage)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.Images = append(out.Images, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "nodeInfo":
			if in.IsNull() {
				in.Skip()
				out.NodeInfo = nil
			} else {
				if out.NodeInfo == nil {
					out.NodeInfo = new(NodeSystemInfo)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV17(in, out.NodeInfo)
			}
		case "phase":
			out.Phase = string(in.String())
		case "volumesAttached":
			if in.IsNull() {
				in.Skip()
				out.VolumesAttached = nil
			} else {
				in.Delim('[')
				if out.VolumesAttached == nil {
					if !in.IsDelim(']') {
						out.VolumesAttached = make([]*AttachedVolume, 0, 8)
					} else {
						out.VolumesAttached = []*AttachedVolume{}
					}
				} else {
					out.VolumesAttached = (out.VolumesAttached)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *AttachedVolume
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(AttachedVolume)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					out.VolumesAttached = append(out.VolumesAttached, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "volumesInUse":
			if in.IsNull() {
				in.Skip()
				out.VolumesInUse = nil
			} else {
				in.Delim('[')
				if out.VolumesInUse == nil {
					if !in.IsDelim(']') {
						out.VolumesInUse = make([]string, 0, 4)
					} else {
						out.VolumesInUse = []string{}
					}
				} else {
					out.VolumesInUse = (out.VolumesInUse)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.VolumesInUse = append(out.VolumesInUse, v7)
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in NodeStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Addresses) != 0 {
		const prefix string = ",\"addresses\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v8, v9 := range in.Addresses {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *v9)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Allocatable) != 0 {
		const prefix string = ",\"allocatable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v10First := true
			for v10Name, v10Value := range in.Allocatable {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				if v10Value == nil {
					out.RawString("null")
				} else {
					out.String(string(*v10Value))
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Capacity) != 0 {
		const prefix string = ",\"capacity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v11First := true
			for v11Name, v11Value := range in.Capacity {
				if v11First {
					v11First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v11Name))
				out.RawByte(':')
				if v11Value == nil {
					out.RawString("null")
				} else {
					out.String(string(*v11Value))
				}
			}
			out.RawByte('}')
		}
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
			for v12, v13 := range in.Conditions {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV14(out, *v13)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Config != nil {
		const prefix string = ",\"config\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV15(out, *in.Config)
	}
	if in.DaemonEndpoints != nil {
		const prefix string = ",\"daemonEndpoints\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV16(out, *in.DaemonEndpoints)
	}
	if len(in.Images) != 0 {
		const prefix string = ",\"images\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.Images {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.NodeInfo != nil {
		const prefix string = ",\"nodeInfo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV17(out, *in.NodeInfo)
	}
	if in.Phase != "" {
		const prefix string = ",\"phase\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Phase))
	}
	if len(in.VolumesAttached) != 0 {
		const prefix string = ",\"volumesAttached\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.VolumesAttached {
				if v16 > 0 {
					out.RawByte(',')
				}
				if v17 == nil {
					out.RawString("null")
				} else {
					(*v17).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.VolumesInUse) != 0 {
		const prefix string = ",\"volumesInUse\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.VolumesInUse {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.String(string(v19))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV17(in *jlexer.Lexer, out *NodeSystemInfo) {
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
		case "architecture":
			if in.IsNull() {
				in.Skip()
				out.Architecture = nil
			} else {
				if out.Architecture == nil {
					out.Architecture = new(string)
				}
				*out.Architecture = string(in.String())
			}
		case "bootID":
			if in.IsNull() {
				in.Skip()
				out.BootID = nil
			} else {
				if out.BootID == nil {
					out.BootID = new(string)
				}
				*out.BootID = string(in.String())
			}
		case "containerRuntimeVersion":
			if in.IsNull() {
				in.Skip()
				out.ContainerRuntimeVersion = nil
			} else {
				if out.ContainerRuntimeVersion == nil {
					out.ContainerRuntimeVersion = new(string)
				}
				*out.ContainerRuntimeVersion = string(in.String())
			}
		case "kernelVersion":
			if in.IsNull() {
				in.Skip()
				out.KernelVersion = nil
			} else {
				if out.KernelVersion == nil {
					out.KernelVersion = new(string)
				}
				*out.KernelVersion = string(in.String())
			}
		case "kubeProxyVersion":
			if in.IsNull() {
				in.Skip()
				out.KubeProxyVersion = nil
			} else {
				if out.KubeProxyVersion == nil {
					out.KubeProxyVersion = new(string)
				}
				*out.KubeProxyVersion = string(in.String())
			}
		case "kubeletVersion":
			if in.IsNull() {
				in.Skip()
				out.KubeletVersion = nil
			} else {
				if out.KubeletVersion == nil {
					out.KubeletVersion = new(string)
				}
				*out.KubeletVersion = string(in.String())
			}
		case "machineID":
			if in.IsNull() {
				in.Skip()
				out.MachineID = nil
			} else {
				if out.MachineID == nil {
					out.MachineID = new(string)
				}
				*out.MachineID = string(in.String())
			}
		case "operatingSystem":
			if in.IsNull() {
				in.Skip()
				out.OperatingSystem = nil
			} else {
				if out.OperatingSystem == nil {
					out.OperatingSystem = new(string)
				}
				*out.OperatingSystem = string(in.String())
			}
		case "osImage":
			if in.IsNull() {
				in.Skip()
				out.OsImage = nil
			} else {
				if out.OsImage == nil {
					out.OsImage = new(string)
				}
				*out.OsImage = string(in.String())
			}
		case "systemUUID":
			if in.IsNull() {
				in.Skip()
				out.SystemUUID = nil
			} else {
				if out.SystemUUID == nil {
					out.SystemUUID = new(string)
				}
				*out.SystemUUID = string(in.String())
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV17(out *jwriter.Writer, in NodeSystemInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"architecture\":"
		out.RawString(prefix[1:])
		if in.Architecture == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Architecture))
		}
	}
	{
		const prefix string = ",\"bootID\":"
		out.RawString(prefix)
		if in.BootID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.BootID))
		}
	}
	{
		const prefix string = ",\"containerRuntimeVersion\":"
		out.RawString(prefix)
		if in.ContainerRuntimeVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ContainerRuntimeVersion))
		}
	}
	{
		const prefix string = ",\"kernelVersion\":"
		out.RawString(prefix)
		if in.KernelVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.KernelVersion))
		}
	}
	{
		const prefix string = ",\"kubeProxyVersion\":"
		out.RawString(prefix)
		if in.KubeProxyVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.KubeProxyVersion))
		}
	}
	{
		const prefix string = ",\"kubeletVersion\":"
		out.RawString(prefix)
		if in.KubeletVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.KubeletVersion))
		}
	}
	{
		const prefix string = ",\"machineID\":"
		out.RawString(prefix)
		if in.MachineID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.MachineID))
		}
	}
	{
		const prefix string = ",\"operatingSystem\":"
		out.RawString(prefix)
		if in.OperatingSystem == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.OperatingSystem))
		}
	}
	{
		const prefix string = ",\"osImage\":"
		out.RawString(prefix)
		if in.OsImage == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.OsImage))
		}
	}
	{
		const prefix string = ",\"systemUUID\":"
		out.RawString(prefix)
		if in.SystemUUID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.SystemUUID))
		}
	}
	out.RawByte('}')
}
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV16(in *jlexer.Lexer, out *NodeDaemonEndpoints) {
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
		case "kubeletEndpoint":
			if in.IsNull() {
				in.Skip()
				out.KubeletEndpoint = nil
			} else {
				if out.KubeletEndpoint == nil {
					out.KubeletEndpoint = new(DaemonEndpoint)
				}
				(*out.KubeletEndpoint).UnmarshalEasyJSON(in)
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV16(out *jwriter.Writer, in NodeDaemonEndpoints) {
	out.RawByte('{')
	first := true
	_ = first
	if in.KubeletEndpoint != nil {
		const prefix string = ",\"kubeletEndpoint\":"
		first = false
		out.RawString(prefix[1:])
		(*in.KubeletEndpoint).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV15(in *jlexer.Lexer, out *NodeConfigStatus) {
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
		case "active":
			if in.IsNull() {
				in.Skip()
				out.Active = nil
			} else {
				if out.Active == nil {
					out.Active = new(NodeConfigSource)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV18(in, out.Active)
			}
		case "assigned":
			if in.IsNull() {
				in.Skip()
				out.Assigned = nil
			} else {
				if out.Assigned == nil {
					out.Assigned = new(NodeConfigSource)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV18(in, out.Assigned)
			}
		case "error":
			out.Error = string(in.String())
		case "lastKnownGood":
			if in.IsNull() {
				in.Skip()
				out.LastKnownGood = nil
			} else {
				if out.LastKnownGood == nil {
					out.LastKnownGood = new(NodeConfigSource)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV18(in, out.LastKnownGood)
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV15(out *jwriter.Writer, in NodeConfigStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Active != nil {
		const prefix string = ",\"active\":"
		first = false
		out.RawString(prefix[1:])
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV18(out, *in.Active)
	}
	if in.Assigned != nil {
		const prefix string = ",\"assigned\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV18(out, *in.Assigned)
	}
	if in.Error != "" {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error))
	}
	if in.LastKnownGood != nil {
		const prefix string = ",\"lastKnownGood\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV18(out, *in.LastKnownGood)
	}
	out.RawByte('}')
}
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV18(in *jlexer.Lexer, out *NodeConfigSource) {
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
		case "configMap":
			if in.IsNull() {
				in.Skip()
				out.ConfigMap = nil
			} else {
				if out.ConfigMap == nil {
					out.ConfigMap = new(ConfigMapNodeConfigSource)
				}
				(*out.ConfigMap).UnmarshalEasyJSON(in)
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV18(out *jwriter.Writer, in NodeConfigSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConfigMap != nil {
		const prefix string = ",\"configMap\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ConfigMap).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV14(in *jlexer.Lexer, out *NodeCondition) {
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
		case "lastHeartbeatTime":
			if in.IsNull() {
				in.Skip()
				out.LastHeartbeatTime = nil
			} else {
				if out.LastHeartbeatTime == nil {
					out.LastHeartbeatTime = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastHeartbeatTime).UnmarshalJSON(data))
				}
			}
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV14(out *jwriter.Writer, in NodeCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastHeartbeatTime != nil {
		const prefix string = ",\"lastHeartbeatTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.LastHeartbeatTime).MarshalJSON())
	}
	if in.LastTransitionTime != nil {
		const prefix string = ",\"lastTransitionTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV13(in *jlexer.Lexer, out *NodeAddress) {
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
		case "address":
			if in.IsNull() {
				in.Skip()
				out.Address = nil
			} else {
				if out.Address == nil {
					out.Address = new(string)
				}
				*out.Address = string(in.String())
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV13(out *jwriter.Writer, in NodeAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix[1:])
		if in.Address == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Address))
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
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *NodeSpec) {
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
		case "configSource":
			if in.IsNull() {
				in.Skip()
				out.ConfigSource = nil
			} else {
				if out.ConfigSource == nil {
					out.ConfigSource = new(NodeConfigSource)
				}
				easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV18(in, out.ConfigSource)
			}
		case "externalID":
			out.ExternalID = string(in.String())
		case "podCIDR":
			out.PodCIDR = string(in.String())
		case "providerID":
			out.ProviderID = string(in.String())
		case "taints":
			if in.IsNull() {
				in.Skip()
				out.Taints = nil
			} else {
				in.Delim('[')
				if out.Taints == nil {
					if !in.IsDelim(']') {
						out.Taints = make([]*Taint, 0, 8)
					} else {
						out.Taints = []*Taint{}
					}
				} else {
					out.Taints = (out.Taints)[:0]
				}
				for !in.IsDelim(']') {
					var v20 *Taint
					if in.IsNull() {
						in.Skip()
						v20 = nil
					} else {
						if v20 == nil {
							v20 = new(Taint)
						}
						easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV19(in, v20)
					}
					out.Taints = append(out.Taints, v20)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "unschedulable":
			out.Unschedulable = bool(in.Bool())
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in NodeSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConfigSource != nil {
		const prefix string = ",\"configSource\":"
		first = false
		out.RawString(prefix[1:])
		easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV18(out, *in.ConfigSource)
	}
	if in.ExternalID != "" {
		const prefix string = ",\"externalID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExternalID))
	}
	if in.PodCIDR != "" {
		const prefix string = ",\"podCIDR\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PodCIDR))
	}
	if in.ProviderID != "" {
		const prefix string = ",\"providerID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ProviderID))
	}
	if len(in.Taints) != 0 {
		const prefix string = ",\"taints\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v21, v22 := range in.Taints {
				if v21 > 0 {
					out.RawByte(',')
				}
				if v22 == nil {
					out.RawString("null")
				} else {
					easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV19(out, *v22)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Unschedulable {
		const prefix string = ",\"unschedulable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Unschedulable))
	}
	out.RawByte('}')
}
func easyjsonCdfae1c8DecodeGithubComKubewardenK8sObjectsApiCoreV19(in *jlexer.Lexer, out *Taint) {
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
		case "effect":
			if in.IsNull() {
				in.Skip()
				out.Effect = nil
			} else {
				if out.Effect == nil {
					out.Effect = new(string)
				}
				*out.Effect = string(in.String())
			}
		case "key":
			if in.IsNull() {
				in.Skip()
				out.Key = nil
			} else {
				if out.Key == nil {
					out.Key = new(string)
				}
				*out.Key = string(in.String())
			}
		case "timeAdded":
			if in.IsNull() {
				in.Skip()
				out.TimeAdded = nil
			} else {
				if out.TimeAdded == nil {
					out.TimeAdded = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.TimeAdded).UnmarshalJSON(data))
				}
			}
		case "value":
			out.Value = string(in.String())
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
func easyjsonCdfae1c8EncodeGithubComKubewardenK8sObjectsApiCoreV19(out *jwriter.Writer, in Taint) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"effect\":"
		out.RawString(prefix[1:])
		if in.Effect == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Effect))
		}
	}
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix)
		if in.Key == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Key))
		}
	}
	if in.TimeAdded != nil {
		const prefix string = ",\"timeAdded\":"
		out.RawString(prefix)
		out.Raw((*in.TimeAdded).MarshalJSON())
	}
	if in.Value != "" {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	out.RawByte('}')
}
