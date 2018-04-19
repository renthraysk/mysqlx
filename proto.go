package mysqlx

import "github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"

func ScalarUint(s *mysqlx_datatypes.Scalar) (uint64, bool) {
	if s == nil || s.Type == nil || *s.Type != mysqlx_datatypes.Scalar_V_UINT || s.VUnsignedInt == nil {
		return 0, false
	}
	return *s.VUnsignedInt, true
}

func ScalarString(s *mysqlx_datatypes.Scalar) (string, bool) {
	if s == nil || s.Type == nil || *s.Type != mysqlx_datatypes.Scalar_V_STRING || s.VString == nil {
		return "", false
	}
	//@TODO Collation?
	return string(s.VString.Value), true
}
