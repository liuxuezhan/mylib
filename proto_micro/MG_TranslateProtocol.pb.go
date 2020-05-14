// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MG_TranslateProtocol.proto

package proto_micro

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ST_TranslationItem_PB struct {
	LanguageID uint32 `protobuf:"varint,1,opt,name=languageID" json:"languageID,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *ST_TranslationItem_PB) Reset()                    { *m = ST_TranslationItem_PB{} }
func (m *ST_TranslationItem_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_TranslationItem_PB) ProtoMessage()               {}
func (*ST_TranslationItem_PB) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *ST_TranslationItem_PB) GetLanguageID() uint32 {
	if m != nil {
		return m.LanguageID
	}
	return 0
}

func (m *ST_TranslationItem_PB) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ST_Translation_PB struct {
	ID         string                   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	TransItems []*ST_TranslationItem_PB `protobuf:"bytes,2,rep,name=transItems" json:"transItems,omitempty"`
}

func (m *ST_Translation_PB) Reset()                    { *m = ST_Translation_PB{} }
func (m *ST_Translation_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_Translation_PB) ProtoMessage()               {}
func (*ST_Translation_PB) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *ST_Translation_PB) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ST_Translation_PB) GetTransItems() []*ST_TranslationItem_PB {
	if m != nil {
		return m.TransItems
	}
	return nil
}

type ST_Translation_Req_PB struct {
	LanguageID uint32 `protobuf:"varint,1,opt,name=languageID" json:"languageID,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *ST_Translation_Req_PB) Reset()                    { *m = ST_Translation_Req_PB{} }
func (m *ST_Translation_Req_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_Translation_Req_PB) ProtoMessage()               {}
func (*ST_Translation_Req_PB) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *ST_Translation_Req_PB) GetLanguageID() uint32 {
	if m != nil {
		return m.LanguageID
	}
	return 0
}

func (m *ST_Translation_Req_PB) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ST_Translation_Respon_PB struct {
	Status     int32  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Result     string `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	LanguageID uint32 `protobuf:"varint,3,opt,name=languageID" json:"languageID,omitempty"`
}

func (m *ST_Translation_Respon_PB) Reset()                    { *m = ST_Translation_Respon_PB{} }
func (m *ST_Translation_Respon_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_Translation_Respon_PB) ProtoMessage()               {}
func (*ST_Translation_Respon_PB) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *ST_Translation_Respon_PB) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ST_Translation_Respon_PB) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *ST_Translation_Respon_PB) GetLanguageID() uint32 {
	if m != nil {
		return m.LanguageID
	}
	return 0
}

func init() {
	proto.RegisterType((*ST_TranslationItem_PB)(nil), "proto_micro.ST_TranslationItem_PB")
	proto.RegisterType((*ST_Translation_PB)(nil), "proto_micro.ST_Translation_PB")
	proto.RegisterType((*ST_Translation_Req_PB)(nil), "proto_micro.ST_Translation_Req_PB")
	proto.RegisterType((*ST_Translation_Respon_PB)(nil), "proto_micro.ST_Translation_Respon_PB")
}

func init() { proto.RegisterFile("MG_TranslateProtocol.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xf2, 0x75, 0x8f, 0x0f,
	0x29, 0x4a, 0xcc, 0x2b, 0xce, 0x49, 0x2c, 0x49, 0x0d, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf,
	0xd1, 0x2b, 0x00, 0x31, 0x84, 0xb8, 0xc1, 0x54, 0x7c, 0x6e, 0x66, 0x72, 0x51, 0xbe, 0x52, 0x20,
	0x97, 0x68, 0x70, 0x08, 0x5c, 0x69, 0x66, 0x7e, 0x9e, 0x67, 0x49, 0x6a, 0x6e, 0x7c, 0x80, 0x93,
	0x90, 0x1c, 0x17, 0x57, 0x4e, 0x62, 0x5e, 0x7a, 0x69, 0x62, 0x7a, 0xaa, 0xa7, 0x8b, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x6f, 0x10, 0x92, 0x88, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a,
	0x5e, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xab, 0x94, 0xce, 0x25, 0x88, 0x6a,
	0x24, 0xc8, 0x38, 0x3e, 0x2e, 0x26, 0xa8, 0x31, 0x9c, 0x41, 0x4c, 0x9e, 0x2e, 0x42, 0x4e, 0x5c,
	0x5c, 0x25, 0x20, 0x15, 0x20, 0xeb, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x94, 0xf4,
	0x90, 0x5c, 0xa6, 0x87, 0xd5, 0x59, 0x41, 0x48, 0xba, 0x30, 0xdd, 0x1e, 0x1f, 0x94, 0x5a, 0x48,
	0x99, 0xdb, 0xb3, 0xb8, 0x24, 0x30, 0x8c, 0x2c, 0x2e, 0x80, 0x78, 0x41, 0x8c, 0x8b, 0xad, 0xb8,
	0x24, 0xb1, 0xa4, 0xb4, 0x18, 0x6c, 0x22, 0x6b, 0x10, 0x94, 0x07, 0x12, 0x2f, 0x4a, 0x2d, 0x2e,
	0xcd, 0x81, 0x19, 0x06, 0xe5, 0xa1, 0xb9, 0x82, 0x19, 0xdd, 0x15, 0x4e, 0xac, 0x1e, 0xcc, 0x0d,
	0x8c, 0x0c, 0x49, 0x6c, 0x60, 0x4f, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x11, 0x59, 0x44,
	0x15, 0xb3, 0x01, 0x00, 0x00,
}