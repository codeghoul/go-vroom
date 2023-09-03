// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: go-vroom/v1/types/vehicle.proto

package types

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

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description   string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Profile       Profile      `protobuf:"varint,3,opt,name=profile,proto3,enum=go_vroom.v1.types.Profile" json:"profile,omitempty"`
	SpeedFactor   float64      `protobuf:"fixed64,4,opt,name=speed_factor,json=speedFactor,proto3" json:"speed_factor,omitempty"`
	Start         *Coordinates `protobuf:"bytes,5,opt,name=start,proto3" json:"start,omitempty"`
	End           *Coordinates `protobuf:"bytes,6,opt,name=end,proto3" json:"end,omitempty"`
	Capacity      []uint32     `protobuf:"varint,7,rep,packed,name=capacity,proto3" json:"capacity,omitempty"`
	Costs         *Cost        `protobuf:"bytes,8,opt,name=costs,proto3" json:"costs,omitempty"`
	MaxTasks      uint32       `protobuf:"varint,9,opt,name=max_tasks,json=maxTasks,proto3" json:"max_tasks,omitempty"`
	MaxTravelTime uint32       `protobuf:"varint,10,opt,name=max_travel_time,json=maxTravelTime,proto3" json:"max_travel_time,omitempty"`
	Skills        []uint32     `protobuf:"varint,11,rep,packed,name=skills,proto3" json:"skills,omitempty"`
	TimeWindow    *TimeWindow  `protobuf:"bytes,12,opt,name=time_window,json=timeWindow,proto3" json:"time_window,omitempty"`
	Breaks        []*Break     `protobuf:"bytes,13,rep,name=breaks,proto3" json:"breaks,omitempty"`
	Steps         []*Step      `protobuf:"bytes,14,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_vroom_v1_types_vehicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_go_vroom_v1_types_vehicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_go_vroom_v1_types_vehicle_proto_rawDescGZIP(), []int{0}
}

func (x *Vehicle) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Vehicle) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Vehicle) GetProfile() Profile {
	if x != nil {
		return x.Profile
	}
	return Profile_CAR
}

func (x *Vehicle) GetSpeedFactor() float64 {
	if x != nil {
		return x.SpeedFactor
	}
	return 0
}

func (x *Vehicle) GetStart() *Coordinates {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Vehicle) GetEnd() *Coordinates {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Vehicle) GetCapacity() []uint32 {
	if x != nil {
		return x.Capacity
	}
	return nil
}

func (x *Vehicle) GetCosts() *Cost {
	if x != nil {
		return x.Costs
	}
	return nil
}

func (x *Vehicle) GetMaxTasks() uint32 {
	if x != nil {
		return x.MaxTasks
	}
	return 0
}

func (x *Vehicle) GetMaxTravelTime() uint32 {
	if x != nil {
		return x.MaxTravelTime
	}
	return 0
}

func (x *Vehicle) GetSkills() []uint32 {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *Vehicle) GetTimeWindow() *TimeWindow {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

func (x *Vehicle) GetBreaks() []*Break {
	if x != nil {
		return x.Breaks
	}
	return nil
}

func (x *Vehicle) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

var File_go_vroom_v1_types_vehicle_proto protoreflect.FileDescriptor

var file_go_vroom_v1_types_vehicle_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x23, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x2d, 0x76, 0x72,
	0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x2d,
	0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62,
	0x72, 0x65, 0x61, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x07, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x5f, 0x76,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x65, 0x64, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x63,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x12, 0x30, 0x0a, 0x06, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x52, 0x06, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65,
	0x70, 0x73, 0x42, 0x4a, 0x0a, 0x11, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x6f, 0x2d, 0x76, 0x72,
	0x6f, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xa2,
	0x02, 0x0c, 0x41, 0x50, 0x49, 0x47, 0x6f, 0x56, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_vroom_v1_types_vehicle_proto_rawDescOnce sync.Once
	file_go_vroom_v1_types_vehicle_proto_rawDescData = file_go_vroom_v1_types_vehicle_proto_rawDesc
)

func file_go_vroom_v1_types_vehicle_proto_rawDescGZIP() []byte {
	file_go_vroom_v1_types_vehicle_proto_rawDescOnce.Do(func() {
		file_go_vroom_v1_types_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_vroom_v1_types_vehicle_proto_rawDescData)
	})
	return file_go_vroom_v1_types_vehicle_proto_rawDescData
}

var file_go_vroom_v1_types_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_vroom_v1_types_vehicle_proto_goTypes = []interface{}{
	(*Vehicle)(nil),     // 0: go_vroom.v1.types.Vehicle
	(Profile)(0),        // 1: go_vroom.v1.types.Profile
	(*Coordinates)(nil), // 2: go_vroom.v1.types.Coordinates
	(*Cost)(nil),        // 3: go_vroom.v1.types.Cost
	(*TimeWindow)(nil),  // 4: go_vroom.v1.types.TimeWindow
	(*Break)(nil),       // 5: go_vroom.v1.types.Break
	(*Step)(nil),        // 6: go_vroom.v1.types.Step
}
var file_go_vroom_v1_types_vehicle_proto_depIdxs = []int32{
	1, // 0: go_vroom.v1.types.Vehicle.profile:type_name -> go_vroom.v1.types.Profile
	2, // 1: go_vroom.v1.types.Vehicle.start:type_name -> go_vroom.v1.types.Coordinates
	2, // 2: go_vroom.v1.types.Vehicle.end:type_name -> go_vroom.v1.types.Coordinates
	3, // 3: go_vroom.v1.types.Vehicle.costs:type_name -> go_vroom.v1.types.Cost
	4, // 4: go_vroom.v1.types.Vehicle.time_window:type_name -> go_vroom.v1.types.TimeWindow
	5, // 5: go_vroom.v1.types.Vehicle.breaks:type_name -> go_vroom.v1.types.Break
	6, // 6: go_vroom.v1.types.Vehicle.steps:type_name -> go_vroom.v1.types.Step
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_go_vroom_v1_types_vehicle_proto_init() }
func file_go_vroom_v1_types_vehicle_proto_init() {
	if File_go_vroom_v1_types_vehicle_proto != nil {
		return
	}
	file_go_vroom_v1_types_time_window_proto_init()
	file_go_vroom_v1_types_coordinates_proto_init()
	file_go_vroom_v1_types_step_proto_init()
	file_go_vroom_v1_types_profile_proto_init()
	file_go_vroom_v1_types_cost_proto_init()
	file_go_vroom_v1_types_break_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_vroom_v1_types_vehicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
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
			RawDescriptor: file_go_vroom_v1_types_vehicle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_vroom_v1_types_vehicle_proto_goTypes,
		DependencyIndexes: file_go_vroom_v1_types_vehicle_proto_depIdxs,
		MessageInfos:      file_go_vroom_v1_types_vehicle_proto_msgTypes,
	}.Build()
	File_go_vroom_v1_types_vehicle_proto = out.File
	file_go_vroom_v1_types_vehicle_proto_rawDesc = nil
	file_go_vroom_v1_types_vehicle_proto_goTypes = nil
	file_go_vroom_v1_types_vehicle_proto_depIdxs = nil
}
