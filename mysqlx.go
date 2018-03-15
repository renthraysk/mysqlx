package mysqlx

import (
	"fmt"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Severity = mysqlx.Error_Severity

const (
	SeverityError Severity = mysqlx.Error_ERROR
	SeverityFatal Severity = mysqlx.Error_FATAL
)

type Error struct {
	Severity Severity
	Code     uint32
	SQLState string
	Msg      string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %d (%s): %s", e.Severity, e.Code, e.SQLState, e.Msg)
}

func newError(b []byte) error {
	var e mysqlx.Error

	if err := proto.Unmarshal(b, &e); err != nil {
		return errors.Wrapf(err, "failed to unmarshal mysqlx.Error %x", b)
	}
	return &Error{
		Severity: e.GetSeverity(),
		Code:     e.GetCode(),
		SQLState: e.GetSqlState(),
		Msg:      e.GetMsg(),
	}
}
