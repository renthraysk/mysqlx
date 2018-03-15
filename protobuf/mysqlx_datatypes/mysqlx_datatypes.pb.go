// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mysqlx_datatypes.proto

package mysqlx_datatypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Scalar_Type int32

const (
	Scalar_V_SINT   Scalar_Type = 1
	Scalar_V_UINT   Scalar_Type = 2
	Scalar_V_NULL   Scalar_Type = 3
	Scalar_V_OCTETS Scalar_Type = 4
	Scalar_V_DOUBLE Scalar_Type = 5
	Scalar_V_FLOAT  Scalar_Type = 6
	Scalar_V_BOOL   Scalar_Type = 7
	Scalar_V_STRING Scalar_Type = 8
)

var Scalar_Type_name = map[int32]string{
	1: "V_SINT",
	2: "V_UINT",
	3: "V_NULL",
	4: "V_OCTETS",
	5: "V_DOUBLE",
	6: "V_FLOAT",
	7: "V_BOOL",
	8: "V_STRING",
}
var Scalar_Type_value = map[string]int32{
	"V_SINT":   1,
	"V_UINT":   2,
	"V_NULL":   3,
	"V_OCTETS": 4,
	"V_DOUBLE": 5,
	"V_FLOAT":  6,
	"V_BOOL":   7,
	"V_STRING": 8,
}

func (x Scalar_Type) Enum() *Scalar_Type {
	p := new(Scalar_Type)
	*p = x
	return p
}
func (x Scalar_Type) String() string {
	return proto.EnumName(Scalar_Type_name, int32(x))
}
func (x *Scalar_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Scalar_Type_value, data, "Scalar_Type")
	if err != nil {
		return err
	}
	*x = Scalar_Type(value)
	return nil
}
func (Scalar_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{0, 0}
}

type Any_Type int32

const (
	Any_SCALAR Any_Type = 1
	Any_OBJECT Any_Type = 2
	Any_ARRAY  Any_Type = 3
)

var Any_Type_name = map[int32]string{
	1: "SCALAR",
	2: "OBJECT",
	3: "ARRAY",
}
var Any_Type_value = map[string]int32{
	"SCALAR": 1,
	"OBJECT": 2,
	"ARRAY":  3,
}

func (x Any_Type) Enum() *Any_Type {
	p := new(Any_Type)
	*p = x
	return p
}
func (x Any_Type) String() string {
	return proto.EnumName(Any_Type_name, int32(x))
}
func (x *Any_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Any_Type_value, data, "Any_Type")
	if err != nil {
		return err
	}
	*x = Any_Type(value)
	return nil
}
func (Any_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{3, 0}
}

// a scalar
type Scalar struct {
	Type *Scalar_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Datatypes.Scalar_Type" json:"type,omitempty"`
	// Types that are valid to be assigned to Types:
	//	*Scalar_VSignedInt
	//	*Scalar_VUnsignedInt
	//	*Scalar_VOctets
	//	*Scalar_VDouble
	//	*Scalar_VFloat
	//	*Scalar_VBool
	//	*Scalar_VString
	Types                isScalar_Types `protobuf_oneof:"types"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Scalar) Reset()         { *m = Scalar{} }
func (m *Scalar) String() string { return proto.CompactTextString(m) }
func (*Scalar) ProtoMessage()    {}
func (*Scalar) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{0}
}
func (m *Scalar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scalar.Unmarshal(m, b)
}
func (m *Scalar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scalar.Marshal(b, m, deterministic)
}
func (dst *Scalar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scalar.Merge(dst, src)
}
func (m *Scalar) XXX_Size() int {
	return xxx_messageInfo_Scalar.Size(m)
}
func (m *Scalar) XXX_DiscardUnknown() {
	xxx_messageInfo_Scalar.DiscardUnknown(m)
}

var xxx_messageInfo_Scalar proto.InternalMessageInfo

type isScalar_Types interface {
	isScalar_Types()
}

type Scalar_VSignedInt struct {
	VSignedInt int64 `protobuf:"zigzag64,2,opt,name=v_signed_int,json=vSignedInt,oneof"`
}
type Scalar_VUnsignedInt struct {
	VUnsignedInt uint64 `protobuf:"varint,3,opt,name=v_unsigned_int,json=vUnsignedInt,oneof"`
}
type Scalar_VOctets struct {
	VOctets *Scalar_Octets `protobuf:"bytes,5,opt,name=v_octets,json=vOctets,oneof"`
}
type Scalar_VDouble struct {
	VDouble float64 `protobuf:"fixed64,6,opt,name=v_double,json=vDouble,oneof"`
}
type Scalar_VFloat struct {
	VFloat float32 `protobuf:"fixed32,7,opt,name=v_float,json=vFloat,oneof"`
}
type Scalar_VBool struct {
	VBool bool `protobuf:"varint,8,opt,name=v_bool,json=vBool,oneof"`
}
type Scalar_VString struct {
	VString *Scalar_String `protobuf:"bytes,9,opt,name=v_string,json=vString,oneof"`
}

func (*Scalar_VSignedInt) isScalar_Types()   {}
func (*Scalar_VUnsignedInt) isScalar_Types() {}
func (*Scalar_VOctets) isScalar_Types()      {}
func (*Scalar_VDouble) isScalar_Types()      {}
func (*Scalar_VFloat) isScalar_Types()       {}
func (*Scalar_VBool) isScalar_Types()        {}
func (*Scalar_VString) isScalar_Types()      {}

func (m *Scalar) GetTypes() isScalar_Types {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Scalar) GetType() Scalar_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Scalar_V_SINT
}

func (m *Scalar) GetVSignedInt() int64 {
	if x, ok := m.GetTypes().(*Scalar_VSignedInt); ok {
		return x.VSignedInt
	}
	return 0
}

func (m *Scalar) GetVUnsignedInt() uint64 {
	if x, ok := m.GetTypes().(*Scalar_VUnsignedInt); ok {
		return x.VUnsignedInt
	}
	return 0
}

func (m *Scalar) GetVOctets() *Scalar_Octets {
	if x, ok := m.GetTypes().(*Scalar_VOctets); ok {
		return x.VOctets
	}
	return nil
}

func (m *Scalar) GetVDouble() float64 {
	if x, ok := m.GetTypes().(*Scalar_VDouble); ok {
		return x.VDouble
	}
	return 0
}

func (m *Scalar) GetVFloat() float32 {
	if x, ok := m.GetTypes().(*Scalar_VFloat); ok {
		return x.VFloat
	}
	return 0
}

func (m *Scalar) GetVBool() bool {
	if x, ok := m.GetTypes().(*Scalar_VBool); ok {
		return x.VBool
	}
	return false
}

func (m *Scalar) GetVString() *Scalar_String {
	if x, ok := m.GetTypes().(*Scalar_VString); ok {
		return x.VString
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Scalar) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Scalar_OneofMarshaler, _Scalar_OneofUnmarshaler, _Scalar_OneofSizer, []interface{}{
		(*Scalar_VSignedInt)(nil),
		(*Scalar_VUnsignedInt)(nil),
		(*Scalar_VOctets)(nil),
		(*Scalar_VDouble)(nil),
		(*Scalar_VFloat)(nil),
		(*Scalar_VBool)(nil),
		(*Scalar_VString)(nil),
	}
}

func _Scalar_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Scalar)
	// types
	switch x := m.Types.(type) {
	case *Scalar_VSignedInt:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeZigzag64(uint64(x.VSignedInt))
	case *Scalar_VUnsignedInt:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.VUnsignedInt))
	case *Scalar_VOctets:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VOctets); err != nil {
			return err
		}
	case *Scalar_VDouble:
		b.EncodeVarint(6<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.VDouble))
	case *Scalar_VFloat:
		b.EncodeVarint(7<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.VFloat)))
	case *Scalar_VBool:
		t := uint64(0)
		if x.VBool {
			t = 1
		}
		b.EncodeVarint(8<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Scalar_VString:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VString); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Scalar.Types has unexpected type %T", x)
	}
	return nil
}

func _Scalar_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Scalar)
	switch tag {
	case 2: // types.v_signed_int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag64()
		m.Types = &Scalar_VSignedInt{int64(x)}
		return true, err
	case 3: // types.v_unsigned_int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Types = &Scalar_VUnsignedInt{x}
		return true, err
	case 5: // types.v_octets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Scalar_Octets)
		err := b.DecodeMessage(msg)
		m.Types = &Scalar_VOctets{msg}
		return true, err
	case 6: // types.v_double
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Types = &Scalar_VDouble{math.Float64frombits(x)}
		return true, err
	case 7: // types.v_float
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Types = &Scalar_VFloat{math.Float32frombits(uint32(x))}
		return true, err
	case 8: // types.v_bool
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Types = &Scalar_VBool{x != 0}
		return true, err
	case 9: // types.v_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Scalar_String)
		err := b.DecodeMessage(msg)
		m.Types = &Scalar_VString{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Scalar_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Scalar)
	// types
	switch x := m.Types.(type) {
	case *Scalar_VSignedInt:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(uint64(x.VSignedInt<<1) ^ uint64((int64(x.VSignedInt) >> 63))))
	case *Scalar_VUnsignedInt:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.VUnsignedInt))
	case *Scalar_VOctets:
		s := proto.Size(x.VOctets)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Scalar_VDouble:
		n += proto.SizeVarint(6<<3 | proto.WireFixed64)
		n += 8
	case *Scalar_VFloat:
		n += proto.SizeVarint(7<<3 | proto.WireFixed32)
		n += 4
	case *Scalar_VBool:
		n += proto.SizeVarint(8<<3 | proto.WireVarint)
		n += 1
	case *Scalar_VString:
		s := proto.Size(x.VString)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// a string with a charset/collation
type Scalar_String struct {
	Value                []byte   `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	Collation            *uint64  `protobuf:"varint,2,opt,name=collation" json:"collation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scalar_String) Reset()         { *m = Scalar_String{} }
func (m *Scalar_String) String() string { return proto.CompactTextString(m) }
func (*Scalar_String) ProtoMessage()    {}
func (*Scalar_String) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{0, 0}
}
func (m *Scalar_String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scalar_String.Unmarshal(m, b)
}
func (m *Scalar_String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scalar_String.Marshal(b, m, deterministic)
}
func (dst *Scalar_String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scalar_String.Merge(dst, src)
}
func (m *Scalar_String) XXX_Size() int {
	return xxx_messageInfo_Scalar_String.Size(m)
}
func (m *Scalar_String) XXX_DiscardUnknown() {
	xxx_messageInfo_Scalar_String.DiscardUnknown(m)
}

var xxx_messageInfo_Scalar_String proto.InternalMessageInfo

func (m *Scalar_String) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Scalar_String) GetCollation() uint64 {
	if m != nil && m.Collation != nil {
		return *m.Collation
	}
	return 0
}

// an opaque octet sequence, with an optional content_type
// See ``Mysqlx.Resultset.ColumnMetadata`` for list of known values.
type Scalar_Octets struct {
	Value                []byte   `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	ContentType          *uint32  `protobuf:"varint,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scalar_Octets) Reset()         { *m = Scalar_Octets{} }
func (m *Scalar_Octets) String() string { return proto.CompactTextString(m) }
func (*Scalar_Octets) ProtoMessage()    {}
func (*Scalar_Octets) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{0, 1}
}
func (m *Scalar_Octets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scalar_Octets.Unmarshal(m, b)
}
func (m *Scalar_Octets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scalar_Octets.Marshal(b, m, deterministic)
}
func (dst *Scalar_Octets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scalar_Octets.Merge(dst, src)
}
func (m *Scalar_Octets) XXX_Size() int {
	return xxx_messageInfo_Scalar_Octets.Size(m)
}
func (m *Scalar_Octets) XXX_DiscardUnknown() {
	xxx_messageInfo_Scalar_Octets.DiscardUnknown(m)
}

var xxx_messageInfo_Scalar_Octets proto.InternalMessageInfo

func (m *Scalar_Octets) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Scalar_Octets) GetContentType() uint32 {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return 0
}

// a object
type Object struct {
	Fld                  []*Object_ObjectField `protobuf:"bytes,1,rep,name=fld" json:"fld,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{1}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (dst *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(dst, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetFld() []*Object_ObjectField {
	if m != nil {
		return m.Fld
	}
	return nil
}

type Object_ObjectField struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value                *Any     `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Object_ObjectField) Reset()         { *m = Object_ObjectField{} }
func (m *Object_ObjectField) String() string { return proto.CompactTextString(m) }
func (*Object_ObjectField) ProtoMessage()    {}
func (*Object_ObjectField) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{1, 0}
}
func (m *Object_ObjectField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object_ObjectField.Unmarshal(m, b)
}
func (m *Object_ObjectField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object_ObjectField.Marshal(b, m, deterministic)
}
func (dst *Object_ObjectField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object_ObjectField.Merge(dst, src)
}
func (m *Object_ObjectField) XXX_Size() int {
	return xxx_messageInfo_Object_ObjectField.Size(m)
}
func (m *Object_ObjectField) XXX_DiscardUnknown() {
	xxx_messageInfo_Object_ObjectField.DiscardUnknown(m)
}

var xxx_messageInfo_Object_ObjectField proto.InternalMessageInfo

func (m *Object_ObjectField) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Object_ObjectField) GetValue() *Any {
	if m != nil {
		return m.Value
	}
	return nil
}

// a Array
type Array struct {
	Value                []*Any   `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Array) Reset()         { *m = Array{} }
func (m *Array) String() string { return proto.CompactTextString(m) }
func (*Array) ProtoMessage()    {}
func (*Array) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{2}
}
func (m *Array) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Array.Unmarshal(m, b)
}
func (m *Array) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Array.Marshal(b, m, deterministic)
}
func (dst *Array) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Array.Merge(dst, src)
}
func (m *Array) XXX_Size() int {
	return xxx_messageInfo_Array.Size(m)
}
func (m *Array) XXX_DiscardUnknown() {
	xxx_messageInfo_Array.DiscardUnknown(m)
}

var xxx_messageInfo_Array proto.InternalMessageInfo

func (m *Array) GetValue() []*Any {
	if m != nil {
		return m.Value
	}
	return nil
}

// a helper to allow all field types
type Any struct {
	Type *Any_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Datatypes.Any_Type" json:"type,omitempty"`
	// Types that are valid to be assigned to Types:
	//	*Any_Scalar
	//	*Any_Obj
	//	*Any_Array
	Types                isAny_Types `protobuf_oneof:"types"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Any) Reset()         { *m = Any{} }
func (m *Any) String() string { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()    {}
func (*Any) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysqlx_datatypes_b62682b2576edded, []int{3}
}
func (m *Any) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Any.Unmarshal(m, b)
}
func (m *Any) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Any.Marshal(b, m, deterministic)
}
func (dst *Any) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Any.Merge(dst, src)
}
func (m *Any) XXX_Size() int {
	return xxx_messageInfo_Any.Size(m)
}
func (m *Any) XXX_DiscardUnknown() {
	xxx_messageInfo_Any.DiscardUnknown(m)
}

var xxx_messageInfo_Any proto.InternalMessageInfo

type isAny_Types interface {
	isAny_Types()
}

type Any_Scalar struct {
	Scalar *Scalar `protobuf:"bytes,2,opt,name=scalar,oneof"`
}
type Any_Obj struct {
	Obj *Object `protobuf:"bytes,3,opt,name=obj,oneof"`
}
type Any_Array struct {
	Array *Array `protobuf:"bytes,4,opt,name=array,oneof"`
}

func (*Any_Scalar) isAny_Types() {}
func (*Any_Obj) isAny_Types()    {}
func (*Any_Array) isAny_Types()  {}

func (m *Any) GetTypes() isAny_Types {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Any) GetType() Any_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Any_SCALAR
}

func (m *Any) GetScalar() *Scalar {
	if x, ok := m.GetTypes().(*Any_Scalar); ok {
		return x.Scalar
	}
	return nil
}

func (m *Any) GetObj() *Object {
	if x, ok := m.GetTypes().(*Any_Obj); ok {
		return x.Obj
	}
	return nil
}

func (m *Any) GetArray() *Array {
	if x, ok := m.GetTypes().(*Any_Array); ok {
		return x.Array
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Any) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Any_OneofMarshaler, _Any_OneofUnmarshaler, _Any_OneofSizer, []interface{}{
		(*Any_Scalar)(nil),
		(*Any_Obj)(nil),
		(*Any_Array)(nil),
	}
}

func _Any_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Any)
	// types
	switch x := m.Types.(type) {
	case *Any_Scalar:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Scalar); err != nil {
			return err
		}
	case *Any_Obj:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Obj); err != nil {
			return err
		}
	case *Any_Array:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Array); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Any.Types has unexpected type %T", x)
	}
	return nil
}

func _Any_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Any)
	switch tag {
	case 2: // types.scalar
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Scalar)
		err := b.DecodeMessage(msg)
		m.Types = &Any_Scalar{msg}
		return true, err
	case 3: // types.obj
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Object)
		err := b.DecodeMessage(msg)
		m.Types = &Any_Obj{msg}
		return true, err
	case 4: // types.array
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Array)
		err := b.DecodeMessage(msg)
		m.Types = &Any_Array{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Any_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Any)
	// types
	switch x := m.Types.(type) {
	case *Any_Scalar:
		s := proto.Size(x.Scalar)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Any_Obj:
		s := proto.Size(x.Obj)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Any_Array:
		s := proto.Size(x.Array)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Scalar)(nil), "Mysqlx.Datatypes.Scalar")
	proto.RegisterType((*Scalar_String)(nil), "Mysqlx.Datatypes.Scalar.String")
	proto.RegisterType((*Scalar_Octets)(nil), "Mysqlx.Datatypes.Scalar.Octets")
	proto.RegisterType((*Object)(nil), "Mysqlx.Datatypes.Object")
	proto.RegisterType((*Object_ObjectField)(nil), "Mysqlx.Datatypes.Object.ObjectField")
	proto.RegisterType((*Array)(nil), "Mysqlx.Datatypes.Array")
	proto.RegisterType((*Any)(nil), "Mysqlx.Datatypes.Any")
	proto.RegisterEnum("Mysqlx.Datatypes.Scalar_Type", Scalar_Type_name, Scalar_Type_value)
	proto.RegisterEnum("Mysqlx.Datatypes.Any_Type", Any_Type_name, Any_Type_value)
}

func init() {
	proto.RegisterFile("mysqlx_datatypes.proto", fileDescriptor_mysqlx_datatypes_b62682b2576edded)
}

var fileDescriptor_mysqlx_datatypes_b62682b2576edded = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xad, 0x9b, 0x8f, 0xb6, 0x37, 0x65, 0x8a, 0x2c, 0x60, 0xa1, 0x0c, 0x08, 0x15, 0x42, 0x41,
	0xa0, 0x20, 0x2a, 0xc4, 0xd3, 0x5e, 0xd2, 0x7d, 0x90, 0xa1, 0xb0, 0x48, 0x4e, 0x3a, 0x89, 0xa7,
	0x28, 0x4d, 0xb3, 0xa9, 0x23, 0x8b, 0x47, 0xe3, 0x46, 0xcb, 0x9f, 0xe0, 0x9f, 0xf1, 0x8f, 0x78,
	0x40, 0xb6, 0x3b, 0x56, 0x58, 0x07, 0x4f, 0x39, 0x3e, 0x39, 0x27, 0xb9, 0xbe, 0xf7, 0xd8, 0xf0,
	0xf0, 0xa2, 0xa9, 0xbe, 0x15, 0x57, 0xc9, 0x2c, 0x65, 0x29, 0x6b, 0x2e, 0xf3, 0xca, 0xbd, 0x5c,
	0x50, 0x46, 0xb1, 0xf9, 0x59, 0xf0, 0xee, 0xfe, 0x35, 0x3f, 0xfc, 0xa1, 0x82, 0x1e, 0x65, 0x69,
	0x91, 0x2e, 0xf0, 0x3b, 0x50, 0x39, 0x67, 0x21, 0xbb, 0xed, 0x6c, 0x8d, 0x9e, 0xb8, 0x7f, 0x6b,
	0x5d, 0xa9, 0x73, 0xe3, 0xe6, 0x32, 0x27, 0x42, 0x8a, 0x87, 0xd0, 0xaf, 0x93, 0x6a, 0x7e, 0x56,
	0xe6, 0xb3, 0x64, 0x5e, 0x32, 0xab, 0x6d, 0x23, 0x07, 0xfb, 0x2d, 0x02, 0x75, 0x24, 0xc8, 0xa3,
	0x92, 0xe1, 0x97, 0xb0, 0x55, 0x27, 0xcb, 0x72, 0x4d, 0xa5, 0xd8, 0xc8, 0x51, 0xfd, 0x16, 0xe9,
	0xd7, 0x93, 0x15, 0xcd, 0x75, 0xbb, 0xd0, 0xad, 0x13, 0x9a, 0xb1, 0x9c, 0x55, 0x96, 0x66, 0x23,
	0xc7, 0x18, 0x3d, 0xbb, 0xb3, 0x84, 0x50, 0xc8, 0xfc, 0x16, 0xe9, 0xd4, 0x12, 0xe2, 0xc7, 0xdc,
	0x3d, 0xa3, 0xcb, 0x69, 0x91, 0x5b, 0xba, 0x8d, 0x1c, 0x24, 0x5e, 0xee, 0x0b, 0x02, 0x3f, 0x82,
	0x4e, 0x9d, 0x9c, 0x16, 0x34, 0x65, 0x56, 0xc7, 0x46, 0x4e, 0xdb, 0x6f, 0x11, 0xbd, 0x3e, 0xe4,
	0x6b, 0xbc, 0x0d, 0x7a, 0x9d, 0x4c, 0x29, 0x2d, 0xac, 0xae, 0x8d, 0x9c, 0xae, 0xdf, 0x22, 0x5a,
	0x3d, 0xa6, 0xb4, 0x90, 0xe5, 0x54, 0x6c, 0x31, 0x2f, 0xcf, 0xac, 0xde, 0x7f, 0xca, 0x89, 0x84,
	0x4c, 0xfc, 0x51, 0xc2, 0xc1, 0x2e, 0xe8, 0x12, 0xe1, 0xfb, 0xa0, 0xd5, 0x69, 0xb1, 0x94, 0x6d,
	0xed, 0x13, 0xb9, 0xc0, 0x3b, 0xd0, 0xcb, 0x68, 0x51, 0xa4, 0x6c, 0x4e, 0x4b, 0xd1, 0x35, 0x95,
	0xdc, 0x10, 0x03, 0x0f, 0xf4, 0xd5, 0xb6, 0x36, 0xbb, 0x9f, 0x43, 0x3f, 0xa3, 0x25, 0xcb, 0x4b,
	0x96, 0x88, 0x89, 0xf1, 0x0f, 0xdc, 0x23, 0xc6, 0x8a, 0xe3, 0xf3, 0x19, 0x5e, 0x80, 0xca, 0x9f,
	0x18, 0x40, 0x3f, 0x49, 0xa2, 0xa3, 0xe3, 0xd8, 0x44, 0x12, 0x4f, 0x38, 0x6e, 0x4b, 0x7c, 0x3c,
	0x09, 0x02, 0x53, 0xc1, 0x7d, 0xe8, 0x9e, 0x24, 0xe1, 0x5e, 0x7c, 0x10, 0x47, 0xa6, 0x2a, 0x57,
	0xfb, 0xe1, 0x64, 0x1c, 0x1c, 0x98, 0x1a, 0x36, 0xa0, 0x73, 0x92, 0x1c, 0x06, 0xa1, 0x17, 0x9b,
	0xba, 0x34, 0x8d, 0xc3, 0x30, 0x30, 0x3b, 0x52, 0x16, 0xc5, 0xe4, 0xe8, 0xf8, 0xa3, 0xd9, 0x1d,
	0x77, 0x40, 0x93, 0x79, 0xfa, 0x8e, 0x40, 0x0f, 0xa7, 0xe7, 0x79, 0xc6, 0xf0, 0x07, 0x50, 0x4e,
	0x8b, 0x99, 0x85, 0x6c, 0xc5, 0x31, 0x46, 0x2f, 0x6e, 0x37, 0x4f, 0xca, 0x56, 0x8f, 0xc3, 0x79,
	0x5e, 0xcc, 0x08, 0x37, 0x0c, 0x02, 0x30, 0xd6, 0x38, 0x6c, 0x82, 0xf2, 0x35, 0x6f, 0x44, 0x03,
	0x7a, 0x84, 0x43, 0xfc, 0xfa, 0xba, 0x29, 0x6d, 0xbb, 0xed, 0x18, 0xa3, 0x07, 0xb7, 0x3f, 0xed,
	0x95, 0xcd, 0xaa, 0x57, 0xc3, 0xf7, 0xa0, 0x79, 0x8b, 0x45, 0xba, 0xe6, 0x92, 0x05, 0xfd, 0xdb,
	0xf5, 0x13, 0x81, 0xe2, 0x95, 0x0d, 0x76, 0xff, 0x38, 0x13, 0x83, 0x8d, 0x9e, 0xf5, 0x03, 0x31,
	0x02, 0xbd, 0x12, 0x99, 0x10, 0x33, 0x31, 0x46, 0xd6, 0x5d, 0x99, 0xe1, 0x11, 0x94, 0x4a, 0xfc,
	0x06, 0x14, 0x3a, 0x3d, 0x17, 0xa7, 0x62, 0xa3, 0x41, 0x36, 0xc3, 0x6f, 0x11, 0x2e, 0xc3, 0x6f,
	0x41, 0x4b, 0xf9, 0x7e, 0x2c, 0x55, 0xe8, 0xb7, 0x37, 0x94, 0xc4, 0x5f, 0xf3, 0x20, 0x0b, 0xdd,
	0xf0, 0xd5, 0x4d, 0x12, 0xa2, 0x3d, 0x2f, 0xf0, 0x88, 0x4c, 0x42, 0x38, 0xfe, 0x74, 0xb0, 0xc7,
	0x93, 0xd0, 0x03, 0xcd, 0x23, 0xc4, 0xfb, 0x62, 0x2a, 0xbf, 0xa7, 0x38, 0x7e, 0x0a, 0x3b, 0x19,
	0xbd, 0x70, 0xc5, 0x2d, 0xe2, 0x66, 0xe7, 0x12, 0x5c, 0xc9, 0x4b, 0x64, 0xba, 0x3c, 0xfd, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x88, 0xe1, 0x27, 0x2b, 0x60, 0x04, 0x00, 0x00,
}
