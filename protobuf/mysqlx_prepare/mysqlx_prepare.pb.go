// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mysqlx_prepare.proto

package mysqlx_prepare

/*
Handling of prepared statments
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/renthraysk/mysqlx/protobuf/mysqlx"
import mysqlx_crud "github.com/renthraysk/mysqlx/protobuf/mysqlx_crud"
import mysqlx_datatypes "github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"
import mysqlx_sql "github.com/renthraysk/mysqlx/protobuf/mysqlx_sql"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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

var Prepare_OneOfMessage_Type_name = map[int32]string{
	0: "FIND",
	1: "INSERT",
	2: "UPDATE",
	4: "DELETE",
	5: "STMT",
}
var Prepare_OneOfMessage_Type_value = map[string]int32{
	"FIND":   0,
	"INSERT": 1,
	"UPDATE": 2,
	"DELETE": 4,
	"STMT":   5,
}

func (x Prepare_OneOfMessage_Type) Enum() *Prepare_OneOfMessage_Type {
	p := new(Prepare_OneOfMessage_Type)
	*p = x
	return p
}
func (x Prepare_OneOfMessage_Type) String() string {
	return proto.EnumName(Prepare_OneOfMessage_Type_name, int32(x))
}
func (x *Prepare_OneOfMessage_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Prepare_OneOfMessage_Type_value, data, "Prepare_OneOfMessage_Type")
	if err != nil {
		return err
	}
	*x = Prepare_OneOfMessage_Type(value)
	return nil
}
func (Prepare_OneOfMessage_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_prepare_e8a26bf7759eb9c1, []int{0, 0, 0}
}

// Prepare a new statement
//
// .. uml::
//
//   client -> server: Prepare
//   alt Success
//     client <- server: Ok
//   else Failure
//     client <- server: Error
//   end
//
// :param stmt_id: client side assigned statement id, which is going to identify the result of preparation
// :param stmt: defines one of following messages to be prepared - Crud.Find, Crud.Insert, Crud.Delete, Crud.Upsert, Sql.StmtExecute
// :Returns: :protobuf:msg:`Mysqlx.Ok|Mysqlx.Error`
type Prepare struct {
	StmtId               *uint32               `protobuf:"varint,1,req,name=stmt_id,json=stmtId" json:"stmt_id,omitempty"`
	Stmt                 *Prepare_OneOfMessage `protobuf:"bytes,2,req,name=stmt" json:"stmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Prepare) Reset()         { *m = Prepare{} }
func (m *Prepare) String() string { return proto.CompactTextString(m) }
func (*Prepare) ProtoMessage()    {}
func (*Prepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_prepare_e8a26bf7759eb9c1, []int{0}
}
func (m *Prepare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prepare.Unmarshal(m, b)
}
func (m *Prepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prepare.Marshal(b, m, deterministic)
}
func (dst *Prepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prepare.Merge(dst, src)
}
func (m *Prepare) XXX_Size() int {
	return xxx_messageInfo_Prepare.Size(m)
}
func (m *Prepare) XXX_DiscardUnknown() {
	xxx_messageInfo_Prepare.DiscardUnknown(m)
}

var xxx_messageInfo_Prepare proto.InternalMessageInfo

func (m *Prepare) GetStmtId() uint32 {
	if m != nil && m.StmtId != nil {
		return *m.StmtId
	}
	return 0
}

func (m *Prepare) GetStmt() *Prepare_OneOfMessage {
	if m != nil {
		return m.Stmt
	}
	return nil
}

type Prepare_OneOfMessage struct {
	Type                 *Prepare_OneOfMessage_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Prepare.Prepare_OneOfMessage_Type" json:"type,omitempty"`
	Find                 *mysqlx_crud.Find          `protobuf:"bytes,2,opt,name=find" json:"find,omitempty"`
	Insert               *mysqlx_crud.Insert        `protobuf:"bytes,3,opt,name=insert" json:"insert,omitempty"`
	Update               *mysqlx_crud.Update        `protobuf:"bytes,4,opt,name=update" json:"update,omitempty"`
	Delete               *mysqlx_crud.Delete        `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	StmtExecute          *mysqlx_sql.StmtExecute    `protobuf:"bytes,6,opt,name=stmt_execute,json=stmtExecute" json:"stmt_execute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Prepare_OneOfMessage) Reset()         { *m = Prepare_OneOfMessage{} }
func (m *Prepare_OneOfMessage) String() string { return proto.CompactTextString(m) }
func (*Prepare_OneOfMessage) ProtoMessage()    {}
func (*Prepare_OneOfMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_prepare_e8a26bf7759eb9c1, []int{0, 0}
}
func (m *Prepare_OneOfMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prepare_OneOfMessage.Unmarshal(m, b)
}
func (m *Prepare_OneOfMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prepare_OneOfMessage.Marshal(b, m, deterministic)
}
func (dst *Prepare_OneOfMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prepare_OneOfMessage.Merge(dst, src)
}
func (m *Prepare_OneOfMessage) XXX_Size() int {
	return xxx_messageInfo_Prepare_OneOfMessage.Size(m)
}
func (m *Prepare_OneOfMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Prepare_OneOfMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Prepare_OneOfMessage proto.InternalMessageInfo

func (m *Prepare_OneOfMessage) GetType() Prepare_OneOfMessage_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Prepare_OneOfMessage_FIND
}

func (m *Prepare_OneOfMessage) GetFind() *mysqlx_crud.Find {
	if m != nil {
		return m.Find
	}
	return nil
}

func (m *Prepare_OneOfMessage) GetInsert() *mysqlx_crud.Insert {
	if m != nil {
		return m.Insert
	}
	return nil
}

func (m *Prepare_OneOfMessage) GetUpdate() *mysqlx_crud.Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *Prepare_OneOfMessage) GetDelete() *mysqlx_crud.Delete {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Prepare_OneOfMessage) GetStmtExecute() *mysqlx_sql.StmtExecute {
	if m != nil {
		return m.StmtExecute
	}
	return nil
}

// Execute already prepared statement
//
// .. uml::
//
//   client -> server: Execute
//   alt Success
//     ... Resultsets...
//     client <- server: StmtExecuteOk
//  else Failure
//     client <- server: Error
//  end
//
// :param stmt_id: client side assigned statement id, must be already prepared
// :param args_list: Arguments to bind to the prepared statement
// :param compact_metadata: send only type information for :protobuf:msg:`Mysqlx.Resultset::ColumnMetadata`, skipping names and others
// :Returns: :protobuf:msg:`Mysqlx.Ok::`
type Execute struct {
	StmtId               *uint32                 `protobuf:"varint,1,req,name=stmt_id,json=stmtId" json:"stmt_id,omitempty"`
	Args                 []*mysqlx_datatypes.Any `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	CompactMetadata      *bool                   `protobuf:"varint,3,opt,name=compact_metadata,json=compactMetadata,def=0" json:"compact_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Execute) Reset()         { *m = Execute{} }
func (m *Execute) String() string { return proto.CompactTextString(m) }
func (*Execute) ProtoMessage()    {}
func (*Execute) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_prepare_e8a26bf7759eb9c1, []int{1}
}
func (m *Execute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Execute.Unmarshal(m, b)
}
func (m *Execute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Execute.Marshal(b, m, deterministic)
}
func (dst *Execute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execute.Merge(dst, src)
}
func (m *Execute) XXX_Size() int {
	return xxx_messageInfo_Execute.Size(m)
}
func (m *Execute) XXX_DiscardUnknown() {
	xxx_messageInfo_Execute.DiscardUnknown(m)
}

var xxx_messageInfo_Execute proto.InternalMessageInfo

const Default_Execute_CompactMetadata bool = false

func (m *Execute) GetStmtId() uint32 {
	if m != nil && m.StmtId != nil {
		return *m.StmtId
	}
	return 0
}

func (m *Execute) GetArgs() []*mysqlx_datatypes.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Execute) GetCompactMetadata() bool {
	if m != nil && m.CompactMetadata != nil {
		return *m.CompactMetadata
	}
	return Default_Execute_CompactMetadata
}

// Deallocate already prepared statement
//
// Deallocating the statement.
//
// .. uml::
//
//   client -> server: Deallocate
//   alt Success
//     client <- server: Ok
//   else Failure
//     client <- server: Error
//   end
//
// :param stmt_id: client side assigned statement id, must be already prepared
// :Returns: :protobuf:msg:`Mysqlx.Ok|Mysqlx.Error`
type Deallocate struct {
	StmtId               *uint32  `protobuf:"varint,1,req,name=stmt_id,json=stmtId" json:"stmt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deallocate) Reset()         { *m = Deallocate{} }
func (m *Deallocate) String() string { return proto.CompactTextString(m) }
func (*Deallocate) ProtoMessage()    {}
func (*Deallocate) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_prepare_e8a26bf7759eb9c1, []int{2}
}
func (m *Deallocate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deallocate.Unmarshal(m, b)
}
func (m *Deallocate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deallocate.Marshal(b, m, deterministic)
}
func (dst *Deallocate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deallocate.Merge(dst, src)
}
func (m *Deallocate) XXX_Size() int {
	return xxx_messageInfo_Deallocate.Size(m)
}
func (m *Deallocate) XXX_DiscardUnknown() {
	xxx_messageInfo_Deallocate.DiscardUnknown(m)
}

var xxx_messageInfo_Deallocate proto.InternalMessageInfo

func (m *Deallocate) GetStmtId() uint32 {
	if m != nil && m.StmtId != nil {
		return *m.StmtId
	}
	return 0
}

func init() {
	proto.RegisterType((*Prepare)(nil), "Mysqlx.Prepare.Prepare")
	proto.RegisterType((*Prepare_OneOfMessage)(nil), "Mysqlx.Prepare.Prepare.OneOfMessage")
	proto.RegisterType((*Execute)(nil), "Mysqlx.Prepare.Execute")
	proto.RegisterType((*Deallocate)(nil), "Mysqlx.Prepare.Deallocate")
	proto.RegisterEnum("Mysqlx.Prepare.Prepare_OneOfMessage_Type", Prepare_OneOfMessage_Type_name, Prepare_OneOfMessage_Type_value)
}

func init() {
	proto.RegisterFile("mysqlx_prepare.proto", fileDescriptor_mysqlx_prepare_e8a26bf7759eb9c1)
}

var fileDescriptor_mysqlx_prepare_e8a26bf7759eb9c1 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x17, 0x47, 0x49, 0xca, 0x49, 0xd6, 0xb9, 0xda, 0x9f, 0x78, 0xb9, 0x0a, 0x61, 0x03,
	0x67, 0x05, 0x53, 0x72, 0x35, 0x02, 0x1b, 0x74, 0xb3, 0x0b, 0x81, 0xa5, 0x2d, 0x8a, 0x7b, 0x1d,
	0x34, 0xeb, 0xa4, 0x64, 0xf8, 0x5f, 0x2c, 0x05, 0x9a, 0x07, 0x18, 0xec, 0x21, 0xf7, 0x04, 0x7b,
	0x81, 0xdd, 0x0e, 0xc9, 0x4a, 0xb7, 0x42, 0x36, 0x76, 0x65, 0xe9, 0xfb, 0x7e, 0xdf, 0xd1, 0x91,
	0x7c, 0xe0, 0x59, 0xb6, 0x93, 0x9b, 0xf4, 0x6e, 0x59, 0x56, 0x58, 0xf2, 0x0a, 0x83, 0xb2, 0x2a,
	0x54, 0x41, 0x8f, 0xe7, 0x46, 0x0d, 0xae, 0x6b, 0x75, 0xd0, 0xab, 0xa9, 0xda, 0x1d, 0xb8, 0x36,
	0x23, 0x37, 0xa9, 0x55, 0x4e, 0xac, 0x92, 0x54, 0x5b, 0x61, 0xa5, 0x17, 0x56, 0x12, 0x5c, 0x71,
	0xb5, 0x2b, 0x51, 0xd6, 0xfa, 0xe8, 0x7b, 0x13, 0x3a, 0xb6, 0x2c, 0xed, 0x43, 0x47, 0xaa, 0x4c,
	0x2d, 0xd7, 0xc2, 0x6b, 0x0c, 0x1d, 0xff, 0x31, 0x6b, 0xeb, 0xed, 0x4c, 0xd0, 0xb7, 0x40, 0xf4,
	0xca, 0x73, 0x86, 0x8e, 0xdf, 0x9d, 0xbc, 0x0a, 0x1e, 0xb6, 0x73, 0xff, 0xbd, 0xca, 0xf1, 0x6a,
	0x35, 0x47, 0x29, 0xf9, 0x2d, 0x32, 0x93, 0x18, 0xfc, 0x74, 0xa0, 0xf7, 0xa7, 0x4c, 0xdf, 0x01,
	0xd1, 0xc7, 0x9b, 0x03, 0x8e, 0x27, 0xe3, 0xff, 0x29, 0x15, 0xc4, 0xbb, 0x12, 0x99, 0x89, 0xd1,
	0xd7, 0x40, 0x56, 0xeb, 0x5c, 0x78, 0xce, 0xb0, 0xe1, 0x77, 0x27, 0x27, 0xfb, 0xf8, 0x47, 0x7d,
	0xd1, 0x8b, 0x75, 0x2e, 0x98, 0xb1, 0xe9, 0x29, 0xb4, 0xd7, 0xb9, 0xc4, 0x4a, 0x79, 0x4d, 0x03,
	0x3e, 0x7d, 0x00, 0xce, 0x8c, 0xc5, 0x2c, 0xa2, 0xe1, 0x6d, 0x29, 0xb8, 0x42, 0x8f, 0x1c, 0x80,
	0x6f, 0x8c, 0xc5, 0x2c, 0xa2, 0x61, 0x81, 0x29, 0x2a, 0xf4, 0x5a, 0x07, 0xe0, 0xd0, 0x58, 0xcc,
	0x22, 0x74, 0x0a, 0x3d, 0xf3, 0xa0, 0x78, 0x87, 0xc9, 0x56, 0xa1, 0xd7, 0x36, 0x91, 0xfe, 0x3e,
	0xb2, 0xd8, 0xa4, 0xc1, 0x42, 0x65, 0x2a, 0xaa, 0x6d, 0xd6, 0x95, 0xbf, 0x37, 0xa3, 0xf7, 0x40,
	0xf4, 0xbd, 0xe9, 0x11, 0x90, 0x8b, 0xd9, 0x65, 0xe8, 0x3e, 0xa2, 0x00, 0xed, 0xd9, 0xe5, 0x22,
	0x62, 0xb1, 0xdb, 0xd0, 0xeb, 0x9b, 0xeb, 0xf0, 0x3c, 0x8e, 0x5c, 0x47, 0xaf, 0xc3, 0xe8, 0x53,
	0x14, 0x47, 0x2e, 0xd1, 0xf4, 0x22, 0x9e, 0xc7, 0x6e, 0x6b, 0x4a, 0xbe, 0xfd, 0x38, 0xf3, 0x47,
	0x5f, 0x1b, 0xd0, 0xb1, 0x15, 0xff, 0xfe, 0x7b, 0xc7, 0x40, 0x78, 0x75, 0x2b, 0x3d, 0x67, 0xd8,
	0xf4, 0xbb, 0x93, 0xe7, 0xfb, 0xf6, 0xc2, 0xfb, 0x51, 0x39, 0xcf, 0x77, 0xcc, 0x20, 0xf4, 0x0c,
	0xdc, 0xa4, 0xc8, 0x4a, 0x9e, 0xa8, 0x65, 0x86, 0x8a, 0xeb, 0x69, 0x32, 0x4f, 0x7c, 0x34, 0x6d,
	0xad, 0x78, 0x2a, 0x91, 0x3d, 0xb1, 0xf6, 0xdc, 0xba, 0xa6, 0x8f, 0xf1, 0xe8, 0x14, 0x20, 0x44,
	0x9e, 0xa6, 0x45, 0xc2, 0xff, 0xd1, 0x89, 0x81, 0xdf, 0x7c, 0x78, 0x09, 0xfd, 0xa4, 0xc8, 0x02,
	0x33, 0xb1, 0x41, 0xf2, 0x25, 0xb0, 0x83, 0xfe, 0x79, 0xbb, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff,
	0xc8, 0xc2, 0xd0, 0xd9, 0x20, 0x03, 0x00, 0x00,
}
