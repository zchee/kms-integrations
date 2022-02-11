// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: fakekms/fault/math.proto

package mathpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fakekms_fault_math_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fakekms_fault_math_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_fakekms_fault_math_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *AddRequest) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fakekms_fault_math_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fakekms_fault_math_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_fakekms_fault_math_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type MultiplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *MultiplyRequest) Reset() {
	*x = MultiplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fakekms_fault_math_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyRequest) ProtoMessage() {}

func (x *MultiplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fakekms_fault_math_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyRequest.ProtoReflect.Descriptor instead.
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return file_fakekms_fault_math_proto_rawDescGZIP(), []int{2}
}

func (x *MultiplyRequest) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MultiplyRequest) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type MultiplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product int64 `protobuf:"varint,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *MultiplyResponse) Reset() {
	*x = MultiplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fakekms_fault_math_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyResponse) ProtoMessage() {}

func (x *MultiplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fakekms_fault_math_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyResponse.ProtoReflect.Descriptor instead.
func (*MultiplyResponse) Descriptor() ([]byte, []int) {
	return file_fakekms_fault_math_proto_rawDescGZIP(), []int{3}
}

func (x *MultiplyResponse) GetProduct() int64 {
	if x != nil {
		return x.Product
	}
	return 0
}

var File_fakekms_fault_math_proto protoreflect.FileDescriptor

var file_fakekms_fault_math_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x61, 0x6b, 0x65, 0x6b, 0x6d, 0x73, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f,
	0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x74, 0x68,
	0x22, 0x28, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x22, 0x1f, 0x0a, 0x0b, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x2d, 0x0a, 0x0f, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x22, 0x2c, 0x0a, 0x10, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0x74, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x10,
	0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12,
	0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29,
	0x5a, 0x27, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x6f, 0x73, 0x73, 0x2f, 0x66, 0x61, 0x6b, 0x65, 0x6b,
	0x6d, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_fakekms_fault_math_proto_rawDescOnce sync.Once
	file_fakekms_fault_math_proto_rawDescData = file_fakekms_fault_math_proto_rawDesc
)

func file_fakekms_fault_math_proto_rawDescGZIP() []byte {
	file_fakekms_fault_math_proto_rawDescOnce.Do(func() {
		file_fakekms_fault_math_proto_rawDescData = protoimpl.X.CompressGZIP(file_fakekms_fault_math_proto_rawDescData)
	})
	return file_fakekms_fault_math_proto_rawDescData
}

var file_fakekms_fault_math_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_fakekms_fault_math_proto_goTypes = []interface{}{
	(*AddRequest)(nil),       // 0: math.AddRequest
	(*AddResponse)(nil),      // 1: math.AddResponse
	(*MultiplyRequest)(nil),  // 2: math.MultiplyRequest
	(*MultiplyResponse)(nil), // 3: math.MultiplyResponse
}
var file_fakekms_fault_math_proto_depIdxs = []int32{
	0, // 0: math.MathService.Add:input_type -> math.AddRequest
	2, // 1: math.MathService.Multiply:input_type -> math.MultiplyRequest
	1, // 2: math.MathService.Add:output_type -> math.AddResponse
	3, // 3: math.MathService.Multiply:output_type -> math.MultiplyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fakekms_fault_math_proto_init() }
func file_fakekms_fault_math_proto_init() {
	if File_fakekms_fault_math_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fakekms_fault_math_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_fakekms_fault_math_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_fakekms_fault_math_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplyRequest); i {
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
		file_fakekms_fault_math_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplyResponse); i {
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
			RawDescriptor: file_fakekms_fault_math_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fakekms_fault_math_proto_goTypes,
		DependencyIndexes: file_fakekms_fault_math_proto_depIdxs,
		MessageInfos:      file_fakekms_fault_math_proto_msgTypes,
	}.Build()
	File_fakekms_fault_math_proto = out.File
	file_fakekms_fault_math_proto_rawDesc = nil
	file_fakekms_fault_math_proto_goTypes = nil
	file_fakekms_fault_math_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
}

type mathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMathServiceClient(cc grpc.ClientConnInterface) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/math.MathService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := c.cc.Invoke(ctx, "/math.MathService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
type MathServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
}

// UnimplementedMathServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMathServiceServer struct {
}

func (*UnimplementedMathServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedMathServiceServer) Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}

func RegisterMathServiceServer(s *grpc.Server, srv MathServiceServer) {
	s.RegisterService(&_MathService_serviceDesc, srv)
}

func _MathService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.MathService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.MathService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "math.MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MathService_Add_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _MathService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fakekms/fault/math.proto",
}
