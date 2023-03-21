// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjson4311d3a4DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(in *jlexer.Lexer, out *CSIStorageCapacity) {
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
		case "capacity":
			if in.IsNull() {
				in.Skip()
				out.Capacity = nil
			} else {
				if out.Capacity == nil {
					out.Capacity = new(resource.Quantity)
				}
				*out.Capacity = resource.Quantity(in.String())
			}
		case "kind":
			out.Kind = string(in.String())
		case "maximumVolumeSize":
			if in.IsNull() {
				in.Skip()
				out.MaximumVolumeSize = nil
			} else {
				if out.MaximumVolumeSize == nil {
					out.MaximumVolumeSize = new(resource.Quantity)
				}
				*out.MaximumVolumeSize = resource.Quantity(in.String())
			}
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
		case "nodeTopology":
			if in.IsNull() {
				in.Skip()
				out.NodeTopology = nil
			} else {
				if out.NodeTopology == nil {
					out.NodeTopology = new(_v1.LabelSelector)
				}
				(*out.NodeTopology).UnmarshalEasyJSON(in)
			}
		case "storageClassName":
			if in.IsNull() {
				in.Skip()
				out.StorageClassName = nil
			} else {
				if out.StorageClassName == nil {
					out.StorageClassName = new(string)
				}
				*out.StorageClassName = string(in.String())
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
func easyjson4311d3a4EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(out *jwriter.Writer, in CSIStorageCapacity) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	if in.Capacity != nil {
		const prefix string = ",\"capacity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Capacity))
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
	if in.MaximumVolumeSize != nil {
		const prefix string = ",\"maximumVolumeSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.MaximumVolumeSize))
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
	if in.NodeTopology != nil {
		const prefix string = ",\"nodeTopology\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.NodeTopology).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"storageClassName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageClassName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.StorageClassName))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSIStorageCapacity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4311d3a4EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSIStorageCapacity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4311d3a4EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSIStorageCapacity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4311d3a4DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSIStorageCapacity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4311d3a4DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(l, v)
}
