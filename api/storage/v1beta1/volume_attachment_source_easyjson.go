// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
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

func easyjson15f4f3f6DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(in *jlexer.Lexer, out *VolumeAttachmentSource) {
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
			(out.InlineVolumeSpec).UnmarshalEasyJSON(in)
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
func easyjson15f4f3f6EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(out *jwriter.Writer, in VolumeAttachmentSource) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"inlineVolumeSpec\":"
		first = false
		out.RawString(prefix[1:])
		(in.InlineVolumeSpec).MarshalEasyJSON(out)
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

// MarshalJSON supports json.Marshaler interface
func (v VolumeAttachmentSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15f4f3f6EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VolumeAttachmentSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15f4f3f6EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VolumeAttachmentSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15f4f3f6DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VolumeAttachmentSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15f4f3f6DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(l, v)
}
