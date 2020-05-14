// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MG_UniqueProtocol.proto

package proto_micro

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 含有待检测名字的包,同时也是含有结果的包
type ST_GS2US_UserName_Check struct {
	Name   string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Result int32  `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
}

func (m *ST_GS2US_UserName_Check) Reset()                    { *m = ST_GS2US_UserName_Check{} }
func (m *ST_GS2US_UserName_Check) String() string            { return proto.CompactTextString(m) }
func (*ST_GS2US_UserName_Check) ProtoMessage()               {}
func (*ST_GS2US_UserName_Check) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *ST_GS2US_UserName_Check) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ST_GS2US_UserName_Check) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ST_GS2US_UserName_Change struct {
	Name    string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=newName" json:"newName,omitempty"`
}

func (m *ST_GS2US_UserName_Change) Reset()                    { *m = ST_GS2US_UserName_Change{} }
func (m *ST_GS2US_UserName_Change) String() string            { return proto.CompactTextString(m) }
func (*ST_GS2US_UserName_Change) ProtoMessage()               {}
func (*ST_GS2US_UserName_Change) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *ST_GS2US_UserName_Change) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ST_GS2US_UserName_Change) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

type ST_GS2US_LeagueName_Check struct {
	Name   string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	SName  string `protobuf:"bytes,2,opt,name=sName" json:"sName,omitempty"`
	Result int32  `protobuf:"varint,3,opt,name=result" json:"result,omitempty"`
}

func (m *ST_GS2US_LeagueName_Check) Reset()                    { *m = ST_GS2US_LeagueName_Check{} }
func (m *ST_GS2US_LeagueName_Check) String() string            { return proto.CompactTextString(m) }
func (*ST_GS2US_LeagueName_Check) ProtoMessage()               {}
func (*ST_GS2US_LeagueName_Check) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *ST_GS2US_LeagueName_Check) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ST_GS2US_LeagueName_Check) GetSName() string {
	if m != nil {
		return m.SName
	}
	return ""
}

func (m *ST_GS2US_LeagueName_Check) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ST_GS2US_LeagueName_Change struct {
	Name     string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	SName    string `protobuf:"bytes,2,opt,name=sName" json:"sName,omitempty"`
	NewName  string `protobuf:"bytes,3,opt,name=newName" json:"newName,omitempty"`
	NewSName string `protobuf:"bytes,4,opt,name=newSName" json:"newSName,omitempty"`
}

func (m *ST_GS2US_LeagueName_Change) Reset()                    { *m = ST_GS2US_LeagueName_Change{} }
func (m *ST_GS2US_LeagueName_Change) String() string            { return proto.CompactTextString(m) }
func (*ST_GS2US_LeagueName_Change) ProtoMessage()               {}
func (*ST_GS2US_LeagueName_Change) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *ST_GS2US_LeagueName_Change) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ST_GS2US_LeagueName_Change) GetSName() string {
	if m != nil {
		return m.SName
	}
	return ""
}

func (m *ST_GS2US_LeagueName_Change) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

func (m *ST_GS2US_LeagueName_Change) GetNewSName() string {
	if m != nil {
		return m.NewSName
	}
	return ""
}

type ST_GS2US_LeagueName_Delete struct {
	Name  string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	SName string `protobuf:"bytes,2,opt,name=sName" json:"sName,omitempty"`
}

func (m *ST_GS2US_LeagueName_Delete) Reset()                    { *m = ST_GS2US_LeagueName_Delete{} }
func (m *ST_GS2US_LeagueName_Delete) String() string            { return proto.CompactTextString(m) }
func (*ST_GS2US_LeagueName_Delete) ProtoMessage()               {}
func (*ST_GS2US_LeagueName_Delete) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *ST_GS2US_LeagueName_Delete) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ST_GS2US_LeagueName_Delete) GetSName() string {
	if m != nil {
		return m.SName
	}
	return ""
}

type ST_LeagueName_PB struct {
	LeagueName  string `protobuf:"bytes,1,opt,name=leagueName" json:"leagueName,omitempty"`
	LeagueLName string `protobuf:"bytes,2,opt,name=leagueLName" json:"leagueLName,omitempty"`
}

func (m *ST_LeagueName_PB) Reset()                    { *m = ST_LeagueName_PB{} }
func (m *ST_LeagueName_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_LeagueName_PB) ProtoMessage()               {}
func (*ST_LeagueName_PB) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *ST_LeagueName_PB) GetLeagueName() string {
	if m != nil {
		return m.LeagueName
	}
	return ""
}

func (m *ST_LeagueName_PB) GetLeagueLName() string {
	if m != nil {
		return m.LeagueLName
	}
	return ""
}

type ST_GS2US_NameList struct {
	ServerID    uint32              `protobuf:"varint,1,opt,name=serverID" json:"serverID,omitempty"`
	UserNames   []string            `protobuf:"bytes,2,rep,name=userNames" json:"userNames,omitempty"`
	LeagueNames []*ST_LeagueName_PB `protobuf:"bytes,3,rep,name=leagueNames" json:"leagueNames,omitempty"`
}

func (m *ST_GS2US_NameList) Reset()                    { *m = ST_GS2US_NameList{} }
func (m *ST_GS2US_NameList) String() string            { return proto.CompactTextString(m) }
func (*ST_GS2US_NameList) ProtoMessage()               {}
func (*ST_GS2US_NameList) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *ST_GS2US_NameList) GetServerID() uint32 {
	if m != nil {
		return m.ServerID
	}
	return 0
}

func (m *ST_GS2US_NameList) GetUserNames() []string {
	if m != nil {
		return m.UserNames
	}
	return nil
}

func (m *ST_GS2US_NameList) GetLeagueNames() []*ST_LeagueName_PB {
	if m != nil {
		return m.LeagueNames
	}
	return nil
}

type ST_GS2US_ServerCheck struct {
	ServerID uint32 `protobuf:"varint,1,opt,name=serverID" json:"serverID,omitempty"`
}

func (m *ST_GS2US_ServerCheck) Reset()                    { *m = ST_GS2US_ServerCheck{} }
func (m *ST_GS2US_ServerCheck) String() string            { return proto.CompactTextString(m) }
func (*ST_GS2US_ServerCheck) ProtoMessage()               {}
func (*ST_GS2US_ServerCheck) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

func (m *ST_GS2US_ServerCheck) GetServerID() uint32 {
	if m != nil {
		return m.ServerID
	}
	return 0
}

func init() {
	proto.RegisterType((*ST_GS2US_UserName_Check)(nil), "proto_micro.ST_GS2US_UserName_Check")
	proto.RegisterType((*ST_GS2US_UserName_Change)(nil), "proto_micro.ST_GS2US_UserName_Change")
	proto.RegisterType((*ST_GS2US_LeagueName_Check)(nil), "proto_micro.ST_GS2US_LeagueName_Check")
	proto.RegisterType((*ST_GS2US_LeagueName_Change)(nil), "proto_micro.ST_GS2US_LeagueName_Change")
	proto.RegisterType((*ST_GS2US_LeagueName_Delete)(nil), "proto_micro.ST_GS2US_LeagueName_Delete")
	proto.RegisterType((*ST_LeagueName_PB)(nil), "proto_micro.ST_LeagueName_PB")
	proto.RegisterType((*ST_GS2US_NameList)(nil), "proto_micro.ST_GS2US_NameList")
	proto.RegisterType((*ST_GS2US_ServerCheck)(nil), "proto_micro.ST_GS2US_ServerCheck")
}

func init() { proto.RegisterFile("MG_UniqueProtocol.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x51, 0x4f, 0xc2, 0x30,
	0x14, 0x85, 0x1d, 0x05, 0x94, 0x4b, 0x4c, 0xb4, 0x21, 0x32, 0x89, 0x9a, 0xa5, 0x4f, 0x7b, 0xe2,
	0x01, 0x7f, 0x80, 0x09, 0xa2, 0x60, 0x82, 0x86, 0xac, 0xf0, 0x68, 0x16, 0x24, 0x37, 0xb8, 0x38,
	0x36, 0x6d, 0x37, 0x79, 0xf1, 0xc1, 0x5f, 0xe0, 0x6f, 0x36, 0xeb, 0x46, 0x2d, 0x66, 0x6a, 0x7c,
	0x5a, 0xcf, 0xb9, 0xd9, 0x77, 0xce, 0x4d, 0x0b, 0xed, 0xdb, 0xa1, 0x3f, 0x8b, 0x82, 0x97, 0x14,
	0x27, 0x22, 0x4e, 0xe2, 0x45, 0x1c, 0x76, 0x9f, 0xb3, 0x03, 0x6d, 0xaa, 0x8f, 0xbf, 0x0a, 0x16,
	0x22, 0x66, 0x57, 0xd0, 0xe6, 0x53, 0x7f, 0xc8, 0x7b, 0x33, 0xee, 0xcf, 0x24, 0x8a, 0xbb, 0xf9,
	0x0a, 0xfd, 0xcb, 0x47, 0x5c, 0x3c, 0x51, 0x0a, 0xd5, 0x4c, 0xd9, 0x96, 0x63, 0xb9, 0x0d, 0x4f,
	0x9d, 0xe9, 0x11, 0xd4, 0x05, 0xca, 0x34, 0x4c, 0xec, 0x8a, 0x63, 0xb9, 0x35, 0xaf, 0x50, 0x6c,
	0x04, 0x76, 0x19, 0x66, 0x1e, 0x2d, 0xb1, 0x94, 0x63, 0xc3, 0x6e, 0x84, 0x6b, 0x65, 0x57, 0x94,
	0xbd, 0x91, 0xec, 0x1e, 0x8e, 0x35, 0x69, 0x8c, 0xf3, 0x65, 0x8a, 0x7f, 0x54, 0x6a, 0x41, 0x4d,
	0x1a, 0xa0, 0x5c, 0x18, 0x45, 0xc9, 0x56, 0xd1, 0x37, 0xe8, 0x94, 0xe3, 0x7f, 0xac, 0x5a, 0xce,
	0x37, 0x16, 0x20, 0x5b, 0x0b, 0xd0, 0x0e, 0xec, 0x45, 0xb8, 0xe6, 0x6a, 0x54, 0x55, 0x23, 0xad,
	0xd9, 0x75, 0x79, 0xfa, 0x00, 0x43, 0x4c, 0xfe, 0x91, 0xce, 0xa6, 0x70, 0xc0, 0xa7, 0x26, 0x61,
	0xd2, 0xa7, 0x67, 0x00, 0xa1, 0x36, 0x0a, 0x86, 0xe1, 0x50, 0x07, 0x9a, 0xb9, 0x1a, 0x1b, 0x3c,
	0xd3, 0x62, 0x1f, 0x16, 0x1c, 0xea, 0x7a, 0x99, 0x33, 0x0e, 0x64, 0x92, 0xed, 0x23, 0x51, 0xbc,
	0xa2, 0xb8, 0x19, 0x28, 0xea, 0xbe, 0xa7, 0x35, 0x3d, 0x81, 0x46, 0x5a, 0xdc, 0xb6, 0xb4, 0x2b,
	0x0e, 0x71, 0x1b, 0xde, 0x97, 0x41, 0x2f, 0x36, 0x89, 0xf9, 0x9c, 0x38, 0xc4, 0x6d, 0xf6, 0x4e,
	0xbb, 0xc6, 0xf3, 0xeb, 0x7e, 0xdf, 0xc2, 0x33, 0xff, 0x60, 0x3d, 0x68, 0xe9, 0x3e, 0x5c, 0x65,
	0xe6, 0xcf, 0xe0, 0x97, 0x4a, 0xfd, 0xda, 0x88, 0xbc, 0x5b, 0x3b, 0x0f, 0x75, 0x95, 0x72, 0xfe,
	0x19, 0x00, 0x00, 0xff, 0xff, 0xba, 0x15, 0x49, 0xcb, 0x06, 0x03, 0x00, 0x00,
}