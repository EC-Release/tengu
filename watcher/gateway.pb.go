// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package watcher

import (
	context "context"
	fmt "fmt"
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

type GatewayInfo struct {
	GtwId                string   `protobuf:"bytes,1,opt,name=gtw_id,json=gtwId,proto3" json:"gtw_id,omitempty"`
	RefId                string   `protobuf:"bytes,3,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	TimeCreated          int64    `protobuf:"varint,4,opt,name=time_created,json=timeCreated,proto3" json:"time_created,omitempty"`
	Active               bool     `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	RoutingUrl           string   `protobuf:"bytes,6,opt,name=routing_url,json=routingUrl,proto3" json:"routing_url,omitempty"`
	Zone                 string   `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayInfo) Reset()         { *m = GatewayInfo{} }
func (m *GatewayInfo) String() string { return proto.CompactTextString(m) }
func (*GatewayInfo) ProtoMessage()    {}
func (*GatewayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *GatewayInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayInfo.Unmarshal(m, b)
}
func (m *GatewayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayInfo.Marshal(b, m, deterministic)
}
func (m *GatewayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayInfo.Merge(m, src)
}
func (m *GatewayInfo) XXX_Size() int {
	return xxx_messageInfo_GatewayInfo.Size(m)
}
func (m *GatewayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayInfo proto.InternalMessageInfo

func (m *GatewayInfo) GetGtwId() string {
	if m != nil {
		return m.GtwId
	}
	return ""
}

func (m *GatewayInfo) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *GatewayInfo) GetTimeCreated() int64 {
	if m != nil {
		return m.TimeCreated
	}
	return 0
}

func (m *GatewayInfo) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *GatewayInfo) GetRoutingUrl() string {
	if m != nil {
		return m.RoutingUrl
	}
	return ""
}

func (m *GatewayInfo) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type BadGatewayList struct {
	List                 []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BadGatewayList) Reset()         { *m = BadGatewayList{} }
func (m *BadGatewayList) String() string { return proto.CompactTextString(m) }
func (*BadGatewayList) ProtoMessage()    {}
func (*BadGatewayList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *BadGatewayList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BadGatewayList.Unmarshal(m, b)
}
func (m *BadGatewayList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BadGatewayList.Marshal(b, m, deterministic)
}
func (m *BadGatewayList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadGatewayList.Merge(m, src)
}
func (m *BadGatewayList) XXX_Size() int {
	return xxx_messageInfo_BadGatewayList.Size(m)
}
func (m *BadGatewayList) XXX_DiscardUnknown() {
	xxx_messageInfo_BadGatewayList.DiscardUnknown(m)
}

var xxx_messageInfo_BadGatewayList proto.InternalMessageInfo

func (m *BadGatewayList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type GoodGatewayList struct {
	List                 []*GatewayInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GoodGatewayList) Reset()         { *m = GoodGatewayList{} }
func (m *GoodGatewayList) String() string { return proto.CompactTextString(m) }
func (*GoodGatewayList) ProtoMessage()    {}
func (*GoodGatewayList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

func (m *GoodGatewayList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodGatewayList.Unmarshal(m, b)
}
func (m *GoodGatewayList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodGatewayList.Marshal(b, m, deterministic)
}
func (m *GoodGatewayList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodGatewayList.Merge(m, src)
}
func (m *GoodGatewayList) XXX_Size() int {
	return xxx_messageInfo_GoodGatewayList.Size(m)
}
func (m *GoodGatewayList) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodGatewayList.DiscardUnknown(m)
}

var xxx_messageInfo_GoodGatewayList proto.InternalMessageInfo

func (m *GoodGatewayList) GetList() []*GatewayInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*GatewayInfo)(nil), "main.GatewayInfo")
	proto.RegisterType((*BadGatewayList)(nil), "main.BadGatewayList")
	proto.RegisterType((*GoodGatewayList)(nil), "main.GoodGatewayList")
}

func init() {
	proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5)
}

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0xc7, 0x1f, 0x3f, 0x49, 0x53, 0xb8, 0xf0, 0x22, 0x2c, 0x8a, 0x2c, 0x16, 0x42, 0x04, 0x52,
	0xa6, 0x0c, 0x65, 0xe9, 0x4c, 0x87, 0x28, 0x12, 0x53, 0x24, 0xe6, 0xc8, 0xd4, 0x4e, 0x64, 0x29,
	0x8d, 0x91, 0x7b, 0x25, 0x82, 0xaf, 0xc4, 0x97, 0x44, 0x7e, 0x19, 0x9a, 0xed, 0xfc, 0xf3, 0xef,
	0xee, 0xfc, 0x37, 0x5c, 0xf6, 0x1c, 0xe5, 0xc4, 0xbf, 0xcb, 0x4f, 0xa3, 0x51, 0xd3, 0x78, 0xcf,
	0xd5, 0x98, 0xff, 0x12, 0x48, 0x2b, 0xcf, 0xeb, 0xb1, 0xd3, 0x74, 0x05, 0x49, 0x8f, 0x53, 0xab,
	0x04, 0x23, 0x19, 0x29, 0xce, 0x9b, 0x45, 0x8f, 0x53, 0x2d, 0x2c, 0x36, 0xb2, 0xb3, 0x38, 0xf2,
	0xd8, 0xc8, 0xae, 0x16, 0xf4, 0x11, 0x2e, 0x50, 0xed, 0x65, 0xbb, 0x33, 0x92, 0xa3, 0x14, 0x2c,
	0xce, 0x48, 0x11, 0x35, 0xa9, 0x65, 0x5b, 0x8f, 0xe8, 0x1d, 0x24, 0x7c, 0x87, 0xea, 0x4b, 0xb2,
	0x45, 0x46, 0x8a, 0xb3, 0x26, 0x9c, 0xe8, 0x03, 0xa4, 0x46, 0x1f, 0x51, 0x8d, 0x7d, 0x7b, 0x34,
	0x03, 0x4b, 0xdc, 0x58, 0x08, 0xe8, 0xdd, 0x0c, 0x94, 0x42, 0xfc, 0xa3, 0x47, 0xc9, 0xfe, 0xbb,
	0x1b, 0x57, 0xe7, 0x4f, 0x70, 0xf5, 0xca, 0x45, 0x78, 0xef, 0x9b, 0x3a, 0xa0, 0xb5, 0x06, 0x75,
	0x40, 0x46, 0xb2, 0xc8, 0x5a, 0xb6, 0xce, 0x37, 0x70, 0x5d, 0x69, 0x3d, 0xd3, 0x9e, 0x4f, 0xb4,
	0x74, 0x7d, 0x53, 0xda, 0xec, 0xe5, 0x49, 0x6e, 0xdf, 0xb9, 0xde, 0xc2, 0x32, 0x40, 0xba, 0x81,
	0x65, 0x25, 0xd1, 0x35, 0xdf, 0x7a, 0x7d, 0xbe, 0xf9, 0x7e, 0x15, 0x86, 0xcc, 0x37, 0xe5, 0xff,
	0x3e, 0x12, 0xf7, 0xbf, 0x2f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xca, 0x89, 0xf3, 0x70,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	GetList(ctx context.Context, in *BadGatewayList, opts ...grpc.CallOption) (*GoodGatewayList, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetList(ctx context.Context, in *BadGatewayList, opts ...grpc.CallOption) (*GoodGatewayList, error) {
	out := new(GoodGatewayList)
	err := c.cc.Invoke(ctx, "/main.Gateway/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	GetList(context.Context, *BadGatewayList) (*GoodGatewayList, error)
}

// UnimplementedGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (*UnimplementedGatewayServer) GetList(ctx context.Context, req *BadGatewayList) (*GoodGatewayList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BadGatewayList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Gateway/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetList(ctx, req.(*BadGatewayList))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _Gateway_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
