#!/bin/bash

# ruleid: openssl-insecure-flags
openssl genpkey -algorithm RSA -out private_key.pem -aes-256-cbc -pass pass:mysecretpass

# ok: openssl-insecure-flags
openssl genpkey -algorithm RSA -out private_key.pem --noencsomeotherflag

# ok: openssl-insecure-flags
openssl genpkey -algorithm RSA -out private_key.pem -aes-256-cbc -pass env:PASSVAR
