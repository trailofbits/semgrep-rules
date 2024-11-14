#!/bin/bash

# ruleid: node-disable-certificate-validation
export NODE_TLS_REJECT_UNAUTHORIZED=0

node app.js
