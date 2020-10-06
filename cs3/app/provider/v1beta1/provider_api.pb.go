// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/app/provider/v1beta1/provider_api.proto

package providerv1beta1

import (
	context "context"
	fmt "fmt"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// REQUIRED.
// View mode.
type OpenFileInAppProviderRequest_ViewMode int32

const (
	OpenFileInAppProviderRequest_VIEW_MODE_INVALID OpenFileInAppProviderRequest_ViewMode = 0
	// The file can be opened but not downloaded.
	OpenFileInAppProviderRequest_VIEW_MODE_VIEW_ONLY OpenFileInAppProviderRequest_ViewMode = 1
	// The file can be downloaded.
	OpenFileInAppProviderRequest_VIEW_MODE_READ_ONLY OpenFileInAppProviderRequest_ViewMode = 2
	// The file can be downloaded and updated.
	OpenFileInAppProviderRequest_VIEW_MODE_READ_WRITE OpenFileInAppProviderRequest_ViewMode = 3
)

var OpenFileInAppProviderRequest_ViewMode_name = map[int32]string{
	0: "VIEW_MODE_INVALID",
	1: "VIEW_MODE_VIEW_ONLY",
	2: "VIEW_MODE_READ_ONLY",
	3: "VIEW_MODE_READ_WRITE",
}

var OpenFileInAppProviderRequest_ViewMode_value = map[string]int32{
	"VIEW_MODE_INVALID":    0,
	"VIEW_MODE_VIEW_ONLY":  1,
	"VIEW_MODE_READ_ONLY":  2,
	"VIEW_MODE_READ_WRITE": 3,
}

func (x OpenFileInAppProviderRequest_ViewMode) String() string {
	return proto.EnumName(OpenFileInAppProviderRequest_ViewMode_name, int32(x))
}

func (OpenFileInAppProviderRequest_ViewMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0, 0}
}

type OpenFileInAppProviderRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The resourceInfo to be opened. The gateway grpc message has a ref instead.
	ResourceInfo *v1beta11.ResourceInfo                `protobuf:"bytes,2,opt,name=resource_info,json=resourceInfo,proto3" json:"resource_info,omitempty"`
	ViewMode     OpenFileInAppProviderRequest_ViewMode `protobuf:"varint,3,opt,name=view_mode,json=viewMode,proto3,enum=cs3.app.provider.v1beta1.OpenFileInAppProviderRequest_ViewMode" json:"view_mode,omitempty"`
	// REQUIRED.
	// The access token this application provider will use when contacting
	// the storage provider to read and write.
	// Service implementors MUST make sure that the access token only grants
	// access to the requested resource.
	// Service implementors should use a ResourceId rather than a filename to grant access, as
	// ResourceIds MUST NOT change when a resource is renamed.
	// The access token MUST be short-lived.
	// TODO(labkode): investigate token derivation techniques.
	AccessToken          string   `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenFileInAppProviderRequest) Reset()         { *m = OpenFileInAppProviderRequest{} }
func (m *OpenFileInAppProviderRequest) String() string { return proto.CompactTextString(m) }
func (*OpenFileInAppProviderRequest) ProtoMessage()    {}
func (*OpenFileInAppProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0}
}

func (m *OpenFileInAppProviderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenFileInAppProviderRequest.Unmarshal(m, b)
}
func (m *OpenFileInAppProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenFileInAppProviderRequest.Marshal(b, m, deterministic)
}
func (m *OpenFileInAppProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenFileInAppProviderRequest.Merge(m, src)
}
func (m *OpenFileInAppProviderRequest) XXX_Size() int {
	return xxx_messageInfo_OpenFileInAppProviderRequest.Size(m)
}
func (m *OpenFileInAppProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenFileInAppProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenFileInAppProviderRequest proto.InternalMessageInfo

func (m *OpenFileInAppProviderRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenFileInAppProviderRequest) GetResourceInfo() *v1beta11.ResourceInfo {
	if m != nil {
		return m.ResourceInfo
	}
	return nil
}

func (m *OpenFileInAppProviderRequest) GetViewMode() OpenFileInAppProviderRequest_ViewMode {
	if m != nil {
		return m.ViewMode
	}
	return OpenFileInAppProviderRequest_VIEW_MODE_INVALID
}

func (m *OpenFileInAppProviderRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type OpenFileInAppProviderResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The url that user agents will render to clients.
	// Usually the rendering happens by using HTML iframes or in separate browser tabs.
	AppProviderUrl       string   `protobuf:"bytes,3,opt,name=app_provider_url,json=appProviderUrl,proto3" json:"app_provider_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenFileInAppProviderResponse) Reset()         { *m = OpenFileInAppProviderResponse{} }
func (m *OpenFileInAppProviderResponse) String() string { return proto.CompactTextString(m) }
func (*OpenFileInAppProviderResponse) ProtoMessage()    {}
func (*OpenFileInAppProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{1}
}

func (m *OpenFileInAppProviderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenFileInAppProviderResponse.Unmarshal(m, b)
}
func (m *OpenFileInAppProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenFileInAppProviderResponse.Marshal(b, m, deterministic)
}
func (m *OpenFileInAppProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenFileInAppProviderResponse.Merge(m, src)
}
func (m *OpenFileInAppProviderResponse) XXX_Size() int {
	return xxx_messageInfo_OpenFileInAppProviderResponse.Size(m)
}
func (m *OpenFileInAppProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenFileInAppProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenFileInAppProviderResponse proto.InternalMessageInfo

func (m *OpenFileInAppProviderResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OpenFileInAppProviderResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenFileInAppProviderResponse) GetAppProviderUrl() string {
	if m != nil {
		return m.AppProviderUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("cs3.app.provider.v1beta1.OpenFileInAppProviderRequest_ViewMode", OpenFileInAppProviderRequest_ViewMode_name, OpenFileInAppProviderRequest_ViewMode_value)
	proto.RegisterType((*OpenFileInAppProviderRequest)(nil), "cs3.app.provider.v1beta1.OpenFileInAppProviderRequest")
	proto.RegisterType((*OpenFileInAppProviderResponse)(nil), "cs3.app.provider.v1beta1.OpenFileInAppProviderResponse")
}

func init() {
	proto.RegisterFile("cs3/app/provider/v1beta1/provider_api.proto", fileDescriptor_c007b70b037097fe)
}

var fileDescriptor_c007b70b037097fe = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0xfe, 0xa5, 0xfd, 0xa9, 0x5a, 0xdd, 0x31, 0x8a, 0x61, 0x5a, 0xa8, 0x3a, 0xa9, 0xf4, 0x54,
	0x01, 0x72, 0xd5, 0x56, 0x82, 0x23, 0x4a, 0xd7, 0x22, 0x45, 0xda, 0x96, 0xc8, 0x8c, 0x4e, 0xa0,
	0x4a, 0x91, 0x97, 0x7a, 0x28, 0xa2, 0x8d, 0xdf, 0xec, 0xa4, 0x13, 0x27, 0xae, 0x9c, 0x39, 0x71,
	0xe6, 0x88, 0xc4, 0x3f, 0xc2, 0x5f, 0x85, 0xe2, 0x38, 0xd9, 0x60, 0x14, 0x01, 0x37, 0xbf, 0xf7,
	0x7d, 0xdf, 0xf3, 0xf7, 0xde, 0xb3, 0xd1, 0xa3, 0x50, 0x8d, 0xfa, 0x0c, 0xa0, 0x0f, 0x52, 0xac,
	0xa3, 0x05, 0x97, 0xfd, 0xf5, 0xe0, 0x8c, 0x27, 0x6c, 0x50, 0x26, 0x02, 0x06, 0x11, 0x01, 0x29,
	0x12, 0x81, 0xed, 0x50, 0x8d, 0x08, 0x03, 0x20, 0x05, 0x46, 0x0c, 0xb9, 0xd5, 0xce, 0xca, 0x48,
	0x08, 0x4b, 0xb5, 0x4a, 0x58, 0x92, 0xaa, 0x5c, 0xd7, 0x7a, 0x9c, 0xa1, 0x2a, 0x11, 0x92, 0xbd,
	0xe1, 0x37, 0x2f, 0x92, 0x5c, 0x89, 0x54, 0x86, 0xbc, 0x60, 0xef, 0x67, 0xec, 0xe4, 0x1d, 0x70,
	0x55, 0x52, 0x74, 0x94, 0xc3, 0xdd, 0x8f, 0x55, 0xd4, 0xf6, 0x80, 0xc7, 0xcf, 0xa3, 0x25, 0x77,
	0x63, 0x07, 0xc0, 0x37, 0x05, 0x29, 0xbf, 0x48, 0xb9, 0x4a, 0xf0, 0x00, 0xd5, 0x04, 0xb0, 0x8b,
	0x94, 0xdb, 0x56, 0xc7, 0xea, 0x35, 0x86, 0xf7, 0x49, 0x66, 0x3b, 0x2f, 0x61, 0x0a, 0x12, 0x4f,
	0x13, 0xa8, 0x21, 0x62, 0x0f, 0xdd, 0x2a, 0x5c, 0x04, 0x51, 0x7c, 0x2e, 0xec, 0x8a, 0x56, 0x3e,
	0xd4, 0x4a, 0x63, 0xfc, 0x46, 0xd3, 0x84, 0x1a, 0x89, 0x1b, 0x9f, 0x0b, 0xba, 0x2d, 0xaf, 0x45,
	0x78, 0x8e, 0xea, 0xeb, 0x88, 0x5f, 0x06, 0x2b, 0xb1, 0xe0, 0x76, 0xb5, 0x63, 0xf5, 0x76, 0x86,
	0xcf, 0xc8, 0xa6, 0xe9, 0x91, 0xdf, 0xb5, 0x43, 0x66, 0x11, 0xbf, 0x3c, 0x12, 0x0b, 0x4e, 0xb7,
	0xd6, 0xe6, 0x84, 0x1f, 0xa0, 0x6d, 0x16, 0x86, 0x5c, 0xa9, 0x20, 0x11, 0x6f, 0x79, 0x6c, 0xff,
	0xdf, 0xb1, 0x7a, 0x75, 0xda, 0xc8, 0x73, 0x27, 0x59, 0xaa, 0xbb, 0x42, 0x5b, 0x85, 0x10, 0xef,
	0xa2, 0x3b, 0x33, 0x77, 0x7a, 0x1a, 0x1c, 0x79, 0x93, 0x69, 0xe0, 0x1e, 0xcf, 0x9c, 0x43, 0x77,
	0xd2, 0xfc, 0x0f, 0xef, 0xa1, 0xbb, 0x57, 0x69, 0x7d, 0xf2, 0x8e, 0x0f, 0x5f, 0x35, 0xad, 0x1f,
	0x01, 0x3a, 0x75, 0x26, 0x39, 0x50, 0xc1, 0x36, 0xba, 0xf7, 0x13, 0x70, 0x4a, 0xdd, 0x93, 0x69,
	0xb3, 0xda, 0xfd, 0x6a, 0xa1, 0xfd, 0x0d, 0x5d, 0x28, 0x10, 0xb1, 0xe2, 0xb8, 0x8f, 0x6a, 0xf9,
	0x9b, 0x30, 0x5b, 0xd9, 0xd3, 0xe3, 0x90, 0x10, 0x96, 0x53, 0x78, 0xa1, 0x61, 0x6a, 0x68, 0xd7,
	0xd6, 0x58, 0xf9, 0xd3, 0x35, 0xf6, 0x50, 0x93, 0x01, 0x04, 0xe5, 0xcb, 0x4d, 0xe5, 0x52, 0x0f,
	0xbf, 0x4e, 0x77, 0xd8, 0x95, 0xa5, 0x97, 0x72, 0x39, 0xfc, 0x64, 0xa1, 0x46, 0x11, 0x3b, 0xbe,
	0x8b, 0x3f, 0x58, 0x68, 0xf7, 0x97, 0xfe, 0xf1, 0x93, 0x7f, 0x5b, 0x5b, 0xeb, 0xe9, 0x5f, 0xeb,
	0xf2, 0x41, 0x8d, 0xdf, 0xa3, 0x76, 0x28, 0x56, 0x1b, 0xd5, 0xe3, 0x66, 0xe9, 0x1b, 0x22, 0x3f,
	0xfb, 0x11, 0xbe, 0xf5, 0xfa, 0x76, 0xc1, 0x32, 0xa4, 0xcf, 0x95, 0xea, 0x81, 0xe3, 0x7f, 0xa9,
	0xd8, 0x07, 0x6a, 0x44, 0x1c, 0x00, 0x52, 0x68, 0xc8, 0x6c, 0x30, 0xce, 0x08, 0xdf, 0x34, 0x34,
	0x77, 0x00, 0xe6, 0x05, 0x34, 0x37, 0xd0, 0x59, 0x4d, 0xff, 0xb3, 0xd1, 0xf7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xec, 0x6c, 0xd9, 0x05, 0x1b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProviderAPIClient is the client API for ProviderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderAPIClient interface {
	// Returns the App provider URL
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	OpenFileInAppProvider(ctx context.Context, in *OpenFileInAppProviderRequest, opts ...grpc.CallOption) (*OpenFileInAppProviderResponse, error)
}

type providerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderAPIClient(cc grpc.ClientConnInterface) ProviderAPIClient {
	return &providerAPIClient{cc}
}

func (c *providerAPIClient) OpenFileInAppProvider(ctx context.Context, in *OpenFileInAppProviderRequest, opts ...grpc.CallOption) (*OpenFileInAppProviderResponse, error) {
	out := new(OpenFileInAppProviderResponse)
	err := c.cc.Invoke(ctx, "/cs3.app.provider.v1beta1.ProviderAPI/OpenFileInAppProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderAPIServer is the server API for ProviderAPI service.
type ProviderAPIServer interface {
	// Returns the App provider URL
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	OpenFileInAppProvider(context.Context, *OpenFileInAppProviderRequest) (*OpenFileInAppProviderResponse, error)
}

// UnimplementedProviderAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProviderAPIServer struct {
}

func (*UnimplementedProviderAPIServer) OpenFileInAppProvider(ctx context.Context, req *OpenFileInAppProviderRequest) (*OpenFileInAppProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenFileInAppProvider not implemented")
}

func RegisterProviderAPIServer(s *grpc.Server, srv ProviderAPIServer) {
	s.RegisterService(&_ProviderAPI_serviceDesc, srv)
}

func _ProviderAPI_OpenFileInAppProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenFileInAppProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).OpenFileInAppProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.app.provider.v1beta1.ProviderAPI/OpenFileInAppProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).OpenFileInAppProvider(ctx, req.(*OpenFileInAppProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.app.provider.v1beta1.ProviderAPI",
	HandlerType: (*ProviderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenFileInAppProvider",
			Handler:    _ProviderAPI_OpenFileInAppProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/app/provider/v1beta1/provider_api.proto",
}
