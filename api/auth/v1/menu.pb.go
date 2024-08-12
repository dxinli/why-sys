// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: auth/v1/menu.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateMenuResponse) Reset() {
	*x = CreateMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuResponse) ProtoMessage() {}

func (x *CreateMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuResponse.ProtoReflect.Descriptor instead.
func (*CreateMenuResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_menu_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMenuResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParentId int64   `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Sort     float64 `protobuf:"fixed64,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Level    uint32  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *ListMenuRequest) Reset() {
	*x = ListMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMenuRequest) ProtoMessage() {}

func (x *ListMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMenuRequest.ProtoReflect.Descriptor instead.
func (*ListMenuRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_menu_proto_rawDescGZIP(), []int{1}
}

func (x *ListMenuRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListMenuRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *ListMenuRequest) GetSort() float64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *ListMenuRequest) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type MenuDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId    int64   `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Sort        float64 `protobuf:"fixed64,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Level       uint32  `protobuf:"varint,5,opt,name=level,proto3" json:"level,omitempty"`
	Path        string  `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Component   string  `protobuf:"bytes,7,opt,name=component,proto3" json:"component,omitempty"`
	Description string  `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Leaf        bool    `protobuf:"varint,9,opt,name=leaf,proto3" json:"leaf,omitempty"`
}

func (x *MenuDetail) Reset() {
	*x = MenuDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuDetail) ProtoMessage() {}

func (x *MenuDetail) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuDetail.ProtoReflect.Descriptor instead.
func (*MenuDetail) Descriptor() ([]byte, []int) {
	return file_auth_v1_menu_proto_rawDescGZIP(), []int{2}
}

func (x *MenuDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuDetail) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *MenuDetail) GetSort() float64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MenuDetail) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *MenuDetail) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MenuDetail) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *MenuDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MenuDetail) GetLeaf() bool {
	if x != nil {
		return x.Leaf
	}
	return false
}

var File_auth_v1_menu_proto protoreflect.FileDescriptor

var file_auth_v1_menu_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x6c, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0xdf, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c,
	0x65, 0x61, 0x66, 0x32, 0xac, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x53, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a,
	0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x3a, 0x01,
	0x2a, 0x12, 0x4f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x18, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x3a,
	0x01, 0x2a, 0x42, 0x18, 0x5a, 0x16, 0x77, 0x68, 0x79, 0x2d, 0x73, 0x79, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_menu_proto_rawDescOnce sync.Once
	file_auth_v1_menu_proto_rawDescData = file_auth_v1_menu_proto_rawDesc
)

func file_auth_v1_menu_proto_rawDescGZIP() []byte {
	file_auth_v1_menu_proto_rawDescOnce.Do(func() {
		file_auth_v1_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_menu_proto_rawDescData)
	})
	return file_auth_v1_menu_proto_rawDescData
}

var file_auth_v1_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_auth_v1_menu_proto_goTypes = []interface{}{
	(*CreateMenuResponse)(nil), // 0: auth.v1.CreateMenuResponse
	(*ListMenuRequest)(nil),    // 1: auth.v1.ListMenuRequest
	(*MenuDetail)(nil),         // 2: auth.v1.MenuDetail
}
var file_auth_v1_menu_proto_depIdxs = []int32{
	2, // 0: auth.v1.Menu.CreateMenu:input_type -> auth.v1.MenuDetail
	1, // 1: auth.v1.Menu.ListMenu:input_type -> auth.v1.ListMenuRequest
	0, // 2: auth.v1.Menu.CreateMenu:output_type -> auth.v1.CreateMenuResponse
	2, // 3: auth.v1.Menu.ListMenu:output_type -> auth.v1.MenuDetail
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_v1_menu_proto_init() }
func file_auth_v1_menu_proto_init() {
	if File_auth_v1_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_v1_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMenuResponse); i {
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
		file_auth_v1_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMenuRequest); i {
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
		file_auth_v1_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuDetail); i {
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
			RawDescriptor: file_auth_v1_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_v1_menu_proto_goTypes,
		DependencyIndexes: file_auth_v1_menu_proto_depIdxs,
		MessageInfos:      file_auth_v1_menu_proto_msgTypes,
	}.Build()
	File_auth_v1_menu_proto = out.File
	file_auth_v1_menu_proto_rawDesc = nil
	file_auth_v1_menu_proto_goTypes = nil
	file_auth_v1_menu_proto_depIdxs = nil
}
