// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: task.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	ExpiresAt int64  `protobuf:"varint,4,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
	Done      bool   `protobuf:"varint,5,opt,name=Done,proto3" json:"Done,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Todo) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *Todo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type TodoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TodoId) Reset() {
	*x = TodoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoId) ProtoMessage() {}

func (x *TodoId) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoId.ProtoReflect.Descriptor instead.
func (*TodoId) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *TodoId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TodoString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *TodoString) Reset() {
	*x = TodoString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoString) ProtoMessage() {}

func (x *TodoString) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoString.ProtoReflect.Descriptor instead.
func (*TodoString) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *TodoString) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Todos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *Todos) Reset() {
	*x = Todos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todos) ProtoMessage() {}

func (x *Todos) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todos.ProtoReflect.Descriptor instead.
func (*Todos) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *Todos) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7c, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0x18,
	0x0a, 0x06, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x05,
	0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9c, 0x02, 0x0a, 0x0b, 0x54,
	0x6f, 0x64, 0x6f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x29, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x1a, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x0c,
	0x53, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x73, 0x12, 0x35, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x77, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_task_proto_goTypes = []interface{}{
	(*Todo)(nil),        // 0: proto.Todo
	(*TodoId)(nil),      // 1: proto.TodoId
	(*TodoString)(nil),  // 2: proto.TodoString
	(*Todos)(nil),       // 3: proto.Todos
	(*Status)(nil),      // 4: proto.Status
	(*empty.Empty)(nil), // 5: google.protobuf.Empty
}
var file_task_proto_depIdxs = []int32{
	0, // 0: proto.Todos.todos:type_name -> proto.Todo
	0, // 1: proto.TodoManager.Add:input_type -> proto.Todo
	0, // 2: proto.TodoManager.UpdateTitle:input_type -> proto.Todo
	1, // 3: proto.TodoManager.UpdateStatus:input_type -> proto.TodoId
	1, // 4: proto.TodoManager.Delete:input_type -> proto.TodoId
	5, // 5: proto.TodoManager.ShowAllTodos:input_type -> google.protobuf.Empty
	5, // 6: proto.TodoManager.ShowDoneTodos:input_type -> google.protobuf.Empty
	0, // 7: proto.TodoManager.Add:output_type -> proto.Todo
	4, // 8: proto.TodoManager.UpdateTitle:output_type -> proto.Status
	4, // 9: proto.TodoManager.UpdateStatus:output_type -> proto.Status
	4, // 10: proto.TodoManager.Delete:output_type -> proto.Status
	3, // 11: proto.TodoManager.ShowAllTodos:output_type -> proto.Todos
	3, // 12: proto.TodoManager.ShowDoneTodos:output_type -> proto.Todos
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoId); i {
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
		file_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoString); i {
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
		file_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todos); i {
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
		file_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoManagerClient is the client API for TodoManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoManagerClient interface {
	Add(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	UpdateTitle(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Status, error)
	UpdateStatus(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*Status, error)
	Delete(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*Status, error)
	ShowAllTodos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Todos, error)
	ShowDoneTodos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Todos, error)
}

type todoManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoManagerClient(cc grpc.ClientConnInterface) TodoManagerClient {
	return &todoManagerClient{cc}
}

func (c *todoManagerClient) Add(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/proto.TodoManager/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) UpdateTitle(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.TodoManager/UpdateTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) UpdateStatus(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.TodoManager/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) Delete(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.TodoManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) ShowAllTodos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Todos, error) {
	out := new(Todos)
	err := c.cc.Invoke(ctx, "/proto.TodoManager/ShowAllTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) ShowDoneTodos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Todos, error) {
	out := new(Todos)
	err := c.cc.Invoke(ctx, "/proto.TodoManager/ShowDoneTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoManagerServer is the server API for TodoManager service.
type TodoManagerServer interface {
	Add(context.Context, *Todo) (*Todo, error)
	UpdateTitle(context.Context, *Todo) (*Status, error)
	UpdateStatus(context.Context, *TodoId) (*Status, error)
	Delete(context.Context, *TodoId) (*Status, error)
	ShowAllTodos(context.Context, *empty.Empty) (*Todos, error)
	ShowDoneTodos(context.Context, *empty.Empty) (*Todos, error)
}

// UnimplementedTodoManagerServer can be embedded to have forward compatible implementations.
type UnimplementedTodoManagerServer struct {
}

func (*UnimplementedTodoManagerServer) Add(context.Context, *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTodoManagerServer) UpdateTitle(context.Context, *Todo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTitle not implemented")
}
func (*UnimplementedTodoManagerServer) UpdateStatus(context.Context, *TodoId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (*UnimplementedTodoManagerServer) Delete(context.Context, *TodoId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedTodoManagerServer) ShowAllTodos(context.Context, *empty.Empty) (*Todos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllTodos not implemented")
}
func (*UnimplementedTodoManagerServer) ShowDoneTodos(context.Context, *empty.Empty) (*Todos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowDoneTodos not implemented")
}

func RegisterTodoManagerServer(s *grpc.Server, srv TodoManagerServer) {
	s.RegisterService(&_TodoManager_serviceDesc, srv)
}

func _TodoManager_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoManager/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).Add(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_UpdateTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).UpdateTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoManager/UpdateTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).UpdateTitle(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoManager/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).UpdateStatus(ctx, req.(*TodoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).Delete(ctx, req.(*TodoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_ShowAllTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).ShowAllTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoManager/ShowAllTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).ShowAllTodos(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_ShowDoneTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).ShowDoneTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoManager/ShowDoneTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).ShowDoneTodos(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TodoManager",
	HandlerType: (*TodoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TodoManager_Add_Handler,
		},
		{
			MethodName: "UpdateTitle",
			Handler:    _TodoManager_UpdateTitle_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _TodoManager_UpdateStatus_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoManager_Delete_Handler,
		},
		{
			MethodName: "ShowAllTodos",
			Handler:    _TodoManager_ShowAllTodos_Handler,
		},
		{
			MethodName: "ShowDoneTodos",
			Handler:    _TodoManager_ShowDoneTodos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}