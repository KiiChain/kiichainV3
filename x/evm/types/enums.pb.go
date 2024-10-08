// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/enums.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PointerType int32

const (
	PointerType_ERC20  PointerType = 0
	PointerType_ERC721 PointerType = 1
	PointerType_NATIVE PointerType = 2
	PointerType_CW20   PointerType = 3
	PointerType_CW721  PointerType = 4
)

var PointerType_name = map[int32]string{
	0: "ERC20",
	1: "ERC721",
	2: "NATIVE",
	3: "CW20",
	4: "CW721",
}

var PointerType_value = map[string]int32{
	"ERC20":  0,
	"ERC721": 1,
	"NATIVE": 2,
	"CW20":   3,
	"CW721":  4,
}

func (x PointerType) String() string {
	return proto.EnumName(PointerType_name, int32(x))
}

func (PointerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9ba0923a26222f98, []int{0}
}

func init() {
	proto.RegisterEnum("evm.PointerType", PointerType_name, PointerType_value)
}

func init() { proto.RegisterFile("evm/enums.proto", fileDescriptor_9ba0923a26222f98) }

var fileDescriptor_9ba0923a26222f98 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2d, 0xcb, 0xd5,
	0x4f, 0xcd, 0x2b, 0xcd, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2d, 0xcb,
	0xd5, 0x72, 0xe5, 0xe2, 0x0e, 0xc8, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x0a, 0xa9, 0x2c, 0x48, 0x15,
	0xe2, 0xe4, 0x62, 0x75, 0x0d, 0x72, 0x36, 0x32, 0x10, 0x60, 0x10, 0xe2, 0xe2, 0x62, 0x73, 0x0d,
	0x72, 0x36, 0x37, 0x32, 0x14, 0x60, 0x04, 0xb1, 0xfd, 0x1c, 0x43, 0x3c, 0xc3, 0x5c, 0x05, 0x98,
	0x84, 0x38, 0xb8, 0x58, 0x9c, 0xc3, 0x8d, 0x0c, 0x04, 0x98, 0x41, 0x8a, 0x9d, 0xc3, 0x41, 0x0a,
	0x58, 0x9c, 0x5c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x2b, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0x3b, 0x33, 0xd3, 0x39, 0x23, 0x31,
	0x33, 0x4f, 0x3f, 0x3b, 0x33, 0x33, 0x19, 0xc4, 0x08, 0x33, 0xd6, 0xaf, 0xd0, 0x07, 0xb9, 0xab,
	0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0xf0, 0x80, 0x47, 0xab, 0x00, 0x00, 0x00,
}
