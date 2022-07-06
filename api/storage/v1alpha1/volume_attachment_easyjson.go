// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"
	_v11 "github.com/kubewarden/k8s-objects/api/core/v1"
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

func easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(in *jlexer.Lexer, out *VolumeAttachment) {
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
					out.Spec = new(VolumeAttachmentSpec)
				}
				easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(VolumeAttachmentStatus)
				}
				easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha12(in, out.Status)
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
func easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(out *jwriter.Writer, in VolumeAttachment) {
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
			easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha11(out, *in.Spec)
		}
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VolumeAttachment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VolumeAttachment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VolumeAttachment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VolumeAttachment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha1(l, v)
}
func easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha12(in *jlexer.Lexer, out *VolumeAttachmentStatus) {
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
		case "attachError":
			if in.IsNull() {
				in.Skip()
				out.AttachError = nil
			} else {
				if out.AttachError == nil {
					out.AttachError = new(VolumeError)
				}
				easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha13(in, out.AttachError)
			}
		case "attached":
			if in.IsNull() {
				in.Skip()
				out.Attached = nil
			} else {
				if out.Attached == nil {
					out.Attached = new(bool)
				}
				*out.Attached = bool(in.Bool())
			}
		case "attachmentMetadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.AttachmentMetadata = make(map[string]string)
				} else {
					out.AttachmentMetadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.AttachmentMetadata)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "detachError":
			if in.IsNull() {
				in.Skip()
				out.DetachError = nil
			} else {
				if out.DetachError == nil {
					out.DetachError = new(VolumeError)
				}
				easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha13(in, out.DetachError)
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
func easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha12(out *jwriter.Writer, in VolumeAttachmentStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AttachError != nil {
		const prefix string = ",\"attachError\":"
		first = false
		out.RawString(prefix[1:])
		easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha13(out, *in.AttachError)
	}
	{
		const prefix string = ",\"attached\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Attached == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Attached))
		}
	}
	if len(in.AttachmentMetadata) != 0 {
		const prefix string = ",\"attachmentMetadata\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.AttachmentMetadata {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				out.String(string(v2Value))
			}
			out.RawByte('}')
		}
	}
	if in.DetachError != nil {
		const prefix string = ",\"detachError\":"
		out.RawString(prefix)
		easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha13(out, *in.DetachError)
	}
	out.RawByte('}')
}
func easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha13(in *jlexer.Lexer, out *VolumeError) {
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
		case "message":
			out.Message = string(in.String())
		case "time":
			if in.IsNull() {
				in.Skip()
				out.Time = nil
			} else {
				if out.Time == nil {
					out.Time = new(_v1.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Time).UnmarshalJSON(data))
				}
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
func easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha13(out *jwriter.Writer, in VolumeError) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	if in.Time != nil {
		const prefix string = ",\"time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.Time).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha11(in *jlexer.Lexer, out *VolumeAttachmentSpec) {
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
		case "attacher":
			if in.IsNull() {
				in.Skip()
				out.Attacher = nil
			} else {
				if out.Attacher == nil {
					out.Attacher = new(string)
				}
				*out.Attacher = string(in.String())
			}
		case "nodeName":
			if in.IsNull() {
				in.Skip()
				out.NodeName = nil
			} else {
				if out.NodeName == nil {
					out.NodeName = new(string)
				}
				*out.NodeName = string(in.String())
			}
		case "source":
			if in.IsNull() {
				in.Skip()
				out.Source = nil
			} else {
				if out.Source == nil {
					out.Source = new(VolumeAttachmentSource)
				}
				easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha14(in, out.Source)
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
func easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha11(out *jwriter.Writer, in VolumeAttachmentSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"attacher\":"
		out.RawString(prefix[1:])
		if in.Attacher == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Attacher))
		}
	}
	{
		const prefix string = ",\"nodeName\":"
		out.RawString(prefix)
		if in.NodeName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NodeName))
		}
	}
	{
		const prefix string = ",\"source\":"
		out.RawString(prefix)
		if in.Source == nil {
			out.RawString("null")
		} else {
			easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha14(out, *in.Source)
		}
	}
	out.RawByte('}')
}
func easyjson56496878DecodeGithubComKubewardenK8sObjectsApiStorageV1alpha14(in *jlexer.Lexer, out *VolumeAttachmentSource) {
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
		case "inlineVolumeSpec":
			if in.IsNull() {
				in.Skip()
				out.InlineVolumeSpec = nil
			} else {
				if out.InlineVolumeSpec == nil {
					out.InlineVolumeSpec = new(_v11.PersistentVolumeSpec)
				}
				(*out.InlineVolumeSpec).UnmarshalEasyJSON(in)
			}
		case "persistentVolumeName":
			out.PersistentVolumeName = string(in.String())
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
func easyjson56496878EncodeGithubComKubewardenK8sObjectsApiStorageV1alpha14(out *jwriter.Writer, in VolumeAttachmentSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.InlineVolumeSpec != nil {
		const prefix string = ",\"inlineVolumeSpec\":"
		first = false
		out.RawString(prefix[1:])
		(*in.InlineVolumeSpec).MarshalEasyJSON(out)
	}
	if in.PersistentVolumeName != "" {
		const prefix string = ",\"persistentVolumeName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PersistentVolumeName))
	}
	out.RawByte('}')
}
