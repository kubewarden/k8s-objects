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

func easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(in *jlexer.Lexer, out *CustomResourceValidation) {
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
		case "openAPIV3Schema":
			if in.IsNull() {
				in.Skip()
				out.OpenAPIV3Schema = nil
			} else {
				if out.OpenAPIV3Schema == nil {
					out.OpenAPIV3Schema = new(JSONSchemaProps)
				}
				easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, out.OpenAPIV3Schema)
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
func easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(out *jwriter.Writer, in CustomResourceValidation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.OpenAPIV3Schema != nil {
		const prefix string = ",\"openAPIV3Schema\":"
		first = false
		out.RawString(prefix[1:])
		easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *in.OpenAPIV3Schema)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomResourceValidation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CustomResourceValidation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomResourceValidation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CustomResourceValidation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(l, v)
}
func easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in *jlexer.Lexer, out *JSONSchemaProps) {
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
		case "$ref":
			out.DollarRef = string(in.String())
		case "$schema":
			out.DollarSchema = string(in.String())
		case "additionalItems":
			(out.AdditionalItems).UnmarshalEasyJSON(in)
		case "additionalProperties":
			(out.AdditionalProperties).UnmarshalEasyJSON(in)
		case "allOf":
			if in.IsNull() {
				in.Skip()
				out.AllOf = nil
			} else {
				in.Delim('[')
				if out.AllOf == nil {
					if !in.IsDelim(']') {
						out.AllOf = make([]*JSONSchemaProps, 0, 8)
					} else {
						out.AllOf = []*JSONSchemaProps{}
					}
				} else {
					out.AllOf = (out.AllOf)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(JSONSchemaProps)
						}
						easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, v1)
					}
					out.AllOf = append(out.AllOf, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "anyOf":
			if in.IsNull() {
				in.Skip()
				out.AnyOf = nil
			} else {
				in.Delim('[')
				if out.AnyOf == nil {
					if !in.IsDelim(']') {
						out.AnyOf = make([]*JSONSchemaProps, 0, 8)
					} else {
						out.AnyOf = []*JSONSchemaProps{}
					}
				} else {
					out.AnyOf = (out.AnyOf)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(JSONSchemaProps)
						}
						easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, v2)
					}
					out.AnyOf = append(out.AnyOf, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "default":
			(out.Default).UnmarshalEasyJSON(in)
		case "definitions":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Definitions = make(map[string]*JSONSchemaProps)
				} else {
					out.Definitions = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(JSONSchemaProps)
						}
						easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, v3)
					}
					(out.Definitions)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "dependencies":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Dependencies = make(map[string]easyjson.RawMessage)
				} else {
					out.Dependencies = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 easyjson.RawMessage
					(v4).UnmarshalEasyJSON(in)
					(out.Dependencies)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "description":
			out.Description = string(in.String())
		case "enum":
			if in.IsNull() {
				in.Skip()
				out.Enum = nil
			} else {
				in.Delim('[')
				if out.Enum == nil {
					if !in.IsDelim(']') {
						out.Enum = make([]easyjson.RawMessage, 0, 2)
					} else {
						out.Enum = []easyjson.RawMessage{}
					}
				} else {
					out.Enum = (out.Enum)[:0]
				}
				for !in.IsDelim(']') {
					var v5 easyjson.RawMessage
					(v5).UnmarshalEasyJSON(in)
					out.Enum = append(out.Enum, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "example":
			(out.Example).UnmarshalEasyJSON(in)
		case "exclusiveMaximum":
			out.ExclusiveMaximum = bool(in.Bool())
		case "exclusiveMinimum":
			out.ExclusiveMinimum = bool(in.Bool())
		case "externalDocs":
			if in.IsNull() {
				in.Skip()
				out.ExternalDocs = nil
			} else {
				if out.ExternalDocs == nil {
					out.ExternalDocs = new(ExternalDocumentation)
				}
				easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(in, out.ExternalDocs)
			}
		case "format":
			out.Format = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "items":
			(out.Items).UnmarshalEasyJSON(in)
		case "maxItems":
			out.MaxItems = int64(in.Int64())
		case "maxLength":
			out.MaxLength = int64(in.Int64())
		case "maxProperties":
			out.MaxProperties = int64(in.Int64())
		case "maximum":
			out.Maximum = float64(in.Float64())
		case "minItems":
			out.MinItems = int64(in.Int64())
		case "minLength":
			out.MinLength = int64(in.Int64())
		case "minProperties":
			out.MinProperties = int64(in.Int64())
		case "minimum":
			out.Minimum = float64(in.Float64())
		case "multipleOf":
			out.MultipleOf = float64(in.Float64())
		case "not":
			if in.IsNull() {
				in.Skip()
				out.Not = nil
			} else {
				if out.Not == nil {
					out.Not = new(JSONSchemaProps)
				}
				easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, out.Not)
			}
		case "nullable":
			out.Nullable = bool(in.Bool())
		case "oneOf":
			if in.IsNull() {
				in.Skip()
				out.OneOf = nil
			} else {
				in.Delim('[')
				if out.OneOf == nil {
					if !in.IsDelim(']') {
						out.OneOf = make([]*JSONSchemaProps, 0, 8)
					} else {
						out.OneOf = []*JSONSchemaProps{}
					}
				} else {
					out.OneOf = (out.OneOf)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(JSONSchemaProps)
						}
						easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, v6)
					}
					out.OneOf = append(out.OneOf, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pattern":
			out.Pattern = string(in.String())
		case "patternProperties":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.PatternProperties = make(map[string]*JSONSchemaProps)
				} else {
					out.PatternProperties = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(JSONSchemaProps)
						}
						easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, v7)
					}
					(out.PatternProperties)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		case "properties":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Properties = make(map[string]*JSONSchemaProps)
				} else {
					out.Properties = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(JSONSchemaProps)
						}
						easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, v8)
					}
					(out.Properties)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
			}
		case "required":
			if in.IsNull() {
				in.Skip()
				out.Required = nil
			} else {
				in.Delim('[')
				if out.Required == nil {
					if !in.IsDelim(']') {
						out.Required = make([]string, 0, 4)
					} else {
						out.Required = []string{}
					}
				} else {
					out.Required = (out.Required)[:0]
				}
				for !in.IsDelim(']') {
					var v9 string
					v9 = string(in.String())
					out.Required = append(out.Required, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "title":
			out.Title = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "uniqueItems":
			out.UniqueItems = bool(in.Bool())
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
func easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out *jwriter.Writer, in JSONSchemaProps) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DollarRef != "" {
		const prefix string = ",\"$ref\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.DollarRef))
	}
	if in.DollarSchema != "" {
		const prefix string = ",\"$schema\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DollarSchema))
	}
	if (in.AdditionalItems).IsDefined() {
		const prefix string = ",\"additionalItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.AdditionalItems).MarshalEasyJSON(out)
	}
	if (in.AdditionalProperties).IsDefined() {
		const prefix string = ",\"additionalProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.AdditionalProperties).MarshalEasyJSON(out)
	}
	if len(in.AllOf) != 0 {
		const prefix string = ",\"allOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.AllOf {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *v11)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.AnyOf) != 0 {
		const prefix string = ",\"anyOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v12, v13 := range in.AnyOf {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *v13)
				}
			}
			out.RawByte(']')
		}
	}
	if (in.Default).IsDefined() {
		const prefix string = ",\"default\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Default).MarshalEasyJSON(out)
	}
	if len(in.Definitions) != 0 {
		const prefix string = ",\"definitions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v14First := true
			for v14Name, v14Value := range in.Definitions {
				if v14First {
					v14First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v14Name))
				out.RawByte(':')
				if v14Value == nil {
					out.RawString("null")
				} else {
					easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *v14Value)
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Dependencies) != 0 {
		const prefix string = ",\"dependencies\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v15First := true
			for v15Name, v15Value := range in.Dependencies {
				if v15First {
					v15First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v15Name))
				out.RawByte(':')
				(v15Value).MarshalEasyJSON(out)
			}
			out.RawByte('}')
		}
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if len(in.Enum) != 0 {
		const prefix string = ",\"enum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.Enum {
				if v16 > 0 {
					out.RawByte(',')
				}
				(v17).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if (in.Example).IsDefined() {
		const prefix string = ",\"example\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Example).MarshalEasyJSON(out)
	}
	if in.ExclusiveMaximum {
		const prefix string = ",\"exclusiveMaximum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ExclusiveMaximum))
	}
	if in.ExclusiveMinimum {
		const prefix string = ",\"exclusiveMinimum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ExclusiveMinimum))
	}
	if in.ExternalDocs != nil {
		const prefix string = ",\"externalDocs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(out, *in.ExternalDocs)
	}
	if in.Format != "" {
		const prefix string = ",\"format\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Format))
	}
	if in.ID != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	if (in.Items).IsDefined() {
		const prefix string = ",\"items\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Items).MarshalEasyJSON(out)
	}
	if in.MaxItems != 0 {
		const prefix string = ",\"maxItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MaxItems))
	}
	if in.MaxLength != 0 {
		const prefix string = ",\"maxLength\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MaxLength))
	}
	if in.MaxProperties != 0 {
		const prefix string = ",\"maxProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MaxProperties))
	}
	if in.Maximum != 0 {
		const prefix string = ",\"maximum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Maximum))
	}
	if in.MinItems != 0 {
		const prefix string = ",\"minItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MinItems))
	}
	if in.MinLength != 0 {
		const prefix string = ",\"minLength\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MinLength))
	}
	if in.MinProperties != 0 {
		const prefix string = ",\"minProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MinProperties))
	}
	if in.Minimum != 0 {
		const prefix string = ",\"minimum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Minimum))
	}
	if in.MultipleOf != 0 {
		const prefix string = ",\"multipleOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.MultipleOf))
	}
	if in.Not != nil {
		const prefix string = ",\"not\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *in.Not)
	}
	if in.Nullable {
		const prefix string = ",\"nullable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Nullable))
	}
	if len(in.OneOf) != 0 {
		const prefix string = ",\"oneOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.OneOf {
				if v18 > 0 {
					out.RawByte(',')
				}
				if v19 == nil {
					out.RawString("null")
				} else {
					easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *v19)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Pattern != "" {
		const prefix string = ",\"pattern\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Pattern))
	}
	if len(in.PatternProperties) != 0 {
		const prefix string = ",\"patternProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v20First := true
			for v20Name, v20Value := range in.PatternProperties {
				if v20First {
					v20First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v20Name))
				out.RawByte(':')
				if v20Value == nil {
					out.RawString("null")
				} else {
					easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *v20Value)
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Properties) != 0 {
		const prefix string = ",\"properties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v21First := true
			for v21Name, v21Value := range in.Properties {
				if v21First {
					v21First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v21Name))
				out.RawByte(':')
				if v21Value == nil {
					out.RawString("null")
				} else {
					easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *v21Value)
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.Required) != 0 {
		const prefix string = ",\"required\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v22, v23 := range in.Required {
				if v22 > 0 {
					out.RawByte(',')
				}
				out.String(string(v23))
			}
			out.RawByte(']')
		}
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
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
	if in.UniqueItems {
		const prefix string = ",\"uniqueItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.UniqueItems))
	}
	out.RawByte('}')
}
func easyjson508aa052DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(in *jlexer.Lexer, out *ExternalDocumentation) {
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
		case "description":
			out.Description = string(in.String())
		case "url":
			out.URL = string(in.String())
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
func easyjson508aa052EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(out *jwriter.Writer, in ExternalDocumentation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Description != "" {
		const prefix string = ",\"description\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Description))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	out.RawByte('}')
}