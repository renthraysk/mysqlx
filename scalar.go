package mysqlx

import (
	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"
)

// scalarUint returns the uint64 value of a pb Scalar
func scalarUint(s *mysqlx_datatypes.Scalar) (uint64, bool) {
	if s == nil || s.Type == nil || *s.Type != mysqlx_datatypes.Scalar_V_UINT || s.VUnsignedInt == nil {
		return 0, false
	}
	return *s.VUnsignedInt, true
}

// scalarString returns the string value of a pb Scalar
func scalarString(s *mysqlx_datatypes.Scalar) (string, collation.Collation, bool) {
	if s == nil || s.Type == nil || *s.Type != mysqlx_datatypes.Scalar_V_STRING || s.VString == nil {
		return "", 0, false
	}
	var col collation.Collation

	if s.VString.Collation != nil {
		col = collation.Collation(*s.VString.Collation)
	}
	return string(s.VString.Value), col, true
}

func scalarValue(s *mysqlx_datatypes.Scalar) (any, bool) {
	if s == nil || s.Type == nil {
		return nil, false
	}
	switch *s.Type {
	case mysqlx_datatypes.Scalar_V_SINT:
		return s.GetVSignedInt(), true
	case mysqlx_datatypes.Scalar_V_UINT:
		return s.GetVUnsignedInt(), true
	case mysqlx_datatypes.Scalar_V_STRING:
		return s.GetVString(), true
	case mysqlx_datatypes.Scalar_V_OCTETS:
		return s.GetVOctets(), true
	case mysqlx_datatypes.Scalar_V_BOOL:
		return s.GetVBool(), true
	case mysqlx_datatypes.Scalar_V_NULL:
		return nil, true
	case mysqlx_datatypes.Scalar_V_DOUBLE:
		return s.GetVDouble(), true
	case mysqlx_datatypes.Scalar_V_FLOAT:
		return s.GetVFloat(), true
	}
	return nil, false
}
