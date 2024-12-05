#!/bin/bash

# ruleid: postgres-insecure-sslmode
psql postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=disable

# ruleid: postgres-insecure-sslmode
psql postgres://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=disable

# ruleid: postgres-insecure-sslmode
psql postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=PREfered

# "This option is deprecated in favor of the sslmode setting."
# ruleid: postgres-insecure-sslmode
psql postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?requiressl=0

# ruleid: postgres-insecure-sslmode
psql "postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?ssl=false"

# ok: postgres-insecure-sslmode
psql postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=require

# ok: postgres-insecure-sslmode
psql "postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=verify-full&something=else"

# ok: postgres-insecure-sslmode
psql "postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=require&something=else"

# ok: postgres-insecure-sslmode
psql 'postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=require&something=else'

echo '
# ok: postgres-insecure-sslmode
postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=require&something=else
'

echo '
# ok: postgres-insecure-sslmode
postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=require
'

# "for compatibility with JDBC connection URIs, instances of ssl=true are translated into sslmode=require."
# ok: postgres-insecure-sslmode
psql "postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?ssl=true"

# ok: postgres-insecure-sslmode
psql "postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?ssl=true&something=else"

# ok: postgres-insecure-sslmode
psql 'postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?requiressl=1'
