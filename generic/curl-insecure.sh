#!/bin/bash

# ruleid: curl-insecure
curl -k https://google.com > /dev/null

# ruleid: curl-insecure
curl --insecure https://google.com > /dev/null

# ok: curl-insecure
curl --ksomeotherflag https://google.com > /dev/null

# ok: curl-insecure
curl https://google.com > /dev/null
