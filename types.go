package mysqlx

import (
	"github.com/renthraysk/mysqlx/collation"
	"github.com/renthraysk/mysqlx/msg"
)

// String wraps a string value with a specific collation for use as an input parameter value
func String(value string, collation collation.Collation) interface{} {
	return &msg.String{Value: value, Collation: collation}
}

// JSON wraps a byte slice value for use as an JSON input parameter value
func JSON(json []byte) interface{} {
	return msg.JSON(json)
}

// XML wraps a byte slice value for use as an XML input parameter value
func XML(xml []byte) interface{} {
	return msg.XML(xml)
}

// Geometry wraps a byte slice value for use as an Geometry input parameter value
func Geometry(geom []byte) interface{} {
	return msg.Geometry(geom)
}
