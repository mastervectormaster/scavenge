// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scavenge/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("scavenge/tx.proto", fileDescriptor_62ec6f1544538068) }

var fileDescriptor_62ec6f1544538068 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4e, 0x4e, 0x2c,
	0x4b, 0xcd, 0x4b, 0x4f, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xc9,
	0x4d, 0x2c, 0x2e, 0x49, 0x2d, 0x2a, 0x4b, 0x4d, 0x2e, 0xc9, 0x2f, 0x82, 0xb0, 0xf5, 0x60, 0xaa,
	0xe0, 0x0c, 0x23, 0x56, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa7, 0xc0, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0xc7, 0x34, 0x51, 0x1f, 0x6e, 0x6f, 0x05, 0x82, 0x59, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0x76, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x78, 0xb8, 0x97, 0x35, 0x9b, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mastervectormaster.scavenge.scavenge.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "scavenge/tx.proto",
}
