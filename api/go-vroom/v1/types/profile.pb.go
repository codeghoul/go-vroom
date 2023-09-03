// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: go-vroom/v1/types/profile.proto

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

type Profile int32

const (
	Profile_CAR   Profile = 0
	Profile_TRUCK Profile = 1
)

// Enum value maps for Profile.
var (
	Profile_name = map[int32]string{
		0: "CAR",
		1: "TRUCK",
	}
	Profile_value = map[string]int32{
		"CAR":   0,
		"TRUCK": 1,
	}
)

func (x Profile) Enum() *Profile {
	p := new(Profile)
	*p = x
	return p
}

func (x Profile) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Profile) Descriptor() protoreflect.EnumDescriptor {
	return file_go_vroom_v1_types_profile_proto_enumTypes[0].Descriptor()
}

func (Profile) Type() protoreflect.EnumType {
	return &file_go_vroom_v1_types_profile_proto_enumTypes[0]
}

func (x Profile) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Profile.Descriptor instead.
func (Profile) EnumDescriptor() ([]byte, []int) {
	return file_go_vroom_v1_types_profile_proto_rawDescGZIP(), []int{0}
}

var File_go_vroom_v1_types_profile_proto protoreflect.FileDescriptor

var file_go_vroom_v1_types_profile_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x67, 0x6f, 0x5f, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2a, 0x1d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x43, 0x41, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x55, 0x43,
	0x4b, 0x10, 0x01, 0x42, 0x4a, 0x0a, 0x11, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x6f, 0x2d, 0x76,
	0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x76, 0x72, 0x6f, 0x6f,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73,
	0xa2, 0x02, 0x0c, 0x41, 0x50, 0x49, 0x47, 0x6f, 0x56, 0x72, 0x6f, 0x6f, 0x6d, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_vroom_v1_types_profile_proto_rawDescOnce sync.Once
	file_go_vroom_v1_types_profile_proto_rawDescData = file_go_vroom_v1_types_profile_proto_rawDesc
)

func file_go_vroom_v1_types_profile_proto_rawDescGZIP() []byte {
	file_go_vroom_v1_types_profile_proto_rawDescOnce.Do(func() {
		file_go_vroom_v1_types_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_vroom_v1_types_profile_proto_rawDescData)
	})
	return file_go_vroom_v1_types_profile_proto_rawDescData
}

var file_go_vroom_v1_types_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_vroom_v1_types_profile_proto_goTypes = []interface{}{
	(Profile)(0), // 0: go_vroom.v1.types.Profile
}
var file_go_vroom_v1_types_profile_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_vroom_v1_types_profile_proto_init() }
func file_go_vroom_v1_types_profile_proto_init() {
	if File_go_vroom_v1_types_profile_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_vroom_v1_types_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_vroom_v1_types_profile_proto_goTypes,
		DependencyIndexes: file_go_vroom_v1_types_profile_proto_depIdxs,
		EnumInfos:         file_go_vroom_v1_types_profile_proto_enumTypes,
	}.Build()
	File_go_vroom_v1_types_profile_proto = out.File
	file_go_vroom_v1_types_profile_proto_rawDesc = nil
	file_go_vroom_v1_types_profile_proto_goTypes = nil
	file_go_vroom_v1_types_profile_proto_depIdxs = nil
}