// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: dish.proto

package orderproto

import (
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

type DishItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DishId string `protobuf:"bytes,1,opt,name=dish_id,json=dishId,proto3" json:"dish_id,omitempty"`
	Dish   string `protobuf:"bytes,2,opt,name=dish,proto3" json:"dish,omitempty"`
	Total  int32  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *DishItem) Reset() {
	*x = DishItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DishItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DishItem) ProtoMessage() {}

func (x *DishItem) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DishItem.ProtoReflect.Descriptor instead.
func (*DishItem) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{0}
}

func (x *DishItem) GetDishId() string {
	if x != nil {
		return x.DishId
	}
	return ""
}

func (x *DishItem) GetDish() string {
	if x != nil {
		return x.Dish
	}
	return ""
}

func (x *DishItem) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetOrderDishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetOrderDishRequest) Reset() {
	*x = GetOrderDishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderDishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDishRequest) ProtoMessage() {}

func (x *GetOrderDishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDishRequest.ProtoReflect.Descriptor instead.
func (*GetOrderDishRequest) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderDishRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetOrderDishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dishes []*DishItem `protobuf:"bytes,1,rep,name=dishes,proto3" json:"dishes,omitempty"`
}

func (x *GetOrderDishResponse) Reset() {
	*x = GetOrderDishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderDishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDishResponse) ProtoMessage() {}

func (x *GetOrderDishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDishResponse.ProtoReflect.Descriptor instead.
func (*GetOrderDishResponse) Descriptor() ([]byte, []int) {
	return file_dish_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderDishResponse) GetDishes() []*DishItem {
	if x != nil {
		return x.Dishes
	}
	return nil
}

var File_dish_proto protoreflect.FileDescriptor

var file_dish_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x17, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x69, 0x73, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x64, 0x69, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x64,
	0x69, 0x73, 0x68, 0x65, 0x73, 0x32, 0x56, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x69, 0x73, 0x68, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a,
	0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x63, 0x2d,
	0x67, 0x6f, 0x65, 0x72, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dish_proto_rawDescOnce sync.Once
	file_dish_proto_rawDescData = file_dish_proto_rawDesc
)

func file_dish_proto_rawDescGZIP() []byte {
	file_dish_proto_rawDescOnce.Do(func() {
		file_dish_proto_rawDescData = protoimpl.X.CompressGZIP(file_dish_proto_rawDescData)
	})
	return file_dish_proto_rawDescData
}

var file_dish_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dish_proto_goTypes = []any{
	(*DishItem)(nil),             // 0: order.DishItem
	(*GetOrderDishRequest)(nil),  // 1: order.GetOrderDishRequest
	(*GetOrderDishResponse)(nil), // 2: order.GetOrderDishResponse
}
var file_dish_proto_depIdxs = []int32{
	0, // 0: order.GetOrderDishResponse.dishes:type_name -> order.DishItem
	1, // 1: order.DishService.GetOrderDish:input_type -> order.GetOrderDishRequest
	2, // 2: order.DishService.GetOrderDish:output_type -> order.GetOrderDishResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dish_proto_init() }
func file_dish_proto_init() {
	if File_dish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dish_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DishItem); i {
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
		file_dish_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrderDishRequest); i {
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
		file_dish_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrderDishResponse); i {
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
			RawDescriptor: file_dish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dish_proto_goTypes,
		DependencyIndexes: file_dish_proto_depIdxs,
		MessageInfos:      file_dish_proto_msgTypes,
	}.Build()
	File_dish_proto = out.File
	file_dish_proto_rawDesc = nil
	file_dish_proto_goTypes = nil
	file_dish_proto_depIdxs = nil
}