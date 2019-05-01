package msg

import (
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_connection"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"

	"github.com/golang/protobuf/proto"
)

// CapabilitySetTLSEnable returns a Msg to send to mysql to enable TLS
func CapabilitySetTLS(buf []byte, enable bool) (MsgBytes, error) {
	cs := &mysqlx_connection.CapabilitiesSet{
		Capabilities: &mysqlx_connection.Capabilities{
			Capabilities: []*mysqlx_connection.Capability{
				{
					Name: proto.String("tls"),
					Value: &mysqlx_datatypes.Any{
						Type: mysqlx_datatypes.Any_SCALAR.Enum(),
						Scalar: &mysqlx_datatypes.Scalar{
							Type:  mysqlx_datatypes.Scalar_V_BOOL.Enum(),
							VBool: &enable,
						},
					},
				},
			},
		},
	}

	buf[4] = byte(mysqlx.ClientMessages_CON_CAPABILITIES_SET)
	b := proto.NewBuffer(buf[:5])
	if err := b.Marshal(cs); err != nil {
		return nil, err
	}
	return MsgBytes(b.Bytes()), nil
}

func CapabilitySetSessionConnectAttrs(buf []byte, attrs map[string]string) (MsgBytes, error) {
	var (
		anyScalar    = mysqlx_datatypes.Any_SCALAR.Enum()
		scalarString = mysqlx_datatypes.Scalar_V_STRING.Enum()
	)

	fld := make([]*mysqlx_datatypes.Object_ObjectField, 0, len(attrs))
	for k, v := range attrs {
		fld = append(fld, &mysqlx_datatypes.Object_ObjectField{
			Key: proto.String(k),
			Value: &mysqlx_datatypes.Any{
				Type: anyScalar,
				Scalar: &mysqlx_datatypes.Scalar{
					Type: scalarString,
					VString: &mysqlx_datatypes.Scalar_String{
						Value: []byte(v),
					},
				},
			},
		})
	}

	cs := &mysqlx_connection.CapabilitiesSet{
		Capabilities: &mysqlx_connection.Capabilities{
			Capabilities: []*mysqlx_connection.Capability{
				{
					Name: proto.String("session_connect_attrs"),
					Value: &mysqlx_datatypes.Any{
						Type: mysqlx_datatypes.Any_OBJECT.Enum(),
						Obj: &mysqlx_datatypes.Object{
							Fld: fld,
						},
					},
				},
			},
		},
	}

	buf[4] = byte(mysqlx.ClientMessages_CON_CAPABILITIES_SET)
	b := proto.NewBuffer(buf[:5])
	if err := b.Marshal(cs); err != nil {
		return nil, err
	}
	return MsgBytes(b.Bytes()), nil
}
