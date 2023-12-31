// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: core/services/ocr2/plugins/directrequestocr/config/config_types.proto

package config

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

type AggregationMethod int32

const (
	AggregationMethod_AGGREGATION_UNSPECIFIED AggregationMethod = 0
	AggregationMethod_AGGREGATION_MODE        AggregationMethod = 1
	AggregationMethod_AGGREGATION_MEDIAN      AggregationMethod = 2
)

// Enum value maps for AggregationMethod.
var (
	AggregationMethod_name = map[int32]string{
		0: "AGGREGATION_UNSPECIFIED",
		1: "AGGREGATION_MODE",
		2: "AGGREGATION_MEDIAN",
	}
	AggregationMethod_value = map[string]int32{
		"AGGREGATION_UNSPECIFIED": 0,
		"AGGREGATION_MODE":        1,
		"AGGREGATION_MEDIAN":      2,
	}
)

func (x AggregationMethod) Enum() *AggregationMethod {
	p := new(AggregationMethod)
	*p = x
	return p
}

func (x AggregationMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AggregationMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_enumTypes[0].Descriptor()
}

func (AggregationMethod) Type() protoreflect.EnumType {
	return &file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_enumTypes[0]
}

func (x AggregationMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AggregationMethod.Descriptor instead.
func (AggregationMethod) EnumDescriptor() ([]byte, []int) {
	return file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescGZIP(), []int{0}
}

type ReportingPluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxQueryLengthBytes       uint32            `protobuf:"varint,1,opt,name=maxQueryLengthBytes,proto3" json:"maxQueryLengthBytes,omitempty"`
	MaxObservationLengthBytes uint32            `protobuf:"varint,2,opt,name=maxObservationLengthBytes,proto3" json:"maxObservationLengthBytes,omitempty"`
	MaxReportLengthBytes      uint32            `protobuf:"varint,3,opt,name=maxReportLengthBytes,proto3" json:"maxReportLengthBytes,omitempty"`
	MaxRequestBatchSize       uint32            `protobuf:"varint,4,opt,name=maxRequestBatchSize,proto3" json:"maxRequestBatchSize,omitempty"`
	DefaultAggregationMethod  AggregationMethod `protobuf:"varint,5,opt,name=defaultAggregationMethod,proto3,enum=config_types.AggregationMethod" json:"defaultAggregationMethod,omitempty"`
	UniqueReports             bool              `protobuf:"varint,6,opt,name=uniqueReports,proto3" json:"uniqueReports,omitempty"`
}

func (x *ReportingPluginConfig) Reset() {
	*x = ReportingPluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportingPluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportingPluginConfig) ProtoMessage() {}

func (x *ReportingPluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportingPluginConfig.ProtoReflect.Descriptor instead.
func (*ReportingPluginConfig) Descriptor() ([]byte, []int) {
	return file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescGZIP(), []int{0}
}

func (x *ReportingPluginConfig) GetMaxQueryLengthBytes() uint32 {
	if x != nil {
		return x.MaxQueryLengthBytes
	}
	return 0
}

func (x *ReportingPluginConfig) GetMaxObservationLengthBytes() uint32 {
	if x != nil {
		return x.MaxObservationLengthBytes
	}
	return 0
}

func (x *ReportingPluginConfig) GetMaxReportLengthBytes() uint32 {
	if x != nil {
		return x.MaxReportLengthBytes
	}
	return 0
}

func (x *ReportingPluginConfig) GetMaxRequestBatchSize() uint32 {
	if x != nil {
		return x.MaxRequestBatchSize
	}
	return 0
}

func (x *ReportingPluginConfig) GetDefaultAggregationMethod() AggregationMethod {
	if x != nil {
		return x.DefaultAggregationMethod
	}
	return AggregationMethod_AGGREGATION_UNSPECIFIED
}

func (x *ReportingPluginConfig) GetUniqueReports() bool {
	if x != nil {
		return x.UniqueReports
	}
	return false
}

var File_core_services_ocr2_plugins_directrequestocr_config_config_types_proto protoreflect.FileDescriptor

var file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6f, 0x63, 0x72, 0x32, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x72, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xf0, 0x02, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x30, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6d, 0x61,
	0x78, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x3c, 0x0a, 0x19, 0x6d, 0x61, 0x78, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x6d, 0x61, 0x78, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x32, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x6d,
	0x61, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x13, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x5b, 0x0a, 0x18, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x18, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2a, 0x5e, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b, 0x0a,
	0x17, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x47,
	0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4d, 0x45, 0x44, 0x49, 0x41, 0x4e, 0x10, 0x02, 0x42, 0x34, 0x5a, 0x32, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x63, 0x72, 0x32, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescOnce sync.Once
	file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescData = file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDesc
)

func file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescGZIP() []byte {
	file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescOnce.Do(func() {
		file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescData)
	})
	return file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDescData
}

var file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_goTypes = []interface{}{
	(AggregationMethod)(0),        // 0: config_types.AggregationMethod
	(*ReportingPluginConfig)(nil), // 1: config_types.ReportingPluginConfig
}
var file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_depIdxs = []int32{
	0, // 0: config_types.ReportingPluginConfig.defaultAggregationMethod:type_name -> config_types.AggregationMethod
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_init() }
func file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_init() {
	if File_core_services_ocr2_plugins_directrequestocr_config_config_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportingPluginConfig); i {
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
			RawDescriptor: file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_goTypes,
		DependencyIndexes: file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_depIdxs,
		EnumInfos:         file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_enumTypes,
		MessageInfos:      file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_msgTypes,
	}.Build()
	File_core_services_ocr2_plugins_directrequestocr_config_config_types_proto = out.File
	file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_rawDesc = nil
	file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_goTypes = nil
	file_core_services_ocr2_plugins_directrequestocr_config_config_types_proto_depIdxs = nil
}
