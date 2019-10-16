// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command.proto

package command

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Command int32

const (
	Command_dummy Command = 0
	// Gateway消息号定义[消息号范围 1-10000]
	Command_CSLogin Command = 1
	Command_SCLogin Command = 2
)

var Command_name = map[int32]string{
	0: "dummy",
	1: "CSLogin",
	2: "SCLogin",
}

var Command_value = map[string]int32{
	"dummy":   0,
	"CSLogin": 1,
	"SCLogin": 2,
}

func (x Command) String() string {
	return proto.EnumName(Command_name, int32(x))
}

func (Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_213c0bb044472049, []int{0}
}

func init() {
	proto.RegisterEnum("command.Command", Command_name, Command_value)
}

func init() { proto.RegisterFile("command.proto", fileDescriptor_213c0bb044472049) }

var fileDescriptor_213c0bb044472049 = []byte{
	// 85 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xb5, 0xf4, 0xb8,
	0xd8, 0x9d, 0x21, 0x4c, 0x21, 0x4e, 0x2e, 0xd6, 0x94, 0xd2, 0xdc, 0xdc, 0x4a, 0x01, 0x06, 0x21,
	0x6e, 0x2e, 0x76, 0xe7, 0x60, 0x9f, 0xfc, 0xf4, 0xcc, 0x3c, 0x01, 0x46, 0x10, 0x27, 0xd8, 0x19,
	0xc2, 0x61, 0x4a, 0x62, 0x03, 0xeb, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x9a, 0xbe,
	0x10, 0x50, 0x00, 0x00, 0x00,
}
