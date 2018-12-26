# mysqlx 

Incomplete.

* DSNs are deliberately not supported. Create a sql.Connector with mysqlx.New() and use go 1.10's sql.OpenDB().
* Beware some protobuf marshalling (StmtExecute, CapabilitiesSet, AuthenticateStart & AuthenticateContinue) and unmarshalling (ColumnMetaData & Row) is hand coded, to reduce allocations.
* TLS is not negotiated, have to explicitly enable using mysqlx.WithTLSConfig() when creating the connector.


## Authentication

Driver can currently authenticate using any of the following authentication mechanisms (mysql_native_password, sha256_password, caching_sha2_password) when using a secure connection (TLS or unix). See authentication_test.go for details.

When not using TLS, sha256 or caching_sha2 does not have any fall back method to populate the cache. So limited to mysql_native_password on insecure connections.

https://dev.mysql.com/doc/refman/8.0/en/caching-sha2-pluggable-authentication.html


## Geometry, JSON & XML parameters.

	j, err := json.Marshal(expected)
    ...
    r, err := db.ExecContext(context.Background(), "INSERT INTO json(json) VALUES(?)", mysqlx.JSON(j))


