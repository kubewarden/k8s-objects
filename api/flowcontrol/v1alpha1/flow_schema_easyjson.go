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

func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(in *jlexer.Lexer, out *FlowSchema) {
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
			(out.Metadata).UnmarshalEasyJSON(in)
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(FlowSchemaSpec)
				}
				easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(FlowSchemaStatus)
				}
				easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha12(in, out.Status)
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(out *jwriter.Writer, in FlowSchema) {
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
	if true {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FlowSchema) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FlowSchema) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FlowSchema) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FlowSchema) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha1(l, v)
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha12(in *jlexer.Lexer, out *FlowSchemaStatus) {
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
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*FlowSchemaCondition, 0, 8)
					} else {
						out.Conditions = []*FlowSchemaCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *FlowSchemaCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(FlowSchemaCondition)
						}
						easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha13(in, v1)
					}
					out.Conditions = append(out.Conditions, v1)
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha12(out *jwriter.Writer, in FlowSchemaStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"conditions\":"
		out.RawString(prefix[1:])
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha13(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha13(in *jlexer.Lexer, out *FlowSchemaCondition) {
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
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastTransitionTime).UnmarshalJSON(data))
			}
		case "message":
			out.Message = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "type":
			out.Type = string(in.String())
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha13(out *jwriter.Writer, in FlowSchemaCondition) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"lastTransitionTime\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.LastTransitionTime).MarshalJSON())
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
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha11(in *jlexer.Lexer, out *FlowSchemaSpec) {
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
		case "distinguisherMethod":
			if in.IsNull() {
				in.Skip()
				out.DistinguisherMethod = nil
			} else {
				if out.DistinguisherMethod == nil {
					out.DistinguisherMethod = new(FlowDistinguisherMethod)
				}
				(*out.DistinguisherMethod).UnmarshalEasyJSON(in)
			}
		case "matchingPrecedence":
			out.MatchingPrecedence = int32(in.Int32())
		case "priorityLevelConfiguration":
			if in.IsNull() {
				in.Skip()
				out.PriorityLevelConfiguration = nil
			} else {
				if out.PriorityLevelConfiguration == nil {
					out.PriorityLevelConfiguration = new(PriorityLevelConfigurationReference)
				}
				easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha14(in, out.PriorityLevelConfiguration)
			}
		case "rules":
			if in.IsNull() {
				in.Skip()
				out.Rules = nil
			} else {
				in.Delim('[')
				if out.Rules == nil {
					if !in.IsDelim(']') {
						out.Rules = make([]*PolicyRulesWithSubjects, 0, 8)
					} else {
						out.Rules = []*PolicyRulesWithSubjects{}
					}
				} else {
					out.Rules = (out.Rules)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *PolicyRulesWithSubjects
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(PolicyRulesWithSubjects)
						}
						easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha15(in, v4)
					}
					out.Rules = append(out.Rules, v4)
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha11(out *jwriter.Writer, in FlowSchemaSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DistinguisherMethod != nil {
		const prefix string = ",\"distinguisherMethod\":"
		first = false
		out.RawString(prefix[1:])
		(*in.DistinguisherMethod).MarshalEasyJSON(out)
	}
	if in.MatchingPrecedence != 0 {
		const prefix string = ",\"matchingPrecedence\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MatchingPrecedence))
	}
	{
		const prefix string = ",\"priorityLevelConfiguration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.PriorityLevelConfiguration == nil {
			out.RawString("null")
		} else {
			easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha14(out, *in.PriorityLevelConfiguration)
		}
	}
	{
		const prefix string = ",\"rules\":"
		out.RawString(prefix)
		if in.Rules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Rules {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha15(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha15(in *jlexer.Lexer, out *PolicyRulesWithSubjects) {
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
		case "nonResourceRules":
			if in.IsNull() {
				in.Skip()
				out.NonResourceRules = nil
			} else {
				in.Delim('[')
				if out.NonResourceRules == nil {
					if !in.IsDelim(']') {
						out.NonResourceRules = make([]*NonResourcePolicyRule, 0, 8)
					} else {
						out.NonResourceRules = []*NonResourcePolicyRule{}
					}
				} else {
					out.NonResourceRules = (out.NonResourceRules)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *NonResourcePolicyRule
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(NonResourcePolicyRule)
						}
						easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha16(in, v7)
					}
					out.NonResourceRules = append(out.NonResourceRules, v7)
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
						out.ResourceRules = make([]*ResourcePolicyRule, 0, 8)
					} else {
						out.ResourceRules = []*ResourcePolicyRule{}
					}
				} else {
					out.ResourceRules = (out.ResourceRules)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *ResourcePolicyRule
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(ResourcePolicyRule)
						}
						easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha17(in, v8)
					}
					out.ResourceRules = append(out.ResourceRules, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "subjects":
			if in.IsNull() {
				in.Skip()
				out.Subjects = nil
			} else {
				in.Delim('[')
				if out.Subjects == nil {
					if !in.IsDelim(']') {
						out.Subjects = make([]*Subject, 0, 8)
					} else {
						out.Subjects = []*Subject{}
					}
				} else {
					out.Subjects = (out.Subjects)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *Subject
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(Subject)
						}
						easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha18(in, v9)
					}
					out.Subjects = append(out.Subjects, v9)
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha15(out *jwriter.Writer, in PolicyRulesWithSubjects) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nonResourceRules\":"
		out.RawString(prefix[1:])
		if in.NonResourceRules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.NonResourceRules {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha16(out, *v11)
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
			for v12, v13 := range in.ResourceRules {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha17(out, *v13)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"subjects\":"
		out.RawString(prefix)
		if in.Subjects == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Subjects {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha18(out, *v15)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha18(in *jlexer.Lexer, out *Subject) {
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
		case "group":
			if in.IsNull() {
				in.Skip()
				out.Group = nil
			} else {
				if out.Group == nil {
					out.Group = new(GroupSubject)
				}
				easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha19(in, out.Group)
			}
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
			}
		case "serviceAccount":
			if in.IsNull() {
				in.Skip()
				out.ServiceAccount = nil
			} else {
				if out.ServiceAccount == nil {
					out.ServiceAccount = new(ServiceAccountSubject)
				}
				easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha110(in, out.ServiceAccount)
			}
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(UserSubject)
				}
				easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha111(in, out.User)
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha18(out *jwriter.Writer, in Subject) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Group != nil {
		const prefix string = ",\"group\":"
		first = false
		out.RawString(prefix[1:])
		easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha19(out, *in.Group)
	}
	{
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Kind))
		}
	}
	if in.ServiceAccount != nil {
		const prefix string = ",\"serviceAccount\":"
		out.RawString(prefix)
		easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha110(out, *in.ServiceAccount)
	}
	if in.User != nil {
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha111(out, *in.User)
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha111(in *jlexer.Lexer, out *UserSubject) {
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
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha111(out *jwriter.Writer, in UserSubject) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha110(in *jlexer.Lexer, out *ServiceAccountSubject) {
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
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "namespace":
			if in.IsNull() {
				in.Skip()
				out.Namespace = nil
			} else {
				if out.Namespace == nil {
					out.Namespace = new(string)
				}
				*out.Namespace = string(in.String())
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha110(out *jwriter.Writer, in ServiceAccountSubject) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		if in.Namespace == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Namespace))
		}
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha19(in *jlexer.Lexer, out *GroupSubject) {
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
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha19(out *jwriter.Writer, in GroupSubject) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha17(in *jlexer.Lexer, out *ResourcePolicyRule) {
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
		case "apiGroups":
			if in.IsNull() {
				in.Skip()
				out.APIGroups = nil
			} else {
				in.Delim('[')
				if out.APIGroups == nil {
					if !in.IsDelim(']') {
						out.APIGroups = make([]string, 0, 4)
					} else {
						out.APIGroups = []string{}
					}
				} else {
					out.APIGroups = (out.APIGroups)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.APIGroups = append(out.APIGroups, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "clusterScope":
			out.ClusterScope = bool(in.Bool())
		case "namespaces":
			if in.IsNull() {
				in.Skip()
				out.Namespaces = nil
			} else {
				in.Delim('[')
				if out.Namespaces == nil {
					if !in.IsDelim(']') {
						out.Namespaces = make([]string, 0, 4)
					} else {
						out.Namespaces = []string{}
					}
				} else {
					out.Namespaces = (out.Namespaces)[:0]
				}
				for !in.IsDelim(']') {
					var v17 string
					v17 = string(in.String())
					out.Namespaces = append(out.Namespaces, v17)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "resources":
			if in.IsNull() {
				in.Skip()
				out.Resources = nil
			} else {
				in.Delim('[')
				if out.Resources == nil {
					if !in.IsDelim(']') {
						out.Resources = make([]string, 0, 4)
					} else {
						out.Resources = []string{}
					}
				} else {
					out.Resources = (out.Resources)[:0]
				}
				for !in.IsDelim(']') {
					var v18 string
					v18 = string(in.String())
					out.Resources = append(out.Resources, v18)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "verbs":
			if in.IsNull() {
				in.Skip()
				out.Verbs = nil
			} else {
				in.Delim('[')
				if out.Verbs == nil {
					if !in.IsDelim(']') {
						out.Verbs = make([]string, 0, 4)
					} else {
						out.Verbs = []string{}
					}
				} else {
					out.Verbs = (out.Verbs)[:0]
				}
				for !in.IsDelim(']') {
					var v19 string
					v19 = string(in.String())
					out.Verbs = append(out.Verbs, v19)
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha17(out *jwriter.Writer, in ResourcePolicyRule) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"apiGroups\":"
		out.RawString(prefix[1:])
		if in.APIGroups == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.APIGroups {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	if in.ClusterScope {
		const prefix string = ",\"clusterScope\":"
		out.RawString(prefix)
		out.Bool(bool(in.ClusterScope))
	}
	{
		const prefix string = ",\"namespaces\":"
		out.RawString(prefix)
		if in.Namespaces == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v22, v23 := range in.Namespaces {
				if v22 > 0 {
					out.RawByte(',')
				}
				out.String(string(v23))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"resources\":"
		out.RawString(prefix)
		if in.Resources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v24, v25 := range in.Resources {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.String(string(v25))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"verbs\":"
		out.RawString(prefix)
		if in.Verbs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.Verbs {
				if v26 > 0 {
					out.RawByte(',')
				}
				out.String(string(v27))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha16(in *jlexer.Lexer, out *NonResourcePolicyRule) {
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
		case "nonResourceURLs":
			if in.IsNull() {
				in.Skip()
				out.NonResourceURLs = nil
			} else {
				in.Delim('[')
				if out.NonResourceURLs == nil {
					if !in.IsDelim(']') {
						out.NonResourceURLs = make([]string, 0, 4)
					} else {
						out.NonResourceURLs = []string{}
					}
				} else {
					out.NonResourceURLs = (out.NonResourceURLs)[:0]
				}
				for !in.IsDelim(']') {
					var v28 string
					v28 = string(in.String())
					out.NonResourceURLs = append(out.NonResourceURLs, v28)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "verbs":
			if in.IsNull() {
				in.Skip()
				out.Verbs = nil
			} else {
				in.Delim('[')
				if out.Verbs == nil {
					if !in.IsDelim(']') {
						out.Verbs = make([]string, 0, 4)
					} else {
						out.Verbs = []string{}
					}
				} else {
					out.Verbs = (out.Verbs)[:0]
				}
				for !in.IsDelim(']') {
					var v29 string
					v29 = string(in.String())
					out.Verbs = append(out.Verbs, v29)
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha16(out *jwriter.Writer, in NonResourcePolicyRule) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nonResourceURLs\":"
		out.RawString(prefix[1:])
		if in.NonResourceURLs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v30, v31 := range in.NonResourceURLs {
				if v30 > 0 {
					out.RawByte(',')
				}
				out.String(string(v31))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"verbs\":"
		out.RawString(prefix)
		if in.Verbs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v32, v33 := range in.Verbs {
				if v32 > 0 {
					out.RawByte(',')
				}
				out.String(string(v33))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonFf54a190DecodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha14(in *jlexer.Lexer, out *PriorityLevelConfigurationReference) {
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
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
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
func easyjsonFf54a190EncodeGithubComKubewardenK8sObjectsApiFlowcontrolV1alpha14(out *jwriter.Writer, in PriorityLevelConfigurationReference) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}
