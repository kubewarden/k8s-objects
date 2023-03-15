// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjsonEdb00d66DecodeGithubComKubewardenK8sObjectsApiStorageV1(in *jlexer.Lexer, out *CSIDriverSpec) {
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
		case "attachRequired":
			out.AttachRequired = bool(in.Bool())
		case "fsGroupPolicy":
			out.FsGroupPolicy = string(in.String())
		case "podInfoOnMount":
			out.PodInfoOnMount = bool(in.Bool())
		case "requiresRepublish":
			out.RequiresRepublish = bool(in.Bool())
		case "storageCapacity":
			out.StorageCapacity = bool(in.Bool())
		case "tokenRequests":
			if in.IsNull() {
				in.Skip()
				out.TokenRequests = nil
			} else {
				in.Delim('[')
				if out.TokenRequests == nil {
					if !in.IsDelim(']') {
						out.TokenRequests = make([]*TokenRequest, 0, 8)
					} else {
						out.TokenRequests = []*TokenRequest{}
					}
				} else {
					out.TokenRequests = (out.TokenRequests)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *TokenRequest
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(TokenRequest)
						}
						easyjsonEdb00d66DecodeGithubComKubewardenK8sObjectsApiStorageV11(in, v1)
					}
					out.TokenRequests = append(out.TokenRequests, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "volumeLifecycleModes":
			if in.IsNull() {
				in.Skip()
				out.VolumeLifecycleModes = nil
			} else {
				in.Delim('[')
				if out.VolumeLifecycleModes == nil {
					if !in.IsDelim(']') {
						out.VolumeLifecycleModes = make([]string, 0, 4)
					} else {
						out.VolumeLifecycleModes = []string{}
					}
				} else {
					out.VolumeLifecycleModes = (out.VolumeLifecycleModes)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.VolumeLifecycleModes = append(out.VolumeLifecycleModes, v2)
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
func easyjsonEdb00d66EncodeGithubComKubewardenK8sObjectsApiStorageV1(out *jwriter.Writer, in CSIDriverSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AttachRequired {
		const prefix string = ",\"attachRequired\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.AttachRequired))
	}
	if in.FsGroupPolicy != "" {
		const prefix string = ",\"fsGroupPolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FsGroupPolicy))
	}
	if in.PodInfoOnMount {
		const prefix string = ",\"podInfoOnMount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.PodInfoOnMount))
	}
	if in.RequiresRepublish {
		const prefix string = ",\"requiresRepublish\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.RequiresRepublish))
	}
	if in.StorageCapacity {
		const prefix string = ",\"storageCapacity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.StorageCapacity))
	}
	if len(in.TokenRequests) != 0 {
		const prefix string = ",\"tokenRequests\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.TokenRequests {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					easyjsonEdb00d66EncodeGithubComKubewardenK8sObjectsApiStorageV11(out, *v4)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.VolumeLifecycleModes) != 0 {
		const prefix string = ",\"volumeLifecycleModes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.VolumeLifecycleModes {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSIDriverSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEdb00d66EncodeGithubComKubewardenK8sObjectsApiStorageV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSIDriverSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEdb00d66EncodeGithubComKubewardenK8sObjectsApiStorageV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSIDriverSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEdb00d66DecodeGithubComKubewardenK8sObjectsApiStorageV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSIDriverSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEdb00d66DecodeGithubComKubewardenK8sObjectsApiStorageV1(l, v)
}
func easyjsonEdb00d66DecodeGithubComKubewardenK8sObjectsApiStorageV11(in *jlexer.Lexer, out *TokenRequest) {
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
		case "audience":
			if in.IsNull() {
				in.Skip()
				out.Audience = nil
			} else {
				if out.Audience == nil {
					out.Audience = new(string)
				}
				*out.Audience = string(in.String())
			}
		case "expirationSeconds":
			out.ExpirationSeconds = int64(in.Int64())
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
func easyjsonEdb00d66EncodeGithubComKubewardenK8sObjectsApiStorageV11(out *jwriter.Writer, in TokenRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"audience\":"
		out.RawString(prefix[1:])
		if in.Audience == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Audience))
		}
	}
	if in.ExpirationSeconds != 0 {
		const prefix string = ",\"expirationSeconds\":"
		out.RawString(prefix)
		out.Int64(int64(in.ExpirationSeconds))
	}
	out.RawByte('}')
}