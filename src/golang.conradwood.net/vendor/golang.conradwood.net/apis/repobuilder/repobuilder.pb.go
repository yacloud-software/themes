// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/repobuilder/repobuilder.proto
// DO NOT EDIT!

/*
Package repobuilder is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/repobuilder/repobuilder.proto

It has these top-level messages:
	TrackerGitRepository
	TrackerLog
	CreateWebRepoRequest
	RepoCreateStatus
	CreateRepoResponse
	RepoStatusRequest
	RepoDomain
	Language
	Choices
	ForkRequest
*/
package repobuilder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"
import auth "golang.conradwood.net/apis/auth"

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

// we create a repository in the gitserver, here we track progress whilst it goes through various stages in the create process
type TrackerGitRepository struct {
	ID                 uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	CreateRequestID    uint64 `protobuf:"varint,2,opt,name=CreateRequestID" json:"CreateRequestID,omitempty"`
	CreateType         uint32 `protobuf:"varint,3,opt,name=CreateType" json:"CreateType,omitempty"`
	RepositoryID       uint64 `protobuf:"varint,4,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	URLHost            string `protobuf:"bytes,5,opt,name=URLHost" json:"URLHost,omitempty"`
	URLPath            string `protobuf:"bytes,6,opt,name=URLPath" json:"URLPath,omitempty"`
	RepositoryCreated  bool   `protobuf:"varint,7,opt,name=RepositoryCreated" json:"RepositoryCreated,omitempty"`
	SourceInstalled    bool   `protobuf:"varint,8,opt,name=SourceInstalled" json:"SourceInstalled,omitempty"`
	PackageID          string `protobuf:"bytes,9,opt,name=PackageID" json:"PackageID,omitempty"`
	PackageName        string `protobuf:"bytes,10,opt,name=PackageName" json:"PackageName,omitempty"`
	ProtoFilename      string `protobuf:"bytes,11,opt,name=ProtoFilename" json:"ProtoFilename,omitempty"`
	ProtoSubmitted     bool   `protobuf:"varint,12,opt,name=ProtoSubmitted" json:"ProtoSubmitted,omitempty"`
	ProtoCommitted     bool   `protobuf:"varint,13,opt,name=ProtoCommitted" json:"ProtoCommitted,omitempty"`
	MinProtoVersion    uint64 `protobuf:"varint,14,opt,name=MinProtoVersion" json:"MinProtoVersion,omitempty"`
	Context            string `protobuf:"bytes,15,opt,name=Context" json:"Context,omitempty"`
	UserID             string `protobuf:"bytes,16,opt,name=UserID" json:"UserID,omitempty"`
	PermissionsCreated bool   `protobuf:"varint,17,opt,name=PermissionsCreated" json:"PermissionsCreated,omitempty"`
	SecureArgsCreated  bool   `protobuf:"varint,18,opt,name=SecureArgsCreated" json:"SecureArgsCreated,omitempty"`
	ServiceID          string `protobuf:"bytes,19,opt,name=ServiceID" json:"ServiceID,omitempty"`
	ServiceUserID      string `protobuf:"bytes,20,opt,name=ServiceUserID" json:"ServiceUserID,omitempty"`
	ServiceToken       string `protobuf:"bytes,21,opt,name=ServiceToken" json:"ServiceToken,omitempty"`
	Finalised          bool   `protobuf:"varint,22,opt,name=Finalised" json:"Finalised,omitempty"`
	PatchRepo          bool   `protobuf:"varint,23,opt,name=PatchRepo" json:"PatchRepo,omitempty"`
	SourceRepositoryID uint64 `protobuf:"varint,24,opt,name=SourceRepositoryID" json:"SourceRepositoryID,omitempty"`
	NotificationSent   bool   `protobuf:"varint,25,opt,name=NotificationSent" json:"NotificationSent,omitempty"`
}

func (m *TrackerGitRepository) Reset()                    { *m = TrackerGitRepository{} }
func (m *TrackerGitRepository) String() string            { return proto.CompactTextString(m) }
func (*TrackerGitRepository) ProtoMessage()               {}
func (*TrackerGitRepository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TrackerGitRepository) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TrackerGitRepository) GetCreateRequestID() uint64 {
	if m != nil {
		return m.CreateRequestID
	}
	return 0
}

func (m *TrackerGitRepository) GetCreateType() uint32 {
	if m != nil {
		return m.CreateType
	}
	return 0
}

func (m *TrackerGitRepository) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *TrackerGitRepository) GetURLHost() string {
	if m != nil {
		return m.URLHost
	}
	return ""
}

func (m *TrackerGitRepository) GetURLPath() string {
	if m != nil {
		return m.URLPath
	}
	return ""
}

func (m *TrackerGitRepository) GetRepositoryCreated() bool {
	if m != nil {
		return m.RepositoryCreated
	}
	return false
}

func (m *TrackerGitRepository) GetSourceInstalled() bool {
	if m != nil {
		return m.SourceInstalled
	}
	return false
}

func (m *TrackerGitRepository) GetPackageID() string {
	if m != nil {
		return m.PackageID
	}
	return ""
}

func (m *TrackerGitRepository) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *TrackerGitRepository) GetProtoFilename() string {
	if m != nil {
		return m.ProtoFilename
	}
	return ""
}

func (m *TrackerGitRepository) GetProtoSubmitted() bool {
	if m != nil {
		return m.ProtoSubmitted
	}
	return false
}

func (m *TrackerGitRepository) GetProtoCommitted() bool {
	if m != nil {
		return m.ProtoCommitted
	}
	return false
}

func (m *TrackerGitRepository) GetMinProtoVersion() uint64 {
	if m != nil {
		return m.MinProtoVersion
	}
	return 0
}

func (m *TrackerGitRepository) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *TrackerGitRepository) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TrackerGitRepository) GetPermissionsCreated() bool {
	if m != nil {
		return m.PermissionsCreated
	}
	return false
}

func (m *TrackerGitRepository) GetSecureArgsCreated() bool {
	if m != nil {
		return m.SecureArgsCreated
	}
	return false
}

func (m *TrackerGitRepository) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *TrackerGitRepository) GetServiceUserID() string {
	if m != nil {
		return m.ServiceUserID
	}
	return ""
}

func (m *TrackerGitRepository) GetServiceToken() string {
	if m != nil {
		return m.ServiceToken
	}
	return ""
}

func (m *TrackerGitRepository) GetFinalised() bool {
	if m != nil {
		return m.Finalised
	}
	return false
}

func (m *TrackerGitRepository) GetPatchRepo() bool {
	if m != nil {
		return m.PatchRepo
	}
	return false
}

func (m *TrackerGitRepository) GetSourceRepositoryID() uint64 {
	if m != nil {
		return m.SourceRepositoryID
	}
	return 0
}

func (m *TrackerGitRepository) GetNotificationSent() bool {
	if m != nil {
		return m.NotificationSent
	}
	return false
}

// log for a reporequest
type TrackerLog struct {
	ID              uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	CreateRequestID uint64 `protobuf:"varint,2,opt,name=CreateRequestID" json:"CreateRequestID,omitempty"`
	CreateType      uint32 `protobuf:"varint,3,opt,name=CreateType" json:"CreateType,omitempty"`
	LogMessage      string `protobuf:"bytes,4,opt,name=LogMessage" json:"LogMessage,omitempty"`
	PublicMessage   string `protobuf:"bytes,5,opt,name=PublicMessage" json:"PublicMessage,omitempty"`
	Occured         uint32 `protobuf:"varint,6,opt,name=Occured" json:"Occured,omitempty"`
	Success         bool   `protobuf:"varint,7,opt,name=Success" json:"Success,omitempty"`
	Task            string `protobuf:"bytes,8,opt,name=Task" json:"Task,omitempty"`
}

func (m *TrackerLog) Reset()                    { *m = TrackerLog{} }
func (m *TrackerLog) String() string            { return proto.CompactTextString(m) }
func (*TrackerLog) ProtoMessage()               {}
func (*TrackerLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TrackerLog) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TrackerLog) GetCreateRequestID() uint64 {
	if m != nil {
		return m.CreateRequestID
	}
	return 0
}

func (m *TrackerLog) GetCreateType() uint32 {
	if m != nil {
		return m.CreateType
	}
	return 0
}

func (m *TrackerLog) GetLogMessage() string {
	if m != nil {
		return m.LogMessage
	}
	return ""
}

func (m *TrackerLog) GetPublicMessage() string {
	if m != nil {
		return m.PublicMessage
	}
	return ""
}

func (m *TrackerLog) GetOccured() uint32 {
	if m != nil {
		return m.Occured
	}
	return 0
}

func (m *TrackerLog) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *TrackerLog) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type CreateWebRepoRequest struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	// A free text description upto 2000 characters
	Description string `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
	// a short name
	Name string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	// which language shall we prepare this for?
	Language common.ProgrammingLanguage `protobuf:"varint,4,opt,name=Language,enum=common.ProgrammingLanguage" json:"Language,omitempty"`
	//
	// which domain shall the repo be served under? (this will be 'git.domain.com')
	// the A-Record for 'git.domain.com' must exist already.
	// Example: Reponame == "stuffrepo", Domain == "foo.com"
	// git clone https://git.foo.com/git/stuffrepo
	Domain string `protobuf:"bytes,5,opt,name=Domain" json:"Domain,omitempty"`
	//
	// the repository name. A "repository name" is somewhat virtual. It is possible to serve the same repository
	// under multiple domains and path if required. this is the 'initial' host & path.
	RepoName string `protobuf:"bytes,6,opt,name=RepoName" json:"RepoName,omitempty"`
	//
	// The Servicename is used to register and find the service. ServiceName and Domain are used to route to a specific service
	ServiceName string `protobuf:"bytes,7,opt,name=ServiceName" json:"ServiceName,omitempty"`
	//
	// list of groups who may /view/ this service. a group which has no viewing rights will neither see its registration, nor have any sort of access to it.
	VisibilityGroupIDs []string `protobuf:"bytes,8,rep,name=VisibilityGroupIDs" json:"VisibilityGroupIDs,omitempty"`
	//
	// list of groupids who may /use/ this service. These users may not have access to the source or binary, but are granted access to view the service and use its
	// api
	AccessGroupIDs []string `protobuf:"bytes,9,rep,name=AccessGroupIDs" json:"AccessGroupIDs,omitempty"`
	//
	// list of groupids who may /develop/ this service
	DeveloperGroupIDs []string `protobuf:"bytes,10,rep,name=DeveloperGroupIDs" json:"DeveloperGroupIDs,omitempty"`
	// domain under which proto will be created, e.g. golang.singingcat.net
	ProtoDomain string `protobuf:"bytes,11,opt,name=ProtoDomain" json:"ProtoDomain,omitempty"`
}

func (m *CreateWebRepoRequest) Reset()                    { *m = CreateWebRepoRequest{} }
func (m *CreateWebRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateWebRepoRequest) ProtoMessage()               {}
func (*CreateWebRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateWebRepoRequest) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *CreateWebRepoRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateWebRepoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateWebRepoRequest) GetLanguage() common.ProgrammingLanguage {
	if m != nil {
		return m.Language
	}
	return common.ProgrammingLanguage_INVALID
}

func (m *CreateWebRepoRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *CreateWebRepoRequest) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *CreateWebRepoRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *CreateWebRepoRequest) GetVisibilityGroupIDs() []string {
	if m != nil {
		return m.VisibilityGroupIDs
	}
	return nil
}

func (m *CreateWebRepoRequest) GetAccessGroupIDs() []string {
	if m != nil {
		return m.AccessGroupIDs
	}
	return nil
}

func (m *CreateWebRepoRequest) GetDeveloperGroupIDs() []string {
	if m != nil {
		return m.DeveloperGroupIDs
	}
	return nil
}

func (m *CreateWebRepoRequest) GetProtoDomain() string {
	if m != nil {
		return m.ProtoDomain
	}
	return ""
}

// if this exists than the repo create request is "complete" (successful or not)
type RepoCreateStatus struct {
	ID              uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	CreateRequestID uint64 `protobuf:"varint,2,opt,name=CreateRequestID" json:"CreateRequestID,omitempty"`
	CreateType      uint32 `protobuf:"varint,3,opt,name=CreateType" json:"CreateType,omitempty"`
	Success         bool   `protobuf:"varint,4,opt,name=Success" json:"Success,omitempty"`
	Error           string `protobuf:"bytes,5,opt,name=Error" json:"Error,omitempty"`
}

func (m *RepoCreateStatus) Reset()                    { *m = RepoCreateStatus{} }
func (m *RepoCreateStatus) String() string            { return proto.CompactTextString(m) }
func (*RepoCreateStatus) ProtoMessage()               {}
func (*RepoCreateStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RepoCreateStatus) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *RepoCreateStatus) GetCreateRequestID() uint64 {
	if m != nil {
		return m.CreateRequestID
	}
	return 0
}

func (m *RepoCreateStatus) GetCreateType() uint32 {
	if m != nil {
		return m.CreateType
	}
	return 0
}

func (m *RepoCreateStatus) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RepoCreateStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CreateRepoResponse struct {
	RequestID uint64 `protobuf:"varint,1,opt,name=RequestID" json:"RequestID,omitempty"`
	Finished  bool   `protobuf:"varint,2,opt,name=Finished" json:"Finished,omitempty"`
	Success   bool   `protobuf:"varint,3,opt,name=Success" json:"Success,omitempty"`
	Error     string `protobuf:"bytes,4,opt,name=Error" json:"Error,omitempty"`
}

func (m *CreateRepoResponse) Reset()                    { *m = CreateRepoResponse{} }
func (m *CreateRepoResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateRepoResponse) ProtoMessage()               {}
func (*CreateRepoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateRepoResponse) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *CreateRepoResponse) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

func (m *CreateRepoResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreateRepoResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RepoStatusRequest struct {
	RequestID uint64 `protobuf:"varint,1,opt,name=RequestID" json:"RequestID,omitempty"`
}

func (m *RepoStatusRequest) Reset()                    { *m = RepoStatusRequest{} }
func (m *RepoStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*RepoStatusRequest) ProtoMessage()               {}
func (*RepoStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RepoStatusRequest) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

type RepoDomain struct {
	Domain string `protobuf:"bytes,1,opt,name=Domain" json:"Domain,omitempty"`
}

func (m *RepoDomain) Reset()                    { *m = RepoDomain{} }
func (m *RepoDomain) String() string            { return proto.CompactTextString(m) }
func (*RepoDomain) ProtoMessage()               {}
func (*RepoDomain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RepoDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type Language struct {
	ID   uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *Language) Reset()                    { *m = Language{} }
func (m *Language) String() string            { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()               {}
func (*Language) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Language) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Language) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Choices struct {
	Domains   []*RepoDomain `protobuf:"bytes,1,rep,name=Domains" json:"Domains,omitempty"`
	Languages []*Language   `protobuf:"bytes,2,rep,name=Languages" json:"Languages,omitempty"`
	Groups    []*auth.Group `protobuf:"bytes,3,rep,name=Groups" json:"Groups,omitempty"`
}

func (m *Choices) Reset()                    { *m = Choices{} }
func (m *Choices) String() string            { return proto.CompactTextString(m) }
func (*Choices) ProtoMessage()               {}
func (*Choices) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Choices) GetDomains() []*RepoDomain {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *Choices) GetLanguages() []*Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

func (m *Choices) GetGroups() []*auth.Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type ForkRequest struct {
	RepositoryID uint64 `protobuf:"varint,1,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Domain       string `protobuf:"bytes,3,opt,name=Domain" json:"Domain,omitempty"`
	RepoName     string `protobuf:"bytes,4,opt,name=RepoName" json:"RepoName,omitempty"`
	//
	// list of groups who may /view/ this service. a group which has no viewing rights will neither see its registration, nor have any sort of access to it.
	VisibilityGroupIDs []string `protobuf:"bytes,5,rep,name=VisibilityGroupIDs" json:"VisibilityGroupIDs,omitempty"`
	//
	// list of groupids who may /use/ this service. These users may not have access to the source or binary, but are granted access to view the service and use its
	// api
	AccessGroupIDs []string `protobuf:"bytes,6,rep,name=AccessGroupIDs" json:"AccessGroupIDs,omitempty"`
	//
	// list of groupids who may /develop/ this service
	DeveloperGroupIDs []string `protobuf:"bytes,7,rep,name=DeveloperGroupIDs" json:"DeveloperGroupIDs,omitempty"`
}

func (m *ForkRequest) Reset()                    { *m = ForkRequest{} }
func (m *ForkRequest) String() string            { return proto.CompactTextString(m) }
func (*ForkRequest) ProtoMessage()               {}
func (*ForkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ForkRequest) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *ForkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ForkRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ForkRequest) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *ForkRequest) GetVisibilityGroupIDs() []string {
	if m != nil {
		return m.VisibilityGroupIDs
	}
	return nil
}

func (m *ForkRequest) GetAccessGroupIDs() []string {
	if m != nil {
		return m.AccessGroupIDs
	}
	return nil
}

func (m *ForkRequest) GetDeveloperGroupIDs() []string {
	if m != nil {
		return m.DeveloperGroupIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*TrackerGitRepository)(nil), "repobuilder.TrackerGitRepository")
	proto.RegisterType((*TrackerLog)(nil), "repobuilder.TrackerLog")
	proto.RegisterType((*CreateWebRepoRequest)(nil), "repobuilder.CreateWebRepoRequest")
	proto.RegisterType((*RepoCreateStatus)(nil), "repobuilder.RepoCreateStatus")
	proto.RegisterType((*CreateRepoResponse)(nil), "repobuilder.CreateRepoResponse")
	proto.RegisterType((*RepoStatusRequest)(nil), "repobuilder.RepoStatusRequest")
	proto.RegisterType((*RepoDomain)(nil), "repobuilder.RepoDomain")
	proto.RegisterType((*Language)(nil), "repobuilder.Language")
	proto.RegisterType((*Choices)(nil), "repobuilder.Choices")
	proto.RegisterType((*ForkRequest)(nil), "repobuilder.ForkRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RepoBuilder service

type RepoBuilderClient interface {
	// get the status of a previously submitted create request
	GetRepoStatus(ctx context.Context, in *RepoStatusRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error)
	// create a web-repo component
	CreateWebRepo(ctx context.Context, in *CreateWebRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error)
	// get choices for "creating a repo" for the current user, for example, this returns the Domains the particular user may create repos in
	GetRepoChoices(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*Choices, error)
	// trigger a run over all pending requests (also done periodically)
	RetriggerAll(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error)
	// limited to 'prober' - this first asks gitserver to reset a repo and then recreates it (more-or-less) from scratch
	RecreateWebRepo(ctx context.Context, in *CreateWebRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error)
	Fork(ctx context.Context, in *ForkRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error)
}

type repoBuilderClient struct {
	cc *grpc.ClientConn
}

func NewRepoBuilderClient(cc *grpc.ClientConn) RepoBuilderClient {
	return &repoBuilderClient{cc}
}

func (c *repoBuilderClient) GetRepoStatus(ctx context.Context, in *RepoStatusRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error) {
	out := new(CreateRepoResponse)
	err := grpc.Invoke(ctx, "/repobuilder.RepoBuilder/GetRepoStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoBuilderClient) CreateWebRepo(ctx context.Context, in *CreateWebRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error) {
	out := new(CreateRepoResponse)
	err := grpc.Invoke(ctx, "/repobuilder.RepoBuilder/CreateWebRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoBuilderClient) GetRepoChoices(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*Choices, error) {
	out := new(Choices)
	err := grpc.Invoke(ctx, "/repobuilder.RepoBuilder/GetRepoChoices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoBuilderClient) RetriggerAll(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/repobuilder.RepoBuilder/RetriggerAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoBuilderClient) RecreateWebRepo(ctx context.Context, in *CreateWebRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error) {
	out := new(CreateRepoResponse)
	err := grpc.Invoke(ctx, "/repobuilder.RepoBuilder/RecreateWebRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoBuilderClient) Fork(ctx context.Context, in *ForkRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error) {
	out := new(CreateRepoResponse)
	err := grpc.Invoke(ctx, "/repobuilder.RepoBuilder/Fork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepoBuilder service

type RepoBuilderServer interface {
	// get the status of a previously submitted create request
	GetRepoStatus(context.Context, *RepoStatusRequest) (*CreateRepoResponse, error)
	// create a web-repo component
	CreateWebRepo(context.Context, *CreateWebRepoRequest) (*CreateRepoResponse, error)
	// get choices for "creating a repo" for the current user, for example, this returns the Domains the particular user may create repos in
	GetRepoChoices(context.Context, *common.Void) (*Choices, error)
	// trigger a run over all pending requests (also done periodically)
	RetriggerAll(context.Context, *common.Void) (*common.Void, error)
	// limited to 'prober' - this first asks gitserver to reset a repo and then recreates it (more-or-less) from scratch
	RecreateWebRepo(context.Context, *CreateWebRepoRequest) (*CreateRepoResponse, error)
	Fork(context.Context, *ForkRequest) (*CreateRepoResponse, error)
}

func RegisterRepoBuilderServer(s *grpc.Server, srv RepoBuilderServer) {
	s.RegisterService(&_RepoBuilder_serviceDesc, srv)
}

func _RepoBuilder_GetRepoStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoBuilderServer).GetRepoStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repobuilder.RepoBuilder/GetRepoStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoBuilderServer).GetRepoStatus(ctx, req.(*RepoStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoBuilder_CreateWebRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWebRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoBuilderServer).CreateWebRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repobuilder.RepoBuilder/CreateWebRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoBuilderServer).CreateWebRepo(ctx, req.(*CreateWebRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoBuilder_GetRepoChoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoBuilderServer).GetRepoChoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repobuilder.RepoBuilder/GetRepoChoices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoBuilderServer).GetRepoChoices(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoBuilder_RetriggerAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoBuilderServer).RetriggerAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repobuilder.RepoBuilder/RetriggerAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoBuilderServer).RetriggerAll(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoBuilder_RecreateWebRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWebRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoBuilderServer).RecreateWebRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repobuilder.RepoBuilder/RecreateWebRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoBuilderServer).RecreateWebRepo(ctx, req.(*CreateWebRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoBuilder_Fork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoBuilderServer).Fork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repobuilder.RepoBuilder/Fork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoBuilderServer).Fork(ctx, req.(*ForkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoBuilder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "repobuilder.RepoBuilder",
	HandlerType: (*RepoBuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRepoStatus",
			Handler:    _RepoBuilder_GetRepoStatus_Handler,
		},
		{
			MethodName: "CreateWebRepo",
			Handler:    _RepoBuilder_CreateWebRepo_Handler,
		},
		{
			MethodName: "GetRepoChoices",
			Handler:    _RepoBuilder_GetRepoChoices_Handler,
		},
		{
			MethodName: "RetriggerAll",
			Handler:    _RepoBuilder_RetriggerAll_Handler,
		},
		{
			MethodName: "RecreateWebRepo",
			Handler:    _RepoBuilder_RecreateWebRepo_Handler,
		},
		{
			MethodName: "Fork",
			Handler:    _RepoBuilder_Fork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/repobuilder/repobuilder.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/repobuilder/repobuilder.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1087 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xcf, 0x6e, 0x1b, 0xb7,
	0x13, 0xc6, 0x5a, 0xb2, 0x25, 0x8d, 0x2c, 0xdb, 0xe1, 0xcf, 0x49, 0xf8, 0x73, 0x0b, 0x57, 0x55,
	0x83, 0x40, 0x30, 0x0a, 0x05, 0x71, 0x0a, 0xb4, 0x57, 0xc7, 0xaa, 0x5d, 0x01, 0x4e, 0x2a, 0xac,
	0x6c, 0xf7, 0xbc, 0x5a, 0xb1, 0x6b, 0xc2, 0xab, 0xa5, 0x4a, 0x72, 0xd3, 0xfa, 0xd2, 0x67, 0x28,
	0x50, 0xa0, 0xa7, 0x3e, 0x44, 0x1f, 0xae, 0xf7, 0x16, 0x43, 0x72, 0xff, 0x49, 0xb2, 0xd3, 0x02,
	0xcd, 0xc5, 0xd6, 0x7c, 0x33, 0xe4, 0x0c, 0xbf, 0xf9, 0x86, 0x4b, 0xf8, 0x2a, 0x12, 0x71, 0x90,
	0x44, 0x83, 0x50, 0x24, 0x32, 0x98, 0xfd, 0x28, 0xc4, 0x6c, 0x90, 0x30, 0xfd, 0x22, 0x58, 0x70,
	0xf5, 0x42, 0xb2, 0x85, 0x98, 0xa6, 0x3c, 0x9e, 0x31, 0x59, 0xfe, 0x3d, 0x58, 0x48, 0xa1, 0x05,
	0x69, 0x97, 0xa0, 0x83, 0xc1, 0x03, 0xdb, 0x84, 0x62, 0x3e, 0x17, 0x89, 0xfb, 0x67, 0x17, 0x1f,
	0x1c, 0x3d, 0x10, 0x1f, 0xa4, 0xfa, 0xc6, 0xfc, 0xb1, 0xb1, 0xbd, 0x5f, 0x1a, 0xb0, 0x7f, 0x29,
	0x83, 0xf0, 0x96, 0xc9, 0x73, 0xae, 0x7d, 0xb6, 0x10, 0x8a, 0x6b, 0x21, 0xef, 0xc8, 0x0e, 0x6c,
	0x8c, 0x86, 0xd4, 0xeb, 0x7a, 0xfd, 0xba, 0xbf, 0x31, 0x1a, 0x92, 0x3e, 0xec, 0x9e, 0x4a, 0x16,
	0x68, 0xe6, 0xb3, 0x1f, 0x52, 0xa6, 0xf4, 0x68, 0x48, 0x37, 0x8c, 0x73, 0x19, 0x26, 0x87, 0x00,
	0x16, 0xba, 0xbc, 0x5b, 0x30, 0x5a, 0xeb, 0x7a, 0xfd, 0x8e, 0x5f, 0x42, 0x48, 0x0f, 0xb6, 0x8b,
	0x3c, 0xa3, 0x21, 0xad, 0x9b, 0x6d, 0x2a, 0x18, 0xa1, 0xd0, 0xb8, 0xf2, 0x2f, 0xbe, 0x11, 0x4a,
	0xd3, 0xcd, 0xae, 0xd7, 0x6f, 0xf9, 0x99, 0xe9, 0x3c, 0xe3, 0x40, 0xdf, 0xd0, 0xad, 0xdc, 0x83,
	0x26, 0xf9, 0x1c, 0x1e, 0x15, 0x7b, 0xd8, 0x7c, 0x33, 0xda, 0xe8, 0x7a, 0xfd, 0xa6, 0xbf, 0xea,
	0xc0, 0xf3, 0x4c, 0x44, 0x2a, 0x43, 0x36, 0x4a, 0x94, 0x0e, 0xe2, 0x98, 0xcd, 0x68, 0xd3, 0xc4,
	0x2e, 0xc3, 0xe4, 0x63, 0x68, 0x8d, 0x83, 0xf0, 0x36, 0x88, 0xd8, 0x68, 0x48, 0x5b, 0x26, 0x67,
	0x01, 0x90, 0x2e, 0xb4, 0x9d, 0xf1, 0x36, 0x98, 0x33, 0x0a, 0xc6, 0x5f, 0x86, 0xc8, 0x33, 0xe8,
	0x8c, 0x91, 0xeb, 0x33, 0x1e, 0xb3, 0x04, 0x63, 0xda, 0x26, 0xa6, 0x0a, 0x92, 0xe7, 0xb0, 0x63,
	0x80, 0x49, 0x3a, 0x9d, 0x73, 0x8d, 0xa5, 0x6f, 0x9b, 0x72, 0x96, 0xd0, 0x3c, 0xee, 0x54, 0xcc,
	0x5d, 0x5c, 0xa7, 0x14, 0x97, 0xa3, 0x78, 0xbe, 0x37, 0x3c, 0x31, 0xe0, 0x35, 0x93, 0x8a, 0x8b,
	0x84, 0xee, 0xd8, 0x7e, 0x2d, 0xc1, 0xc8, 0xe8, 0xa9, 0x48, 0x34, 0xfb, 0x49, 0xd3, 0x5d, 0xcb,
	0xa8, 0x33, 0xc9, 0x13, 0xd8, 0xba, 0x52, 0x4c, 0x8e, 0x86, 0x74, 0xcf, 0x38, 0x9c, 0x45, 0x06,
	0x40, 0xc6, 0x4c, 0xce, 0xb9, 0xc2, 0xf5, 0x2a, 0xa3, 0xfa, 0x91, 0xa9, 0x63, 0x8d, 0x07, 0x3b,
	0x33, 0x61, 0x61, 0x2a, 0xd9, 0x89, 0x8c, 0xf2, 0x70, 0x62, 0x3b, 0xb3, 0xe2, 0x40, 0xbe, 0x27,
	0x4c, 0xbe, 0xe3, 0x21, 0xf2, 0xfd, 0x3f, 0xcb, 0x77, 0x0e, 0x20, 0x9b, 0xce, 0x70, 0xa5, 0xed,
	0x5b, 0x36, 0x2b, 0x20, 0x6a, 0xcc, 0x01, 0x97, 0xe2, 0x96, 0x25, 0xf4, 0xb1, 0x09, 0xaa, 0x60,
	0x98, 0xe7, 0x8c, 0x27, 0x41, 0xcc, 0x15, 0x9b, 0xd1, 0x27, 0xa6, 0x9a, 0x02, 0xb0, 0x5d, 0xd7,
	0xe1, 0x0d, 0x2a, 0x87, 0x3e, 0xb5, 0xde, 0x1c, 0x40, 0x06, 0xac, 0x4c, 0x2a, 0x4a, 0xa6, 0x86,
	0xe0, 0x35, 0x1e, 0x72, 0x04, 0x7b, 0x6f, 0x85, 0xe6, 0xdf, 0xf3, 0x30, 0xd0, 0x5c, 0x24, 0x13,
	0x96, 0x68, 0xfa, 0x7f, 0xb3, 0xe9, 0x0a, 0xde, 0xfb, 0xd3, 0x03, 0x70, 0x23, 0x79, 0x21, 0xa2,
	0x0f, 0x38, 0x88, 0x87, 0x00, 0x17, 0x22, 0x7a, 0xc3, 0x94, 0x0a, 0x22, 0x66, 0xc6, 0xb0, 0xe5,
	0x97, 0x10, 0x23, 0xdc, 0x74, 0x1a, 0xf3, 0x30, 0x0b, 0xd9, 0x74, 0xc2, 0x2d, 0x83, 0x28, 0x9f,
	0x6f, 0x43, 0xec, 0xe1, 0xcc, 0x0c, 0x64, 0xc7, 0xcf, 0x4c, 0xf4, 0x4c, 0xd2, 0x30, 0x64, 0x4a,
	0xb9, 0x31, 0xcc, 0x4c, 0x42, 0xa0, 0x7e, 0x19, 0xa8, 0x5b, 0x33, 0x71, 0x2d, 0xdf, 0xfc, 0xee,
	0xfd, 0x5a, 0x83, 0x7d, 0x5b, 0xdc, 0x77, 0x6c, 0x8a, 0xe4, 0xb9, 0x83, 0xac, 0x10, 0xd0, 0x85,
	0xf6, 0x90, 0xa9, 0x50, 0xf2, 0x05, 0x52, 0x66, 0x0e, 0xdf, 0xf2, 0xcb, 0x10, 0x6e, 0x6f, 0x86,
	0xb1, 0x66, 0xb7, 0x37, 0x53, 0xf8, 0x25, 0x34, 0x2f, 0x82, 0x24, 0x4a, 0xb3, 0xa3, 0xee, 0x1c,
	0x7f, 0x34, 0x70, 0xb7, 0xe6, 0x58, 0x8a, 0x48, 0x06, 0xf3, 0x39, 0x4f, 0xa2, 0x2c, 0xc4, 0xcf,
	0x83, 0x71, 0x08, 0x86, 0x62, 0x1e, 0xf0, 0xc4, 0x1d, 0xdf, 0x59, 0xe4, 0x00, 0x9a, 0x58, 0xa5,
	0x49, 0x64, 0x6f, 0xa2, 0xdc, 0xc6, 0x12, 0x9d, 0xd4, 0x8c, 0xbb, 0x61, 0x4b, 0x2c, 0x41, 0x28,
	0xa0, 0x6b, 0xae, 0xf8, 0x94, 0xc7, 0x5c, 0xdf, 0x9d, 0x4b, 0x91, 0x2e, 0x46, 0x43, 0x45, 0x9b,
	0xdd, 0x5a, 0xbf, 0xe5, 0xaf, 0xf1, 0xe0, 0xd8, 0x9f, 0x18, 0xee, 0xf2, 0xd8, 0x96, 0x89, 0x5d,
	0x42, 0x71, 0xd4, 0x86, 0xec, 0x1d, 0x8b, 0xc5, 0x82, 0xc9, 0x3c, 0x14, 0x4c, 0xe8, 0xaa, 0xc3,
	0x5c, 0x5e, 0x78, 0x15, 0xb8, 0x03, 0xb6, 0xdd, 0xe5, 0x55, 0x40, 0xbd, 0xdf, 0x3d, 0xd8, 0xc3,
	0x63, 0xd9, 0xce, 0x4c, 0x74, 0xa0, 0x53, 0xf5, 0x01, 0x25, 0x59, 0x92, 0x4c, 0xbd, 0x2a, 0x99,
	0x7d, 0xd8, 0xfc, 0x5a, 0x4a, 0x21, 0x5d, 0x17, 0xac, 0xd1, 0xfb, 0x19, 0x48, 0x96, 0x02, 0x05,
	0xa3, 0x16, 0x22, 0x51, 0x0c, 0x67, 0xb7, 0xa8, 0xc4, 0x96, 0x59, 0x00, 0xd8, 0xb8, 0x33, 0x9e,
	0x70, 0x75, 0xc3, 0x66, 0xa6, 0xcc, 0xa6, 0x9f, 0xdb, 0xe5, 0xfc, 0xb5, 0x7b, 0xf2, 0xd7, 0xcb,
	0xf9, 0x5f, 0xda, 0x6f, 0x8e, 0xe5, 0x25, 0x13, 0xec, 0x83, 0xe9, 0x7b, 0xcf, 0x00, 0x70, 0x89,
	0x53, 0x51, 0xa1, 0x2e, 0xaf, 0xac, 0xae, 0xde, 0xa0, 0x90, 0xeb, 0x0a, 0xdd, 0x99, 0xbc, 0x37,
	0x0a, 0x79, 0xf7, 0x7e, 0xf3, 0xa0, 0x71, 0x7a, 0x23, 0x78, 0xc8, 0x14, 0x79, 0x09, 0x0d, 0xbb,
	0x8b, 0xa2, 0x5e, 0xb7, 0xd6, 0x6f, 0x1f, 0x3f, 0x1d, 0x94, 0x5f, 0x18, 0x45, 0x76, 0x3f, 0x8b,
	0x23, 0xaf, 0xa0, 0x95, 0xa5, 0x53, 0x74, 0xc3, 0x2c, 0x7a, 0x5c, 0x59, 0x94, 0x0f, 0x46, 0x11,
	0x47, 0x3e, 0x83, 0x2d, 0xa3, 0x24, 0xe4, 0x0a, 0x57, 0xb4, 0x07, 0xe6, 0x61, 0x61, 0x30, 0xdf,
	0xb9, 0x7a, 0x7f, 0x79, 0xd0, 0x3e, 0x13, 0xf2, 0x36, 0x23, 0x67, 0xf9, 0xeb, 0xef, 0xad, 0xf9,
	0xfa, 0xaf, 0x39, 0x60, 0x89, 0xa8, 0xda, 0xbd, 0x63, 0x58, 0x5f, 0x1a, 0xc3, 0xf5, 0x43, 0xb6,
	0xf9, 0x2f, 0x86, 0x6c, 0xeb, 0x9f, 0x0f, 0x59, 0xe3, 0x9e, 0x21, 0x3b, 0xfe, 0xa3, 0x06, 0x6d,
	0x2c, 0xe9, 0xb5, 0xa5, 0x92, 0x8c, 0xa1, 0x73, 0xce, 0x74, 0x21, 0x1b, 0x72, 0xb8, 0xd2, 0x9e,
	0x8a, 0x9e, 0x0e, 0x3e, 0xa9, 0xf8, 0xd7, 0xe8, 0x7d, 0x02, 0x9d, 0xca, 0xcd, 0x49, 0x3e, 0x5d,
	0xb3, 0xa2, 0x7a, 0xab, 0xbe, 0x7f, 0xd3, 0x2f, 0x60, 0xc7, 0x95, 0x99, 0xe9, 0x6a, 0x3b, 0xbb,
	0x30, 0xaf, 0x05, 0x9f, 0x1d, 0xec, 0x57, 0x37, 0x70, 0x31, 0x47, 0xd8, 0x5e, 0x2d, 0x79, 0x14,
	0x31, 0x79, 0x12, 0xc7, 0x4b, 0x6b, 0x2a, 0x16, 0xb9, 0x82, 0x5d, 0x9f, 0x85, 0xff, 0x79, 0xe1,
	0x27, 0x50, 0x47, 0xc1, 0x11, 0x5a, 0x09, 0x2c, 0x69, 0xf0, 0xbd, 0x5b, 0xbc, 0xee, 0xc3, 0xf3,
	0x84, 0xe9, 0xf2, 0x03, 0xda, 0x3d, 0xa9, 0xf1, 0x0d, 0x5d, 0x5e, 0x3c, 0xdd, 0x32, 0xcf, 0xe8,
	0x57, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x2c, 0x14, 0x71, 0xeb, 0x0b, 0x00, 0x00,
}
