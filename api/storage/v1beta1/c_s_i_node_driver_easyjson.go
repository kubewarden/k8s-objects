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

func easyjson27007753DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(in *jlexer.Lexer, out *CSINodeDriver) {
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
		case "nodeID":
			if in.IsNull() {
				in.Skip()
				out.NodeID = nil
			} else {
				if out.NodeID == nil {
					out.NodeID = new(string)
				}
				*out.NodeID = string(in.String())
			}
		case "topologyKeys":
			if in.IsNull() {
				in.Skip()
				out.TopologyKeys = nil
			} else {
				in.Delim('[')
				if out.TopologyKeys == nil {
					if !in.IsDelim(']') {
						out.TopologyKeys = make([]string, 0, 4)
					} else {
						out.TopologyKeys = []string{}
					}
				} else {
					out.TopologyKeys = (out.TopologyKeys)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.TopologyKeys = append(out.TopologyKeys, v1)
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
func easyjson27007753EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(out *jwriter.Writer, in CSINodeDriver) {
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
		const prefix string = ",\"nodeID\":"
		out.RawString(prefix)
		if in.NodeID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.NodeID))
		}
	}
	{
		const prefix string = ",\"topologyKeys\":"
		out.RawString(prefix)
		if in.TopologyKeys == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.TopologyKeys {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSINodeDriver) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson27007753EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSINodeDriver) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson27007753EncodeGithubComKubewardenK8sObjectsApiStorageV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSINodeDriver) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson27007753DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSINodeDriver) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson27007753DecodeGithubComKubewardenK8sObjectsApiStorageV1beta1(l, v)
}
