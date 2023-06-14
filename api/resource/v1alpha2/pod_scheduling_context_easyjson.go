// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha2

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

func easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(in *jlexer.Lexer, out *PodSchedulingContext) {
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
					out.Spec = new(PodSchedulingContextSpec)
				}
				easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha21(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(PodSchedulingContextStatus)
				}
				easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha22(in, out.Status)
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
func easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(out *jwriter.Writer, in PodSchedulingContext) {
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
			easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha21(out, *in.Spec)
		}
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha22(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodSchedulingContext) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodSchedulingContext) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodSchedulingContext) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodSchedulingContext) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha2(l, v)
}
func easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha22(in *jlexer.Lexer, out *PodSchedulingContextStatus) {
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
		case "resourceClaims":
			if in.IsNull() {
				in.Skip()
				out.ResourceClaims = nil
			} else {
				in.Delim('[')
				if out.ResourceClaims == nil {
					if !in.IsDelim(']') {
						out.ResourceClaims = make([]*ResourceClaimSchedulingStatus, 0, 8)
					} else {
						out.ResourceClaims = []*ResourceClaimSchedulingStatus{}
					}
				} else {
					out.ResourceClaims = (out.ResourceClaims)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *ResourceClaimSchedulingStatus
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(ResourceClaimSchedulingStatus)
						}
						easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha23(in, v1)
					}
					out.ResourceClaims = append(out.ResourceClaims, v1)
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
func easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha22(out *jwriter.Writer, in PodSchedulingContextStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.ResourceClaims) != 0 {
		const prefix string = ",\"resourceClaims\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.ResourceClaims {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha23(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha23(in *jlexer.Lexer, out *ResourceClaimSchedulingStatus) {
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
		case "name":
			out.Name = string(in.String())
		case "unsuitableNodes":
			if in.IsNull() {
				in.Skip()
				out.UnsuitableNodes = nil
			} else {
				in.Delim('[')
				if out.UnsuitableNodes == nil {
					if !in.IsDelim(']') {
						out.UnsuitableNodes = make([]string, 0, 4)
					} else {
						out.UnsuitableNodes = []string{}
					}
				} else {
					out.UnsuitableNodes = (out.UnsuitableNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.UnsuitableNodes = append(out.UnsuitableNodes, v4)
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
func easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha23(out *jwriter.Writer, in ResourceClaimSchedulingStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if len(in.UnsuitableNodes) != 0 {
		const prefix string = ",\"unsuitableNodes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.UnsuitableNodes {
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
func easyjson81f1465aDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha21(in *jlexer.Lexer, out *PodSchedulingContextSpec) {
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
		case "potentialNodes":
			if in.IsNull() {
				in.Skip()
				out.PotentialNodes = nil
			} else {
				in.Delim('[')
				if out.PotentialNodes == nil {
					if !in.IsDelim(']') {
						out.PotentialNodes = make([]string, 0, 4)
					} else {
						out.PotentialNodes = []string{}
					}
				} else {
					out.PotentialNodes = (out.PotentialNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.PotentialNodes = append(out.PotentialNodes, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "selectedNode":
			out.SelectedNode = string(in.String())
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
func easyjson81f1465aEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha21(out *jwriter.Writer, in PodSchedulingContextSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.PotentialNodes) != 0 {
		const prefix string = ",\"potentialNodes\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v8, v9 := range in.PotentialNodes {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.SelectedNode != "" {
		const prefix string = ",\"selectedNode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SelectedNode))
	}
	out.RawByte('}')
}
