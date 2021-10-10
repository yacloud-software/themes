// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/deployminator/deployminator.proto
// DO NOT EDIT!

/*
Package deployminator is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/deployminator/deployminator.proto

It has these top-level messages:
	Version
	Application
	DeploymentDescriptor
	InstanceDef
	Argument
	NewBuildRequest
	DeployRequest
	UndeployRequest
*/
package deployminator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"
import deploymonkey "golang.conradwood.net/apis/deploymonkey"

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

type Version struct {
	ID           uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	BuildNumber  uint64 `protobuf:"varint,2,opt,name=BuildNumber" json:"BuildNumber,omitempty"`
	RepositoryID uint64 `protobuf:"varint,3,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	Branch       string `protobuf:"bytes,4,opt,name=Branch" json:"Branch,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Version) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Version) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *Version) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *Version) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

//
// an application is the definition of a "a binary in a repository".
// it may exist in many different versions
type Application struct {
	ID           uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Binary       string `protobuf:"bytes,2,opt,name=Binary" json:"Binary,omitempty"`
	RepositoryID uint64 `protobuf:"varint,3,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	DownloadURL  string `protobuf:"bytes,4,opt,name=DownloadURL" json:"DownloadURL,omitempty"`
}

func (m *Application) Reset()                    { *m = Application{} }
func (m *Application) String() string            { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()               {}
func (*Application) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Application) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Application) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *Application) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *Application) GetDownloadURL() string {
	if m != nil {
		return m.DownloadURL
	}
	return ""
}

// a deployment descriptor links together all the bits required for a deployment, like instances, args, etc..
type DeploymentDescriptor struct {
	ID          uint64       `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Application *Application `protobuf:"bytes,2,opt,name=Application" json:"Application,omitempty"`
	BuildNumber uint64       `protobuf:"varint,3,opt,name=BuildNumber" json:"BuildNumber,omitempty"`
	Branch      string       `protobuf:"bytes,4,opt,name=Branch" json:"Branch,omitempty"`
}

func (m *DeploymentDescriptor) Reset()                    { *m = DeploymentDescriptor{} }
func (m *DeploymentDescriptor) String() string            { return proto.CompactTextString(m) }
func (*DeploymentDescriptor) ProtoMessage()               {}
func (*DeploymentDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeploymentDescriptor) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *DeploymentDescriptor) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *DeploymentDescriptor) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *DeploymentDescriptor) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

// Define how many instances to run where
type InstanceDef struct {
	ID                        uint64                `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	DeploymentID              *DeploymentDescriptor `protobuf:"bytes,2,opt,name=DeploymentID" json:"DeploymentID,omitempty"`
	MachineGroup              string                `protobuf:"bytes,3,opt,name=MachineGroup" json:"MachineGroup,omitempty"`
	Instances                 uint32                `protobuf:"varint,4,opt,name=Instances" json:"Instances,omitempty"`
	InstanceCountIsPerMachine bool                  `protobuf:"varint,5,opt,name=InstanceCountIsPerMachine" json:"InstanceCountIsPerMachine,omitempty"`
}

func (m *InstanceDef) Reset()                    { *m = InstanceDef{} }
func (m *InstanceDef) String() string            { return proto.CompactTextString(m) }
func (*InstanceDef) ProtoMessage()               {}
func (*InstanceDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InstanceDef) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *InstanceDef) GetDeploymentID() *DeploymentDescriptor {
	if m != nil {
		return m.DeploymentID
	}
	return nil
}

func (m *InstanceDef) GetMachineGroup() string {
	if m != nil {
		return m.MachineGroup
	}
	return ""
}

func (m *InstanceDef) GetInstances() uint32 {
	if m != nil {
		return m.Instances
	}
	return 0
}

func (m *InstanceDef) GetInstanceCountIsPerMachine() bool {
	if m != nil {
		return m.InstanceCountIsPerMachine
	}
	return false
}

type Argument struct {
	ID           uint64                `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	DeploymentID *DeploymentDescriptor `protobuf:"bytes,2,opt,name=DeploymentID" json:"DeploymentID,omitempty"`
	Argument     string                `protobuf:"bytes,3,opt,name=Argument" json:"Argument,omitempty"`
}

func (m *Argument) Reset()                    { *m = Argument{} }
func (m *Argument) String() string            { return proto.CompactTextString(m) }
func (*Argument) ProtoMessage()               {}
func (*Argument) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Argument) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Argument) GetDeploymentID() *DeploymentDescriptor {
	if m != nil {
		return m.DeploymentID
	}
	return nil
}

func (m *Argument) GetArgument() string {
	if m != nil {
		return m.Argument
	}
	return ""
}

type NewBuildRequest struct {
	DeployFile    []byte `protobuf:"bytes,1,opt,name=DeployFile,proto3" json:"DeployFile,omitempty"`
	RepositoryID  uint64 `protobuf:"varint,2,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	Branch        string `protobuf:"bytes,3,opt,name=Branch" json:"Branch,omitempty"`
	BuildNumber   uint64 `protobuf:"varint,4,opt,name=BuildNumber" json:"BuildNumber,omitempty"`
	CommitID      string `protobuf:"bytes,5,opt,name=CommitID" json:"CommitID,omitempty"`
	ArtefactName  string `protobuf:"bytes,6,opt,name=ArtefactName" json:"ArtefactName,omitempty"`
	BuildserverID string `protobuf:"bytes,7,opt,name=BuildserverID" json:"BuildserverID,omitempty"`
}

func (m *NewBuildRequest) Reset()                    { *m = NewBuildRequest{} }
func (m *NewBuildRequest) String() string            { return proto.CompactTextString(m) }
func (*NewBuildRequest) ProtoMessage()               {}
func (*NewBuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NewBuildRequest) GetDeployFile() []byte {
	if m != nil {
		return m.DeployFile
	}
	return nil
}

func (m *NewBuildRequest) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *NewBuildRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *NewBuildRequest) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *NewBuildRequest) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

func (m *NewBuildRequest) GetArtefactName() string {
	if m != nil {
		return m.ArtefactName
	}
	return ""
}

func (m *NewBuildRequest) GetBuildserverID() string {
	if m != nil {
		return m.BuildserverID
	}
	return ""
}

type DeployRequest struct {
	Version uint64   `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Tags    []string `protobuf:"bytes,2,rep,name=Tags" json:"Tags,omitempty"`
}

func (m *DeployRequest) Reset()                    { *m = DeployRequest{} }
func (m *DeployRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()               {}
func (*DeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeployRequest) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *DeployRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type UndeployRequest struct {
	Version uint64   `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Tags    []string `protobuf:"bytes,2,rep,name=Tags" json:"Tags,omitempty"`
}

func (m *UndeployRequest) Reset()                    { *m = UndeployRequest{} }
func (m *UndeployRequest) String() string            { return proto.CompactTextString(m) }
func (*UndeployRequest) ProtoMessage()               {}
func (*UndeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UndeployRequest) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *UndeployRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Version)(nil), "deployminator.Version")
	proto.RegisterType((*Application)(nil), "deployminator.Application")
	proto.RegisterType((*DeploymentDescriptor)(nil), "deployminator.DeploymentDescriptor")
	proto.RegisterType((*InstanceDef)(nil), "deployminator.InstanceDef")
	proto.RegisterType((*Argument)(nil), "deployminator.Argument")
	proto.RegisterType((*NewBuildRequest)(nil), "deployminator.NewBuildRequest")
	proto.RegisterType((*DeployRequest)(nil), "deployminator.DeployRequest")
	proto.RegisterType((*UndeployRequest)(nil), "deployminator.UndeployRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Deployminator service

type DeployminatorClient interface {
	// parse a yamlfile and deploy the new version defined in it
	NewBuildAvailable(ctx context.Context, in *NewBuildRequest, opts ...grpc.CallOption) (*common.Void, error)
	// deploy a specific version
	DeployVersion(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*common.Void, error)
	// deployments
	ListDeployments(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*deploymonkey.DeploymentList, error)
	// stop an application
	UndeployVersion(ctx context.Context, in *UndeployRequest, opts ...grpc.CallOption) (*common.Void, error)
	// called by autodeployer if it startsup
	AutodeployerStartup(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error)
	// called by autodeployer if it shutsdown
	AutodeployerShutdown(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error)
}

type deployminatorClient struct {
	cc *grpc.ClientConn
}

func NewDeployminatorClient(cc *grpc.ClientConn) DeployminatorClient {
	return &deployminatorClient{cc}
}

func (c *deployminatorClient) NewBuildAvailable(ctx context.Context, in *NewBuildRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/deployminator.Deployminator/NewBuildAvailable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployminatorClient) DeployVersion(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/deployminator.Deployminator/DeployVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployminatorClient) ListDeployments(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*deploymonkey.DeploymentList, error) {
	out := new(deploymonkey.DeploymentList)
	err := grpc.Invoke(ctx, "/deployminator.Deployminator/ListDeployments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployminatorClient) UndeployVersion(ctx context.Context, in *UndeployRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/deployminator.Deployminator/UndeployVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployminatorClient) AutodeployerStartup(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/deployminator.Deployminator/AutodeployerStartup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployminatorClient) AutodeployerShutdown(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/deployminator.Deployminator/AutodeployerShutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Deployminator service

type DeployminatorServer interface {
	// parse a yamlfile and deploy the new version defined in it
	NewBuildAvailable(context.Context, *NewBuildRequest) (*common.Void, error)
	// deploy a specific version
	DeployVersion(context.Context, *DeployRequest) (*common.Void, error)
	// deployments
	ListDeployments(context.Context, *common.Void) (*deploymonkey.DeploymentList, error)
	// stop an application
	UndeployVersion(context.Context, *UndeployRequest) (*common.Void, error)
	// called by autodeployer if it startsup
	AutodeployerStartup(context.Context, *common.Void) (*common.Void, error)
	// called by autodeployer if it shutsdown
	AutodeployerShutdown(context.Context, *common.Void) (*common.Void, error)
}

func RegisterDeployminatorServer(s *grpc.Server, srv DeployminatorServer) {
	s.RegisterService(&_Deployminator_serviceDesc, srv)
}

func _Deployminator_NewBuildAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployminatorServer).NewBuildAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployminator.Deployminator/NewBuildAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployminatorServer).NewBuildAvailable(ctx, req.(*NewBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployminator_DeployVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployminatorServer).DeployVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployminator.Deployminator/DeployVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployminatorServer).DeployVersion(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployminator_ListDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployminatorServer).ListDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployminator.Deployminator/ListDeployments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployminatorServer).ListDeployments(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployminator_UndeployVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployminatorServer).UndeployVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployminator.Deployminator/UndeployVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployminatorServer).UndeployVersion(ctx, req.(*UndeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployminator_AutodeployerStartup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployminatorServer).AutodeployerStartup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployminator.Deployminator/AutodeployerStartup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployminatorServer).AutodeployerStartup(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployminator_AutodeployerShutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployminatorServer).AutodeployerShutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployminator.Deployminator/AutodeployerShutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployminatorServer).AutodeployerShutdown(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deployminator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deployminator.Deployminator",
	HandlerType: (*DeployminatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBuildAvailable",
			Handler:    _Deployminator_NewBuildAvailable_Handler,
		},
		{
			MethodName: "DeployVersion",
			Handler:    _Deployminator_DeployVersion_Handler,
		},
		{
			MethodName: "ListDeployments",
			Handler:    _Deployminator_ListDeployments_Handler,
		},
		{
			MethodName: "UndeployVersion",
			Handler:    _Deployminator_UndeployVersion_Handler,
		},
		{
			MethodName: "AutodeployerStartup",
			Handler:    _Deployminator_AutodeployerStartup_Handler,
		},
		{
			MethodName: "AutodeployerShutdown",
			Handler:    _Deployminator_AutodeployerShutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/deployminator/deployminator.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/deployminator/deployminator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x56, 0x7e, 0x9a, 0x36, 0x27, 0xc9, 0xad, 0xee, 0xdc, 0xea, 0xca, 0xb5, 0xaa, 0x2a, 0xd7,
	0xbd, 0x48, 0x11, 0x0b, 0x57, 0xb4, 0x88, 0x05, 0xb4, 0x42, 0x49, 0x2d, 0x90, 0xa5, 0x52, 0x21,
	0x43, 0xbb, 0x45, 0x53, 0x7b, 0x9a, 0x8c, 0xb0, 0x67, 0xcc, 0x78, 0xdc, 0x28, 0x82, 0x55, 0x1f,
	0x03, 0x76, 0x3c, 0x03, 0x4f, 0xc1, 0xfb, 0xc0, 0x1a, 0xf9, 0x2f, 0xf1, 0x38, 0x69, 0x41, 0xac,
	0x58, 0xb5, 0xe7, 0xe4, 0x7c, 0x73, 0xbe, 0xef, 0x3b, 0xc7, 0x33, 0x70, 0x34, 0xe6, 0x3e, 0x66,
	0x63, 0xd3, 0xe5, 0x4c, 0x60, 0x6f, 0xca, 0xb9, 0x67, 0x32, 0x22, 0xf7, 0x71, 0x48, 0xa3, 0x7d,
	0x8f, 0x84, 0x3e, 0x9f, 0x05, 0x94, 0x61, 0xc9, 0x85, 0x1a, 0x99, 0xa1, 0xe0, 0x92, 0xa3, 0x9e,
	0x92, 0xd4, 0xcd, 0x3b, 0x0e, 0x73, 0x79, 0x10, 0x70, 0x96, 0xff, 0xc9, 0xe0, 0xfa, 0xe3, 0x9f,
	0x37, 0xe7, 0xec, 0x2d, 0x99, 0x29, 0x41, 0x86, 0x35, 0xa6, 0xb0, 0x7e, 0x41, 0x44, 0x44, 0x39,
	0x43, 0x7f, 0x41, 0xdd, 0xb6, 0xb4, 0x5a, 0xbf, 0x36, 0x68, 0x3a, 0x75, 0xdb, 0x42, 0x7d, 0xe8,
	0x8c, 0x62, 0xea, 0x7b, 0x67, 0x71, 0x70, 0x49, 0x84, 0x56, 0x4f, 0x7f, 0x28, 0xa7, 0x90, 0x01,
	0x5d, 0x87, 0x84, 0x3c, 0xa2, 0x92, 0x8b, 0x99, 0x6d, 0x69, 0x8d, 0xb4, 0x44, 0xc9, 0xa1, 0x7f,
	0xa1, 0x35, 0x12, 0x98, 0xb9, 0x13, 0xad, 0xd9, 0xaf, 0x0d, 0xda, 0x4e, 0x1e, 0x19, 0xef, 0xa1,
	0x33, 0x0c, 0x43, 0x9f, 0xba, 0x58, 0xae, 0x6a, 0x9e, 0xc0, 0x28, 0xc3, 0x62, 0x96, 0xf6, 0x4d,
	0x60, 0x69, 0xf4, 0x4b, 0x2d, 0xfb, 0xd0, 0xb1, 0xf8, 0x94, 0xf9, 0x1c, 0x7b, 0xe7, 0xce, 0x69,
	0xde, 0xb7, 0x9c, 0x32, 0xbe, 0xd6, 0x60, 0xcb, 0xca, 0xcc, 0x20, 0x4c, 0x5a, 0x24, 0x72, 0x05,
	0x0d, 0x25, 0x17, 0x4b, 0x34, 0x26, 0x0a, 0xcb, 0x94, 0x4b, 0xe7, 0x40, 0x37, 0xd5, 0x21, 0x96,
	0x2a, 0x46, 0xf7, 0x3f, 0xde, 0x6c, 0xb7, 0x62, 0xca, 0xe4, 0xa3, 0x87, 0x9f, 0x6f, 0xb6, 0x77,
	0x94, 0xc2, 0x37, 0x78, 0x51, 0x68, 0x52, 0xcf, 0x51, 0x0c, 0xa8, 0xb8, 0xdd, 0x58, 0x76, 0xfb,
	0x36, 0x27, 0x3f, 0xd5, 0xa1, 0x63, 0xb3, 0x48, 0x62, 0xe6, 0x12, 0x8b, 0x5c, 0x2d, 0x69, 0xf8,
	0x00, 0xdd, 0x85, 0x56, 0xdb, 0xca, 0x45, 0xec, 0x55, 0x44, 0xac, 0xb2, 0x63, 0x74, 0xa8, 0xa8,
	0xb9, 0xa7, 0xaa, 0xf1, 0xe6, 0x08, 0x6f, 0x8e, 0x48, 0x64, 0x29, 0xdd, 0x92, 0x81, 0xbd, 0xc0,
	0xee, 0x84, 0x32, 0xf2, 0x5c, 0xf0, 0x38, 0x4c, 0x85, 0xb5, 0x1d, 0x25, 0x87, 0x76, 0xa0, 0x5d,
	0x08, 0x88, 0x52, 0x71, 0x3d, 0x67, 0x91, 0x40, 0x47, 0xb0, 0x5d, 0x04, 0x27, 0x3c, 0x66, 0xd2,
	0x8e, 0x5e, 0x12, 0x91, 0xe3, 0xb5, 0xb5, 0x7e, 0x6d, 0xb0, 0xe1, 0xdc, 0x5e, 0x60, 0x7c, 0xa9,
	0xc1, 0xc6, 0x50, 0x8c, 0xe3, 0x84, 0xce, 0x1f, 0x66, 0x8d, 0xbe, 0x60, 0x96, 0xdb, 0x32, 0x8f,
	0x8d, 0x6f, 0x35, 0xd8, 0x3c, 0x23, 0xd3, 0x74, 0xfe, 0x0e, 0x79, 0x17, 0x93, 0x48, 0xa2, 0x5d,
	0x80, 0x0c, 0xff, 0x8c, 0xfa, 0x24, 0x55, 0xd1, 0x75, 0x4a, 0x99, 0xa5, 0x6f, 0xa3, 0x7e, 0xe7,
	0xe7, 0xd8, 0x28, 0x2f, 0x51, 0x75, 0xfd, 0x9a, 0xcb, 0xeb, 0xa7, 0xc3, 0xc6, 0x09, 0x0f, 0x02,
	0x9a, 0xf8, 0xb4, 0x96, 0xb1, 0x2d, 0xe2, 0xa4, 0xf3, 0x50, 0x48, 0x72, 0x85, 0x5d, 0x79, 0x86,
	0x03, 0xa2, 0xb5, 0xb2, 0x21, 0x97, 0x73, 0xe8, 0x7f, 0xe8, 0xa5, 0xc7, 0x45, 0x44, 0x5c, 0x13,
	0x61, 0x5b, 0xda, 0x7a, 0x5a, 0xa4, 0x26, 0x8d, 0x63, 0xe8, 0x65, 0x8a, 0x0a, 0xd1, 0xda, 0xfc,
	0x82, 0xca, 0xe7, 0x36, 0xbf, 0xaf, 0x10, 0x34, 0x5f, 0xe3, 0x71, 0xa4, 0xd5, 0xfb, 0x8d, 0x41,
	0xdb, 0x49, 0xff, 0x37, 0x9e, 0xc2, 0xe6, 0x39, 0xf3, 0x7e, 0xff, 0x80, 0x83, 0xef, 0xf5, 0x82,
	0x40, 0x3e, 0x4b, 0x34, 0x84, 0xbf, 0x8b, 0x41, 0x0c, 0xaf, 0x31, 0xf5, 0xf1, 0xa5, 0x4f, 0xd0,
	0x6e, 0x65, 0x45, 0x2a, 0xa3, 0xd2, 0xbb, 0x66, 0x7e, 0x43, 0x5f, 0x70, 0xea, 0xa1, 0x27, 0xc5,
	0x99, 0x45, 0xe7, 0x9d, 0x95, 0x1b, 0xb6, 0x1a, 0x7c, 0x0c, 0x9b, 0xa7, 0x34, 0x92, 0x8b, 0xcd,
	0x89, 0x90, 0x52, 0xa0, 0xcf, 0x0f, 0xcb, 0xee, 0xf5, 0x45, 0x61, 0x02, 0x43, 0x25, 0x47, 0x8a,
	0xee, 0x55, 0xf2, 0x15, 0xc7, 0x2a, 0xfd, 0x1f, 0xc0, 0x3f, 0xc3, 0x58, 0xf2, 0xac, 0x84, 0x88,
	0x57, 0x12, 0x0b, 0x19, 0x87, 0x15, 0x0e, 0x2a, 0xe4, 0x00, 0xb6, 0x14, 0xc8, 0x24, 0x96, 0x1e,
	0x9f, 0xb2, 0xbb, 0x30, 0xa3, 0x3d, 0xf8, 0x8f, 0x11, 0x59, 0x7e, 0xc3, 0x92, 0xf7, 0x4b, 0x65,
	0x7a, 0xd9, 0x4a, 0x1f, 0xad, 0xc3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x88, 0x88, 0x57,
	0x6f, 0x07, 0x00, 0x00,
}
