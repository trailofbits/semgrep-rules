#!/bin/bash

# ruleid: redis-unencrypted-transport
redis-cli -u redis://user:password@host:port/dbnum PING

# ok: redis-unencrypted-transport
redis-cli -u rediss://user:password@host:port/dbnum PING
