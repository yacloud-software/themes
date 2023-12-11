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
	Host string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
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

func init() {
	proto.RegisterType((*ThemeResponse)(nil), "themes.ThemeResponse")
	proto.RegisterType((*HostThemeRequest)(nil), "themes.HostThemeRequest")
	proto.RegisterType((*Image)(nil), "themes.Image")
	proto.RegisterType((*Text)(nil), "themes.Text")
	proto.RegisterType((*CSS)(nil), "themes.CSS")
	proto.RegisterType((*URLRequest)(nil), "themes.URLRequest")
	proto.RegisterType((*ServeResponse)(nil), "themes.ServeResponse")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/themes/themes.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/themes/themes.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x9a, 0x34, 0xa4, 0x93, 0x04, 0xd0, 0x4a, 0x40, 0x64, 0x21, 0x54, 0x2c, 0x84, 0x7a,
	0x72, 0x24, 0x97, 0x08, 0x71, 0x40, 0xa8, 0x0d, 0xaa, 0x13, 0x29, 0x45, 0xc8, 0x4e, 0x2f, 0xdc,
	0x96, 0x64, 0xe4, 0x58, 0x8a, 0x77, 0x8d, 0x77, 0x1b, 0x9a, 0xbf, 0xc2, 0x3f, 0xe1, 0xdf, 0xa1,
	0x1d, 0x7f, 0x24, 0x2e, 0x8a, 0xdb, 0x93, 0x67, 0xdf, 0xbe, 0xb7, 0x33, 0x6f, 0x67, 0xbc, 0x30,
	0x4a, 0x52, 0xa9, 0xa5, 0x1a, 0x86, 0x72, 0xcd, 0x45, 0xe8, 0x2c, 0xa4, 0x48, 0xf9, 0xf2, 0xb7,
	0x94, 0x4b, 0x47, 0xa0, 0x1e, 0xf2, 0x24, 0x52, 0x43, 0xbd, 0xc2, 0x18, 0x8b, 0x8f, 0x43, 0x7c,
	0xd6, 0xce, 0x56, 0x96, 0x53, 0xa3, 0x5b, 0xc8, 0x38, 0x96, 0x22, 0xff, 0x64, 0x3a, 0xcb, 0xad,
	0xe1, 0xaf, 0xdc, 0x30, 0x49, 0xe5, 0xdd, 0xb6, 0x0c, 0x32, 0x8d, 0xfd, 0xb7, 0x01, 0xfd, 0xb9,
	0x49, 0xe7, 0xa3, 0x4a, 0xa4, 0x50, 0xc8, 0xde, 0x41, 0x3f, 0x88, 0xf9, 0x7a, 0x3d, 0x93, 0xa1,
	0xfc, 0xc6, 0x63, 0x1c, 0x34, 0x4e, 0x1b, 0x67, 0x27, 0x7e, 0x15, 0x64, 0xa7, 0xd0, 0xbd, 0xe2,
	0x9b, 0xe9, 0x42, 0x0a, 0xe2, 0x1c, 0x11, 0x67, 0x1f, 0x62, 0x6f, 0x00, 0x26, 0xc8, 0x97, 0x98,
	0xce, 0xf1, 0x4e, 0x0f, 0x9a, 0x44, 0xd8, 0x43, 0x98, 0x0d, 0xbd, 0xb1, 0x4c, 0x13, 0x99, 0x72,
	0x8d, 0x63, 0xa5, 0x06, 0x2d, 0x62, 0x54, 0x30, 0xf6, 0x1a, 0x4e, 0xa8, 0x38, 0xca, 0x71, 0x4c,
	0x84, 0x1d, 0x60, 0xbf, 0x87, 0xe7, 0x13, 0xa9, 0x74, 0x5e, 0xfe, 0xaf, 0x5b, 0x54, 0x9a, 0x31,
	0x68, 0x19, 0x2c, 0x2f, 0x9a, 0x62, 0x3b, 0x80, 0xe3, 0x69, 0xcc, 0x43, 0x64, 0x16, 0x74, 0xae,
	0xa2, 0x35, 0x8a, 0x9d, 0xab, 0x72, 0x6d, 0xf6, 0xae, 0xa3, 0x18, 0xe7, 0xdb, 0xa4, 0x70, 0x53,
	0xae, 0xcd, 0xa1, 0x5f, 0xb9, 0xe6, 0x64, 0xa2, 0xe7, 0x53, 0x6c, 0x5b, 0xd0, 0x22, 0x1b, 0x2c,
	0xfb, 0x16, 0x09, 0x4d, 0x6c, 0x8f, 0xa0, 0x39, 0x0e, 0x82, 0xda, 0x74, 0xc5, 0x91, 0x47, 0x7b,
	0x47, 0x7e, 0x00, 0xb8, 0xf1, 0x67, 0x35, 0x4e, 0x0c, 0xf6, 0x9d, 0xeb, 0x55, 0x5e, 0x20, 0xc5,
	0xb6, 0x80, 0x7e, 0x80, 0xe9, 0x66, 0xd7, 0xc0, 0x73, 0xe8, 0x14, 0x31, 0x89, 0xbb, 0xee, 0x2b,
	0xa7, 0xec, 0x7a, 0x85, 0xea, 0x97, 0x44, 0x76, 0x06, 0xcf, 0xa6, 0xea, 0xe2, 0x56, 0xaf, 0x64,
	0x1a, 0x69, 0xae, 0xa3, 0x4d, 0x76, 0x0b, 0x1d, 0xff, 0x3e, 0xec, 0xfe, 0x69, 0x42, 0x9b, 0xae,
	0x5c, 0xb1, 0xcf, 0xd0, 0xf3, 0x50, 0x4f, 0xe6, 0xd7, 0x33, 0x02, 0xd8, 0xcb, 0xff, 0xf2, 0x90,
	0x15, 0xeb, 0x85, 0x93, 0xcf, 0x77, 0x75, 0xd2, 0x2e, 0xe0, 0xa9, 0x87, 0x59, 0xfb, 0x2e, 0xb7,
	0xe4, 0x6f, 0x50, 0x10, 0xef, 0xf7, 0xf5, 0xd0, 0x11, 0x2e, 0x3c, 0xf1, 0x50, 0x9b, 0xa9, 0xac,
	0xd1, 0xf6, 0x8b, 0x9d, 0x6c, 0x0a, 0x46, 0x00, 0x1e, 0xea, 0x7c, 0x54, 0x1f, 0x2f, 0xfb, 0x08,
	0x7d, 0x63, 0x76, 0x37, 0xc0, 0x87, 0x95, 0xbd, 0xb2, 0x58, 0xc3, 0x1b, 0x42, 0xdb, 0x43, 0x6d,
	0x06, 0xe2, 0xb0, 0xa2, 0x5b, 0xec, 0x18, 0xda, 0x27, 0xe8, 0xd0, 0xf5, 0xdd, 0xf8, 0xb3, 0x87,
	0xaf, 0xb4, 0xd2, 0xd0, 0xcb, 0x2f, 0x60, 0x09, 0xd4, 0xfb, 0x2f, 0x80, 0xf9, 0xfb, 0x73, 0xee,
	0x8f, 0xb7, 0x0f, 0x3e, 0x44, 0x3f, 0xdb, 0xf4, 0x2c, 0x9c, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0xcc, 0x84, 0x4a, 0xbb, 0x04, 0x00, 0x00,
}


