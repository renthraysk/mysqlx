
path = github.com/renthraysk/mysqlx/protobuf
protos := $(wildcard *.proto)

# This is fragile
map := Mmysqlx.proto=$(path)/mysqlx,Mmysqlx_datatypes.proto=$(path)/mysqlx_datatypes,Mmysqlx_expr.proto=$(path)/mysqlx_expr,Mmysqlx_prepare.proto=$(path)/mysqlx_prepare,Mmysqlx_crud.proto=$(path)/mysqlx_crud,Mmysqlx_sql.proto=$(path)/mysqlx_sql

%.pb.go: %.proto
	mkdir -p $(basename $<)
	protoc --go_out=import_path=$(basename $<),$(map):$(basename $<) $< 

build: $(protos:.proto=.pb.go)

clean:
	rm -rf $(protos:.proto=)
