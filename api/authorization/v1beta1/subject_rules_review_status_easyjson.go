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

func easyjsonAe3f4c9aDecodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(in *jlexer.Lexer, out *SubjectRulesReviewStatus) {
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
		case "evaluationError":
			out.EvaluationError = string(in.String())
		case "incomplete":
			if in.IsNull() {
				in.Skip()
				out.Incomplete = nil
			} else {
				if out.Incomplete == nil {
					out.Incomplete = new(bool)
				}
				*out.Incomplete = bool(in.Bool())
			}
		case "nonResourceRules":
			if in.IsNull() {
				in.Skip()
				out.NonResourceRules = nil
			} else {
				in.Delim('[')
				if out.NonResourceRules == nil {
					if !in.IsDelim(']') {
						out.NonResourceRules = make([]*NonResourceRule, 0, 8)
					} else {
						out.NonResourceRules = []*NonResourceRule{}
					}
				} else {
					out.NonResourceRules = (out.NonResourceRules)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *NonResourceRule
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(NonResourceRule)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.NonResourceRules = append(out.NonResourceRules, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "resourceRules":
			if in.IsNull() {
				in.Skip()
				out.ResourceRules = nil
			} else {
				in.Delim('[')
				if out.ResourceRules == nil {
					if !in.IsDelim(']') {
						out.ResourceRules = make([]*ResourceRule, 0, 8)
					} else {
						out.ResourceRules = []*ResourceRule{}
					}
				} else {
					out.ResourceRules = (out.ResourceRules)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *ResourceRule
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(ResourceRule)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.ResourceRules = append(out.ResourceRules, v2)
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
func easyjsonAe3f4c9aEncodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(out *jwriter.Writer, in SubjectRulesReviewStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.EvaluationError != "" {
		const prefix string = ",\"evaluationError\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.EvaluationError))
	}
	{
		const prefix string = ",\"incomplete\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Incomplete == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Incomplete))
		}
	}
	{
		const prefix string = ",\"nonResourceRules\":"
		out.RawString(prefix)
		if in.NonResourceRules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.NonResourceRules {
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
	{
		const prefix string = ",\"resourceRules\":"
		out.RawString(prefix)
		if in.ResourceRules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.ResourceRules {
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
func (v SubjectRulesReviewStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe3f4c9aEncodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubjectRulesReviewStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe3f4c9aEncodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubjectRulesReviewStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe3f4c9aDecodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubjectRulesReviewStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe3f4c9aDecodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(l, v)
}
