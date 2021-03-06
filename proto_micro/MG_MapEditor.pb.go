// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MG_MapEditor.proto

package proto_micro

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 主城地图信息
type ST_MapData_PB struct {
	CityBorderX     uint32              `protobuf:"varint,1,opt,name=cityBorderX" json:"cityBorderX,omitempty"`
	CityBorderY     uint32              `protobuf:"varint,2,opt,name=cityBorderY" json:"cityBorderY,omitempty"`
	ObstacleList    []*ST_Vector2Int_PB `protobuf:"bytes,3,rep,name=obstacleList" json:"obstacleList,omitempty"`
	UnreachableList []*ST_Vector2Int_PB `protobuf:"bytes,4,rep,name=unreachableList" json:"unreachableList,omitempty"`
}

func (m *ST_MapData_PB) Reset()                    { *m = ST_MapData_PB{} }
func (m *ST_MapData_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_MapData_PB) ProtoMessage()               {}
func (*ST_MapData_PB) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *ST_MapData_PB) GetCityBorderX() uint32 {
	if m != nil {
		return m.CityBorderX
	}
	return 0
}

func (m *ST_MapData_PB) GetCityBorderY() uint32 {
	if m != nil {
		return m.CityBorderY
	}
	return 0
}

func (m *ST_MapData_PB) GetObstacleList() []*ST_Vector2Int_PB {
	if m != nil {
		return m.ObstacleList
	}
	return nil
}

func (m *ST_MapData_PB) GetUnreachableList() []*ST_Vector2Int_PB {
	if m != nil {
		return m.UnreachableList
	}
	return nil
}

// 大地图格子
type ST_TileData_PB struct {
	Px             uint32   `protobuf:"varint,1,opt,name=px" json:"px,omitempty"`
	Py             uint32   `protobuf:"varint,2,opt,name=py" json:"py,omitempty"`
	IsAvaliable    bool     `protobuf:"varint,3,opt,name=isAvaliable" json:"isAvaliable,omitempty"`
	ResourceLevel  uint32   `protobuf:"varint,4,opt,name=resourceLevel" json:"resourceLevel,omitempty"`
	BelongCityList []uint32 `protobuf:"varint,5,rep,packed,name=belongCityList" json:"belongCityList,omitempty"`
}

func (m *ST_TileData_PB) Reset()                    { *m = ST_TileData_PB{} }
func (m *ST_TileData_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_TileData_PB) ProtoMessage()               {}
func (*ST_TileData_PB) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *ST_TileData_PB) GetPx() uint32 {
	if m != nil {
		return m.Px
	}
	return 0
}

func (m *ST_TileData_PB) GetPy() uint32 {
	if m != nil {
		return m.Py
	}
	return 0
}

func (m *ST_TileData_PB) GetIsAvaliable() bool {
	if m != nil {
		return m.IsAvaliable
	}
	return false
}

func (m *ST_TileData_PB) GetResourceLevel() uint32 {
	if m != nil {
		return m.ResourceLevel
	}
	return 0
}

func (m *ST_TileData_PB) GetBelongCityList() []uint32 {
	if m != nil {
		return m.BelongCityList
	}
	return nil
}

// 单个地块内装饰物信息
type ST_BlockDecoList_PB struct {
	BlockName string            `protobuf:"bytes,1,opt,name=blockName" json:"blockName,omitempty"`
	DecoList  []*ST_DecoData_PB `protobuf:"bytes,2,rep,name=decoList" json:"decoList,omitempty"`
}

func (m *ST_BlockDecoList_PB) Reset()                    { *m = ST_BlockDecoList_PB{} }
func (m *ST_BlockDecoList_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_BlockDecoList_PB) ProtoMessage()               {}
func (*ST_BlockDecoList_PB) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *ST_BlockDecoList_PB) GetBlockName() string {
	if m != nil {
		return m.BlockName
	}
	return ""
}

func (m *ST_BlockDecoList_PB) GetDecoList() []*ST_DecoData_PB {
	if m != nil {
		return m.DecoList
	}
	return nil
}

// 装饰物
type ST_DecoData_PB struct {
	DecoName   string `protobuf:"bytes,1,opt,name=decoName" json:"decoName,omitempty"`
	Px         uint32 `protobuf:"varint,2,opt,name=px" json:"px,omitempty"`
	Py         uint32 `protobuf:"varint,3,opt,name=py" json:"py,omitempty"`
	Lx         uint32 `protobuf:"varint,4,opt,name=lx" json:"lx,omitempty"`
	Ly         uint32 `protobuf:"varint,5,opt,name=ly" json:"ly,omitempty"`
	Sx         uint32 `protobuf:"varint,6,opt,name=sx" json:"sx,omitempty"`
	Sy         uint32 `protobuf:"varint,7,opt,name=sy" json:"sy,omitempty"`
	Sz         uint32 `protobuf:"varint,8,opt,name=sz" json:"sz,omitempty"`
	Rotation   uint32 `protobuf:"varint,9,opt,name=rotation" json:"rotation,omitempty"`
	IsObstacle bool   `protobuf:"varint,10,opt,name=isObstacle" json:"isObstacle,omitempty"`
}

func (m *ST_DecoData_PB) Reset()                    { *m = ST_DecoData_PB{} }
func (m *ST_DecoData_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_DecoData_PB) ProtoMessage()               {}
func (*ST_DecoData_PB) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *ST_DecoData_PB) GetDecoName() string {
	if m != nil {
		return m.DecoName
	}
	return ""
}

func (m *ST_DecoData_PB) GetPx() uint32 {
	if m != nil {
		return m.Px
	}
	return 0
}

func (m *ST_DecoData_PB) GetPy() uint32 {
	if m != nil {
		return m.Py
	}
	return 0
}

func (m *ST_DecoData_PB) GetLx() uint32 {
	if m != nil {
		return m.Lx
	}
	return 0
}

func (m *ST_DecoData_PB) GetLy() uint32 {
	if m != nil {
		return m.Ly
	}
	return 0
}

func (m *ST_DecoData_PB) GetSx() uint32 {
	if m != nil {
		return m.Sx
	}
	return 0
}

func (m *ST_DecoData_PB) GetSy() uint32 {
	if m != nil {
		return m.Sy
	}
	return 0
}

func (m *ST_DecoData_PB) GetSz() uint32 {
	if m != nil {
		return m.Sz
	}
	return 0
}

func (m *ST_DecoData_PB) GetRotation() uint32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

func (m *ST_DecoData_PB) GetIsObstacle() bool {
	if m != nil {
		return m.IsObstacle
	}
	return false
}

// 大地图城市信息
type ST_WorldCityData_PB struct {
	CityID          uint32 `protobuf:"varint,1,opt,name=cityID" json:"cityID,omitempty"`
	CityLevel       uint32 `protobuf:"varint,2,opt,name=cityLevel" json:"cityLevel,omitempty"`
	Px              uint32 `protobuf:"varint,3,opt,name=px" json:"px,omitempty"`
	Py              uint32 `protobuf:"varint,4,opt,name=py" json:"py,omitempty"`
	ResAreaX        int32  `protobuf:"varint,5,opt,name=resAreaX" json:"resAreaX,omitempty"`
	ResAreaY        int32  `protobuf:"varint,6,opt,name=resAreaY" json:"resAreaY,omitempty"`
	ResDeltaX       int32  `protobuf:"varint,7,opt,name=resDeltaX" json:"resDeltaX,omitempty"`
	ResDeltaY       int32  `protobuf:"varint,8,opt,name=resDeltaY" json:"resDeltaY,omitempty"`
	TerritoryAreaX  int32  `protobuf:"varint,9,opt,name=territoryAreaX" json:"territoryAreaX,omitempty"`
	TerritoryAreaY  int32  `protobuf:"varint,10,opt,name=territoryAreaY" json:"territoryAreaY,omitempty"`
	TerritoryDeltaX int32  `protobuf:"varint,11,opt,name=territoryDeltaX" json:"territoryDeltaX,omitempty"`
	TerritoryDeltaY int32  `protobuf:"varint,12,opt,name=territoryDeltaY" json:"territoryDeltaY,omitempty"`
}

func (m *ST_WorldCityData_PB) Reset()                    { *m = ST_WorldCityData_PB{} }
func (m *ST_WorldCityData_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_WorldCityData_PB) ProtoMessage()               {}
func (*ST_WorldCityData_PB) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *ST_WorldCityData_PB) GetCityID() uint32 {
	if m != nil {
		return m.CityID
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetCityLevel() uint32 {
	if m != nil {
		return m.CityLevel
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetPx() uint32 {
	if m != nil {
		return m.Px
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetPy() uint32 {
	if m != nil {
		return m.Py
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetResAreaX() int32 {
	if m != nil {
		return m.ResAreaX
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetResAreaY() int32 {
	if m != nil {
		return m.ResAreaY
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetResDeltaX() int32 {
	if m != nil {
		return m.ResDeltaX
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetResDeltaY() int32 {
	if m != nil {
		return m.ResDeltaY
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetTerritoryAreaX() int32 {
	if m != nil {
		return m.TerritoryAreaX
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetTerritoryAreaY() int32 {
	if m != nil {
		return m.TerritoryAreaY
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetTerritoryDeltaX() int32 {
	if m != nil {
		return m.TerritoryDeltaX
	}
	return 0
}

func (m *ST_WorldCityData_PB) GetTerritoryDeltaY() int32 {
	if m != nil {
		return m.TerritoryDeltaY
	}
	return 0
}

// 大地图信息
type ST_WorldMapData_PB struct {
	WorldBorderX uint32                 `protobuf:"varint,1,opt,name=worldBorderX" json:"worldBorderX,omitempty"`
	WorldBorderY uint32                 `protobuf:"varint,2,opt,name=worldBorderY" json:"worldBorderY,omitempty"`
	BlockBorderX uint32                 `protobuf:"varint,3,opt,name=blockBorderX" json:"blockBorderX,omitempty"`
	BlockBorderY uint32                 `protobuf:"varint,4,opt,name=blockBorderY" json:"blockBorderY,omitempty"`
	TileList     []*ST_TileData_PB      `protobuf:"bytes,5,rep,name=tileList" json:"tileList,omitempty"`
	DecoList     []*ST_DecoData_PB      `protobuf:"bytes,6,rep,name=decoList" json:"decoList,omitempty"`
	CityList     []*ST_WorldCityData_PB `protobuf:"bytes,7,rep,name=cityList" json:"cityList,omitempty"`
	BlockList    []*ST_BlockDecoList_PB `protobuf:"bytes,8,rep,name=blockList" json:"blockList,omitempty"`
}

func (m *ST_WorldMapData_PB) Reset()                    { *m = ST_WorldMapData_PB{} }
func (m *ST_WorldMapData_PB) String() string            { return proto.CompactTextString(m) }
func (*ST_WorldMapData_PB) ProtoMessage()               {}
func (*ST_WorldMapData_PB) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *ST_WorldMapData_PB) GetWorldBorderX() uint32 {
	if m != nil {
		return m.WorldBorderX
	}
	return 0
}

func (m *ST_WorldMapData_PB) GetWorldBorderY() uint32 {
	if m != nil {
		return m.WorldBorderY
	}
	return 0
}

func (m *ST_WorldMapData_PB) GetBlockBorderX() uint32 {
	if m != nil {
		return m.BlockBorderX
	}
	return 0
}

func (m *ST_WorldMapData_PB) GetBlockBorderY() uint32 {
	if m != nil {
		return m.BlockBorderY
	}
	return 0
}

func (m *ST_WorldMapData_PB) GetTileList() []*ST_TileData_PB {
	if m != nil {
		return m.TileList
	}
	return nil
}

func (m *ST_WorldMapData_PB) GetDecoList() []*ST_DecoData_PB {
	if m != nil {
		return m.DecoList
	}
	return nil
}

func (m *ST_WorldMapData_PB) GetCityList() []*ST_WorldCityData_PB {
	if m != nil {
		return m.CityList
	}
	return nil
}

func (m *ST_WorldMapData_PB) GetBlockList() []*ST_BlockDecoList_PB {
	if m != nil {
		return m.BlockList
	}
	return nil
}

func init() {
	proto.RegisterType((*ST_MapData_PB)(nil), "proto_micro.ST_MapData_PB")
	proto.RegisterType((*ST_TileData_PB)(nil), "proto_micro.ST_TileData_PB")
	proto.RegisterType((*ST_BlockDecoList_PB)(nil), "proto_micro.ST_BlockDecoList_PB")
	proto.RegisterType((*ST_DecoData_PB)(nil), "proto_micro.ST_DecoData_PB")
	proto.RegisterType((*ST_WorldCityData_PB)(nil), "proto_micro.ST_WorldCityData_PB")
	proto.RegisterType((*ST_WorldMapData_PB)(nil), "proto_micro.ST_WorldMapData_PB")
}

func init() { proto.RegisterFile("MG_MapEditor.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0xd9, 0xae, 0xd3, 0x64, 0xd3, 0xa4, 0xd2, 0xfe, 0x24, 0x58, 0x95, 0x3f, 0xb2, 0x2c,
	0x84, 0x72, 0xea, 0xa1, 0x1c, 0xb8, 0x20, 0xa4, 0xa6, 0x41, 0xa5, 0x52, 0x03, 0x91, 0x13, 0x41,
	0x7c, 0xb2, 0x1c, 0x67, 0x05, 0x16, 0x9b, 0x6c, 0xb4, 0xde, 0x96, 0xb8, 0x27, 0x24, 0x5e, 0x81,
	0x07, 0xe0, 0xb1, 0xb8, 0xf1, 0x2a, 0x68, 0xc7, 0x6b, 0xc7, 0xde, 0xe6, 0x00, 0xa7, 0x64, 0xbe,
	0xf9, 0x66, 0xbd, 0xdf, 0x37, 0x3b, 0x83, 0xf0, 0xf8, 0x32, 0x1a, 0xc7, 0x9b, 0x37, 0xcb, 0x54,
	0x72, 0x71, 0xba, 0x11, 0x5c, 0x72, 0xdc, 0x85, 0x9f, 0x68, 0x95, 0x26, 0x82, 0x9f, 0x3c, 0x1c,
	0x5f, 0x46, 0x17, 0x7c, 0xb5, 0xe2, 0xeb, 0x89, 0x42, 0x13, 0xce, 0x0a, 0x96, 0xff, 0xdb, 0x42,
	0xbd, 0xe9, 0x4c, 0x15, 0x8f, 0x62, 0x19, 0x47, 0x93, 0x21, 0xf6, 0x50, 0x37, 0x49, 0x65, 0x3e,
	0xe4, 0x62, 0x49, 0xc5, 0x9c, 0x58, 0x9e, 0x35, 0xe8, 0x05, 0x75, 0xa8, 0xc9, 0x08, 0x89, 0x6d,
	0x32, 0x42, 0x7c, 0x8e, 0x8e, 0xf8, 0x22, 0x93, 0x71, 0xc2, 0xe8, 0x75, 0x9a, 0x49, 0xe2, 0x78,
	0xce, 0xa0, 0x7b, 0xf6, 0xe4, 0xb4, 0x76, 0xa5, 0xd3, 0xe9, 0x2c, 0xfa, 0x40, 0x13, 0xc9, 0xc5,
	0xd9, 0xd5, 0x5a, 0x46, 0x93, 0x61, 0xd0, 0x28, 0xc1, 0x97, 0xe8, 0xf8, 0x66, 0x2d, 0x68, 0x9c,
	0x7c, 0x8e, 0x17, 0xfa, 0x94, 0x83, 0xbf, 0x39, 0xc5, 0xac, 0xf2, 0x7f, 0x5a, 0xa8, 0x3f, 0x9d,
	0x45, 0xb3, 0x94, 0xd1, 0x52, 0x62, 0x1f, 0xd9, 0x9b, 0xad, 0x56, 0x66, 0x6f, 0xb6, 0x10, 0xe7,
	0x5a, 0x87, 0xbd, 0xc9, 0x95, 0xc0, 0x34, 0x3b, 0xbf, 0x8d, 0x59, 0xaa, 0x4e, 0x21, 0x8e, 0x67,
	0x0d, 0xda, 0x41, 0x1d, 0xc2, 0xcf, 0x50, 0x4f, 0xd0, 0x8c, 0xdf, 0x88, 0x84, 0x5e, 0xd3, 0x5b,
	0xca, 0xc8, 0x01, 0x14, 0x37, 0x41, 0xfc, 0x1c, 0xf5, 0x17, 0x94, 0xf1, 0xf5, 0xa7, 0x8b, 0x54,
	0xe6, 0x20, 0xc1, 0xf5, 0x9c, 0x41, 0x2f, 0x30, 0x50, 0x9f, 0xa1, 0xff, 0xa7, 0xb3, 0x68, 0xc8,
	0x78, 0xf2, 0x65, 0x44, 0x13, 0xae, 0x30, 0x75, 0xcd, 0xc7, 0xa8, 0xb3, 0x50, 0xd8, 0xbb, 0x78,
	0x45, 0xe1, 0xb6, 0x9d, 0x60, 0x07, 0xe0, 0x97, 0xa8, 0xbd, 0xd4, 0x64, 0x62, 0x83, 0x33, 0x8f,
	0x4c, 0x67, 0xd4, 0x61, 0x5a, 0x73, 0x50, 0x91, 0xfd, 0x5f, 0x85, 0x21, 0xb5, 0x24, 0x3e, 0x29,
	0xce, 0xaa, 0x7d, 0xa8, 0x8a, 0xb5, 0x59, 0xb6, 0x61, 0x96, 0x53, 0x99, 0xd5, 0x47, 0x36, 0xdb,
	0x6a, 0xfd, 0x36, 0x83, 0x3c, 0xcb, 0x89, 0xab, 0x63, 0xc8, 0x67, 0x5b, 0xd2, 0x2a, 0xe2, 0x0c,
	0xf2, 0x59, 0x4e, 0x0e, 0x75, 0x5c, 0xe4, 0xef, 0x48, 0x5b, 0xc7, 0x77, 0xea, 0x2e, 0x82, 0xcb,
	0x58, 0xa6, 0x7c, 0x4d, 0x3a, 0x80, 0x56, 0x31, 0x7e, 0x8a, 0x50, 0x9a, 0xbd, 0xd7, 0xcf, 0x84,
	0x20, 0xe8, 0x4b, 0x0d, 0xf1, 0xbf, 0x3b, 0xe0, 0xe4, 0x47, 0x2e, 0xd8, 0x52, 0xb9, 0x5b, 0xea,
	0x7b, 0x80, 0x5a, 0xea, 0x79, 0x5e, 0x8d, 0x74, 0xd3, 0x75, 0xa4, 0x1c, 0x56, 0xff, 0x8a, 0x16,
	0x16, 0x12, 0x77, 0x80, 0x56, 0xee, 0x18, 0xca, 0x0f, 0x2a, 0xe5, 0xea, 0xa6, 0x34, 0x3b, 0x17,
	0x34, 0x9e, 0x83, 0x5e, 0x37, 0xa8, 0xe2, 0x5a, 0x2e, 0x04, 0xed, 0xbb, 0x5c, 0xa8, 0xbe, 0x2a,
	0x68, 0x36, 0xa2, 0x4c, 0xc6, 0x73, 0x30, 0xc2, 0x0d, 0x76, 0x40, 0x3d, 0x1b, 0x82, 0x2d, 0xb5,
	0x6c, 0xa8, 0x9e, 0x94, 0xa4, 0x42, 0xa8, 0x39, 0xcf, 0x8b, 0x2f, 0x77, 0x80, 0x62, 0xa0, 0xf7,
	0x78, 0x21, 0xb8, 0x65, 0xf2, 0x42, 0x3c, 0x40, 0xc7, 0x15, 0xa2, 0x6f, 0xd4, 0x05, 0xa2, 0x09,
	0xdf, 0x67, 0x86, 0xe4, 0x68, 0x1f, 0x33, 0xf4, 0x7f, 0x38, 0x08, 0x97, 0x5d, 0xa8, 0x2d, 0x16,
	0x1f, 0x1d, 0x7d, 0x55, 0x50, 0x73, 0xb3, 0x34, 0x30, 0x83, 0x53, 0xee, 0x96, 0x06, 0xa6, 0x38,
	0x30, 0x05, 0xe5, 0x39, 0x45, 0x83, 0x1a, 0x98, 0xc1, 0x09, 0x75, 0xd3, 0x1a, 0x98, 0x1a, 0x20,
	0x99, 0xea, 0xd5, 0xe2, 0xee, 0x1f, 0xa0, 0xda, 0xd2, 0x08, 0x2a, 0x72, 0x63, 0xf2, 0x5a, 0xff,
	0x30, 0x79, 0xf8, 0x15, 0x6a, 0x27, 0xe5, 0x26, 0x38, 0x84, 0x42, 0xcf, 0x2c, 0x34, 0x9f, 0x6e,
	0x50, 0x55, 0xe0, 0xd7, 0x7a, 0x1d, 0x40, 0x79, 0x7b, 0x7f, 0xb9, 0xb9, 0x43, 0x82, 0x5d, 0xc9,
	0xd0, 0x7d, 0xeb, 0x7c, 0xb3, 0xfe, 0x5b, 0xb4, 0xa0, 0xe4, 0xc5, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x92, 0xc5, 0x3d, 0xe4, 0x34, 0x06, 0x00, 0x00,
}
