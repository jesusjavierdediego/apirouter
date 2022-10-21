// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.19.1
// source: apirouter.proto

package protobuf

import (
	context "context"
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

type RecordSetStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []string `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *RecordSetStream) Reset() {
	*x = RecordSetStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apirouter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordSetStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSetStream) ProtoMessage() {}

func (x *RecordSetStream) ProtoReflect() protoreflect.Message {
	mi := &file_apirouter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSetStream.ProtoReflect.Descriptor instead.
func (*RecordSetStream) Descriptor() ([]byte, []int) {
	return file_apirouter_proto_rawDescGZIP(), []int{0}
}

func (x *RecordSetStream) GetRecords() []string {
	if x != nil {
		return x.Records
	}
	return nil
}

type RecordNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numberrecords int64 `protobuf:"varint,1,opt,name=numberrecords,proto3" json:"numberrecords,omitempty"`
}

func (x *RecordNumber) Reset() {
	*x = RecordNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apirouter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordNumber) ProtoMessage() {}

func (x *RecordNumber) ProtoReflect() protoreflect.Message {
	mi := &file_apirouter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordNumber.ProtoReflect.Descriptor instead.
func (*RecordNumber) Descriptor() ([]byte, []int) {
	return file_apirouter_proto_rawDescGZIP(), []int{1}
}

func (x *RecordNumber) GetNumberrecords() int64 {
	if x != nil {
		return x.Numberrecords
	}
	return 0
}

type DBCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database   string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (x *DBCollection) Reset() {
	*x = DBCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apirouter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBCollection) ProtoMessage() {}

func (x *DBCollection) ProtoReflect() protoreflect.Message {
	mi := &file_apirouter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBCollection.ProtoReflect.Descriptor instead.
func (*DBCollection) Descriptor() ([]byte, []int) {
	return file_apirouter_proto_rawDescGZIP(), []int{2}
}

func (x *DBCollection) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *DBCollection) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apirouter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_apirouter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_apirouter_proto_rawDescGZIP(), []int{3}
}

var File_apirouter_proto protoreflect.FileDescriptor

var file_apirouter_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x34,
	0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x22, 0x4a, 0x0a, 0x0c, 0x44, 0x42, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xf9, 0x02, 0x0a, 0x13, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x42, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x12, 0x56, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x42, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x5a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x42, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x56, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x42, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0e, 0x48, 0x01, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apirouter_proto_rawDescOnce sync.Once
	file_apirouter_proto_rawDescData = file_apirouter_proto_rawDesc
)

func file_apirouter_proto_rawDescGZIP() []byte {
	file_apirouter_proto_rawDescOnce.Do(func() {
		file_apirouter_proto_rawDescData = protoimpl.X.CompressGZIP(file_apirouter_proto_rawDescData)
	})
	return file_apirouter_proto_rawDescData
}

var file_apirouter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apirouter_proto_goTypes = []interface{}{
	(*RecordSetStream)(nil), // 0: apirouterproto.RecordSetStream
	(*RecordNumber)(nil),    // 1: apirouterproto.RecordNumber
	(*DBCollection)(nil),    // 2: apirouterproto.DBCollection
	(*Empty)(nil),           // 3: apirouterproto.Empty
}
var file_apirouter_proto_depIdxs = []int32{
	2, // 0: apirouterproto.RecordStreamService.GetTotalRecordSet:input_type -> apirouterproto.DBCollection
	2, // 1: apirouterproto.RecordStreamService.GetTotalRecordNumber:input_type -> apirouterproto.DBCollection
	2, // 2: apirouterproto.RecordStreamService.GetCompletedRecordNumber:input_type -> apirouterproto.DBCollection
	2, // 3: apirouterproto.RecordStreamService.GetErrorRecordNumber:input_type -> apirouterproto.DBCollection
	0, // 4: apirouterproto.RecordStreamService.GetTotalRecordSet:output_type -> apirouterproto.RecordSetStream
	1, // 5: apirouterproto.RecordStreamService.GetTotalRecordNumber:output_type -> apirouterproto.RecordNumber
	1, // 6: apirouterproto.RecordStreamService.GetCompletedRecordNumber:output_type -> apirouterproto.RecordNumber
	1, // 7: apirouterproto.RecordStreamService.GetErrorRecordNumber:output_type -> apirouterproto.RecordNumber
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apirouter_proto_init() }
func file_apirouter_proto_init() {
	if File_apirouter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apirouter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordSetStream); i {
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
		file_apirouter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordNumber); i {
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
		file_apirouter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBCollection); i {
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
		file_apirouter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_apirouter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apirouter_proto_goTypes,
		DependencyIndexes: file_apirouter_proto_depIdxs,
		MessageInfos:      file_apirouter_proto_msgTypes,
	}.Build()
	File_apirouter_proto = out.File
	file_apirouter_proto_rawDesc = nil
	file_apirouter_proto_goTypes = nil
	file_apirouter_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecordStreamServiceClient is the client API for RecordStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecordStreamServiceClient interface {
	GetTotalRecordSet(ctx context.Context, in *DBCollection, opts ...grpc.CallOption) (RecordStreamService_GetTotalRecordSetClient, error)
	GetTotalRecordNumber(ctx context.Context, in *DBCollection, opts ...grpc.CallOption) (RecordStreamService_GetTotalRecordNumberClient, error)
	GetCompletedRecordNumber(ctx context.Context, in *DBCollection, opts ...grpc.CallOption) (RecordStreamService_GetCompletedRecordNumberClient, error)
	GetErrorRecordNumber(ctx context.Context, in *DBCollection, opts ...grpc.CallOption) (RecordStreamService_GetErrorRecordNumberClient, error)
}

type recordStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordStreamServiceClient(cc grpc.ClientConnInterface) RecordStreamServiceClient {
	return &recordStreamServiceClient{cc}
}

func (c *recordStreamServiceClient) GetTotalRecordSet(ctx context.Context, in *DBCollection, opts ...grpc.CallOption) (RecordStreamService_GetTotalRecordSetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecordStreamService_serviceDesc.Streams[0], "/apirouterproto.RecordStreamService/GetTotalRecordSet", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordStreamServiceGetTotalRecordSetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecordStreamService_GetTotalRecordSetClient interface {
	Recv() (*RecordSetStream, error)
	grpc.ClientStream
}

type recordStreamServiceGetTotalRecordSetClient struct {
	grpc.ClientStream
}

func (x *recordStreamServiceGetTotalRecordSetClient) Recv() (*RecordSetStream, error) {
	m := new(RecordSetStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recordStreamServiceClient) GetTotalRecordNumber(ctx context.Context, in *DBCollection, opts ...grpc.CallOption) (RecordStreamService_GetTotalRecordNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecordStreamService_serviceDesc.Streams[1], "/apirouterproto.RecordStreamService/GetTotalRecordNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordStreamServiceGetTotalRecordNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecordStreamService_GetTotalRecordNumberClient interface {
	Recv() (*RecordNumber, error)
	grpc.ClientStream
}

type recordStreamServiceGetTotalRecordNumberClient struct {
	grpc.ClientStream
}

func (x *recordStreamServiceGetTotalRecordNumberClient) Recv() (*RecordNumber, error) {
	m := new(RecordNumber)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recordStreamServiceClient) GetCompletedRecordNumber(ctx context.Context, in *DBCollection, opts ...grpc.CallOption) (RecordStreamService_GetCompletedRecordNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecordStreamService_serviceDesc.Streams[2], "/apirouterproto.RecordStreamService/GetCompletedRecordNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordStreamServiceGetCompletedRecordNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecordStreamService_GetCompletedRecordNumberClient interface {
	Recv() (*RecordNumber, error)
	grpc.ClientStream
}

type recordStreamServiceGetCompletedRecordNumberClient struct {
	grpc.ClientStream
}

func (x *recordStreamServiceGetCompletedRecordNumberClient) Recv() (*RecordNumber, error) {
	m := new(RecordNumber)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recordStreamServiceClient) GetErrorRecordNumber(ctx context.Context, in *DBCollection, opts ...grpc.CallOption) (RecordStreamService_GetErrorRecordNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecordStreamService_serviceDesc.Streams[3], "/apirouterproto.RecordStreamService/GetErrorRecordNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordStreamServiceGetErrorRecordNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecordStreamService_GetErrorRecordNumberClient interface {
	Recv() (*RecordNumber, error)
	grpc.ClientStream
}

type recordStreamServiceGetErrorRecordNumberClient struct {
	grpc.ClientStream
}

func (x *recordStreamServiceGetErrorRecordNumberClient) Recv() (*RecordNumber, error) {
	m := new(RecordNumber)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecordStreamServiceServer is the server API for RecordStreamService service.
type RecordStreamServiceServer interface {
	GetTotalRecordSet(*DBCollection, RecordStreamService_GetTotalRecordSetServer) error
	GetTotalRecordNumber(*DBCollection, RecordStreamService_GetTotalRecordNumberServer) error
	GetCompletedRecordNumber(*DBCollection, RecordStreamService_GetCompletedRecordNumberServer) error
	GetErrorRecordNumber(*DBCollection, RecordStreamService_GetErrorRecordNumberServer) error
}

// UnimplementedRecordStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecordStreamServiceServer struct {
}

func (*UnimplementedRecordStreamServiceServer) GetTotalRecordSet(*DBCollection, RecordStreamService_GetTotalRecordSetServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTotalRecordSet not implemented")
}
func (*UnimplementedRecordStreamServiceServer) GetTotalRecordNumber(*DBCollection, RecordStreamService_GetTotalRecordNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTotalRecordNumber not implemented")
}
func (*UnimplementedRecordStreamServiceServer) GetCompletedRecordNumber(*DBCollection, RecordStreamService_GetCompletedRecordNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCompletedRecordNumber not implemented")
}
func (*UnimplementedRecordStreamServiceServer) GetErrorRecordNumber(*DBCollection, RecordStreamService_GetErrorRecordNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method GetErrorRecordNumber not implemented")
}

func RegisterRecordStreamServiceServer(s *grpc.Server, srv RecordStreamServiceServer) {
	s.RegisterService(&_RecordStreamService_serviceDesc, srv)
}

func _RecordStreamService_GetTotalRecordSet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DBCollection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordStreamServiceServer).GetTotalRecordSet(m, &recordStreamServiceGetTotalRecordSetServer{stream})
}

type RecordStreamService_GetTotalRecordSetServer interface {
	Send(*RecordSetStream) error
	grpc.ServerStream
}

type recordStreamServiceGetTotalRecordSetServer struct {
	grpc.ServerStream
}

func (x *recordStreamServiceGetTotalRecordSetServer) Send(m *RecordSetStream) error {
	return x.ServerStream.SendMsg(m)
}

func _RecordStreamService_GetTotalRecordNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DBCollection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordStreamServiceServer).GetTotalRecordNumber(m, &recordStreamServiceGetTotalRecordNumberServer{stream})
}

type RecordStreamService_GetTotalRecordNumberServer interface {
	Send(*RecordNumber) error
	grpc.ServerStream
}

type recordStreamServiceGetTotalRecordNumberServer struct {
	grpc.ServerStream
}

func (x *recordStreamServiceGetTotalRecordNumberServer) Send(m *RecordNumber) error {
	return x.ServerStream.SendMsg(m)
}

func _RecordStreamService_GetCompletedRecordNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DBCollection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordStreamServiceServer).GetCompletedRecordNumber(m, &recordStreamServiceGetCompletedRecordNumberServer{stream})
}

type RecordStreamService_GetCompletedRecordNumberServer interface {
	Send(*RecordNumber) error
	grpc.ServerStream
}

type recordStreamServiceGetCompletedRecordNumberServer struct {
	grpc.ServerStream
}

func (x *recordStreamServiceGetCompletedRecordNumberServer) Send(m *RecordNumber) error {
	return x.ServerStream.SendMsg(m)
}

func _RecordStreamService_GetErrorRecordNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DBCollection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordStreamServiceServer).GetErrorRecordNumber(m, &recordStreamServiceGetErrorRecordNumberServer{stream})
}

type RecordStreamService_GetErrorRecordNumberServer interface {
	Send(*RecordNumber) error
	grpc.ServerStream
}

type recordStreamServiceGetErrorRecordNumberServer struct {
	grpc.ServerStream
}

func (x *recordStreamServiceGetErrorRecordNumberServer) Send(m *RecordNumber) error {
	return x.ServerStream.SendMsg(m)
}

var _RecordStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apirouterproto.RecordStreamService",
	HandlerType: (*RecordStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTotalRecordSet",
			Handler:       _RecordStreamService_GetTotalRecordSet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTotalRecordNumber",
			Handler:       _RecordStreamService_GetTotalRecordNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetCompletedRecordNumber",
			Handler:       _RecordStreamService_GetCompletedRecordNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetErrorRecordNumber",
			Handler:       _RecordStreamService_GetErrorRecordNumber_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apirouter.proto",
}