// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjson44021553DecodeGithubComKubewardenK8sObjectsApiStorageV1(in *jlexer.Lexer, out *VolumeAttachmentStatus) {
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
				easyjson44021553DecodeGithubComKubewardenK8sObjectsApiStorageV11(in, out.AttachError)
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
				easyjson44021553DecodeGithubComKubewardenK8sObjectsApiStorageV11(in, out.DetachError)
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
func easyjson44021553EncodeGithubComKubewardenK8sObjectsApiStorageV1(out *jwriter.Writer, in VolumeAttachmentStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AttachError != nil {
		const prefix string = ",\"attachError\":"
		first = false
		out.RawString(prefix[1:])
		easyjson44021553EncodeGithubComKubewardenK8sObjectsApiStorageV11(out, *in.AttachError)
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
		easyjson44021553EncodeGithubComKubewardenK8sObjectsApiStorageV11(out, *in.DetachError)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VolumeAttachmentStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson44021553EncodeGithubComKubewardenK8sObjectsApiStorageV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VolumeAttachmentStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson44021553EncodeGithubComKubewardenK8sObjectsApiStorageV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VolumeAttachmentStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson44021553DecodeGithubComKubewardenK8sObjectsApiStorageV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VolumeAttachmentStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson44021553DecodeGithubComKubewardenK8sObjectsApiStorageV1(l, v)
}
func easyjson44021553DecodeGithubComKubewardenK8sObjectsApiStorageV11(in *jlexer.Lexer, out *VolumeError) {
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
func easyjson44021553EncodeGithubComKubewardenK8sObjectsApiStorageV11(out *jwriter.Writer, in VolumeError) {
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
