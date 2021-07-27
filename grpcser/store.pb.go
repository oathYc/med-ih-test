// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: store.proto

package grpcser

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protocol "ihtest/protocol"
	reflect "reflect"
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

var File_store_proto protoreflect.FileDescriptor

var file_store_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x1a, 0x0c, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4a, 0x0a, 0x08, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x53, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x4d, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x4d,
	0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x69, 0x68, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_store_proto_goTypes = []interface{}{
	(*protocol.StoreListReq)(nil), // 0: MedTest.StoreListReq
	(*protocol.StoreListRes)(nil), // 1: MedTest.StoreListRes
}
var file_store_proto_depIdxs = []int32{
	0, // 0: StoreGRPCSer.StoreSer.GetStoreList:input_type -> MedTest.StoreListReq
	1, // 1: StoreGRPCSer.StoreSer.GetStoreList:output_type -> MedTest.StoreListRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_proto_init() }
func file_store_proto_init() {
	if File_store_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_proto_goTypes,
		DependencyIndexes: file_store_proto_depIdxs,
	}.Build()
	File_store_proto = out.File
	file_store_proto_rawDesc = nil
	file_store_proto_goTypes = nil
	file_store_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreSerClient is the client API for StoreSer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreSerClient interface {
	// 定义一个接口
	GetStoreList(ctx context.Context, in *protocol.StoreListReq, opts ...grpc.CallOption) (*protocol.StoreListRes, error)
}

type storeSerClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreSerClient(cc grpc.ClientConnInterface) StoreSerClient {
	return &storeSerClient{cc}
}

func (c *storeSerClient) GetStoreList(ctx context.Context, in *protocol.StoreListReq, opts ...grpc.CallOption) (*protocol.StoreListRes, error) {
	out := new(protocol.StoreListRes)
	err := c.cc.Invoke(ctx, "/StoreGRPCSer.StoreSer/GetStoreList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreSerServer is the server API for StoreSer service.
type StoreSerServer interface {
	// 定义一个接口
	GetStoreList(context.Context, *protocol.StoreListReq) (*protocol.StoreListRes, error)
}

// UnimplementedStoreSerServer can be embedded to have forward compatible implementations.
type UnimplementedStoreSerServer struct {
}

func (*UnimplementedStoreSerServer) GetStoreList(context.Context, *protocol.StoreListReq) (*protocol.StoreListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoreList not implemented")
}

func RegisterStoreSerServer(s *grpc.Server, srv StoreSerServer) {
	s.RegisterService(&_StoreSer_serviceDesc, srv)
}

func _StoreSer_GetStoreList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.StoreListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSerServer).GetStoreList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StoreGRPCSer.StoreSer/GetStoreList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSerServer).GetStoreList(ctx, req.(*protocol.StoreListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreSer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StoreGRPCSer.StoreSer",
	HandlerType: (*StoreSerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStoreList",
			Handler:    _StoreSer_GetStoreList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store.proto",
}
