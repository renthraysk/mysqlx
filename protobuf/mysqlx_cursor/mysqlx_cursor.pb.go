// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mysqlx_cursor.proto

package mysqlx_cursor

/*
Handling of Cursors
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/renthraysk/mysqlx/protobuf/mysqlx"
import mysqlx_prepare "github.com/renthraysk/mysqlx/protobuf/mysqlx_prepare"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Open_OneOfMessage_Type int32

const (
	Open_OneOfMessage_PREPARE_EXECUTE Open_OneOfMessage_Type = 0
)

var Open_OneOfMessage_Type_name = map[int32]string{
	0: "PREPARE_EXECUTE",
}
var Open_OneOfMessage_Type_value = map[string]int32{
	"PREPARE_EXECUTE": 0,
}

func (x Open_OneOfMessage_Type) Enum() *Open_OneOfMessage_Type {
	p := new(Open_OneOfMessage_Type)
	*p = x
	return p
}
func (x Open_OneOfMessage_Type) String() string {
	return proto.EnumName(Open_OneOfMessage_Type_name, int32(x))
}
func (x *Open_OneOfMessage_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Open_OneOfMessage_Type_value, data, "Open_OneOfMessage_Type")
	if err != nil {
		return err
	}
	*x = Open_OneOfMessage_Type(value)
	return nil
}
func (Open_OneOfMessage_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_cursor_d119b3ddea7269e7, []int{0, 0, 0}
}

// Open a cursor
//
// .. uml::
//
//   client -> server: Open
//   alt Success
//     ... none or partial Resultsets or full Resultsets ...
//     client <- server: StmtExecuteOk
//  else Failure
//     client <- server: Error
//  end
//
// :param cursor_id: client side assigned cursor id, the ID is going to represent new cursor and assigned to it statement
// :param stmt: statement which resultset is going to be iterated through the cursor
// :param fetch_rows: number of rows which should be retrieved from sequential cursor
// :Returns: :protobuf:msg:`Mysqlx.Ok::`
type Open struct {
	CursorId             *uint32            `protobuf:"varint,1,req,name=cursor_id,json=cursorId" json:"cursor_id,omitempty"`
	Stmt                 *Open_OneOfMessage `protobuf:"bytes,4,req,name=stmt" json:"stmt,omitempty"`
	FetchRows            *uint64            `protobuf:"varint,5,opt,name=fetch_rows,json=fetchRows" json:"fetch_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Open) Reset()         { *m = Open{} }
func (m *Open) String() string { return proto.CompactTextString(m) }
func (*Open) ProtoMessage()    {}
func (*Open) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_cursor_d119b3ddea7269e7, []int{0}
}
func (m *Open) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Open.Unmarshal(m, b)
}
func (m *Open) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Open.Marshal(b, m, deterministic)
}
func (dst *Open) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Open.Merge(dst, src)
}
func (m *Open) XXX_Size() int {
	return xxx_messageInfo_Open.Size(m)
}
func (m *Open) XXX_DiscardUnknown() {
	xxx_messageInfo_Open.DiscardUnknown(m)
}

var xxx_messageInfo_Open proto.InternalMessageInfo

func (m *Open) GetCursorId() uint32 {
	if m != nil && m.CursorId != nil {
		return *m.CursorId
	}
	return 0
}

func (m *Open) GetStmt() *Open_OneOfMessage {
	if m != nil {
		return m.Stmt
	}
	return nil
}

func (m *Open) GetFetchRows() uint64 {
	if m != nil && m.FetchRows != nil {
		return *m.FetchRows
	}
	return 0
}

type Open_OneOfMessage struct {
	Type                 *Open_OneOfMessage_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Cursor.Open_OneOfMessage_Type" json:"type,omitempty"`
	PrepareExecute       *mysqlx_prepare.Execute `protobuf:"bytes,2,opt,name=prepare_execute,json=prepareExecute" json:"prepare_execute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Open_OneOfMessage) Reset()         { *m = Open_OneOfMessage{} }
func (m *Open_OneOfMessage) String() string { return proto.CompactTextString(m) }
func (*Open_OneOfMessage) ProtoMessage()    {}
func (*Open_OneOfMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_cursor_d119b3ddea7269e7, []int{0, 0}
}
func (m *Open_OneOfMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Open_OneOfMessage.Unmarshal(m, b)
}
func (m *Open_OneOfMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Open_OneOfMessage.Marshal(b, m, deterministic)
}
func (dst *Open_OneOfMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Open_OneOfMessage.Merge(dst, src)
}
func (m *Open_OneOfMessage) XXX_Size() int {
	return xxx_messageInfo_Open_OneOfMessage.Size(m)
}
func (m *Open_OneOfMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Open_OneOfMessage.DiscardUnknown(m)
}

var xxx_messageInfo_Open_OneOfMessage proto.InternalMessageInfo

func (m *Open_OneOfMessage) GetType() Open_OneOfMessage_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Open_OneOfMessage_PREPARE_EXECUTE
}

func (m *Open_OneOfMessage) GetPrepareExecute() *mysqlx_prepare.Execute {
	if m != nil {
		return m.PrepareExecute
	}
	return nil
}

// Fetch next portion of data from a cursor
//
// .. uml::
//
//   client -> server: Fetch
//   alt Success
//     ... none or partial Resultsets or full Resultsets ...
//     client <- server: StmtExecuteOk
//   else
//    client <- server: Error
//   end
//
// :param cursor_id: client side assigned cursor id, must be already open
// :param fetch_rows: number of rows which should be retrieved from sequential cursor
type Fetch struct {
	CursorId             *uint32  `protobuf:"varint,1,req,name=cursor_id,json=cursorId" json:"cursor_id,omitempty"`
	FetchRows            *uint64  `protobuf:"varint,5,opt,name=fetch_rows,json=fetchRows" json:"fetch_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fetch) Reset()         { *m = Fetch{} }
func (m *Fetch) String() string { return proto.CompactTextString(m) }
func (*Fetch) ProtoMessage()    {}
func (*Fetch) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_cursor_d119b3ddea7269e7, []int{1}
}
func (m *Fetch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fetch.Unmarshal(m, b)
}
func (m *Fetch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fetch.Marshal(b, m, deterministic)
}
func (dst *Fetch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fetch.Merge(dst, src)
}
func (m *Fetch) XXX_Size() int {
	return xxx_messageInfo_Fetch.Size(m)
}
func (m *Fetch) XXX_DiscardUnknown() {
	xxx_messageInfo_Fetch.DiscardUnknown(m)
}

var xxx_messageInfo_Fetch proto.InternalMessageInfo

func (m *Fetch) GetCursorId() uint32 {
	if m != nil && m.CursorId != nil {
		return *m.CursorId
	}
	return 0
}

func (m *Fetch) GetFetchRows() uint64 {
	if m != nil && m.FetchRows != nil {
		return *m.FetchRows
	}
	return 0
}

// Close cursor
//
// .. uml::
//
//   client -> server: Close
//   alt Success
//     client <- server: Ok
//   else Failure
//     client <- server: Error
//   end
//
// :param cursor_id: client side assigned cursor id, must be allocated/open
// :Returns: :protobuf:msg:`Mysqlx.Ok|Mysqlx.Error`
type Close struct {
	CursorId             *uint32  `protobuf:"varint,1,req,name=cursor_id,json=cursorId" json:"cursor_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_cursor_d119b3ddea7269e7, []int{2}
}
func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (dst *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(dst, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func (m *Close) GetCursorId() uint32 {
	if m != nil && m.CursorId != nil {
		return *m.CursorId
	}
	return 0
}

func init() {
	proto.RegisterType((*Open)(nil), "Mysqlx.Cursor.Open")
	proto.RegisterType((*Open_OneOfMessage)(nil), "Mysqlx.Cursor.Open.OneOfMessage")
	proto.RegisterType((*Fetch)(nil), "Mysqlx.Cursor.Fetch")
	proto.RegisterType((*Close)(nil), "Mysqlx.Cursor.Close")
	proto.RegisterEnum("Mysqlx.Cursor.Open_OneOfMessage_Type", Open_OneOfMessage_Type_name, Open_OneOfMessage_Type_value)
}

func init() { proto.RegisterFile("mysqlx_cursor.proto", fileDescriptor_mysqlx_cursor_d119b3ddea7269e7) }

var fileDescriptor_mysqlx_cursor_d119b3ddea7269e7 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xdd, 0x4a, 0x02, 0x41,
	0x14, 0x6e, 0x97, 0x11, 0xf2, 0xf8, 0x17, 0x63, 0xe0, 0xa6, 0x04, 0xcb, 0x42, 0xb0, 0xf4, 0x33,
	0x84, 0x74, 0x53, 0x57, 0x95, 0x6c, 0xe0, 0x85, 0x28, 0x83, 0x41, 0x77, 0x8b, 0xad, 0xc7, 0x7e,
	0x50, 0x67, 0x9b, 0x19, 0x51, 0xdf, 0xa0, 0x47, 0xa9, 0x57, 0xeb, 0x29, 0xc2, 0x99, 0x09, 0xea,
	0xc6, 0x2e, 0xe7, 0x3b, 0xdf, 0xcf, 0x7c, 0xe7, 0x40, 0x7d, 0xb6, 0x56, 0x6f, 0xd3, 0x55, 0x9a,
	0x2d, 0xa4, 0x12, 0x92, 0xe5, 0x52, 0x68, 0x41, 0x2b, 0x3d, 0x03, 0xb2, 0x8e, 0x01, 0x9b, 0x65,
	0xcb, 0xb1, 0xc3, 0xe6, 0xbe, 0x53, 0xe4, 0x12, 0xf3, 0x91, 0x44, 0x8b, 0x46, 0x9f, 0x3e, 0x90,
	0x7e, 0x8e, 0x73, 0xda, 0x82, 0xa2, 0xf5, 0x4a, 0x5f, 0xc6, 0x81, 0x17, 0xfa, 0x71, 0x85, 0xef,
	0x5a, 0xa0, 0x3b, 0xa6, 0x17, 0x40, 0x94, 0x9e, 0xe9, 0x80, 0x84, 0x7e, 0x5c, 0x6a, 0x87, 0xec,
	0x4f, 0x0e, 0xdb, 0xe8, 0x59, 0x7f, 0x8e, 0xfd, 0x49, 0x0f, 0x95, 0x1a, 0x3d, 0x21, 0x37, 0x6c,
	0x7a, 0x08, 0x30, 0x41, 0x9d, 0x3d, 0xa7, 0x52, 0x2c, 0x55, 0x50, 0x08, 0xbd, 0x98, 0xf0, 0xa2,
	0x41, 0xb8, 0x58, 0xaa, 0xe6, 0x87, 0x07, 0xe5, 0xdf, 0x2a, 0x7a, 0x09, 0x44, 0xaf, 0x73, 0x34,
	0xe9, 0xd5, 0xf6, 0xd1, 0x7f, 0x29, 0x6c, 0xb8, 0xce, 0x91, 0x1b, 0x09, 0xbd, 0x86, 0x9a, 0xeb,
	0x95, 0xe2, 0x0a, 0xb3, 0x85, 0xc6, 0xc0, 0x0f, 0xbd, 0xb8, 0xd4, 0x6e, 0xfc, 0xb8, 0x0c, 0x5c,
	0xed, 0xc4, 0x8e, 0x79, 0xd5, 0xf1, 0xdd, 0x3b, 0x6a, 0x01, 0xd9, 0xf8, 0xd1, 0x3a, 0xd4, 0x06,
	0x3c, 0x19, 0xdc, 0xf0, 0x24, 0x4d, 0x1e, 0x92, 0xce, 0xfd, 0x30, 0xd9, 0xdb, 0xb9, 0x22, 0xef,
	0x5f, 0xe7, 0x27, 0x51, 0x17, 0x0a, 0x77, 0x9b, 0xdf, 0x6f, 0xdf, 0xd5, 0xf6, 0xd6, 0xc6, 0xea,
	0x2c, 0x3a, 0x86, 0x42, 0x67, 0x2a, 0x14, 0x6e, 0xb5, 0x32, 0xdc, 0xd3, 0xdb, 0x03, 0x68, 0x64,
	0x62, 0xc6, 0xcc, 0xf9, 0x58, 0xf6, 0xca, 0xdc, 0x41, 0x1f, 0x17, 0x93, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0x0b, 0xf8, 0x0e, 0x06, 0x02, 0x00, 0x00,
}