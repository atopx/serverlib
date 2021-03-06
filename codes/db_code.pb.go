// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: db_code.proto

package codes

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

type DatabaseTypeCode int32

const (
	DatabaseTypeCode_DB_TYPE_NONE          DatabaseTypeCode = 0
	DatabaseTypeCode_DB_TYPE_MYSQL         DatabaseTypeCode = 1
	DatabaseTypeCode_DB_TYPE_POSTGRESQL    DatabaseTypeCode = 2
	DatabaseTypeCode_DB_TYPE_ORACLE        DatabaseTypeCode = 3
	DatabaseTypeCode_DB_TYPE_SQLLITE       DatabaseTypeCode = 4
	DatabaseTypeCode_DB_TYPE_MSSQL         DatabaseTypeCode = 5
	DatabaseTypeCode_DB_TYPE_TIDB          DatabaseTypeCode = 6
	DatabaseTypeCode_DB_TYPE_CLIECKHOUSE   DatabaseTypeCode = 100
	DatabaseTypeCode_DB_TYPE_HBASE         DatabaseTypeCode = 101
	DatabaseTypeCode_DB_TYPE_MONGODB       DatabaseTypeCode = 102
	DatabaseTypeCode_DB_TYPE_ELASTICSERACH DatabaseTypeCode = 103
	DatabaseTypeCode_DB_TYPE_REDIS         DatabaseTypeCode = 201
	DatabaseTypeCode_DB_TYPE_MEMCACHED     DatabaseTypeCode = 202
	DatabaseTypeCode_DB_TYPE_KAFKA         DatabaseTypeCode = 301
	DatabaseTypeCode_DB_TYPE_RABBITMQ      DatabaseTypeCode = 302
	DatabaseTypeCode_DB_TYPE_ZEROMQ        DatabaseTypeCode = 303
	DatabaseTypeCode_DB_TYPE_ROCKETMQ      DatabaseTypeCode = 304
	DatabaseTypeCode_DB_TYPE_ACTIVEMQ      DatabaseTypeCode = 305
	DatabaseTypeCode_DB_TYPE_UNKNOWN       DatabaseTypeCode = 1000
)

// Enum value maps for DatabaseTypeCode.
var (
	DatabaseTypeCode_name = map[int32]string{
		0:    "DB_TYPE_NONE",
		1:    "DB_TYPE_MYSQL",
		2:    "DB_TYPE_POSTGRESQL",
		3:    "DB_TYPE_ORACLE",
		4:    "DB_TYPE_SQLLITE",
		5:    "DB_TYPE_MSSQL",
		6:    "DB_TYPE_TIDB",
		100:  "DB_TYPE_CLIECKHOUSE",
		101:  "DB_TYPE_HBASE",
		102:  "DB_TYPE_MONGODB",
		103:  "DB_TYPE_ELASTICSERACH",
		201:  "DB_TYPE_REDIS",
		202:  "DB_TYPE_MEMCACHED",
		301:  "DB_TYPE_KAFKA",
		302:  "DB_TYPE_RABBITMQ",
		303:  "DB_TYPE_ZEROMQ",
		304:  "DB_TYPE_ROCKETMQ",
		305:  "DB_TYPE_ACTIVEMQ",
		1000: "DB_TYPE_UNKNOWN",
	}
	DatabaseTypeCode_value = map[string]int32{
		"DB_TYPE_NONE":          0,
		"DB_TYPE_MYSQL":         1,
		"DB_TYPE_POSTGRESQL":    2,
		"DB_TYPE_ORACLE":        3,
		"DB_TYPE_SQLLITE":       4,
		"DB_TYPE_MSSQL":         5,
		"DB_TYPE_TIDB":          6,
		"DB_TYPE_CLIECKHOUSE":   100,
		"DB_TYPE_HBASE":         101,
		"DB_TYPE_MONGODB":       102,
		"DB_TYPE_ELASTICSERACH": 103,
		"DB_TYPE_REDIS":         201,
		"DB_TYPE_MEMCACHED":     202,
		"DB_TYPE_KAFKA":         301,
		"DB_TYPE_RABBITMQ":      302,
		"DB_TYPE_ZEROMQ":        303,
		"DB_TYPE_ROCKETMQ":      304,
		"DB_TYPE_ACTIVEMQ":      305,
		"DB_TYPE_UNKNOWN":       1000,
	}
)

func (x DatabaseTypeCode) Enum() *DatabaseTypeCode {
	p := new(DatabaseTypeCode)
	*p = x
	return p
}

func (x DatabaseTypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatabaseTypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_db_code_proto_enumTypes[0].Descriptor()
}

func (DatabaseTypeCode) Type() protoreflect.EnumType {
	return &file_db_code_proto_enumTypes[0]
}

func (x DatabaseTypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatabaseTypeCode.Descriptor instead.
func (DatabaseTypeCode) EnumDescriptor() ([]byte, []int) {
	return file_db_code_proto_rawDescGZIP(), []int{0}
}

var File_db_code_proto protoreflect.FileDescriptor

var file_db_code_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x2a, 0xa9, 0x03, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x42, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x42,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x51, 0x4c,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52,
	0x41, 0x43, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x51, 0x4c, 0x4c, 0x49, 0x54, 0x45, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x44,
	0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x53, 0x53, 0x51, 0x4c, 0x10, 0x05, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x44, 0x42, 0x10, 0x06,
	0x12, 0x17, 0x0a, 0x13, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45,
	0x43, 0x4b, 0x48, 0x4f, 0x55, 0x53, 0x45, 0x10, 0x64, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x42, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x42, 0x41, 0x53, 0x45, 0x10, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x47, 0x4f, 0x44, 0x42, 0x10,
	0x66, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4c, 0x41,
	0x53, 0x54, 0x49, 0x43, 0x53, 0x45, 0x52, 0x41, 0x43, 0x48, 0x10, 0x67, 0x12, 0x12, 0x0a, 0x0d,
	0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x53, 0x10, 0xc9, 0x01,
	0x12, 0x16, 0x0a, 0x11, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x43,
	0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0xca, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x44, 0x42, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4b, 0x41, 0x46, 0x4b, 0x41, 0x10, 0xad, 0x02, 0x12, 0x15, 0x0a, 0x10,
	0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x41, 0x42, 0x42, 0x49, 0x54, 0x4d, 0x51,
	0x10, 0xae, 0x02, 0x12, 0x13, 0x0a, 0x0e, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x5a,
	0x45, 0x52, 0x4f, 0x4d, 0x51, 0x10, 0xaf, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x44, 0x42, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x51, 0x10, 0xb0, 0x02, 0x12,
	0x15, 0x0a, 0x10, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x4d, 0x51, 0x10, 0xb1, 0x02, 0x12, 0x14, 0x0a, 0x0f, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xe8, 0x07, 0x42, 0x22, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x74, 0x6f, 0x70, 0x78,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x69, 0x62, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_code_proto_rawDescOnce sync.Once
	file_db_code_proto_rawDescData = file_db_code_proto_rawDesc
)

func file_db_code_proto_rawDescGZIP() []byte {
	file_db_code_proto_rawDescOnce.Do(func() {
		file_db_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_code_proto_rawDescData)
	})
	return file_db_code_proto_rawDescData
}

var file_db_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_db_code_proto_goTypes = []interface{}{
	(DatabaseTypeCode)(0), // 0: serverlib.codes.DatabaseTypeCode
}
var file_db_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_db_code_proto_init() }
func file_db_code_proto_init() {
	if File_db_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_db_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_code_proto_goTypes,
		DependencyIndexes: file_db_code_proto_depIdxs,
		EnumInfos:         file_db_code_proto_enumTypes,
	}.Build()
	File_db_code_proto = out.File
	file_db_code_proto_rawDesc = nil
	file_db_code_proto_goTypes = nil
	file_db_code_proto_depIdxs = nil
}
