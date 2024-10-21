// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/themes/themes.proto
// DO NOT EDIT!

/*
Package themes is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/themes/themes.proto

It has these top-level messages:
	ThemeResponse
	HostThemeRequest
	Image
	Text
	CSS
	URLRequest
	ServeResponse
	HTMLTemplate
*/
package themes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "golang.conradwood.net/apis/common"
import h2gproxy "golang.conradwood.net/apis/h2gproxy"

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

type ThemeResponse struct {
	SmallLogoName string `protobuf:"bytes,1,opt,name=SmallLogoName" json:"SmallLogoName,omitempty"`
	FavIconName   string `protobuf:"bytes,2,opt,name=FavIconName" json:"FavIconName,omitempty"`
	HeaderText    string `protobuf:"bytes,3,opt,name=HeaderText" json:"HeaderText,omitempty"`
	CorporateCss  string `protobuf:"bytes,4,opt,name=CorporateCss" json:"CorporateCss,omitempty"`
	ThemeName     string `protobuf:"bytes,5,opt,name=ThemeName" json:"ThemeName,omitempty"`
}

func (m *ThemeResponse) Reset()                    { *m = ThemeResponse{} }
func (m *ThemeResponse) String() string            { return proto.CompactTextString(m) }
func (*ThemeResponse) ProtoMessage()               {}
func (*ThemeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ThemeResponse) GetSmallLogoName() string {
	if m != nil {
		return m.SmallLogoName
	}
	return ""
}

func (m *ThemeResponse) GetFavIconName() string {
	if m != nil {
		return m.FavIconName
	}
	return ""
}

func (m *ThemeResponse) GetHeaderText() string {
	if m != nil {
		return m.HeaderText
	}
	return ""
}

func (m *ThemeResponse) GetCorporateCss() string {
	if m != nil {
		return m.CorporateCss
	}
	return ""
}

func (m *ThemeResponse) GetThemeName() string {
	if m != nil {
		return m.ThemeName
	}
	return ""
}

type HostThemeRequest struct {
	Host      string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
	UserAgent string `protobuf:"bytes,2,opt,name=UserAgent" json:"UserAgent,omitempty"`
}

func (m *HostThemeRequest) Reset()                    { *m = HostThemeRequest{} }
func (m *HostThemeRequest) String() string            { return proto.CompactTextString(m) }
func (*HostThemeRequest) ProtoMessage()               {}
func (*HostThemeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HostThemeRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HostThemeRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

type Image struct {
	Filename string `protobuf:"bytes,1,opt,name=Filename" json:"Filename,omitempty"`
	MimeType string `protobuf:"bytes,2,opt,name=MimeType" json:"MimeType,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Image) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Image) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Text struct {
	Text string `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
}

func (m *Text) Reset()                    { *m = Text{} }
func (m *Text) String() string            { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()               {}
func (*Text) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CSS struct {
	Filename string `protobuf:"bytes,1,opt,name=Filename" json:"Filename,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *CSS) Reset()                    { *m = CSS{} }
func (m *CSS) String() string            { return proto.CompactTextString(m) }
func (*CSS) ProtoMessage()               {}
func (*CSS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CSS) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *CSS) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type URLRequest struct {
	Host string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=Path" json:"Path,omitempty"`
}

func (m *URLRequest) Reset()                    { *m = URLRequest{} }
func (m *URLRequest) String() string            { return proto.CompactTextString(m) }
func (*URLRequest) ProtoMessage()               {}
func (*URLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *URLRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *URLRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ServeResponse struct {
	Response        *h2gproxy.ServeResponse `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
	IsAuthoritative bool                    `protobuf:"varint,2,opt,name=IsAuthoritative" json:"IsAuthoritative,omitempty"`
}

func (m *ServeResponse) Reset()                    { *m = ServeResponse{} }
func (m *ServeResponse) String() string            { return proto.CompactTextString(m) }
func (*ServeResponse) ProtoMessage()               {}
func (*ServeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServeResponse) GetResponse() *h2gproxy.ServeResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ServeResponse) GetIsAuthoritative() bool {
	if m != nil {
		return m.IsAuthoritative
	}
	return false
}

type HTMLTemplate struct {
	Prefix        []byte `protobuf:"bytes,1,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	SectionPrefix []byte `protobuf:"bytes,2,opt,name=SectionPrefix,proto3" json:"SectionPrefix,omitempty"`
	SectionSuffix []byte `protobuf:"bytes,3,opt,name=SectionSuffix,proto3" json:"SectionSuffix,omitempty"`
	Suffix        []byte `protobuf:"bytes,4,opt,name=Suffix,proto3" json:"Suffix,omitempty"`
}

func (m *HTMLTemplate) Reset()                    { *m = HTMLTemplate{} }
func (m *HTMLTemplate) String() string            { return proto.CompactTextString(m) }
func (*HTMLTemplate) ProtoMessage()               {}
func (*HTMLTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *HTMLTemplate) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *HTMLTemplate) GetSectionPrefix() []byte {
	if m != nil {
		return m.SectionPrefix
	}
	return nil
}

func (m *HTMLTemplate) GetSectionSuffix() []byte {
	if m != nil {
		return m.SectionSuffix
	}
	return nil
}

func (m *HTMLTemplate) GetSuffix() []byte {
	if m != nil {
		return m.Suffix
	}
	return nil
}

func init() {
	proto.RegisterType((*ThemeResponse)(nil), "themes.ThemeResponse")
	proto.RegisterType((*HostThemeRequest)(nil), "themes.HostThemeRequest")
	proto.RegisterType((*Image)(nil), "themes.Image")
	proto.RegisterType((*Text)(nil), "themes.Text")
	proto.RegisterType((*CSS)(nil), "themes.CSS")
	proto.RegisterType((*URLRequest)(nil), "themes.URLRequest")
	proto.RegisterType((*ServeResponse)(nil), "themes.ServeResponse")
	proto.RegisterType((*HTMLTemplate)(nil), "themes.HTMLTemplate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Themes service

type ThemesClient interface {
	// DEPRECATED
	GetHTMLTheme(ctx context.Context, in *h2gproxy.ServeRequest, opts ...grpc.CallOption) (*ThemeResponse, error)
	// get some defaults for a given webrequest. this proto is expected to grow
	GetThemeByHost(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*ThemeResponse, error)
	// get logo for this site
	GetLogo(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*Image, error)
	GetFavIcon(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*Image, error)
	GetHeaderText(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*Text, error)
	GetCSS(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*CSS, error)
	ServeURL(ctx context.Context, in *h2gproxy.ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error)
	//
	// if you have an html snippet, use this
	GetHTMLTemplate(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*HTMLTemplate, error)
}

type themesClient struct {
	cc *grpc.ClientConn
}

func NewThemesClient(cc *grpc.ClientConn) ThemesClient {
	return &themesClient{cc}
}

func (c *themesClient) GetHTMLTheme(ctx context.Context, in *h2gproxy.ServeRequest, opts ...grpc.CallOption) (*ThemeResponse, error) {
	out := new(ThemeResponse)
	err := grpc.Invoke(ctx, "/themes.Themes/GetHTMLTheme", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themesClient) GetThemeByHost(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*ThemeResponse, error) {
	out := new(ThemeResponse)
	err := grpc.Invoke(ctx, "/themes.Themes/GetThemeByHost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themesClient) GetLogo(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/themes.Themes/GetLogo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themesClient) GetFavIcon(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/themes.Themes/GetFavIcon", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themesClient) GetHeaderText(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := grpc.Invoke(ctx, "/themes.Themes/GetHeaderText", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themesClient) GetCSS(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*CSS, error) {
	out := new(CSS)
	err := grpc.Invoke(ctx, "/themes.Themes/GetCSS", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themesClient) ServeURL(ctx context.Context, in *h2gproxy.ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error) {
	out := new(ServeResponse)
	err := grpc.Invoke(ctx, "/themes.Themes/ServeURL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themesClient) GetHTMLTemplate(ctx context.Context, in *HostThemeRequest, opts ...grpc.CallOption) (*HTMLTemplate, error) {
	out := new(HTMLTemplate)
	err := grpc.Invoke(ctx, "/themes.Themes/GetHTMLTemplate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Themes service

type ThemesServer interface {
	// DEPRECATED
	GetHTMLTheme(context.Context, *h2gproxy.ServeRequest) (*ThemeResponse, error)
	// get some defaults for a given webrequest. this proto is expected to grow
	GetThemeByHost(context.Context, *HostThemeRequest) (*ThemeResponse, error)
	// get logo for this site
	GetLogo(context.Context, *HostThemeRequest) (*Image, error)
	GetFavIcon(context.Context, *HostThemeRequest) (*Image, error)
	GetHeaderText(context.Context, *HostThemeRequest) (*Text, error)
	GetCSS(context.Context, *HostThemeRequest) (*CSS, error)
	ServeURL(context.Context, *h2gproxy.ServeRequest) (*ServeResponse, error)
	//
	// if you have an html snippet, use this
	GetHTMLTemplate(context.Context, *HostThemeRequest) (*HTMLTemplate, error)
}

func RegisterThemesServer(s *grpc.Server, srv ThemesServer) {
	s.RegisterService(&_Themes_serviceDesc, srv)
}

func _Themes_GetHTMLTheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(h2gproxy.ServeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemesServer).GetHTMLTheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themes.Themes/GetHTMLTheme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemesServer).GetHTMLTheme(ctx, req.(*h2gproxy.ServeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Themes_GetThemeByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemesServer).GetThemeByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themes.Themes/GetThemeByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemesServer).GetThemeByHost(ctx, req.(*HostThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Themes_GetLogo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemesServer).GetLogo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themes.Themes/GetLogo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemesServer).GetLogo(ctx, req.(*HostThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Themes_GetFavIcon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemesServer).GetFavIcon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themes.Themes/GetFavIcon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemesServer).GetFavIcon(ctx, req.(*HostThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Themes_GetHeaderText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemesServer).GetHeaderText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themes.Themes/GetHeaderText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemesServer).GetHeaderText(ctx, req.(*HostThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Themes_GetCSS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemesServer).GetCSS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themes.Themes/GetCSS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemesServer).GetCSS(ctx, req.(*HostThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Themes_ServeURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(h2gproxy.ServeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemesServer).ServeURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themes.Themes/ServeURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemesServer).ServeURL(ctx, req.(*h2gproxy.ServeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Themes_GetHTMLTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemesServer).GetHTMLTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themes.Themes/GetHTMLTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemesServer).GetHTMLTemplate(ctx, req.(*HostThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Themes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "themes.Themes",
	HandlerType: (*ThemesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHTMLTheme",
			Handler:    _Themes_GetHTMLTheme_Handler,
		},
		{
			MethodName: "GetThemeByHost",
			Handler:    _Themes_GetThemeByHost_Handler,
		},
		{
			MethodName: "GetLogo",
			Handler:    _Themes_GetLogo_Handler,
		},
		{
			MethodName: "GetFavIcon",
			Handler:    _Themes_GetFavIcon_Handler,
		},
		{
			MethodName: "GetHeaderText",
			Handler:    _Themes_GetHeaderText_Handler,
		},
		{
			MethodName: "GetCSS",
			Handler:    _Themes_GetCSS_Handler,
		},
		{
			MethodName: "ServeURL",
			Handler:    _Themes_ServeURL_Handler,
		},
		{
			MethodName: "GetHTMLTemplate",
			Handler:    _Themes_GetHTMLTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/themes/themes.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/themes/themes.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0x16, 0x81, 0x50, 0x32, 0x31, 0x4d, 0xb5, 0x6a, 0x53, 0x64, 0x55, 0x55, 0x6a, 0xf5, 0x90,
	0x93, 0x91, 0x48, 0x51, 0xd5, 0x43, 0x55, 0x11, 0xa2, 0x00, 0x12, 0xa9, 0x22, 0x1b, 0x2e, 0xbd,
	0x6d, 0x61, 0x62, 0x2c, 0x61, 0xaf, 0x6b, 0x2f, 0x14, 0x5e, 0xa1, 0x6f, 0xd4, 0x87, 0xe9, 0xbb,
	0x54, 0x3b, 0x5e, 0x1b, 0x3b, 0x15, 0xa4, 0xa7, 0x9d, 0xf9, 0xe6, 0x9b, 0x9d, 0x9f, 0x9d, 0x59,
	0xe8, 0x46, 0xb1, 0x90, 0x22, 0x69, 0x7b, 0x62, 0xc9, 0x43, 0xcf, 0x9e, 0x89, 0x30, 0xe6, 0xf3,
	0x9f, 0x42, 0xcc, 0xed, 0x10, 0x65, 0x9b, 0x47, 0x7e, 0xd2, 0x96, 0x0b, 0x0c, 0x30, 0x3b, 0x6c,
	0xe2, 0xb3, 0x7a, 0xaa, 0x99, 0xf6, 0x01, 0xbf, 0x99, 0x08, 0x02, 0x11, 0xea, 0x23, 0xf5, 0x33,
	0x3b, 0x07, 0xf8, 0x8b, 0x8e, 0x17, 0xc5, 0x62, 0xb3, 0xcd, 0x85, 0xd4, 0xc7, 0xfa, 0x5d, 0x81,
	0xe6, 0x44, 0x85, 0x73, 0x30, 0x89, 0x44, 0x98, 0x20, 0x7b, 0x0f, 0x4d, 0x37, 0xe0, 0xcb, 0xe5,
	0x58, 0x78, 0xe2, 0x2b, 0x0f, 0xb0, 0x55, 0xb9, 0xa8, 0x5c, 0x9e, 0x38, 0x65, 0x90, 0x5d, 0xc0,
	0xe9, 0x2d, 0x5f, 0x8f, 0x66, 0x22, 0x24, 0xce, 0x11, 0x71, 0x8a, 0x10, 0x7b, 0x0b, 0x30, 0x44,
	0x3e, 0xc7, 0x78, 0x82, 0x1b, 0xd9, 0xaa, 0x12, 0xa1, 0x80, 0x30, 0x0b, 0x8c, 0xbe, 0x88, 0x23,
	0x11, 0x73, 0x89, 0xfd, 0x24, 0x69, 0xd5, 0x88, 0x51, 0xc2, 0xd8, 0x1b, 0x38, 0xa1, 0xe4, 0x28,
	0xc6, 0x31, 0x11, 0x76, 0x80, 0x75, 0x03, 0x2f, 0x86, 0x22, 0x91, 0x3a, 0xfd, 0x1f, 0x2b, 0x4c,
	0x24, 0x63, 0x50, 0x53, 0x98, 0x4e, 0x9a, 0x64, 0x75, 0xcb, 0x34, 0xc1, 0xb8, 0xe7, 0x61, 0x28,
	0x75, 0xa6, 0x3b, 0xc0, 0x72, 0xe1, 0x78, 0x14, 0x70, 0x0f, 0x99, 0x09, 0x8d, 0x5b, 0x7f, 0x89,
	0xe1, 0xae, 0xe6, 0x5c, 0x57, 0xb6, 0x3b, 0x3f, 0xc0, 0xc9, 0x36, 0xca, 0x6a, 0xcd, 0x75, 0x15,
	0xf2, 0x86, 0x4b, 0x4e, 0x25, 0x1a, 0x0e, 0xc9, 0x96, 0x09, 0x35, 0x2a, 0x92, 0xa5, 0x67, 0x96,
	0x8e, 0x92, 0xad, 0x2e, 0x54, 0xfb, 0xae, 0x7b, 0x30, 0x5c, 0x76, 0xe5, 0x51, 0xe1, 0xca, 0x0f,
	0x00, 0x53, 0x67, 0x7c, 0xa8, 0x4e, 0x06, 0xb5, 0x7b, 0x2e, 0x17, 0x3a, 0x41, 0x92, 0xad, 0x10,
	0x9a, 0x2e, 0xc6, 0xeb, 0xdd, 0xf3, 0x5e, 0x41, 0x23, 0x93, 0xc9, 0xf9, 0xb4, 0xf3, 0xda, 0xce,
	0x67, 0xa2, 0x44, 0x75, 0x72, 0x22, 0xbb, 0x84, 0xb3, 0x51, 0xd2, 0x5b, 0xc9, 0x85, 0x88, 0x7d,
	0xc9, 0xa5, 0xbf, 0x4e, 0xbb, 0xd0, 0x70, 0x1e, 0xc3, 0xd6, 0xaf, 0x0a, 0x18, 0xc3, 0xc9, 0xdd,
	0x78, 0x82, 0x41, 0xb4, 0xe4, 0x12, 0xd9, 0x39, 0xd4, 0xef, 0x63, 0x7c, 0xf0, 0x37, 0x14, 0xcd,
	0x70, 0xb4, 0x46, 0x63, 0x86, 0x33, 0xe9, 0x8b, 0x50, 0x9b, 0xd3, 0x5a, 0xcb, 0x60, 0x81, 0xe5,
	0xae, 0x1e, 0x14, 0xab, 0x5a, 0x62, 0xa5, 0xa0, 0x8a, 0xa1, 0xcd, 0xb5, 0x34, 0x46, 0xaa, 0x75,
	0xfe, 0x54, 0xa1, 0x4e, 0xd3, 0x91, 0xb0, 0xcf, 0x60, 0x0c, 0x50, 0x52, 0x66, 0x0a, 0x60, 0xe7,
	0xff, 0x14, 0x4d, 0x7d, 0x35, 0x5f, 0xd9, 0x7a, 0x15, 0xcb, 0x4b, 0xd1, 0x83, 0xe7, 0x03, 0x4c,
	0x27, 0xed, 0x7a, 0x4b, 0xcd, 0x6e, 0x65, 0xc4, 0xc7, 0x23, 0xb8, 0xef, 0x8a, 0x0e, 0x3c, 0x1b,
	0xa0, 0x54, 0x0b, 0x74, 0xc0, 0xb7, 0x99, 0x59, 0xd2, 0x91, 0xec, 0x02, 0x0c, 0x50, 0xea, 0xad,
	0xfa, 0x7f, 0xb7, 0x8f, 0xd0, 0x54, 0xc5, 0xee, 0x76, 0x6d, 0xbf, 0xa7, 0x91, 0x27, 0xab, 0x78,
	0x6d, 0xa8, 0x0f, 0x50, 0xaa, 0xe9, 0xdc, 0xef, 0x71, 0x9a, 0x59, 0x14, 0xed, 0x13, 0x34, 0xa8,
	0x7d, 0x53, 0x67, 0xfc, 0x74, 0x4b, 0xcb, 0x83, 0xd8, 0x83, 0xb3, 0xec, 0x45, 0xb2, 0x59, 0xd9,
	0x1f, 0xf4, 0x65, 0x6e, 0x29, 0xf0, 0xaf, 0xbf, 0x80, 0x19, 0xa2, 0x2c, 0xfe, 0x77, 0xea, 0xaf,
	0xd3, 0xd4, 0x6f, 0xef, 0x9e, 0xfc, 0x76, 0xbf, 0xd7, 0xe9, 0x13, 0xbc, 0xfa, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x09, 0x9c, 0xf5, 0x0c, 0xa9, 0x05, 0x00, 0x00,
}
