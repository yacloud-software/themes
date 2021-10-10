// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/sensorapi/sensorapi.proto
// DO NOT EDIT!

/*
Package sensorapi is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/sensorapi/sensorapi.proto

It has these top-level messages:
	RecentSensorValueRequest
	SensorValueResponse
*/
package sensorapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import singingcat "golang.singingcat.net/apis/singingcat"
import scweb "golang.singingcat.net/apis/scweb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//  get a recent sensor value.
type RecentSensorValueRequest struct {
	SensorID uint64 `protobuf:"varint,1,opt,name=SensorID" json:"SensorID,omitempty"`
}

func (m *RecentSensorValueRequest) Reset()                    { *m = RecentSensorValueRequest{} }
func (m *RecentSensorValueRequest) String() string            { return proto.CompactTextString(m) }
func (*RecentSensorValueRequest) ProtoMessage()               {}
func (*RecentSensorValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RecentSensorValueRequest) GetSensorID() uint64 {
	if m != nil {
		return m.SensorID
	}
	return 0
}

type SensorValueResponse struct {
	Value     float64               `protobuf:"fixed64,1,opt,name=Value" json:"Value,omitempty"`
	Type      singingcat.SensorType `protobuf:"varint,2,opt,name=Type,enum=singingcat.SensorType" json:"Type,omitempty"`
	Timestamp uint32                `protobuf:"varint,3,opt,name=Timestamp" json:"Timestamp,omitempty"`
	Unit      *scweb.SensorUnit     `protobuf:"bytes,4,opt,name=Unit" json:"Unit,omitempty"`
}

func (m *SensorValueResponse) Reset()                    { *m = SensorValueResponse{} }
func (m *SensorValueResponse) String() string            { return proto.CompactTextString(m) }
func (*SensorValueResponse) ProtoMessage()               {}
func (*SensorValueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SensorValueResponse) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *SensorValueResponse) GetType() singingcat.SensorType {
	if m != nil {
		return m.Type
	}
	return singingcat.SensorType_SENSORTYPE_UNDEFINED
}

func (m *SensorValueResponse) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SensorValueResponse) GetUnit() *scweb.SensorUnit {
	if m != nil {
		return m.Unit
	}
	return nil
}

func init() {
	proto.RegisterType((*RecentSensorValueRequest)(nil), "sensorapi.RecentSensorValueRequest")
	proto.RegisterType((*SensorValueResponse)(nil), "sensorapi.SensorValueResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SensorAPIService service

type SensorAPIServiceClient interface {
	// get the most recent sensorvalue for a sensor
	GetMostRecent(ctx context.Context, in *RecentSensorValueRequest, opts ...grpc.CallOption) (*SensorValueResponse, error)
}

type sensorAPIServiceClient struct {
	cc *grpc.ClientConn
}

func NewSensorAPIServiceClient(cc *grpc.ClientConn) SensorAPIServiceClient {
	return &sensorAPIServiceClient{cc}
}

func (c *sensorAPIServiceClient) GetMostRecent(ctx context.Context, in *RecentSensorValueRequest, opts ...grpc.CallOption) (*SensorValueResponse, error) {
	out := new(SensorValueResponse)
	err := grpc.Invoke(ctx, "/sensorapi.SensorAPIService/GetMostRecent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SensorAPIService service

type SensorAPIServiceServer interface {
	// get the most recent sensorvalue for a sensor
	GetMostRecent(context.Context, *RecentSensorValueRequest) (*SensorValueResponse, error)
}

func RegisterSensorAPIServiceServer(s *grpc.Server, srv SensorAPIServiceServer) {
	s.RegisterService(&_SensorAPIService_serviceDesc, srv)
}

func _SensorAPIService_GetMostRecent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecentSensorValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorAPIServiceServer).GetMostRecent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensorapi.SensorAPIService/GetMostRecent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorAPIServiceServer).GetMostRecent(ctx, req.(*RecentSensorValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SensorAPIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sensorapi.SensorAPIService",
	HandlerType: (*SensorAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMostRecent",
			Handler:    _SensorAPIService_GetMostRecent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.singingcat.net/apis/sensorapi/sensorapi.proto",
}

func init() {
	proto.RegisterFile("golang.singingcat.net/apis/sensorapi/sensorapi.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0x5f, 0x4b, 0xfb, 0x30,
	0x14, 0x25, 0xbf, 0x5f, 0x15, 0x77, 0x65, 0xa2, 0x51, 0xa4, 0x14, 0x91, 0x32, 0xff, 0x50, 0x04,
	0x3b, 0xa8, 0xb2, 0x77, 0x45, 0x90, 0x3d, 0x08, 0x92, 0x55, 0xdf, 0xd3, 0x7a, 0xa9, 0x81, 0x2d,
	0x89, 0x4d, 0xa6, 0xf8, 0x5d, 0xfc, 0xb0, 0xd2, 0x64, 0xb4, 0x15, 0x74, 0x2f, 0xe5, 0xde, 0x73,
	0xef, 0x39, 0xcd, 0x39, 0x17, 0xae, 0x2b, 0x35, 0xe7, 0xb2, 0x4a, 0x8d, 0x90, 0x95, 0x90, 0x55,
	0xc9, 0x6d, 0x2a, 0xd1, 0x8e, 0xb9, 0x16, 0x66, 0x6c, 0x50, 0x1a, 0x55, 0x73, 0x2d, 0xba, 0x2a,
	0xd5, 0xb5, 0xb2, 0x8a, 0x0e, 0x5a, 0x20, 0x9a, 0xac, 0x13, 0x68, 0xb1, 0x5e, 0xe9, 0x25, 0xa2,
	0xf1, 0x3a, 0x5e, 0xf9, 0x81, 0x85, 0xff, 0x5e, 0xbe, 0x14, 0x9e, 0x30, 0x9a, 0x40, 0xc8, 0xb0,
	0x44, 0x69, 0x67, 0xee, 0xdf, 0xcf, 0x7c, 0xbe, 0x44, 0x86, 0x6f, 0x4b, 0x34, 0x96, 0x46, 0xb0,
	0xe5, 0xd1, 0xe9, 0x5d, 0x48, 0x62, 0x92, 0x04, 0xac, 0xed, 0x47, 0x5f, 0x04, 0xf6, 0x7f, 0x50,
	0x8c, 0x56, 0xd2, 0x20, 0x3d, 0x80, 0x0d, 0x07, 0x38, 0x02, 0x61, 0xbe, 0xa1, 0x17, 0x10, 0xe4,
	0x9f, 0x1a, 0xc3, 0x7f, 0x31, 0x49, 0x76, 0xb2, 0xc3, 0xfe, 0xf3, 0xbc, 0x48, 0x33, 0x65, 0x6e,
	0x87, 0x1e, 0xc1, 0x20, 0x17, 0x0b, 0x34, 0x96, 0x2f, 0x74, 0xf8, 0x3f, 0x26, 0xc9, 0x90, 0x75,
	0x00, 0x3d, 0x83, 0xe0, 0x49, 0x0a, 0x1b, 0x06, 0x31, 0x49, 0xb6, 0xb3, 0xbd, 0xd4, 0xd9, 0x59,
	0x89, 0x34, 0x03, 0xe6, 0xc6, 0xd9, 0x2b, 0xec, 0x7a, 0xec, 0xe6, 0x71, 0x3a, 0xc3, 0xfa, 0x5d,
	0x94, 0x48, 0x73, 0x18, 0xde, 0xa3, 0x7d, 0x50, 0xc6, 0x7a, 0xc7, 0xf4, 0x24, 0xed, 0x2e, 0xf0,
	0x57, 0x08, 0xd1, 0x71, 0x6f, 0xe9, 0x17, 0xc3, 0xb7, 0xe7, 0x70, 0x2a, 0xd1, 0xf6, 0x1d, 0xad,
	0x4e, 0xd0, 0x64, 0xde, 0x71, 0x8b, 0x4d, 0x97, 0xf7, 0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x55, 0xfa, 0x92, 0xe3, 0x1b, 0x02, 0x00, 0x00,
}
