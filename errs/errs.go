package errs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/renthraysk/mysqlx/protobuf/mysqlx"
	"google.golang.org/protobuf/proto"
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

type ErrCode uint32

func (e ErrCode) Is(err error) bool {
	var er *Error

	if errors.As(err, &er) {
		return e == er.Code
	}
	return false
}

// Error represents a mysqlx Error
type Error struct {
	Severity Severity
	Code     ErrCode
	SQLState string
	Msg      string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %d (%s): %s", e.Severity, e.Code, e.SQLState, e.Msg)
}

func New(b []byte) error {
	var e mysqlx.Error

	if err := proto.Unmarshal(b, &e); err != nil {
		return fmt.Errorf("failed to unmarshal mysqlx.Error %x: %w", b, err)
	}
	return &Error{
		Severity: Severity(e.GetSeverity()),
		Code:     ErrCode(e.GetCode()),
		SQLState: e.GetSqlState(),
		Msg:      e.GetMsg(),
	}
}

func IsFatal(err error) bool {
	var e *Error
	ok := errors.As(err, &e)
	return ok && e.Severity == SeverityFatal
}

type Errors map[int]error

func (e Errors) Error() string {
	var s strings.Builder
	for i, err := range e {
		fmt.Fprintf(&s, "* %d: %s\n", i, err)
	}
	return s.String()
}
