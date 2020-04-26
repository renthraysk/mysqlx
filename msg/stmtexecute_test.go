package msg

import (
	"math"
	"reflect"
	"testing"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx_datatypes"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx_sql"

	"github.com/golang/protobuf/proto"
)

var _ Args = (*StmtExecute)(nil)

func TestSerialization(t *testing.T) {
	var out mysqlx_sql.StmtExecute
	var b [1024]byte

	tests := map[string]struct {
		Stmt string
		Args []interface{}
	}{
		"emptystring": {"SELECT ?", []interface{}{""}},
		"string":      {"SELECT ?", []interface{}{"abc"}},
		"int":         {"SELECT ?", []interface{}{int64(42)}},
		"bool":        {"SELECT ?", []interface{}{true}},
		"float32":     {"SELECT ?", []interface{}{float32(math.Pi)}},
		"float64":     {"SELECT ?", []interface{}{float64(math.Pi)}},
	}

	for name, in := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewStmtExecute(b[:0], in.Stmt)
			AppendArgValues(&s, in.Args...)

			if err := proto.Unmarshal(s[headerSize:], &out); err != nil {
				t.Fatalf("failed to unmarshal: %s", err)
			}
			if string(out.Stmt) != in.Stmt {
				t.Fatalf("failed to unmarshal Stmt expected %s got %s", in.Stmt, out.Stmt)
			}
			if len(out.Args) != len(in.Args) {
				t.Fatalf("failed to unmarshal Args expected %d got %d", len(in.Args), len(out.Args))
			}
			for i, any := range out.Args {
				if v, ok := AnyToInterface(any); !ok || !reflect.DeepEqual(v, in.Args[i]) {
					t.Fatalf("failed to unmarshal Arg %d expected %T(%v) got %T(%v)", i, in.Args[i], in.Args[i], v, v)
				}
			}
		})
	}
}

func ScalarToInterface(s *mysqlx_datatypes.Scalar) (interface{}, bool) {
	if s != nil {
		switch s.GetType() {
		case mysqlx_datatypes.Scalar_V_SINT:
			return s.GetVSignedInt(), s.VSignedInt != nil
		case mysqlx_datatypes.Scalar_V_UINT:
			return s.GetVUnsignedInt(), s.VUnsignedInt != nil
		case mysqlx_datatypes.Scalar_V_FLOAT:
			return s.GetVFloat(), s.VFloat != nil
		case mysqlx_datatypes.Scalar_V_DOUBLE:
			return s.GetVDouble(), s.VDouble != nil
		case mysqlx_datatypes.Scalar_V_BOOL:
			return s.GetVBool(), s.VBool != nil
		case mysqlx_datatypes.Scalar_V_STRING:
			if s.VString != nil {
				return string(s.VString.Value), true
			}
		case mysqlx_datatypes.Scalar_V_NULL:
			return nil, true
		case mysqlx_datatypes.Scalar_V_OCTETS:
			if s.VOctets != nil {
				return s.VOctets.Value, true
			}
		}
	}
	return nil, false
}

func AnyToInterface(a *mysqlx_datatypes.Any) (interface{}, bool) {
	if a != nil {
		switch a.GetType() {
		case mysqlx_datatypes.Any_SCALAR:
			return ScalarToInterface(a.Scalar)
		}
	}
	return nil, false
}
