#!/bin/bash

# ruleid: node-disable-certificate-validation
export NODE_TLS_REJECT_UNAUTHORIZED=0

# ok: node-disable-certificate-validation
export NODE_TLS_REJECT_UNAUTHORIZED=1

# ok: node-disable-certificate-validation
export NODE_TLS_REJECT_UNAUTHORIZED=false

node app.js
