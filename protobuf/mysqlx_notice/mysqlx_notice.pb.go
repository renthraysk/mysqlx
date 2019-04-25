// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mysqlx_notice.proto

package mysqlx_notice

/*
Notices

A notice

* is sent from the server to the client
* may be global or relate to the current message sequence
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/renthraysk/mysqlx/protobuf/mysqlx"
import mysqlx_datatypes "github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Frame_Scope int32

const (
	Frame_GLOBAL Frame_Scope = 1
	Frame_LOCAL  Frame_Scope = 2
)

var Frame_Scope_name = map[int32]string{
	1: "GLOBAL",
	2: "LOCAL",
}
var Frame_Scope_value = map[string]int32{
	"GLOBAL": 1,
	"LOCAL":  2,
}

func (x Frame_Scope) Enum() *Frame_Scope {
	p := new(Frame_Scope)
	*p = x
	return p
}
func (x Frame_Scope) String() string {
	return proto.EnumName(Frame_Scope_name, int32(x))
}
func (x *Frame_Scope) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Frame_Scope_value, data, "Frame_Scope")
	if err != nil {
		return err
	}
	*x = Frame_Scope(value)
	return nil
}
func (Frame_Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{0, 0}
}

type Frame_Type int32

const (
	Frame_WARNING                         Frame_Type = 1
	Frame_SESSION_VARIABLE_CHANGED        Frame_Type = 2
	Frame_SESSION_STATE_CHANGED           Frame_Type = 3
	Frame_GROUP_REPLICATION_STATE_CHANGED Frame_Type = 4
	Frame_SERVER_HELLO                    Frame_Type = 5
)

var Frame_Type_name = map[int32]string{
	1: "WARNING",
	2: "SESSION_VARIABLE_CHANGED",
	3: "SESSION_STATE_CHANGED",
	4: "GROUP_REPLICATION_STATE_CHANGED",
	5: "SERVER_HELLO",
}
var Frame_Type_value = map[string]int32{
	"WARNING":                         1,
	"SESSION_VARIABLE_CHANGED":        2,
	"SESSION_STATE_CHANGED":           3,
	"GROUP_REPLICATION_STATE_CHANGED": 4,
	"SERVER_HELLO":                    5,
}

func (x Frame_Type) Enum() *Frame_Type {
	p := new(Frame_Type)
	*p = x
	return p
}
func (x Frame_Type) String() string {
	return proto.EnumName(Frame_Type_name, int32(x))
}
func (x *Frame_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Frame_Type_value, data, "Frame_Type")
	if err != nil {
		return err
	}
	*x = Frame_Type(value)
	return nil
}
func (Frame_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{0, 1}
}

type Warning_Level int32

const (
	Warning_NOTE    Warning_Level = 1
	Warning_WARNING Warning_Level = 2
	Warning_ERROR   Warning_Level = 3
)

var Warning_Level_name = map[int32]string{
	1: "NOTE",
	2: "WARNING",
	3: "ERROR",
}
var Warning_Level_value = map[string]int32{
	"NOTE":    1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x Warning_Level) Enum() *Warning_Level {
	p := new(Warning_Level)
	*p = x
	return p
}
func (x Warning_Level) String() string {
	return proto.EnumName(Warning_Level_name, int32(x))
}
func (x *Warning_Level) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Warning_Level_value, data, "Warning_Level")
	if err != nil {
		return err
	}
	*x = Warning_Level(value)
	return nil
}
func (Warning_Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{1, 0}
}

type SessionStateChanged_Parameter int32

const (
	SessionStateChanged_CURRENT_SCHEMA         SessionStateChanged_Parameter = 1
	SessionStateChanged_ACCOUNT_EXPIRED        SessionStateChanged_Parameter = 2
	SessionStateChanged_GENERATED_INSERT_ID    SessionStateChanged_Parameter = 3
	SessionStateChanged_ROWS_AFFECTED          SessionStateChanged_Parameter = 4
	SessionStateChanged_ROWS_FOUND             SessionStateChanged_Parameter = 5
	SessionStateChanged_ROWS_MATCHED           SessionStateChanged_Parameter = 6
	SessionStateChanged_TRX_COMMITTED          SessionStateChanged_Parameter = 7
	SessionStateChanged_TRX_ROLLEDBACK         SessionStateChanged_Parameter = 9
	SessionStateChanged_PRODUCED_MESSAGE       SessionStateChanged_Parameter = 10
	SessionStateChanged_CLIENT_ID_ASSIGNED     SessionStateChanged_Parameter = 11
	SessionStateChanged_GENERATED_DOCUMENT_IDS SessionStateChanged_Parameter = 12
)

var SessionStateChanged_Parameter_name = map[int32]string{
	1:  "CURRENT_SCHEMA",
	2:  "ACCOUNT_EXPIRED",
	3:  "GENERATED_INSERT_ID",
	4:  "ROWS_AFFECTED",
	5:  "ROWS_FOUND",
	6:  "ROWS_MATCHED",
	7:  "TRX_COMMITTED",
	9:  "TRX_ROLLEDBACK",
	10: "PRODUCED_MESSAGE",
	11: "CLIENT_ID_ASSIGNED",
	12: "GENERATED_DOCUMENT_IDS",
}
var SessionStateChanged_Parameter_value = map[string]int32{
	"CURRENT_SCHEMA":         1,
	"ACCOUNT_EXPIRED":        2,
	"GENERATED_INSERT_ID":    3,
	"ROWS_AFFECTED":          4,
	"ROWS_FOUND":             5,
	"ROWS_MATCHED":           6,
	"TRX_COMMITTED":          7,
	"TRX_ROLLEDBACK":         9,
	"PRODUCED_MESSAGE":       10,
	"CLIENT_ID_ASSIGNED":     11,
	"GENERATED_DOCUMENT_IDS": 12,
}

func (x SessionStateChanged_Parameter) Enum() *SessionStateChanged_Parameter {
	p := new(SessionStateChanged_Parameter)
	*p = x
	return p
}
func (x SessionStateChanged_Parameter) String() string {
	return proto.EnumName(SessionStateChanged_Parameter_name, int32(x))
}
func (x *SessionStateChanged_Parameter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SessionStateChanged_Parameter_value, data, "SessionStateChanged_Parameter")
	if err != nil {
		return err
	}
	*x = SessionStateChanged_Parameter(value)
	return nil
}
func (SessionStateChanged_Parameter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{3, 0}
}

type GroupReplicationStateChanged_Type int32

const (
	GroupReplicationStateChanged_MEMBERSHIP_QUORUM_LOSS GroupReplicationStateChanged_Type = 1
	GroupReplicationStateChanged_MEMBERSHIP_VIEW_CHANGE GroupReplicationStateChanged_Type = 2
	GroupReplicationStateChanged_MEMBER_ROLE_CHANGE     GroupReplicationStateChanged_Type = 3
	GroupReplicationStateChanged_MEMBER_STATE_CHANGE    GroupReplicationStateChanged_Type = 4
)

var GroupReplicationStateChanged_Type_name = map[int32]string{
	1: "MEMBERSHIP_QUORUM_LOSS",
	2: "MEMBERSHIP_VIEW_CHANGE",
	3: "MEMBER_ROLE_CHANGE",
	4: "MEMBER_STATE_CHANGE",
}
var GroupReplicationStateChanged_Type_value = map[string]int32{
	"MEMBERSHIP_QUORUM_LOSS": 1,
	"MEMBERSHIP_VIEW_CHANGE": 2,
	"MEMBER_ROLE_CHANGE":     3,
	"MEMBER_STATE_CHANGE":    4,
}

func (x GroupReplicationStateChanged_Type) Enum() *GroupReplicationStateChanged_Type {
	p := new(GroupReplicationStateChanged_Type)
	*p = x
	return p
}
func (x GroupReplicationStateChanged_Type) String() string {
	return proto.EnumName(GroupReplicationStateChanged_Type_name, int32(x))
}
func (x *GroupReplicationStateChanged_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GroupReplicationStateChanged_Type_value, data, "GroupReplicationStateChanged_Type")
	if err != nil {
		return err
	}
	*x = GroupReplicationStateChanged_Type(value)
	return nil
}
func (GroupReplicationStateChanged_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{4, 0}
}

// Common Frame for all Notices
//
// ===================================================== =====
// .type                                                 value
// ===================================================== =====
// :protobuf:msg:`Mysqlx.Notice::Warning`                1
// :protobuf:msg:`Mysqlx.Notice::SessionVariableChanged` 2
// :protobuf:msg:`Mysqlx.Notice::SessionStateChanged`    3
// ===================================================== =====
//
// :param type: the type of the payload
// :param payload: the payload of the notification
// :param scope: global or local notification
//
type Frame struct {
	Type                 *uint32      `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Scope                *Frame_Scope `protobuf:"varint,2,opt,name=scope,enum=Mysqlx.Notice.Frame_Scope,def=1" json:"scope,omitempty"`
	Payload              []byte       `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{0}
}
func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (dst *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(dst, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

const Default_Frame_Scope Frame_Scope = Frame_GLOBAL

func (m *Frame) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Frame) GetScope() Frame_Scope {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return Default_Frame_Scope
}

func (m *Frame) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Server-side warnings and notes
//
// ``.scope`` == ``local``
//   ``.level``, ``.code`` and ``.msg`` map the content of
//
//   .. code-block:: sql
//
//     SHOW WARNINGS
//
// ``.scope`` == ``global``
//   (undefined) will be used for global, unstructured messages like:
//
//   * server is shutting down
//   * a node disconnected from group
//   * schema or table dropped
//
// ========================================== =======================
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =======================
// ``.type``                                  1
// ``.scope``                                 ``local`` or ``global``
// ========================================== =======================
//
// :param level: warning level: Note or Warning
// :param code: warning code
// :param msg: warning message
type Warning struct {
	Level                *Warning_Level `protobuf:"varint,1,opt,name=level,enum=Mysqlx.Notice.Warning_Level,def=2" json:"level,omitempty"`
	Code                 *uint32        `protobuf:"varint,2,req,name=code" json:"code,omitempty"`
	Msg                  *string        `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Warning) Reset()         { *m = Warning{} }
func (m *Warning) String() string { return proto.CompactTextString(m) }
func (*Warning) ProtoMessage()    {}
func (*Warning) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{1}
}
func (m *Warning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Warning.Unmarshal(m, b)
}
func (m *Warning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Warning.Marshal(b, m, deterministic)
}
func (dst *Warning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Warning.Merge(dst, src)
}
func (m *Warning) XXX_Size() int {
	return xxx_messageInfo_Warning.Size(m)
}
func (m *Warning) XXX_DiscardUnknown() {
	xxx_messageInfo_Warning.DiscardUnknown(m)
}

var xxx_messageInfo_Warning proto.InternalMessageInfo

const Default_Warning_Level Warning_Level = Warning_WARNING

func (m *Warning) GetLevel() Warning_Level {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return Default_Warning_Level
}

func (m *Warning) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Warning) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// Notify clients about changes to the current session variables
//
// Every change to a variable that is accessible through:
//
// .. code-block:: sql
//
//   SHOW SESSION VARIABLES
//
// ========================================== =========
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =========
// ``.type``                                  2
// ``.scope``                                 ``local``
// ========================================== =========
//
// :param namespace: namespace that param belongs to
// :param param: name of the variable
// :param value: the changed value of param
type SessionVariableChanged struct {
	Param                *string                  `protobuf:"bytes,1,req,name=param" json:"param,omitempty"`
	Value                *mysqlx_datatypes.Scalar `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SessionVariableChanged) Reset()         { *m = SessionVariableChanged{} }
func (m *SessionVariableChanged) String() string { return proto.CompactTextString(m) }
func (*SessionVariableChanged) ProtoMessage()    {}
func (*SessionVariableChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{2}
}
func (m *SessionVariableChanged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionVariableChanged.Unmarshal(m, b)
}
func (m *SessionVariableChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionVariableChanged.Marshal(b, m, deterministic)
}
func (dst *SessionVariableChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionVariableChanged.Merge(dst, src)
}
func (m *SessionVariableChanged) XXX_Size() int {
	return xxx_messageInfo_SessionVariableChanged.Size(m)
}
func (m *SessionVariableChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionVariableChanged.DiscardUnknown(m)
}

var xxx_messageInfo_SessionVariableChanged proto.InternalMessageInfo

func (m *SessionVariableChanged) GetParam() string {
	if m != nil && m.Param != nil {
		return *m.Param
	}
	return ""
}

func (m *SessionVariableChanged) GetValue() *mysqlx_datatypes.Scalar {
	if m != nil {
		return m.Value
	}
	return nil
}

// Notify clients about changes to the internal session state
//
// ========================================== =========
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== =========
// ``.type``                                  3
// ``.scope``                                 ``local``
// ========================================== =========
//
// :param param: parameter key
// :param value: updated value
type SessionStateChanged struct {
	Param                *SessionStateChanged_Parameter `protobuf:"varint,1,req,name=param,enum=Mysqlx.Notice.SessionStateChanged_Parameter" json:"param,omitempty"`
	Value                []*mysqlx_datatypes.Scalar     `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SessionStateChanged) Reset()         { *m = SessionStateChanged{} }
func (m *SessionStateChanged) String() string { return proto.CompactTextString(m) }
func (*SessionStateChanged) ProtoMessage()    {}
func (*SessionStateChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{3}
}
func (m *SessionStateChanged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionStateChanged.Unmarshal(m, b)
}
func (m *SessionStateChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionStateChanged.Marshal(b, m, deterministic)
}
func (dst *SessionStateChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionStateChanged.Merge(dst, src)
}
func (m *SessionStateChanged) XXX_Size() int {
	return xxx_messageInfo_SessionStateChanged.Size(m)
}
func (m *SessionStateChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionStateChanged.DiscardUnknown(m)
}

var xxx_messageInfo_SessionStateChanged proto.InternalMessageInfo

func (m *SessionStateChanged) GetParam() SessionStateChanged_Parameter {
	if m != nil && m.Param != nil {
		return *m.Param
	}
	return SessionStateChanged_CURRENT_SCHEMA
}

func (m *SessionStateChanged) GetValue() []*mysqlx_datatypes.Scalar {
	if m != nil {
		return m.Value
	}
	return nil
}

// Notify clients about group replication state changes
//
// ========================================== ==========
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== ==========
// ``.type``                                  4
// ``.scope``                                 ``global``
// ========================================== ==========
//
// :param type: type of group replication event
// :param view_id: view identifier
type GroupReplicationStateChanged struct {
	Type                 *uint32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	ViewId               *string  `protobuf:"bytes,2,opt,name=view_id,json=viewId" json:"view_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupReplicationStateChanged) Reset()         { *m = GroupReplicationStateChanged{} }
func (m *GroupReplicationStateChanged) String() string { return proto.CompactTextString(m) }
func (*GroupReplicationStateChanged) ProtoMessage()    {}
func (*GroupReplicationStateChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{4}
}
func (m *GroupReplicationStateChanged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupReplicationStateChanged.Unmarshal(m, b)
}
func (m *GroupReplicationStateChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupReplicationStateChanged.Marshal(b, m, deterministic)
}
func (dst *GroupReplicationStateChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupReplicationStateChanged.Merge(dst, src)
}
func (m *GroupReplicationStateChanged) XXX_Size() int {
	return xxx_messageInfo_GroupReplicationStateChanged.Size(m)
}
func (m *GroupReplicationStateChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupReplicationStateChanged.DiscardUnknown(m)
}

var xxx_messageInfo_GroupReplicationStateChanged proto.InternalMessageInfo

func (m *GroupReplicationStateChanged) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *GroupReplicationStateChanged) GetViewId() string {
	if m != nil && m.ViewId != nil {
		return *m.ViewId
	}
	return ""
}

// Notify clients about connection to X Protocol server
//
// ========================================== ==========
// :protobuf:msg:`Mysqlx.Notice::Frame` field value
// ========================================== ==========
// ``.type``                                  5
// ``.scope``                                 ``global``
// ========================================== ==========
//
type ServerHello struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerHello) Reset()         { *m = ServerHello{} }
func (m *ServerHello) String() string { return proto.CompactTextString(m) }
func (*ServerHello) ProtoMessage()    {}
func (*ServerHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_notice_3ee079b89c444548, []int{5}
}
func (m *ServerHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerHello.Unmarshal(m, b)
}
func (m *ServerHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerHello.Marshal(b, m, deterministic)
}
func (dst *ServerHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerHello.Merge(dst, src)
}
func (m *ServerHello) XXX_Size() int {
	return xxx_messageInfo_ServerHello.Size(m)
}
func (m *ServerHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerHello.DiscardUnknown(m)
}

var xxx_messageInfo_ServerHello proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Frame)(nil), "Mysqlx.Notice.Frame")
	proto.RegisterType((*Warning)(nil), "Mysqlx.Notice.Warning")
	proto.RegisterType((*SessionVariableChanged)(nil), "Mysqlx.Notice.SessionVariableChanged")
	proto.RegisterType((*SessionStateChanged)(nil), "Mysqlx.Notice.SessionStateChanged")
	proto.RegisterType((*GroupReplicationStateChanged)(nil), "Mysqlx.Notice.GroupReplicationStateChanged")
	proto.RegisterType((*ServerHello)(nil), "Mysqlx.Notice.ServerHello")
	proto.RegisterEnum("Mysqlx.Notice.Frame_Scope", Frame_Scope_name, Frame_Scope_value)
	proto.RegisterEnum("Mysqlx.Notice.Frame_Type", Frame_Type_name, Frame_Type_value)
	proto.RegisterEnum("Mysqlx.Notice.Warning_Level", Warning_Level_name, Warning_Level_value)
	proto.RegisterEnum("Mysqlx.Notice.SessionStateChanged_Parameter", SessionStateChanged_Parameter_name, SessionStateChanged_Parameter_value)
	proto.RegisterEnum("Mysqlx.Notice.GroupReplicationStateChanged_Type", GroupReplicationStateChanged_Type_name, GroupReplicationStateChanged_Type_value)
}

func init() { proto.RegisterFile("mysqlx_notice.proto", fileDescriptor_mysqlx_notice_3ee079b89c444548) }

var fileDescriptor_mysqlx_notice_3ee079b89c444548 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x72, 0xeb, 0x34,
	0x14, 0x1d, 0x3b, 0x71, 0x43, 0x6e, 0xda, 0x22, 0xd4, 0x47, 0x9a, 0x97, 0xe9, 0x40, 0xc7, 0x6c,
	0xca, 0x0c, 0xe3, 0x61, 0xba, 0x62, 0xca, 0x4a, 0xb1, 0xd5, 0xc4, 0x83, 0x6d, 0x05, 0xc9, 0x69,
	0xdf, 0x0a, 0x8d, 0x5e, 0x22, 0x4a, 0x18, 0x27, 0x0e, 0x8e, 0x5b, 0xe8, 0x9a, 0x1f, 0x60, 0xc7,
	0x82, 0x5f, 0xe1, 0x8b, 0xd8, 0xf2, 0x01, 0x30, 0xb2, 0x9d, 0xbe, 0x12, 0xba, 0x78, 0x3b, 0xdd,
	0x7b, 0xcf, 0x91, 0x8e, 0xcf, 0xbd, 0xd7, 0x70, 0xb2, 0x7a, 0xdc, 0xfe, 0x94, 0xfd, 0x22, 0xd7,
	0x79, 0xb9, 0x9c, 0x6b, 0x6f, 0x53, 0xe4, 0x65, 0x8e, 0x8f, 0xe2, 0x2a, 0xe9, 0x25, 0x55, 0x72,
	0x78, 0x58, 0x63, 0xea, 0xe2, 0xb0, 0xdf, 0x30, 0x16, 0xaa, 0x54, 0xe5, 0xe3, 0x46, 0x6f, 0xeb,
	0xbc, 0xfb, 0x87, 0x0d, 0xce, 0x75, 0xa1, 0x56, 0x1a, 0x63, 0x68, 0x9b, 0xc2, 0xc0, 0x3a, 0xb7,
	0x2f, 0x8e, 0x78, 0x75, 0xc6, 0x5f, 0x81, 0xb3, 0x9d, 0xe7, 0x1b, 0x3d, 0xb0, 0xcf, 0xad, 0x8b,
	0xe3, 0xcb, 0xa1, 0xf7, 0x9f, 0x27, 0xbc, 0x8a, 0xe8, 0x09, 0x83, 0xb8, 0x3a, 0x18, 0x47, 0x6c,
	0x44, 0x22, 0x5e, 0x13, 0xf0, 0x00, 0x3a, 0x1b, 0xf5, 0x98, 0xe5, 0x6a, 0x31, 0x68, 0x9d, 0x5b,
	0x17, 0x87, 0x7c, 0x17, 0xba, 0x9f, 0x80, 0x53, 0x31, 0x30, 0x40, 0xc3, 0x41, 0x16, 0xee, 0x82,
	0x13, 0x31, 0x9f, 0x44, 0xc8, 0x76, 0x7f, 0xb5, 0xa0, 0x9d, 0x9a, 0xc7, 0x7b, 0xd0, 0xb9, 0x25,
	0x3c, 0x09, 0x93, 0x31, 0xb2, 0xf0, 0x19, 0x0c, 0x04, 0x15, 0x22, 0x64, 0x89, 0xbc, 0x21, 0x3c,
	0x24, 0xa3, 0x88, 0x4a, 0x7f, 0x42, 0x92, 0x31, 0x0d, 0x90, 0x8d, 0x5f, 0xc3, 0xc7, 0xbb, 0xaa,
	0x48, 0x49, 0xfa, 0xae, 0xd4, 0xc2, 0x9f, 0xc1, 0xa7, 0x63, 0xce, 0x66, 0x53, 0xc9, 0xe9, 0x34,
	0x0a, 0x7d, 0x92, 0xfe, 0x1f, 0xd4, 0xc6, 0x08, 0x0e, 0x05, 0xe5, 0x37, 0x94, 0xcb, 0x09, 0x8d,
	0x22, 0x86, 0x9c, 0xab, 0xf6, 0x6f, 0x7f, 0x7d, 0xd9, 0x73, 0x7f, 0xb7, 0xa0, 0x73, 0xab, 0x8a,
	0xf5, 0x72, 0x7d, 0x87, 0xbf, 0x06, 0x27, 0xd3, 0x0f, 0x3a, 0x1b, 0x58, 0x95, 0x17, 0x67, 0x7b,
	0x5e, 0x34, 0x30, 0x2f, 0x32, 0x98, 0xab, 0x9d, 0x72, 0x5e, 0x73, 0x8c, 0xb9, 0xf3, 0x7c, 0x61,
	0x7c, 0xac, 0xcc, 0x35, 0x67, 0x8c, 0xa0, 0xb5, 0xda, 0xde, 0x0d, 0x5a, 0xe7, 0xf6, 0x45, 0x97,
	0x9b, 0xa3, 0xfb, 0x39, 0x38, 0x15, 0x1d, 0x7f, 0x00, 0xed, 0x84, 0xa5, 0x14, 0x59, 0xcf, 0x4d,
	0xb0, 0x8d, 0x4b, 0x94, 0x73, 0xc6, 0x51, 0xcb, 0xfd, 0x0e, 0xfa, 0x42, 0x6f, 0xb7, 0xcb, 0x7c,
	0x7d, 0xa3, 0x8a, 0xa5, 0x7a, 0x9b, 0x69, 0xff, 0x07, 0xb5, 0xbe, 0xd3, 0x0b, 0xfc, 0x0a, 0x9c,
	0x8d, 0x2a, 0xd4, 0xaa, 0x6a, 0x64, 0x97, 0xd7, 0x01, 0xf6, 0xc0, 0x79, 0x50, 0xd9, 0x7d, 0xdd,
	0xc9, 0xde, 0xe5, 0x60, 0xa7, 0x3e, 0x78, 0x9a, 0x07, 0x31, 0x57, 0x99, 0x2a, 0x78, 0x0d, 0x73,
	0xff, 0xb1, 0xe1, 0xa4, 0x79, 0x40, 0x94, 0xaa, 0x7c, 0xba, 0x7d, 0xf4, 0xfc, 0xf6, 0xe3, 0xcb,
	0x2f, 0xf6, 0x5c, 0x78, 0x81, 0xe2, 0x4d, 0x0d, 0x5e, 0x97, 0xba, 0x78, 0x41, 0x4b, 0xeb, 0x7d,
	0xb4, 0xfc, 0x6d, 0x41, 0xf7, 0xe9, 0x12, 0x8c, 0xe1, 0xd8, 0x9f, 0x71, 0x4e, 0x93, 0x54, 0x0a,
	0x7f, 0x42, 0x63, 0x82, 0x2c, 0x7c, 0x02, 0x1f, 0x12, 0xdf, 0x67, 0xb3, 0x24, 0x95, 0xf4, 0xcd,
	0x34, 0xe4, 0xd5, 0x50, 0x9c, 0xc2, 0xc9, 0x98, 0x26, 0x94, 0x93, 0x94, 0x06, 0x32, 0x4c, 0x04,
	0xe5, 0xa9, 0x0c, 0xcd, 0x48, 0x7c, 0x04, 0x47, 0x9c, 0xdd, 0x0a, 0x49, 0xae, 0xaf, 0xa9, 0x9f,
	0x56, 0x03, 0x70, 0x0c, 0x50, 0xa5, 0xae, 0xd9, 0x2c, 0x09, 0x90, 0x63, 0x06, 0xa2, 0x8a, 0x63,
	0x92, 0xfa, 0x13, 0x1a, 0xa0, 0x03, 0x43, 0x4a, 0xf9, 0x1b, 0xe9, 0xb3, 0x38, 0x0e, 0x53, 0x43,
	0xea, 0x18, 0x25, 0x26, 0xc5, 0x59, 0x14, 0xd1, 0x60, 0x44, 0xfc, 0x6f, 0x50, 0x17, 0xbf, 0x02,
	0x34, 0xe5, 0x2c, 0x98, 0xf9, 0x34, 0x90, 0x31, 0x15, 0x82, 0x8c, 0x29, 0x02, 0xdc, 0x07, 0xec,
	0x47, 0xa1, 0x91, 0x1c, 0x06, 0x92, 0x08, 0x11, 0x8e, 0x13, 0x1a, 0xa0, 0x1e, 0x1e, 0x42, 0xff,
	0x9d, 0xc4, 0x80, 0xf9, 0xb3, 0xb8, 0xc6, 0x08, 0x74, 0xe8, 0xfe, 0x69, 0xc1, 0xd9, 0xb8, 0xc8,
	0xef, 0x37, 0x5c, 0x6f, 0xb2, 0xe5, 0x5c, 0x95, 0xfb, 0xad, 0x78, 0x69, 0x61, 0x4f, 0xa1, 0xf3,
	0xb0, 0xd4, 0x3f, 0xcb, 0xe5, 0xa2, 0x6a, 0x74, 0x97, 0x1f, 0x98, 0x30, 0x5c, 0xb8, 0x79, 0xb3,
	0x54, 0x43, 0xe8, 0xc7, 0x34, 0x1e, 0x51, 0x2e, 0x26, 0xe1, 0x54, 0x7e, 0x3b, 0x63, 0x7c, 0x16,
	0xcb, 0x88, 0x09, 0x81, 0xac, 0xbd, 0xda, 0x4d, 0x48, 0x6f, 0x9b, 0x15, 0x41, 0xb6, 0xf9, 0x82,
	0xba, 0x66, 0x3e, 0x77, 0xb7, 0x3a, 0xa8, 0x65, 0x4c, 0x6e, 0xf2, 0xcf, 0x77, 0x0a, 0xb5, 0xdd,
	0x23, 0xe8, 0x09, 0x5d, 0x3c, 0xe8, 0x62, 0xa2, 0xb3, 0x2c, 0x1f, 0xbd, 0x86, 0xd3, 0x79, 0xbe,
	0xf2, 0xaa, 0xbf, 0x90, 0x37, 0xff, 0xd1, 0x6b, 0xfe, 0x4b, 0x6f, 0xef, 0xbf, 0xff, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x47, 0x65, 0xb8, 0x86, 0xcd, 0x04, 0x00, 0x00,
}
