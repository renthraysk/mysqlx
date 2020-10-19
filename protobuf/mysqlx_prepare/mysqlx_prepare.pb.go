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
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: mysqlx_prepare.proto

// ifdef PROTOBUF_LITE: option optimize_for = LITE_RUNTIME;

//*
//@namespace Mysqlx::Prepare
//@brief Handling of prepared statments

package mysqlx_prepare

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/renthraysk/mysqlx/protobuf/mysqlx"
	mysqlx_crud "github.com/renthraysk/mysqlx/protobuf/mysqlx_crud"
	mysqlx_datatypes "github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"
	mysqlx_sql "github.com/renthraysk/mysqlx/protobuf/mysqlx_sql"
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

// Determine which of optional fields was set by the client
// (Workaround for missing "oneof" keyword in pb2.5)
type Prepare_OneOfMessage_Type int32

const (
	Prepare_OneOfMessage_FIND   Prepare_OneOfMessage_Type = 0
	Prepare_OneOfMessage_INSERT Prepare_OneOfMessage_Type = 1
	Prepare_OneOfMessage_UPDATE Prepare_OneOfMessage_Type = 2
	Prepare_OneOfMessage_DELETE Prepare_OneOfMessage_Type = 4
	Prepare_OneOfMessage_STMT   Prepare_OneOfMessage_Type = 5
)

// Enum value maps for Prepare_OneOfMessage_Type.
var (
	Prepare_OneOfMessage_Type_name = map[int32]string{
		0: "FIND",
		1: "INSERT",
		2: "UPDATE",
		4: "DELETE",
		5: "STMT",
	}
	Prepare_OneOfMessage_Type_value = map[string]int32{
		"FIND":   0,
		"INSERT": 1,
		"UPDATE": 2,
		"DELETE": 4,
		"STMT":   5,
	}
)

func (x Prepare_OneOfMessage_Type) Enum() *Prepare_OneOfMessage_Type {
	p := new(Prepare_OneOfMessage_Type)
	*p = x
	return p
}

func (x Prepare_OneOfMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Prepare_OneOfMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_mysqlx_prepare_proto_enumTypes[0].Descriptor()
}

func (Prepare_OneOfMessage_Type) Type() protoreflect.EnumType {
	return &file_mysqlx_prepare_proto_enumTypes[0]
}

func (x Prepare_OneOfMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Prepare_OneOfMessage_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Prepare_OneOfMessage_Type(num)
	return nil
}

// Deprecated: Use Prepare_OneOfMessage_Type.Descriptor instead.
func (Prepare_OneOfMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_mysqlx_prepare_proto_rawDescGZIP(), []int{0, 0, 0}
}

//*
//Prepare a new statement
//
//@startuml
//client -> server: Prepare
//alt Success
//client <- server: Ok
//else Failure
//client <- server: Error
//end
//@enduml
//
//@returns @ref Mysqlx::Ok or @ref Mysqlx::Error
type Prepare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* client side assigned statement id, which is going to identify
	//the result of preparation
	StmtId *uint32 `protobuf:"varint,1,req,name=stmt_id,json=stmtId" json:"stmt_id,omitempty"`
	//* defines one of following messages to be prepared:
	//Crud::Find, Crud::Insert, Crud::Delete, Crud::Upsert, Sql::StmtExecute
	Stmt *Prepare_OneOfMessage `protobuf:"bytes,2,req,name=stmt" json:"stmt,omitempty"`
}

func (x *Prepare) Reset() {
	*x = Prepare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_prepare_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prepare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prepare) ProtoMessage() {}

func (x *Prepare) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_prepare_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prepare.ProtoReflect.Descriptor instead.
func (*Prepare) Descriptor() ([]byte, []int) {
	return file_mysqlx_prepare_proto_rawDescGZIP(), []int{0}
}

func (x *Prepare) GetStmtId() uint32 {
	if x != nil && x.StmtId != nil {
		return *x.StmtId
	}
	return 0
}

func (x *Prepare) GetStmt() *Prepare_OneOfMessage {
	if x != nil {
		return x.Stmt
	}
	return nil
}

//*
//Execute already prepared statement
//
//@startuml
//
//client -> server: Execute
//alt Success
//... Resultsets...
//client <- server: StmtExecuteOk
//else Failure
//client <- server: Error
//end
//@enduml
//@returns @ref Mysqlx::Ok
type Execute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* client side assigned statement id, must be already prepared
	StmtId *uint32 `protobuf:"varint,1,req,name=stmt_id,json=stmtId" json:"stmt_id,omitempty"`
	//* Arguments to bind to the prepared statement
	Args []*mysqlx_datatypes.Any `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	//* send only type information for
	//@ref Mysqlx::Resultset::ColumnMetaData, skipping names and others
	CompactMetadata *bool `protobuf:"varint,3,opt,name=compact_metadata,json=compactMetadata,def=0" json:"compact_metadata,omitempty"`
}

// Default values for Execute fields.
const (
	Default_Execute_CompactMetadata = bool(false)
)

func (x *Execute) Reset() {
	*x = Execute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_prepare_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Execute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Execute) ProtoMessage() {}

func (x *Execute) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_prepare_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Execute.ProtoReflect.Descriptor instead.
func (*Execute) Descriptor() ([]byte, []int) {
	return file_mysqlx_prepare_proto_rawDescGZIP(), []int{1}
}

func (x *Execute) GetStmtId() uint32 {
	if x != nil && x.StmtId != nil {
		return *x.StmtId
	}
	return 0
}

func (x *Execute) GetArgs() []*mysqlx_datatypes.Any {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Execute) GetCompactMetadata() bool {
	if x != nil && x.CompactMetadata != nil {
		return *x.CompactMetadata
	}
	return Default_Execute_CompactMetadata
}

//*
//Deallocate already prepared statement
//
//@startuml
//client -> server: Deallocate
//alt Success
//client <- server: Ok
//else Failure
//client <- server: Error
//end
//@enduml
//
//@returns @ref Mysqlx::Ok or @ref Mysqlx::Error
type Deallocate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* client side assigned statement id, must be already prepared
	StmtId *uint32 `protobuf:"varint,1,req,name=stmt_id,json=stmtId" json:"stmt_id,omitempty"`
}

func (x *Deallocate) Reset() {
	*x = Deallocate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_prepare_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deallocate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deallocate) ProtoMessage() {}

func (x *Deallocate) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_prepare_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deallocate.ProtoReflect.Descriptor instead.
func (*Deallocate) Descriptor() ([]byte, []int) {
	return file_mysqlx_prepare_proto_rawDescGZIP(), []int{2}
}

func (x *Deallocate) GetStmtId() uint32 {
	if x != nil && x.StmtId != nil {
		return *x.StmtId
	}
	return 0
}

type Prepare_OneOfMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        *Prepare_OneOfMessage_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Prepare.Prepare_OneOfMessage_Type" json:"type,omitempty"`
	Find        *mysqlx_crud.Find          `protobuf:"bytes,2,opt,name=find" json:"find,omitempty"`
	Insert      *mysqlx_crud.Insert        `protobuf:"bytes,3,opt,name=insert" json:"insert,omitempty"`
	Update      *mysqlx_crud.Update        `protobuf:"bytes,4,opt,name=update" json:"update,omitempty"`
	Delete      *mysqlx_crud.Delete        `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	StmtExecute *mysqlx_sql.StmtExecute    `protobuf:"bytes,6,opt,name=stmt_execute,json=stmtExecute" json:"stmt_execute,omitempty"`
}

func (x *Prepare_OneOfMessage) Reset() {
	*x = Prepare_OneOfMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_prepare_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prepare_OneOfMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prepare_OneOfMessage) ProtoMessage() {}

func (x *Prepare_OneOfMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_prepare_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prepare_OneOfMessage.ProtoReflect.Descriptor instead.
func (*Prepare_OneOfMessage) Descriptor() ([]byte, []int) {
	return file_mysqlx_prepare_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Prepare_OneOfMessage) GetType() Prepare_OneOfMessage_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Prepare_OneOfMessage_FIND
}

func (x *Prepare_OneOfMessage) GetFind() *mysqlx_crud.Find {
	if x != nil {
		return x.Find
	}
	return nil
}

func (x *Prepare_OneOfMessage) GetInsert() *mysqlx_crud.Insert {
	if x != nil {
		return x.Insert
	}
	return nil
}

func (x *Prepare_OneOfMessage) GetUpdate() *mysqlx_crud.Update {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *Prepare_OneOfMessage) GetDelete() *mysqlx_crud.Delete {
	if x != nil {
		return x.Delete
	}
	return nil
}

func (x *Prepare_OneOfMessage) GetStmtExecute() *mysqlx_sql.StmtExecute {
	if x != nil {
		return x.StmtExecute
	}
	return nil
}

var File_mysqlx_prepare_proto protoreflect.FileDescriptor

var file_mysqlx_prepare_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x1a, 0x0c, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x73, 0x71, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x63,
	0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x79, 0x73, 0x71, 0x6c,
	0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdc, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x74, 0x6d, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x74, 0x6d, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x73, 0x74, 0x6d, 0x74, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x4f, 0x6e,
	0x65, 0x4f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x73, 0x74, 0x6d, 0x74,
	0x1a, 0xf7, 0x02, 0x0a, 0x0c, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x25, 0x0a, 0x04, 0x66, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x43, 0x72, 0x75, 0x64, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x52, 0x04, 0x66, 0x69, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x69, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78,
	0x2e, 0x43, 0x72, 0x75, 0x64, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x06, 0x69, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x43, 0x72,
	0x75, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x43, 0x72, 0x75, 0x64, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x3a,
	0x0a, 0x0c, 0x73, 0x74, 0x6d, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x53, 0x71,
	0x6c, 0x2e, 0x53, 0x74, 0x6d, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x0b, 0x73,
	0x74, 0x6d, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x22, 0x3e, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x54, 0x4d, 0x54, 0x10, 0x05, 0x3a, 0x04, 0x88, 0xea, 0x30, 0x28,
	0x22, 0x85, 0x01, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x74, 0x6d, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x74, 0x6d, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x12, 0x30, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73,
	0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x3a, 0x04, 0x88, 0xea, 0x30, 0x29, 0x22, 0x2b, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x6d, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x6d, 0x74, 0x49, 0x64, 0x3a,
	0x04, 0x88, 0xea, 0x30, 0x2a, 0x42, 0x4f, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x73,
	0x71, 0x6c, 0x2e, 0x63, 0x6a, 0x2e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6e,
	0x74, 0x68, 0x72, 0x61, 0x79, 0x73, 0x6b, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x70,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
}

var (
	file_mysqlx_prepare_proto_rawDescOnce sync.Once
	file_mysqlx_prepare_proto_rawDescData = file_mysqlx_prepare_proto_rawDesc
)

func file_mysqlx_prepare_proto_rawDescGZIP() []byte {
	file_mysqlx_prepare_proto_rawDescOnce.Do(func() {
		file_mysqlx_prepare_proto_rawDescData = protoimpl.X.CompressGZIP(file_mysqlx_prepare_proto_rawDescData)
	})
	return file_mysqlx_prepare_proto_rawDescData
}

var file_mysqlx_prepare_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mysqlx_prepare_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mysqlx_prepare_proto_goTypes = []interface{}{
	(Prepare_OneOfMessage_Type)(0), // 0: Mysqlx.Prepare.Prepare.OneOfMessage.Type
	(*Prepare)(nil),                // 1: Mysqlx.Prepare.Prepare
	(*Execute)(nil),                // 2: Mysqlx.Prepare.Execute
	(*Deallocate)(nil),             // 3: Mysqlx.Prepare.Deallocate
	(*Prepare_OneOfMessage)(nil),   // 4: Mysqlx.Prepare.Prepare.OneOfMessage
	(*mysqlx_datatypes.Any)(nil),   // 5: Mysqlx.Datatypes.Any
	(*mysqlx_crud.Find)(nil),       // 6: Mysqlx.Crud.Find
	(*mysqlx_crud.Insert)(nil),     // 7: Mysqlx.Crud.Insert
	(*mysqlx_crud.Update)(nil),     // 8: Mysqlx.Crud.Update
	(*mysqlx_crud.Delete)(nil),     // 9: Mysqlx.Crud.Delete
	(*mysqlx_sql.StmtExecute)(nil), // 10: Mysqlx.Sql.StmtExecute
}
var file_mysqlx_prepare_proto_depIdxs = []int32{
	4,  // 0: Mysqlx.Prepare.Prepare.stmt:type_name -> Mysqlx.Prepare.Prepare.OneOfMessage
	5,  // 1: Mysqlx.Prepare.Execute.args:type_name -> Mysqlx.Datatypes.Any
	0,  // 2: Mysqlx.Prepare.Prepare.OneOfMessage.type:type_name -> Mysqlx.Prepare.Prepare.OneOfMessage.Type
	6,  // 3: Mysqlx.Prepare.Prepare.OneOfMessage.find:type_name -> Mysqlx.Crud.Find
	7,  // 4: Mysqlx.Prepare.Prepare.OneOfMessage.insert:type_name -> Mysqlx.Crud.Insert
	8,  // 5: Mysqlx.Prepare.Prepare.OneOfMessage.update:type_name -> Mysqlx.Crud.Update
	9,  // 6: Mysqlx.Prepare.Prepare.OneOfMessage.delete:type_name -> Mysqlx.Crud.Delete
	10, // 7: Mysqlx.Prepare.Prepare.OneOfMessage.stmt_execute:type_name -> Mysqlx.Sql.StmtExecute
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_mysqlx_prepare_proto_init() }
func file_mysqlx_prepare_proto_init() {
	if File_mysqlx_prepare_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mysqlx_prepare_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prepare); i {
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
		file_mysqlx_prepare_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Execute); i {
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
		file_mysqlx_prepare_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deallocate); i {
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
		file_mysqlx_prepare_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prepare_OneOfMessage); i {
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
			RawDescriptor: file_mysqlx_prepare_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mysqlx_prepare_proto_goTypes,
		DependencyIndexes: file_mysqlx_prepare_proto_depIdxs,
		EnumInfos:         file_mysqlx_prepare_proto_enumTypes,
		MessageInfos:      file_mysqlx_prepare_proto_msgTypes,
	}.Build()
	File_mysqlx_prepare_proto = out.File
	file_mysqlx_prepare_proto_rawDesc = nil
	file_mysqlx_prepare_proto_goTypes = nil
	file_mysqlx_prepare_proto_depIdxs = nil
}
