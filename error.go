package mysqlx

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
)

// Severity level of an Error
type Severity mysqlx.Error_Severity

// Imported consts from the protobuf to so users not have to import the protobuf
const (
	SeverityError Severity = Severity(mysqlx.Error_ERROR)
	SeverityFatal Severity = Severity(mysqlx.Error_FATAL)
)

func (s Severity) String() string {
	switch s {
	case SeverityError:
		return "Error"
	case SeverityFatal:
		return "Fatal"
	}
	return "Unknown"
}

// Error represents a mysqlx Error
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
		Severity: Severity(e.GetSeverity()),
		Code:     e.GetCode(),
		SQLState: e.GetSqlState(),
		Msg:      e.GetMsg(),
	}
}
