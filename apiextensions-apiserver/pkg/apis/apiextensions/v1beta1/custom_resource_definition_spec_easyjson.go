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

func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(in *jlexer.Lexer, out *CustomResourceDefinitionSpec) {
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
		case "additionalPrinterColumns":
			if in.IsNull() {
				in.Skip()
				out.AdditionalPrinterColumns = nil
			} else {
				in.Delim('[')
				if out.AdditionalPrinterColumns == nil {
					if !in.IsDelim(']') {
						out.AdditionalPrinterColumns = make([]*CustomResourceColumnDefinition, 0, 8)
					} else {
						out.AdditionalPrinterColumns = []*CustomResourceColumnDefinition{}
					}
				} else {
					out.AdditionalPrinterColumns = (out.AdditionalPrinterColumns)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *CustomResourceColumnDefinition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(CustomResourceColumnDefinition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.AdditionalPrinterColumns = append(out.AdditionalPrinterColumns, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "conversion":
			if in.IsNull() {
				in.Skip()
				out.Conversion = nil
			} else {
				if out.Conversion == nil {
					out.Conversion = new(CustomResourceConversion)
				}
				(*out.Conversion).UnmarshalEasyJSON(in)
			}
		case "group":
			if in.IsNull() {
				in.Skip()
				out.Group = nil
			} else {
				if out.Group == nil {
					out.Group = new(string)
				}
				*out.Group = string(in.String())
			}
		case "names":
			if in.IsNull() {
				in.Skip()
				out.Names = nil
			} else {
				if out.Names == nil {
					out.Names = new(CustomResourceDefinitionNames)
				}
				(*out.Names).UnmarshalEasyJSON(in)
			}
		case "preserveUnknownFields":
			out.PreserveUnknownFields = bool(in.Bool())
		case "scope":
			if in.IsNull() {
				in.Skip()
				out.Scope = nil
			} else {
				if out.Scope == nil {
					out.Scope = new(string)
				}
				*out.Scope = string(in.String())
			}
		case "subresources":
			if in.IsNull() {
				in.Skip()
				out.Subresources = nil
			} else {
				if out.Subresources == nil {
					out.Subresources = new(CustomResourceSubresources)
				}
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, out.Subresources)
			}
		case "validation":
			if in.IsNull() {
				in.Skip()
				out.Validation = nil
			} else {
				if out.Validation == nil {
					out.Validation = new(CustomResourceValidation)
				}
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(in, out.Validation)
			}
		case "version":
			out.Version = string(in.String())
		case "versions":
			if in.IsNull() {
				in.Skip()
				out.Versions = nil
			} else {
				in.Delim('[')
				if out.Versions == nil {
					if !in.IsDelim(']') {
						out.Versions = make([]*CustomResourceDefinitionVersion, 0, 8)
					} else {
						out.Versions = []*CustomResourceDefinitionVersion{}
					}
				} else {
					out.Versions = (out.Versions)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *CustomResourceDefinitionVersion
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(CustomResourceDefinitionVersion)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta13(in, v2)
					}
					out.Versions = append(out.Versions, v2)
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(out *jwriter.Writer, in CustomResourceDefinitionSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"additionalPrinterColumns\":"
		out.RawString(prefix[1:])
		if in.AdditionalPrinterColumns == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.AdditionalPrinterColumns {
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
	if in.Conversion != nil {
		const prefix string = ",\"conversion\":"
		out.RawString(prefix)
		(*in.Conversion).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"group\":"
		out.RawString(prefix)
		if in.Group == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Group))
		}
	}
	{
		const prefix string = ",\"names\":"
		out.RawString(prefix)
		if in.Names == nil {
			out.RawString("null")
		} else {
			(*in.Names).MarshalEasyJSON(out)
		}
	}
	if in.PreserveUnknownFields {
		const prefix string = ",\"preserveUnknownFields\":"
		out.RawString(prefix)
		out.Bool(bool(in.PreserveUnknownFields))
	}
	{
		const prefix string = ",\"scope\":"
		out.RawString(prefix)
		if in.Scope == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Scope))
		}
	}
	if in.Subresources != nil {
		const prefix string = ",\"subresources\":"
		out.RawString(prefix)
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *in.Subresources)
	}
	if in.Validation != nil {
		const prefix string = ",\"validation\":"
		out.RawString(prefix)
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(out, *in.Validation)
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"versions\":"
		out.RawString(prefix)
		if in.Versions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Versions {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta13(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomResourceDefinitionSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CustomResourceDefinitionSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomResourceDefinitionSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CustomResourceDefinitionSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(l, v)
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta13(in *jlexer.Lexer, out *CustomResourceDefinitionVersion) {
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
		case "additionalPrinterColumns":
			if in.IsNull() {
				in.Skip()
				out.AdditionalPrinterColumns = nil
			} else {
				in.Delim('[')
				if out.AdditionalPrinterColumns == nil {
					if !in.IsDelim(']') {
						out.AdditionalPrinterColumns = make([]*CustomResourceColumnDefinition, 0, 8)
					} else {
						out.AdditionalPrinterColumns = []*CustomResourceColumnDefinition{}
					}
				} else {
					out.AdditionalPrinterColumns = (out.AdditionalPrinterColumns)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *CustomResourceColumnDefinition
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(CustomResourceColumnDefinition)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.AdditionalPrinterColumns = append(out.AdditionalPrinterColumns, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
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
		case "schema":
			if in.IsNull() {
				in.Skip()
				out.Schema = nil
			} else {
				if out.Schema == nil {
					out.Schema = new(CustomResourceValidation)
				}
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(in, out.Schema)
			}
		case "served":
			if in.IsNull() {
				in.Skip()
				out.Served = nil
			} else {
				if out.Served == nil {
					out.Served = new(bool)
				}
				*out.Served = bool(in.Bool())
			}
		case "storage":
			if in.IsNull() {
				in.Skip()
				out.Storage = nil
			} else {
				if out.Storage == nil {
					out.Storage = new(bool)
				}
				*out.Storage = bool(in.Bool())
			}
		case "subresources":
			if in.IsNull() {
				in.Skip()
				out.Subresources = nil
			} else {
				if out.Subresources == nil {
					out.Subresources = new(CustomResourceSubresources)
				}
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in, out.Subresources)
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta13(out *jwriter.Writer, in CustomResourceDefinitionVersion) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"additionalPrinterColumns\":"
		out.RawString(prefix[1:])
		if in.AdditionalPrinterColumns == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.AdditionalPrinterColumns {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	if in.Schema != nil {
		const prefix string = ",\"schema\":"
		out.RawString(prefix)
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(out, *in.Schema)
	}
	{
		const prefix string = ",\"served\":"
		out.RawString(prefix)
		if in.Served == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Served))
		}
	}
	{
		const prefix string = ",\"storage\":"
		out.RawString(prefix)
		if in.Storage == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Storage))
		}
	}
	if in.Subresources != nil {
		const prefix string = ",\"subresources\":"
		out.RawString(prefix)
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out, *in.Subresources)
	}
	out.RawByte('}')
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(in *jlexer.Lexer, out *CustomResourceValidation) {
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, out.OpenAPIV3Schema)
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta12(out *jwriter.Writer, in CustomResourceValidation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.OpenAPIV3Schema != nil {
		const prefix string = ",\"openAPIV3Schema\":"
		first = false
		out.RawString(prefix[1:])
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *in.OpenAPIV3Schema)
	}
	out.RawByte('}')
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in *jlexer.Lexer, out *JSONSchemaProps) {
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
			if m, ok := out.AdditionalItems.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.AdditionalItems.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.AdditionalItems = in.Interface()
			}
		case "additionalProperties":
			if m, ok := out.AdditionalProperties.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.AdditionalProperties.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.AdditionalProperties = in.Interface()
			}
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
					var v10 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v10)
					}
					out.AllOf = append(out.AllOf, v10)
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
					var v11 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v11)
					}
					out.AnyOf = append(out.AnyOf, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "default":
			if m, ok := out.Default.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Default.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Default = in.Interface()
			}
		case "definitions":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Definitions = make(map[string]JSONSchemaProps)
				} else {
					out.Definitions = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v12 JSONSchemaProps
					easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, &v12)
					(out.Definitions)[key] = v12
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
					out.Dependencies = make(map[string]JSONSchemaPropsOrStringArray)
				} else {
					out.Dependencies = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v13 JSONSchemaPropsOrStringArray
					if m, ok := v13.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v13.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v13 = in.Interface()
					}
					(out.Dependencies)[key] = v13
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
						out.Enum = make([]JSON, 0, 4)
					} else {
						out.Enum = []JSON{}
					}
				} else {
					out.Enum = (out.Enum)[:0]
				}
				for !in.IsDelim(']') {
					var v14 JSON
					if m, ok := v14.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v14.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v14 = in.Interface()
					}
					out.Enum = append(out.Enum, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "example":
			if m, ok := out.Example.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Example.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Example = in.Interface()
			}
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta15(in, out.ExternalDocs)
			}
		case "format":
			out.Format = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "items":
			if m, ok := out.Items.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Items.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Items = in.Interface()
			}
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
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, out.Not)
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
					var v15 *JSONSchemaProps
					if in.IsNull() {
						in.Skip()
						v15 = nil
					} else {
						if v15 == nil {
							v15 = new(JSONSchemaProps)
						}
						easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, v15)
					}
					out.OneOf = append(out.OneOf, v15)
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
					out.PatternProperties = make(map[string]JSONSchemaProps)
				} else {
					out.PatternProperties = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v16 JSONSchemaProps
					easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, &v16)
					(out.PatternProperties)[key] = v16
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
					out.Properties = make(map[string]JSONSchemaProps)
				} else {
					out.Properties = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v17 JSONSchemaProps
					easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(in, &v17)
					(out.Properties)[key] = v17
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
					var v18 string
					v18 = string(in.String())
					out.Required = append(out.Required, v18)
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
		case "x-kubernetes-embedded-resource":
			out.XKubernetesEmbeddedResource = bool(in.Bool())
		case "x-kubernetes-int-or-string":
			out.XKubernetesIntOrString = bool(in.Bool())
		case "x-kubernetes-list-map-keys":
			if in.IsNull() {
				in.Skip()
				out.XKubernetesListMapKeys = nil
			} else {
				in.Delim('[')
				if out.XKubernetesListMapKeys == nil {
					if !in.IsDelim(']') {
						out.XKubernetesListMapKeys = make([]string, 0, 4)
					} else {
						out.XKubernetesListMapKeys = []string{}
					}
				} else {
					out.XKubernetesListMapKeys = (out.XKubernetesListMapKeys)[:0]
				}
				for !in.IsDelim(']') {
					var v19 string
					v19 = string(in.String())
					out.XKubernetesListMapKeys = append(out.XKubernetesListMapKeys, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "x-kubernetes-list-type":
			out.XKubernetesListType = string(in.String())
		case "x-kubernetes-map-type":
			out.XKubernetesMapType = string(in.String())
		case "x-kubernetes-preserve-unknown-fields":
			out.XKubernetesPreserveUnknownFields = bool(in.Bool())
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out *jwriter.Writer, in JSONSchemaProps) {
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
	if in.AdditionalItems != nil {
		const prefix string = ",\"additionalItems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.AdditionalItems.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.AdditionalItems.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.AdditionalItems))
		}
	}
	if in.AdditionalProperties != nil {
		const prefix string = ",\"additionalProperties\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.AdditionalProperties.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.AdditionalProperties.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.AdditionalProperties))
		}
	}
	{
		const prefix string = ",\"allOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.AllOf == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.AllOf {
				if v20 > 0 {
					out.RawByte(',')
				}
				if v21 == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v21)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"anyOf\":"
		out.RawString(prefix)
		if in.AnyOf == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v22, v23 := range in.AnyOf {
				if v22 > 0 {
					out.RawByte(',')
				}
				if v23 == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v23)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Default != nil {
		const prefix string = ",\"default\":"
		out.RawString(prefix)
		if m, ok := in.Default.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Default.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Default))
		}
	}
	if len(in.Definitions) != 0 {
		const prefix string = ",\"definitions\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v24First := true
			for v24Name, v24Value := range in.Definitions {
				if v24First {
					v24First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v24Name))
				out.RawByte(':')
				easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, v24Value)
			}
			out.RawByte('}')
		}
	}
	if len(in.Dependencies) != 0 {
		const prefix string = ",\"dependencies\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v25First := true
			for v25Name, v25Value := range in.Dependencies {
				if v25First {
					v25First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v25Name))
				out.RawByte(':')
				if m, ok := v25Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v25Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v25Value))
				}
			}
			out.RawByte('}')
		}
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"enum\":"
		out.RawString(prefix)
		if in.Enum == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.Enum {
				if v26 > 0 {
					out.RawByte(',')
				}
				if m, ok := v27.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v27.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v27))
				}
			}
			out.RawByte(']')
		}
	}
	if in.Example != nil {
		const prefix string = ",\"example\":"
		out.RawString(prefix)
		if m, ok := in.Example.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Example.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Example))
		}
	}
	if in.ExclusiveMaximum {
		const prefix string = ",\"exclusiveMaximum\":"
		out.RawString(prefix)
		out.Bool(bool(in.ExclusiveMaximum))
	}
	if in.ExclusiveMinimum {
		const prefix string = ",\"exclusiveMinimum\":"
		out.RawString(prefix)
		out.Bool(bool(in.ExclusiveMinimum))
	}
	if in.ExternalDocs != nil {
		const prefix string = ",\"externalDocs\":"
		out.RawString(prefix)
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta15(out, *in.ExternalDocs)
	}
	if in.Format != "" {
		const prefix string = ",\"format\":"
		out.RawString(prefix)
		out.String(string(in.Format))
	}
	if in.ID != "" {
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	if in.Items != nil {
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		if m, ok := in.Items.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Items.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Items))
		}
	}
	if in.MaxItems != 0 {
		const prefix string = ",\"maxItems\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxItems))
	}
	if in.MaxLength != 0 {
		const prefix string = ",\"maxLength\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxLength))
	}
	if in.MaxProperties != 0 {
		const prefix string = ",\"maxProperties\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxProperties))
	}
	if in.Maximum != 0 {
		const prefix string = ",\"maximum\":"
		out.RawString(prefix)
		out.Float64(float64(in.Maximum))
	}
	if in.MinItems != 0 {
		const prefix string = ",\"minItems\":"
		out.RawString(prefix)
		out.Int64(int64(in.MinItems))
	}
	if in.MinLength != 0 {
		const prefix string = ",\"minLength\":"
		out.RawString(prefix)
		out.Int64(int64(in.MinLength))
	}
	if in.MinProperties != 0 {
		const prefix string = ",\"minProperties\":"
		out.RawString(prefix)
		out.Int64(int64(in.MinProperties))
	}
	if in.Minimum != 0 {
		const prefix string = ",\"minimum\":"
		out.RawString(prefix)
		out.Float64(float64(in.Minimum))
	}
	if in.MultipleOf != 0 {
		const prefix string = ",\"multipleOf\":"
		out.RawString(prefix)
		out.Float64(float64(in.MultipleOf))
	}
	if in.Not != nil {
		const prefix string = ",\"not\":"
		out.RawString(prefix)
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *in.Not)
	}
	if in.Nullable {
		const prefix string = ",\"nullable\":"
		out.RawString(prefix)
		out.Bool(bool(in.Nullable))
	}
	{
		const prefix string = ",\"oneOf\":"
		out.RawString(prefix)
		if in.OneOf == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v28, v29 := range in.OneOf {
				if v28 > 0 {
					out.RawByte(',')
				}
				if v29 == nil {
					out.RawString("null")
				} else {
					easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, *v29)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Pattern != "" {
		const prefix string = ",\"pattern\":"
		out.RawString(prefix)
		out.String(string(in.Pattern))
	}
	if len(in.PatternProperties) != 0 {
		const prefix string = ",\"patternProperties\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v30First := true
			for v30Name, v30Value := range in.PatternProperties {
				if v30First {
					v30First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v30Name))
				out.RawByte(':')
				easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, v30Value)
			}
			out.RawByte('}')
		}
	}
	if len(in.Properties) != 0 {
		const prefix string = ",\"properties\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v31First := true
			for v31Name, v31Value := range in.Properties {
				if v31First {
					v31First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v31Name))
				out.RawByte(':')
				easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta14(out, v31Value)
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"required\":"
		out.RawString(prefix)
		if in.Required == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v32, v33 := range in.Required {
				if v32 > 0 {
					out.RawByte(',')
				}
				out.String(string(v33))
			}
			out.RawByte(']')
		}
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	if in.UniqueItems {
		const prefix string = ",\"uniqueItems\":"
		out.RawString(prefix)
		out.Bool(bool(in.UniqueItems))
	}
	if in.XKubernetesEmbeddedResource {
		const prefix string = ",\"x-kubernetes-embedded-resource\":"
		out.RawString(prefix)
		out.Bool(bool(in.XKubernetesEmbeddedResource))
	}
	if in.XKubernetesIntOrString {
		const prefix string = ",\"x-kubernetes-int-or-string\":"
		out.RawString(prefix)
		out.Bool(bool(in.XKubernetesIntOrString))
	}
	{
		const prefix string = ",\"x-kubernetes-list-map-keys\":"
		out.RawString(prefix)
		if in.XKubernetesListMapKeys == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v34, v35 := range in.XKubernetesListMapKeys {
				if v34 > 0 {
					out.RawByte(',')
				}
				out.String(string(v35))
			}
			out.RawByte(']')
		}
	}
	if in.XKubernetesListType != "" {
		const prefix string = ",\"x-kubernetes-list-type\":"
		out.RawString(prefix)
		out.String(string(in.XKubernetesListType))
	}
	if in.XKubernetesMapType != "" {
		const prefix string = ",\"x-kubernetes-map-type\":"
		out.RawString(prefix)
		out.String(string(in.XKubernetesMapType))
	}
	if in.XKubernetesPreserveUnknownFields {
		const prefix string = ",\"x-kubernetes-preserve-unknown-fields\":"
		out.RawString(prefix)
		out.Bool(bool(in.XKubernetesPreserveUnknownFields))
	}
	out.RawByte('}')
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta15(in *jlexer.Lexer, out *ExternalDocumentation) {
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta15(out *jwriter.Writer, in ExternalDocumentation) {
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
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(in *jlexer.Lexer, out *CustomResourceSubresources) {
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
		case "scale":
			if in.IsNull() {
				in.Skip()
				out.Scale = nil
			} else {
				if out.Scale == nil {
					out.Scale = new(CustomResourceSubresourceScale)
				}
				easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta16(in, out.Scale)
			}
		case "status":
			if m, ok := out.Status.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Status.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Status = in.Interface()
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta11(out *jwriter.Writer, in CustomResourceSubresources) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Scale != nil {
		const prefix string = ",\"scale\":"
		first = false
		out.RawString(prefix[1:])
		easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta16(out, *in.Scale)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Status.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Status.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Status))
		}
	}
	out.RawByte('}')
}
func easyjson60898210DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta16(in *jlexer.Lexer, out *CustomResourceSubresourceScale) {
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
		case "labelSelectorPath":
			out.LabelSelectorPath = string(in.String())
		case "specReplicasPath":
			if in.IsNull() {
				in.Skip()
				out.SpecReplicasPath = nil
			} else {
				if out.SpecReplicasPath == nil {
					out.SpecReplicasPath = new(string)
				}
				*out.SpecReplicasPath = string(in.String())
			}
		case "statusReplicasPath":
			if in.IsNull() {
				in.Skip()
				out.StatusReplicasPath = nil
			} else {
				if out.StatusReplicasPath == nil {
					out.StatusReplicasPath = new(string)
				}
				*out.StatusReplicasPath = string(in.String())
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
func easyjson60898210EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta16(out *jwriter.Writer, in CustomResourceSubresourceScale) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LabelSelectorPath != "" {
		const prefix string = ",\"labelSelectorPath\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.LabelSelectorPath))
	}
	{
		const prefix string = ",\"specReplicasPath\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.SpecReplicasPath == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.SpecReplicasPath))
		}
	}
	{
		const prefix string = ",\"statusReplicasPath\":"
		out.RawString(prefix)
		if in.StatusReplicasPath == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.StatusReplicasPath))
		}
	}
	out.RawByte('}')
}