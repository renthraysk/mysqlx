//
// Copyright (c) 2015, 2020, Oracle and/or its affiliates. All rights reserved.
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: mysqlx_cursor.proto

// ifdef PROTOBUF_LITE: option optimize_for = LITE_RUNTIME;

//*
//@namespace Mysqlx::Cursor
//@brief Handling of Cursors

package mysqlx_cursor

import (
	_ "github.com/renthraysk/mysqlx/protobuf/mysqlx"
	mysqlx_prepare "github.com/renthraysk/mysqlx/protobuf/mysqlx_prepare"
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

type Open_OneOfMessage_Type int32

const (
	Open_OneOfMessage_PREPARE_EXECUTE Open_OneOfMessage_Type = 0
)

// Enum value maps for Open_OneOfMessage_Type.
var (
	Open_OneOfMessage_Type_name = map[int32]string{
		0: "PREPARE_EXECUTE",
	}
	Open_OneOfMessage_Type_value = map[string]int32{
		"PREPARE_EXECUTE": 0,
	}
)

func (x Open_OneOfMessage_Type) Enum() *Open_OneOfMessage_Type {
	p := new(Open_OneOfMessage_Type)
	*p = x
	return p
}

func (x Open_OneOfMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Open_OneOfMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_mysqlx_cursor_proto_enumTypes[0].Descriptor()
}

func (Open_OneOfMessage_Type) Type() protoreflect.EnumType {
	return &file_mysqlx_cursor_proto_enumTypes[0]
}

func (x Open_OneOfMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Open_OneOfMessage_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Open_OneOfMessage_Type(num)
	return nil
}

// Deprecated: Use Open_OneOfMessage_Type.Descriptor instead.
func (Open_OneOfMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_mysqlx_cursor_proto_rawDescGZIP(), []int{0, 0, 0}
}

//*
//Open a cursor
//
//@startuml
//client -> server: Open
//alt Success
//... none or partial Resultsets or full Resultsets ...
//client <- server: StmtExecuteOk
//else Failure
//client <- server: Error
//end alt
//@enduml
//
//@returns @ref Mysqlx::Ok
type Open struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* client side assigned cursor id, the ID is going to represent
	//new cursor and assigned to it statement
	CursorId *uint32 `protobuf:"varint,1,req,name=cursor_id,json=cursorId" json:"cursor_id,omitempty"`
	//* statement which resultset is going to be iterated through the cursor
	Stmt *Open_OneOfMessage `protobuf:"bytes,4,req,name=stmt" json:"stmt,omitempty"`
	//* number of rows which should be retrieved from sequential cursor
	FetchRows *uint64 `protobuf:"varint,5,opt,name=fetch_rows,json=fetchRows" json:"fetch_rows,omitempty"`
}

func (x *Open) Reset() {
	*x = Open{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_cursor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Open) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Open) ProtoMessage() {}

func (x *Open) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_cursor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Open.ProtoReflect.Descriptor instead.
func (*Open) Descriptor() ([]byte, []int) {
	return file_mysqlx_cursor_proto_rawDescGZIP(), []int{0}
}

func (x *Open) GetCursorId() uint32 {
	if x != nil && x.CursorId != nil {
		return *x.CursorId
	}
	return 0
}

func (x *Open) GetStmt() *Open_OneOfMessage {
	if x != nil {
		return x.Stmt
	}
	return nil
}

func (x *Open) GetFetchRows() uint64 {
	if x != nil && x.FetchRows != nil {
		return *x.FetchRows
	}
	return 0
}

//*
//Fetch next portion of data from a cursor
//
//@startuml
//client -> server: Fetch
//alt Success
//... none or partial Resultsets or full Resultsets ...
//client <- server: StmtExecuteOk
//else
//client <- server: Error
//end
//@enduml
type Fetch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* client side assigned cursor id, must be already open
	CursorId *uint32 `protobuf:"varint,1,req,name=cursor_id,json=cursorId" json:"cursor_id,omitempty"`
	//* number of rows which should be retrieved from sequential cursor
	FetchRows *uint64 `protobuf:"varint,5,opt,name=fetch_rows,json=fetchRows" json:"fetch_rows,omitempty"`
}

func (x *Fetch) Reset() {
	*x = Fetch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_cursor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fetch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fetch) ProtoMessage() {}

func (x *Fetch) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_cursor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fetch.ProtoReflect.Descriptor instead.
func (*Fetch) Descriptor() ([]byte, []int) {
	return file_mysqlx_cursor_proto_rawDescGZIP(), []int{1}
}

func (x *Fetch) GetCursorId() uint32 {
	if x != nil && x.CursorId != nil {
		return *x.CursorId
	}
	return 0
}

func (x *Fetch) GetFetchRows() uint64 {
	if x != nil && x.FetchRows != nil {
		return *x.FetchRows
	}
	return 0
}

//*
//Close cursor
//
//@startuml
//client -> server: Close
//alt Success
//client <- server: Ok
//else Failure
//client <- server: Error
//end
//@enduml
//
//@returns @ref Mysqlx::Ok or @ref Mysqlx::Error
type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* client side assigned cursor id, must be allocated/open
	CursorId *uint32 `protobuf:"varint,1,req,name=cursor_id,json=cursorId" json:"cursor_id,omitempty"`
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_cursor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_cursor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_mysqlx_cursor_proto_rawDescGZIP(), []int{2}
}

func (x *Close) GetCursorId() uint32 {
	if x != nil && x.CursorId != nil {
		return *x.CursorId
	}
	return 0
}

type Open_OneOfMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           *Open_OneOfMessage_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Cursor.Open_OneOfMessage_Type" json:"type,omitempty"`
	PrepareExecute *mysqlx_prepare.Execute `protobuf:"bytes,2,opt,name=prepare_execute,json=prepareExecute" json:"prepare_execute,omitempty"`
}

func (x *Open_OneOfMessage) Reset() {
	*x = Open_OneOfMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_cursor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Open_OneOfMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Open_OneOfMessage) ProtoMessage() {}

func (x *Open_OneOfMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_cursor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Open_OneOfMessage.ProtoReflect.Descriptor instead.
func (*Open_OneOfMessage) Descriptor() ([]byte, []int) {
	return file_mysqlx_cursor_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Open_OneOfMessage) GetType() Open_OneOfMessage_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Open_OneOfMessage_PREPARE_EXECUTE
}

func (x *Open_OneOfMessage) GetPrepareExecute() *mysqlx_prepare.Execute {
	if x != nil {
		return x.PrepareExecute
	}
	return nil
}

var File_mysqlx_cursor_proto protoreflect.FileDescriptor

var file_mysqlx_cursor_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x1a, 0x0c, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x70, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x04, 0x4f, 0x70, 0x65,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x34,
	0x0a, 0x04, 0x73, 0x74, 0x6d, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x4d,
	0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2e, 0x4f, 0x70, 0x65,
	0x6e, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04,
	0x73, 0x74, 0x6d, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x6f,
	0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x6f, 0x77, 0x73, 0x1a, 0xa8, 0x01, 0x0a, 0x0c, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x40, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c,
	0x78, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x22, 0x1b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x45,
	0x50, 0x41, 0x52, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10, 0x00, 0x3a, 0x04,
	0x88, 0xea, 0x30, 0x2b, 0x22, 0x49, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x77, 0x73, 0x3a, 0x04, 0x88, 0xea, 0x30, 0x2d, 0x22,
	0x2a, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x49, 0x64, 0x3a, 0x04, 0x88, 0xea, 0x30, 0x2c, 0x42, 0x4e, 0x0a, 0x17, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x63, 0x6a, 0x2e, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6e, 0x74, 0x68, 0x72, 0x61, 0x79, 0x73, 0x6b, 0x2f, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x78, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
}

var (
	file_mysqlx_cursor_proto_rawDescOnce sync.Once
	file_mysqlx_cursor_proto_rawDescData = file_mysqlx_cursor_proto_rawDesc
)

func file_mysqlx_cursor_proto_rawDescGZIP() []byte {
	file_mysqlx_cursor_proto_rawDescOnce.Do(func() {
		file_mysqlx_cursor_proto_rawDescData = protoimpl.X.CompressGZIP(file_mysqlx_cursor_proto_rawDescData)
	})
	return file_mysqlx_cursor_proto_rawDescData
}

var file_mysqlx_cursor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mysqlx_cursor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mysqlx_cursor_proto_goTypes = []interface{}{
	(Open_OneOfMessage_Type)(0),    // 0: Mysqlx.Cursor.Open.OneOfMessage.Type
	(*Open)(nil),                   // 1: Mysqlx.Cursor.Open
	(*Fetch)(nil),                  // 2: Mysqlx.Cursor.Fetch
	(*Close)(nil),                  // 3: Mysqlx.Cursor.Close
	(*Open_OneOfMessage)(nil),      // 4: Mysqlx.Cursor.Open.OneOfMessage
	(*mysqlx_prepare.Execute)(nil), // 5: Mysqlx.Prepare.Execute
}
var file_mysqlx_cursor_proto_depIdxs = []int32{
	4, // 0: Mysqlx.Cursor.Open.stmt:type_name -> Mysqlx.Cursor.Open.OneOfMessage
	0, // 1: Mysqlx.Cursor.Open.OneOfMessage.type:type_name -> Mysqlx.Cursor.Open.OneOfMessage.Type
	5, // 2: Mysqlx.Cursor.Open.OneOfMessage.prepare_execute:type_name -> Mysqlx.Prepare.Execute
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mysqlx_cursor_proto_init() }
func file_mysqlx_cursor_proto_init() {
	if File_mysqlx_cursor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mysqlx_cursor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Open); i {
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
		file_mysqlx_cursor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fetch); i {
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
		file_mysqlx_cursor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
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
		file_mysqlx_cursor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Open_OneOfMessage); i {
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
			RawDescriptor: file_mysqlx_cursor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mysqlx_cursor_proto_goTypes,
		DependencyIndexes: file_mysqlx_cursor_proto_depIdxs,
		EnumInfos:         file_mysqlx_cursor_proto_enumTypes,
		MessageInfos:      file_mysqlx_cursor_proto_msgTypes,
	}.Build()
	File_mysqlx_cursor_proto = out.File
	file_mysqlx_cursor_proto_rawDesc = nil
	file_mysqlx_cursor_proto_goTypes = nil
	file_mysqlx_cursor_proto_depIdxs = nil
}
