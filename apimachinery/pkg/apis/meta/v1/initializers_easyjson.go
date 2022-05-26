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

func easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(in *jlexer.Lexer, out *Initializers) {
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
		case "pending":
			if in.IsNull() {
				in.Skip()
				out.Pending = nil
			} else {
				in.Delim('[')
				if out.Pending == nil {
					if !in.IsDelim(']') {
						out.Pending = make([]*Initializer, 0, 8)
					} else {
						out.Pending = []*Initializer{}
					}
				} else {
					out.Pending = (out.Pending)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Initializer
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Initializer)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Pending = append(out.Pending, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "result":
			if in.IsNull() {
				in.Skip()
				out.Result = nil
			} else {
				if out.Result == nil {
					out.Result = new(Status)
				}
				easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV11(in, out.Result)
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
func easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(out *jwriter.Writer, in Initializers) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pending\":"
		out.RawString(prefix[1:])
		if in.Pending == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Pending {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Result != nil {
		const prefix string = ",\"result\":"
		out.RawString(prefix)
		easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV11(out, *in.Result)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Initializers) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Initializers) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Initializers) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Initializers) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(l, v)
}
func easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV11(in *jlexer.Lexer, out *Status) {
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
		case "code":
			out.Code = int32(in.Int32())
		case "details":
			if in.IsNull() {
				in.Skip()
				out.Details = nil
			} else {
				if out.Details == nil {
					out.Details = new(StatusDetails)
				}
				easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV12(in, out.Details)
			}
		case "kind":
			out.Kind = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(ListMeta)
				}
				easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV13(in, out.Metadata)
			}
		case "reason":
			out.Reason = string(in.String())
		case "status":
			out.Status = string(in.String())
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
func easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV11(out *jwriter.Writer, in Status) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	if in.Code != 0 {
		const prefix string = ",\"code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Code))
	}
	if in.Details != nil {
		const prefix string = ",\"details\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV12(out, *in.Details)
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
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV13(out, *in.Metadata)
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
	if in.Status != "" {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	out.RawByte('}')
}
func easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV13(in *jlexer.Lexer, out *ListMeta) {
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
		case "continue":
			out.Continue = string(in.String())
		case "remainingItemCount":
			out.RemainingItemCount = int64(in.Int64())
		case "resourceVersion":
			out.ResourceVersion = string(in.String())
		case "selfLink":
			out.SelfLink = string(in.String())
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
func easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV13(out *jwriter.Writer, in ListMeta) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Continue != "" {
		const prefix string = ",\"continue\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Continue))
	}
	if in.RemainingItemCount != 0 {
		const prefix string = ",\"remainingItemCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RemainingItemCount))
	}
	if in.ResourceVersion != "" {
		const prefix string = ",\"resourceVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ResourceVersion))
	}
	if in.SelfLink != "" {
		const prefix string = ",\"selfLink\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SelfLink))
	}
	out.RawByte('}')
}
func easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV12(in *jlexer.Lexer, out *StatusDetails) {
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
		case "causes":
			if in.IsNull() {
				in.Skip()
				out.Causes = nil
			} else {
				in.Delim('[')
				if out.Causes == nil {
					if !in.IsDelim(']') {
						out.Causes = make([]*StatusCause, 0, 8)
					} else {
						out.Causes = []*StatusCause{}
					}
				} else {
					out.Causes = (out.Causes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *StatusCause
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(StatusCause)
						}
						easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV14(in, v4)
					}
					out.Causes = append(out.Causes, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "group":
			out.Group = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "retryAfterSeconds":
			out.RetryAfterSeconds = int32(in.Int32())
		case "uid":
			out.UID = string(in.String())
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
func easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV12(out *jwriter.Writer, in StatusDetails) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"causes\":"
		out.RawString(prefix[1:])
		if in.Causes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Causes {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV14(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Group != "" {
		const prefix string = ",\"group\":"
		out.RawString(prefix)
		out.String(string(in.Group))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.RetryAfterSeconds != 0 {
		const prefix string = ",\"retryAfterSeconds\":"
		out.RawString(prefix)
		out.Int32(int32(in.RetryAfterSeconds))
	}
	if in.UID != "" {
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.String(string(in.UID))
	}
	out.RawByte('}')
}
func easyjson30d800e3DecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV14(in *jlexer.Lexer, out *StatusCause) {
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
		case "field":
			out.Field = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "reason":
			out.Reason = string(in.String())
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
func easyjson30d800e3EncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV14(out *jwriter.Writer, in StatusCause) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Field != "" {
		const prefix string = ",\"field\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Field))
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
	out.RawByte('}')
}
