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

func easyjson94b78b83DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *NodeStatus) {
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
						(*v1).UnmarshalEasyJSON(in)
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
						(*v4).UnmarshalEasyJSON(in)
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
				(*out.Config).UnmarshalEasyJSON(in)
			}
		case "daemonEndpoints":
			if in.IsNull() {
				in.Skip()
				out.DaemonEndpoints = nil
			} else {
				if out.DaemonEndpoints == nil {
					out.DaemonEndpoints = new(NodeDaemonEndpoints)
				}
				(*out.DaemonEndpoints).UnmarshalEasyJSON(in)
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
				easyjson94b78b83DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.NodeInfo)
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
func easyjson94b78b83EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in NodeStatus) {
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
					(*v9).MarshalEasyJSON(out)
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
					(*v13).MarshalEasyJSON(out)
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
		(*in.Config).MarshalEasyJSON(out)
	}
	if in.DaemonEndpoints != nil {
		const prefix string = ",\"daemonEndpoints\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.DaemonEndpoints).MarshalEasyJSON(out)
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
		easyjson94b78b83EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.NodeInfo)
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

// MarshalJSON supports json.Marshaler interface
func (v NodeStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson94b78b83EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NodeStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson94b78b83EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NodeStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson94b78b83DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NodeStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson94b78b83DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson94b78b83DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *NodeSystemInfo) {
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
func easyjson94b78b83EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in NodeSystemInfo) {
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
