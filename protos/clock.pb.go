// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: clock.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetHealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHealthRequest) Reset() {
	*x = GetHealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthRequest) ProtoMessage() {}

func (x *GetHealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthRequest.ProtoReflect.Descriptor instead.
func (*GetHealthRequest) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{0}
}

type GetHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsHealthy bool `protobuf:"varint,1,opt,name=isHealthy,proto3" json:"isHealthy,omitempty"`
}

func (x *GetHealthResponse) Reset() {
	*x = GetHealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthResponse) ProtoMessage() {}

func (x *GetHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthResponse.ProtoReflect.Descriptor instead.
func (*GetHealthResponse) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{1}
}

func (x *GetHealthResponse) GetIsHealthy() bool {
	if x != nil {
		return x.IsHealthy
	}
	return false
}

type GetTimestampRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTimestampRequest) Reset() {
	*x = GetTimestampRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimestampRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimestampRequest) ProtoMessage() {}

func (x *GetTimestampRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimestampRequest.ProtoReflect.Descriptor instead.
func (*GetTimestampRequest) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{2}
}

type GetTimestampResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetTimestampResponse) Reset() {
	*x = GetTimestampResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimestampResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimestampResponse) ProtoMessage() {}

func (x *GetTimestampResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimestampResponse.ProtoReflect.Descriptor instead.
func (*GetTimestampResponse) Descriptor() ([]byte, []int) {
	return file_clock_proto_rawDescGZIP(), []int{3}
}

func (x *GetTimestampResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_clock_proto protoreflect.FileDescriptor

var file_clock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x34, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xb9, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x12, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12,
	0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x12, 0x57, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clock_proto_rawDescOnce sync.Once
	file_clock_proto_rawDescData = file_clock_proto_rawDesc
)

func file_clock_proto_rawDescGZIP() []byte {
	file_clock_proto_rawDescOnce.Do(func() {
		file_clock_proto_rawDescData = protoimpl.X.CompressGZIP(file_clock_proto_rawDescData)
	})
	return file_clock_proto_rawDescData
}

var file_clock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_clock_proto_goTypes = []interface{}{
	(*GetHealthRequest)(nil),     // 0: clock.GetHealthRequest
	(*GetHealthResponse)(nil),    // 1: clock.GetHealthResponse
	(*GetTimestampRequest)(nil),  // 2: clock.GetTimestampRequest
	(*GetTimestampResponse)(nil), // 3: clock.GetTimestampResponse
}
var file_clock_proto_depIdxs = []int32{
	0, // 0: clock.ClockService.GetHealth:input_type -> clock.GetHealthRequest
	2, // 1: clock.ClockService.GetTimestamp:input_type -> clock.GetTimestampRequest
	1, // 2: clock.ClockService.GetHealth:output_type -> clock.GetHealthResponse
	3, // 3: clock.ClockService.GetTimestamp:output_type -> clock.GetTimestampResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clock_proto_init() }
func file_clock_proto_init() {
	if File_clock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHealthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHealthResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimestampRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimestampResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_clock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clock_proto_goTypes,
		DependencyIndexes: file_clock_proto_depIdxs,
		MessageInfos:      file_clock_proto_msgTypes,
	}.Build()
	File_clock_proto = out.File
	file_clock_proto_rawDesc = nil
	file_clock_proto_goTypes = nil
	file_clock_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClockServiceClient is the client API for ClockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClockServiceClient interface {
	GetHealth(ctx context.Context, in *GetHealthRequest, opts ...grpc.CallOption) (*GetHealthResponse, error)
	GetTimestamp(ctx context.Context, in *GetTimestampRequest, opts ...grpc.CallOption) (*GetTimestampResponse, error)
}

type clockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClockServiceClient(cc grpc.ClientConnInterface) ClockServiceClient {
	return &clockServiceClient{cc}
}

func (c *clockServiceClient) GetHealth(ctx context.Context, in *GetHealthRequest, opts ...grpc.CallOption) (*GetHealthResponse, error) {
	out := new(GetHealthResponse)
	err := c.cc.Invoke(ctx, "/clock.ClockService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockServiceClient) GetTimestamp(ctx context.Context, in *GetTimestampRequest, opts ...grpc.CallOption) (*GetTimestampResponse, error) {
	out := new(GetTimestampResponse)
	err := c.cc.Invoke(ctx, "/clock.ClockService/GetTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClockServiceServer is the server API for ClockService service.
type ClockServiceServer interface {
	GetHealth(context.Context, *GetHealthRequest) (*GetHealthResponse, error)
	GetTimestamp(context.Context, *GetTimestampRequest) (*GetTimestampResponse, error)
}

// UnimplementedClockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClockServiceServer struct {
}

func (*UnimplementedClockServiceServer) GetHealth(context.Context, *GetHealthRequest) (*GetHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (*UnimplementedClockServiceServer) GetTimestamp(context.Context, *GetTimestampRequest) (*GetTimestampResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimestamp not implemented")
}

func RegisterClockServiceServer(s *grpc.Server, srv ClockServiceServer) {
	s.RegisterService(&_ClockService_serviceDesc, srv)
}

func _ClockService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clock.ClockService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServiceServer).GetHealth(ctx, req.(*GetHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClockService_GetTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimestampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServiceServer).GetTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clock.ClockService/GetTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServiceServer).GetTimestamp(ctx, req.(*GetTimestampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clock.ClockService",
	HandlerType: (*ClockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _ClockService_GetHealth_Handler,
		},
		{
			MethodName: "GetTimestamp",
			Handler:    _ClockService_GetTimestamp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clock.proto",
}
