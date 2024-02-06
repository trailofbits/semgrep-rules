#!/bin/bash

# ruleid: tar-insecure-flags
tar -xvf --absolute-paths archive.tar

# ruleid: tar-insecure-flags
tar -xvf -P archive.tar

# ok: tar-insecure-flags
tar -xvf --Psomeotherflag archive.tar

# ok: tar-insecure-flags
tar -xvf archive.tar

# Unbounded ellipsis span in pattern wraps to next command in some scenarios
# ok: tar-insecure-flags
wget https://git.kernel.org/torvalds/t/linux-6.8-rc1.tar.gz -O - | tar -xz -C / && mv linux-6.8-rc1 linux

# This shouldn't match, but the ellipsis wrap picks up the p flag in this command and associates it with the previous command
wget https://github.com/trailofbits/semgrep-rules -P /ToB/
