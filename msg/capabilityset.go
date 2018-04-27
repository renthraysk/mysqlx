package msg

import (
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"
	"github.com/renthraysk/mysqlx/slice"

	"github.com/golang/protobuf/proto"
)

const (
	tagCapabilityName  = 1
	tagCapabilityValue = 2
)

var setTLSEnable = []byte{
	0, 0, 0, 0, byte(mysqlx.ClientMessages_CON_CAPABILITIES_SET),
	1<<3 | proto.WireBytes, 17,
	1<<3 | proto.WireBytes, 15,
	tagCapabilityName<<3 | proto.WireBytes, 3, 't', 'l', 's',
	tagCapabilityValue<<3 | proto.WireBytes, 8,
	tagAnyType<<3 | proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
	tagAnyScalar<<3 | proto.WireBytes, 4,
	tagScalarType<<3 | proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_BOOL),
	tagScalarBool<<3 | proto.WireVarint, 1,
}

// NewCapabilitySetTLSEnable returns a Msg to send to mysql to enable TLS
func NewCapabilitySetTLSEnable(buf []byte) MsgBytes {
	_, b := slice.Allocate(buf, len(setTLSEnable))
	copy(b, setTLSEnable)
	return MsgBytes(b)
}
