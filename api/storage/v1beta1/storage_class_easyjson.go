// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
	_v11 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
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

func easyjsonEd5e2e24DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(in *jlexer.Lexer, out *StorageClass) {
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
		case "allowVolumeExpansion":
			out.AllowVolumeExpansion = bool(in.Bool())
		case "allowedTopologies":
			if in.IsNull() {
				in.Skip()
				out.AllowedTopologies = nil
			} else {
				in.Delim('[')
				if out.AllowedTopologies == nil {
					if !in.IsDelim(']') {
						out.AllowedTopologies = make([]*_v1.TopologySelectorTerm, 0, 8)
					} else {
						out.AllowedTopologies = []*_v1.TopologySelectorTerm{}
					}
				} else {
					out.AllowedTopologies = (out.AllowedTopologies)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *_v1.TopologySelectorTerm
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(_v1.TopologySelectorTerm)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.AllowedTopologies = append(out.AllowedTopologies, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
					out.Metadata = new(_v11.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "mountOptions":
			if in.IsNull() {
				in.Skip()
				out.MountOptions = nil
			} else {
				in.Delim('[')
				if out.MountOptions == nil {
					if !in.IsDelim(']') {
						out.MountOptions = make([]string, 0, 4)
					} else {
						out.MountOptions = []string{}
					}
				} else {
					out.MountOptions = (out.MountOptions)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.MountOptions = append(out.MountOptions, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "parameters":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Parameters = make(map[string]string)
				} else {
					out.Parameters = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 string
					v3 = string(in.String())
					(out.Parameters)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "provisioner":
			if in.IsNull() {
				in.Skip()
				out.Provisioner = nil
			} else {
				if out.Provisioner == nil {
					out.Provisioner = new(string)
				}
				*out.Provisioner = string(in.String())
			}
		case "reclaimPolicy":
			out.ReclaimPolicy = string(in.String())
		case "volumeBindingMode":
			out.VolumeBindingMode = string(in.String())
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
func easyjsonEd5e2e24EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(out *jwriter.Writer, in StorageClass) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllowVolumeExpansion {
		const prefix string = ",\"allowVolumeExpansion\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.AllowVolumeExpansion))
	}
	if len(in.AllowedTopologies) != 0 {
		const prefix string = ",\"allowedTopologies\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v4, v5 := range in.AllowedTopologies {
				if v4 > 0 {
					out.RawByte(',')
				}
				if v5 == nil {
					out.RawString("null")
				} else {
					(*v5).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
	if len(in.MountOptions) != 0 {
		const prefix string = ",\"mountOptions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.MountOptions {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.Parameters) != 0 {
		const prefix string = ",\"parameters\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v8First := true
			for v8Name, v8Value := range in.Parameters {
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
		const prefix string = ",\"provisioner\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Provisioner == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Provisioner))
		}
	}
	if in.ReclaimPolicy != "" {
		const prefix string = ",\"reclaimPolicy\":"
		out.RawString(prefix)
		out.String(string(in.ReclaimPolicy))
	}
	if in.VolumeBindingMode != "" {
		const prefix string = ",\"volumeBindingMode\":"
		out.RawString(prefix)
		out.String(string(in.VolumeBindingMode))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StorageClass) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd5e2e24EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StorageClass) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd5e2e24EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StorageClass) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd5e2e24DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StorageClass) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd5e2e24DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(l, v)
}
