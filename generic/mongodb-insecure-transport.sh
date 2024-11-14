#!/bin/bash

# ruleid: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/"

# ruleid: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=true&tlsAllowInvalidCertificates=true"

# ok: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?tls=true"

# ok: mongodb-insecure-transport
mongo "mongodb://user:pass@db0.example.com,db1.example.com,db2.example.com/?ssl=true"
