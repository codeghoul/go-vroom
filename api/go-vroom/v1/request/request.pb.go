// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: go-vroom/v1/request/request.proto

package request

import (
	types "go-vroom/api/go-vroom/v1/types"
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

type SolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs      []*types.Job      `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	Shipments []*types.Shipment `protobuf:"bytes,2,rep,name=shipments,proto3" json:"shipments,omitempty"`
	Vehicles  []*types.Vehicle  `protobuf:"bytes,3,rep,name=vehicles,proto3" json:"vehicles,omitempty"`
}

func (x *SolveRequest) Reset() {
	*x = SolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_vroom_v1_request_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveRequest) ProtoMessage() {}

func (x *SolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_vroom_v1_request_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveRequest.ProtoReflect.Descriptor instead.
func (*SolveRequest) Descriptor() ([]byte, []int) {
	return file_go_vroom_v1_request_request_proto_rawDescGZIP(), []int{0}
}

func (x *SolveRequest) GetJobs() []*types.Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *SolveRequest) GetShipments() []*types.Shipment {
	if x != nil {
		return x.Shipments
	}
	return nil
}

func (x *SolveRequest) GetVehicles() []*types.Vehicle {
	if x != nil {
		return x.Vehicles
	}
	return nil
}

var File_go_vroom_v1_request_request_proto protoreflect.FileDescriptor

var file_go_vroom_v1_request_request_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f,
	0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0c, 0x53, 0x6f, 0x6c,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6a, 0x6f, 0x62,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x52,
	0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x36, 0x0a, 0x08, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x08,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x50, 0x0a, 0x13, 0x67, 0x6f, 0x2d, 0x76,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x01, 0x5a, 0x28, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x3b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0xa2, 0x02, 0x0c, 0x41, 0x50,
	0x49, 0x47, 0x6f, 0x56, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_go_vroom_v1_request_request_proto_rawDescOnce sync.Once
	file_go_vroom_v1_request_request_proto_rawDescData = file_go_vroom_v1_request_request_proto_rawDesc
)

func file_go_vroom_v1_request_request_proto_rawDescGZIP() []byte {
	file_go_vroom_v1_request_request_proto_rawDescOnce.Do(func() {
		file_go_vroom_v1_request_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_vroom_v1_request_request_proto_rawDescData)
	})
	return file_go_vroom_v1_request_request_proto_rawDescData
}

var file_go_vroom_v1_request_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_vroom_v1_request_request_proto_goTypes = []interface{}{
	(*SolveRequest)(nil),   // 0: go_vroom.v1.request.SolveRequest
	(*types.Job)(nil),      // 1: go_vroom.v1.types.Job
	(*types.Shipment)(nil), // 2: go_vroom.v1.types.Shipment
	(*types.Vehicle)(nil),  // 3: go_vroom.v1.types.Vehicle
}
var file_go_vroom_v1_request_request_proto_depIdxs = []int32{
	1, // 0: go_vroom.v1.request.SolveRequest.jobs:type_name -> go_vroom.v1.types.Job
	2, // 1: go_vroom.v1.request.SolveRequest.shipments:type_name -> go_vroom.v1.types.Shipment
	3, // 2: go_vroom.v1.request.SolveRequest.vehicles:type_name -> go_vroom.v1.types.Vehicle
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_vroom_v1_request_request_proto_init() }
func file_go_vroom_v1_request_request_proto_init() {
	if File_go_vroom_v1_request_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_vroom_v1_request_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveRequest); i {
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
			RawDescriptor: file_go_vroom_v1_request_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_vroom_v1_request_request_proto_goTypes,
		DependencyIndexes: file_go_vroom_v1_request_request_proto_depIdxs,
		MessageInfos:      file_go_vroom_v1_request_request_proto_msgTypes,
	}.Build()
	File_go_vroom_v1_request_request_proto = out.File
	file_go_vroom_v1_request_request_proto_rawDesc = nil
	file_go_vroom_v1_request_request_proto_goTypes = nil
	file_go_vroom_v1_request_request_proto_depIdxs = nil
}