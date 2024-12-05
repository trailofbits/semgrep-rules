#!/bin/bash

# ruleid: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/"

# ruleid: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=true&tlsAllowInvalidCertificates=true"

# ruleid: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=true&tlsAllowInvalidCertificates=true&something=else"

# we want to be lenient here
# ruleid: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=true&tlsAllowInvalidCertificates=truebutbug"

# ruleid: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=true&tlsAllowInvalidCertificates=truebutbug&something=else"

# we want to be strict here
# ruleid: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=truebutbug"

# ruleid: mongodb-insecure-transport
mongo 'mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=truebutbug&something=else'

# ok: mongodb-insecure-transport
mongo 'mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=true'

# ok: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=true&something=else"

# ok: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?ssl=true"

echo '
# ok: mongodb-insecure-transport
mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?ssl=true
'

echo '
# ruleid: mongodb-insecure-transport
mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?ssl=truebutbug
'

echo '
# ok: mongodb-insecure-transport
mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?ssl=true&something=else
'

echo '
# ruleid: mongodb-insecure-transport
mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?ssl=truebutbug&something=else
'
