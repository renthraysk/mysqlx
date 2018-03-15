# mysqlx 

Incomplete, and needs refactoring in places.

DSNs are deliberately not supported.
Beware some protobuf serialization and unserialization is hand coded, to reduce allocations.


## Required mysql plugins.

INSTALL PLUGIN 'mysqlx' SONAME 'mysqlx.so';
INSTALL PLUGIN 'mysqlx_cache_cleaner' SONAME 'mysqlx.so';

Driver can currently authenticate using any of the following authentication plugins (mysql_native_password, sha256_password, caching_sha2_password) when using a secure connection (TLS or unix). See authentication_test.go for details.

When not using TLS, sha256 or caching_sha2 does not have any fall back to populate the cache. So practically limited to mysql_native_password on insecure connections.

https://dev.mysql.com/doc/refman/8.0/en/caching-sha2-pluggable-authentication.html
