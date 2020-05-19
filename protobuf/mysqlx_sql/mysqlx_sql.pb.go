//
// Copyright (c) 2015, 2019, Oracle and/or its affiliates. All rights reserved.
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License, version 2.0,
// as published by the Free Software Foundation.
//
// This program is also distributed with certain software (including
// but not limited to OpenSSL) that is licensed under separate terms,
// as designated in a particular file or component or in included license
// documentation.  The authors of MySQL hereby grant you an additional
// permission to link the program and your derivative works with the
// separately licensed software that they have included with MySQL.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License, version 2.0, for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: mysqlx_sql.proto

// ifdef PROTOBUF_LITE: option optimize_for = LITE_RUNTIME;

// Messages of the MySQL Package

package mysqlx_sql

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/renthraysk/mysqlx/protobuf/mysqlx"
	mysqlx_datatypes "github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"
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

// execute a statement in the given namespace
//
// .. uml::
//
//   client -> server: StmtExecute
//   ... zero or more Resultsets ...
//   server --> client: StmtExecuteOk
//
// Notices:
//   This message may generate a notice containing WARNINGs generated by its execution.
//   This message may generate a notice containing INFO messages generated by its execution.
//
// :param namespace: namespace of the statement to be executed
// :param stmt: statement that shall be executed.
// :param args: values for wildcard replacements
// :param compact_metadata: send only type information for :protobuf:msg:`Mysqlx.Resultset::ColumnMetadata`, skipping names and others
// :returns:
//    * zero or one :protobuf:msg:`Mysqlx.Resultset::` followed by :protobuf:msg:`Mysqlx.Sql::StmtExecuteOk`
type StmtExecute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace       *string                 `protobuf:"bytes,3,opt,name=namespace,def=sql" json:"namespace,omitempty"`
	Stmt            []byte                  `protobuf:"bytes,1,req,name=stmt" json:"stmt,omitempty"`
	Args            []*mysqlx_datatypes.Any `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	CompactMetadata *bool                   `protobuf:"varint,4,opt,name=compact_metadata,json=compactMetadata,def=0" json:"compact_metadata,omitempty"`
}

// Default values for StmtExecute fields.
const (
	Default_StmtExecute_Namespace       = string("sql")
	Default_StmtExecute_CompactMetadata = bool(false)
)

func (x *StmtExecute) Reset() {
	*x = StmtExecute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_sql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StmtExecute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StmtExecute) ProtoMessage() {}

func (x *StmtExecute) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_sql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StmtExecute.ProtoReflect.Descriptor instead.
func (*StmtExecute) Descriptor() ([]byte, []int) {
	return file_mysqlx_sql_proto_rawDescGZIP(), []int{0}
}

func (x *StmtExecute) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return Default_StmtExecute_Namespace
}

func (x *StmtExecute) GetStmt() []byte {
	if x != nil {
		return x.Stmt
	}
	return nil
}

func (x *StmtExecute) GetArgs() []*mysqlx_datatypes.Any {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *StmtExecute) GetCompactMetadata() bool {
	if x != nil && x.CompactMetadata != nil {
		return *x.CompactMetadata
	}
	return Default_StmtExecute_CompactMetadata
}

// statement executed successful
type StmtExecuteOk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StmtExecuteOk) Reset() {
	*x = StmtExecuteOk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_sql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StmtExecuteOk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StmtExecuteOk) ProtoMessage() {}

func (x *StmtExecuteOk) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_sql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StmtExecuteOk.ProtoReflect.Descriptor instead.
func (*StmtExecuteOk) Descriptor() ([]byte, []int) {
	return file_mysqlx_sql_proto_rawDescGZIP(), []int{1}
}

var File_mysqlx_sql_proto protoreflect.FileDescriptor

var file_mysqlx_sql_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x53, 0x71, 0x6c, 0x1a, 0x0c,
	0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x6d, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x03, 0x73, 0x71, 0x6c, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x6d, 0x74, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x74, 0x6d, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x79, 0x73, 0x71,
	0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x04, 0x88, 0xea, 0x30, 0x0c, 0x22, 0x15,
	0x0a, 0x0d, 0x53, 0x74, 0x6d, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4f, 0x6b, 0x3a,
	0x04, 0x90, 0xea, 0x30, 0x11, 0x42, 0x4b, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x73,
	0x71, 0x6c, 0x2e, 0x63, 0x6a, 0x2e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6e,
	0x74, 0x68, 0x72, 0x61, 0x79, 0x73, 0x6b, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x73,
	0x71, 0x6c,
}

var (
	file_mysqlx_sql_proto_rawDescOnce sync.Once
	file_mysqlx_sql_proto_rawDescData = file_mysqlx_sql_proto_rawDesc
)

func file_mysqlx_sql_proto_rawDescGZIP() []byte {
	file_mysqlx_sql_proto_rawDescOnce.Do(func() {
		file_mysqlx_sql_proto_rawDescData = protoimpl.X.CompressGZIP(file_mysqlx_sql_proto_rawDescData)
	})
	return file_mysqlx_sql_proto_rawDescData
}

var file_mysqlx_sql_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mysqlx_sql_proto_goTypes = []interface{}{
	(*StmtExecute)(nil),          // 0: Mysqlx.Sql.StmtExecute
	(*StmtExecuteOk)(nil),        // 1: Mysqlx.Sql.StmtExecuteOk
	(*mysqlx_datatypes.Any)(nil), // 2: Mysqlx.Datatypes.Any
}
var file_mysqlx_sql_proto_depIdxs = []int32{
	2, // 0: Mysqlx.Sql.StmtExecute.args:type_name -> Mysqlx.Datatypes.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mysqlx_sql_proto_init() }
func file_mysqlx_sql_proto_init() {
	if File_mysqlx_sql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mysqlx_sql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StmtExecute); i {
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
		file_mysqlx_sql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StmtExecuteOk); i {
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
			RawDescriptor: file_mysqlx_sql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mysqlx_sql_proto_goTypes,
		DependencyIndexes: file_mysqlx_sql_proto_depIdxs,
		MessageInfos:      file_mysqlx_sql_proto_msgTypes,
	}.Build()
	File_mysqlx_sql_proto = out.File
	file_mysqlx_sql_proto_rawDesc = nil
	file_mysqlx_sql_proto_goTypes = nil
	file_mysqlx_sql_proto_depIdxs = nil
}
