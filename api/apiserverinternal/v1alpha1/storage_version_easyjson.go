// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"
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

func easyjsonEacd27c4DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(in *jlexer.Lexer, out *StorageVersion) {
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
					out.Spec = new(easyjson.RawMessage)
				}
				(*out.Spec).UnmarshalEasyJSON(in)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(StorageVersionStatus)
				}
				easyjsonEacd27c4DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha11(in, out.Status)
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
func easyjsonEacd27c4EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(out *jwriter.Writer, in StorageVersion) {
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
	{
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Spec == nil {
			out.RawString("null")
		} else {
			(*in.Spec).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		if in.Status == nil {
			out.RawString("null")
		} else {
			easyjsonEacd27c4EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha11(out, *in.Status)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StorageVersion) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEacd27c4EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StorageVersion) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEacd27c4EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StorageVersion) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEacd27c4DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StorageVersion) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEacd27c4DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(l, v)
}
func easyjsonEacd27c4DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha11(in *jlexer.Lexer, out *StorageVersionStatus) {
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
		case "commonEncodingVersion":
			out.CommonEncodingVersion = string(in.String())
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*StorageVersionCondition, 0, 8)
					} else {
						out.Conditions = []*StorageVersionCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *StorageVersionCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(StorageVersionCondition)
						}
						easyjsonEacd27c4DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha12(in, v1)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "storageVersions":
			if in.IsNull() {
				in.Skip()
				out.StorageVersions = nil
			} else {
				in.Delim('[')
				if out.StorageVersions == nil {
					if !in.IsDelim(']') {
						out.StorageVersions = make([]*ServerStorageVersion, 0, 8)
					} else {
						out.StorageVersions = []*ServerStorageVersion{}
					}
				} else {
					out.StorageVersions = (out.StorageVersions)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *ServerStorageVersion
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(ServerStorageVersion)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.StorageVersions = append(out.StorageVersions, v2)
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
func easyjsonEacd27c4EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha11(out *jwriter.Writer, in StorageVersionStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CommonEncodingVersion != "" {
		const prefix string = ",\"commonEncodingVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.CommonEncodingVersion))
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
			for v3, v4 := range in.Conditions {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					easyjsonEacd27c4EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha12(out, *v4)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.StorageVersions) != 0 {
		const prefix string = ",\"storageVersions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.StorageVersions {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonEacd27c4DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha12(in *jlexer.Lexer, out *StorageVersionCondition) {
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
		case "message":
			out.Message = string(in.String())
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
		case "reason":
			if in.IsNull() {
				in.Skip()
				out.Reason = nil
			} else {
				if out.Reason == nil {
					out.Reason = new(string)
				}
				*out.Reason = string(in.String())
			}
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
func easyjsonEacd27c4EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha12(out *jwriter.Writer, in StorageVersionCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LastTransitionTime != nil {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
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
	{
		const prefix string = ",\"reason\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Reason == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Reason))
		}
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
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