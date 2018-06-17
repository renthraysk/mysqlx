package msg

import "github.com/renthraysk/mysqlx/protobuf/mysqlx_resultset"

type ContentType uint32

const (
	ContentTypePlain    ContentType = 0
	ContentTypeGeometry             = ContentType(mysqlx_resultset.ContentType_BYTES_GEOMETRY)
	ContentTypeJSON                 = ContentType(mysqlx_resultset.ContentType_BYTES_JSON)
	ContentTypeXML                  = ContentType(mysqlx_resultset.ContentType_BYTES_XML)
)
