package mysqlx

import (
	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/msg"
)

func Null() any { return msg.Null() }

// String wraps a string value with a specific collation for use as an input parameter value
func String(value string, collation collation.Collation) any {
	return &msg.String{Value: value, Collation: collation}
}

// JSON wraps a byte slice value for use as an JSON input parameter value
func JSON[T []byte | string](json T) any {
	return msg.JSON(json)
}

// XML wraps a byte slice value for use as an XML input parameter value
func XML[T []byte | string](xml T) any {
	return msg.XML(xml)
}

// Geometry wraps a byte slice value for use as an Geometry input parameter value
func Geometry[T []byte | string](geom T) any {
	return msg.Geometry(geom)
}
