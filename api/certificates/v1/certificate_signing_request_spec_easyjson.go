// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	strfmt "github.com/go-openapi/strfmt"
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

func easyjson7bd069bfDecodeGithubComKubewardenK8sObjectsApiCertificatesV1(in *jlexer.Lexer, out *CertificateSigningRequestSpec) {
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
		case "expirationSeconds":
			out.ExpirationSeconds = int32(in.Int32())
		case "extra":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Extra = make(map[string][]string)
				} else {
					out.Extra = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 []string
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						in.Delim('[')
						if v1 == nil {
							if !in.IsDelim(']') {
								v1 = make([]string, 0, 4)
							} else {
								v1 = []string{}
							}
						} else {
							v1 = (v1)[:0]
						}
						for !in.IsDelim(']') {
							var v2 string
							v2 = string(in.String())
							v1 = append(v1, v2)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Extra)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "groups":
			if in.IsNull() {
				in.Skip()
				out.Groups = nil
			} else {
				in.Delim('[')
				if out.Groups == nil {
					if !in.IsDelim(']') {
						out.Groups = make([]string, 0, 4)
					} else {
						out.Groups = []string{}
					}
				} else {
					out.Groups = (out.Groups)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Groups = append(out.Groups, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "request":
			if in.IsNull() {
				in.Skip()
				out.Request = nil
			} else {
				if out.Request == nil {
					out.Request = new(strfmt.Base64)
				}
				if in.IsNull() {
					in.Skip()
					*out.Request = nil
				} else {
					*out.Request = in.Bytes()
				}
			}
		case "signerName":
			if in.IsNull() {
				in.Skip()
				out.SignerName = nil
			} else {
				if out.SignerName == nil {
					out.SignerName = new(string)
				}
				*out.SignerName = string(in.String())
			}
		case "uid":
			out.UID = string(in.String())
		case "usages":
			if in.IsNull() {
				in.Skip()
				out.Usages = nil
			} else {
				in.Delim('[')
				if out.Usages == nil {
					if !in.IsDelim(']') {
						out.Usages = make([]string, 0, 4)
					} else {
						out.Usages = []string{}
					}
				} else {
					out.Usages = (out.Usages)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.Usages = append(out.Usages, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "username":
			out.Username = string(in.String())
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
func easyjson7bd069bfEncodeGithubComKubewardenK8sObjectsApiCertificatesV1(out *jwriter.Writer, in CertificateSigningRequestSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ExpirationSeconds != 0 {
		const prefix string = ",\"expirationSeconds\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.ExpirationSeconds))
	}
	if len(in.Extra) != 0 {
		const prefix string = ",\"extra\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v6First := true
			for v6Name, v6Value := range in.Extra {
				if v6First {
					v6First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v6Name))
				out.RawByte(':')
				if v6Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v7, v8 := range v6Value {
						if v7 > 0 {
							out.RawByte(',')
						}
						out.String(string(v8))
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"groups\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Groups == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Groups {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"request\":"
		out.RawString(prefix)
		if in.Request == nil {
			out.RawString("null")
		} else {
			out.Base64Bytes(*in.Request)
		}
	}
	{
		const prefix string = ",\"signerName\":"
		out.RawString(prefix)
		if in.SignerName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.SignerName))
		}
	}
	if in.UID != "" {
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.String(string(in.UID))
	}
	{
		const prefix string = ",\"usages\":"
		out.RawString(prefix)
		if in.Usages == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.Usages {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.String(string(v14))
			}
			out.RawByte(']')
		}
	}
	if in.Username != "" {
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CertificateSigningRequestSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7bd069bfEncodeGithubComKubewardenK8sObjectsApiCertificatesV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CertificateSigningRequestSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7bd069bfEncodeGithubComKubewardenK8sObjectsApiCertificatesV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CertificateSigningRequestSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7bd069bfDecodeGithubComKubewardenK8sObjectsApiCertificatesV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CertificateSigningRequestSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7bd069bfDecodeGithubComKubewardenK8sObjectsApiCertificatesV1(l, v)
}
