// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

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

func easyjson847fb0b7DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(in *jlexer.Lexer, out *StorageVersionStatus) {
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
						(*v1).UnmarshalEasyJSON(in)
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
func easyjson847fb0b7EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(out *jwriter.Writer, in StorageVersionStatus) {
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
					(*v4).MarshalEasyJSON(out)
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

// MarshalJSON supports json.Marshaler interface
func (v StorageVersionStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson847fb0b7EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StorageVersionStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson847fb0b7EncodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StorageVersionStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson847fb0b7DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StorageVersionStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson847fb0b7DecodeGithubComKubewardenK8sObjectsApiApiserverinternalV1alpha1(l, v)
}
