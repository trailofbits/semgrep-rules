#!/bin/bash

# ruleid: curl-unencrypted-url
curl http://google.com > /dev/null

# ruleid: curl-unencrypted-url
curl ftp://google.com > /dev/null

# ok: curl-unencrypted-url
curl https://google.com > /dev/null

# ok: curl-unencrypted-url
curl http://localhost > /dev/null

# ok: curl-unencrypted-url
curl http://127.0.0.1 > /dev/null