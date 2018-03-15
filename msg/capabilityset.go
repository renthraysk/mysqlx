package msg

import (
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"

	"github.com/golang/protobuf/proto"
)

const (
	tagCapabilityName  = 1
	tagCapabilityValue = 2
)

func NewCapabilitySetTLSEnable(buf []byte) Msg {
	buf = append(buf, 0, 0, 0, 0, byte(mysqlx.ClientMessages_CON_CAPABILITIES_SET),
		1<<3|proto.WireBytes, 17,
		1<<3|proto.WireBytes, 15,
		tagCapabilityName<<3|proto.WireBytes, 3, 't', 'l', 's',
		tagCapabilityValue<<3|proto.WireBytes, 8,
		tagAnyType<<3|proto.WireVarint, byte(mysqlx_datatypes.Any_SCALAR),
		tagAnyScalar<<3|proto.WireBytes, 4,
		tagScalarType<<3|proto.WireVarint, byte(mysqlx_datatypes.Scalar_V_BOOL),
		tagScalarBool<<3|proto.WireVarint, 1)
	return msg(buf[len(buf)-24:])
}
