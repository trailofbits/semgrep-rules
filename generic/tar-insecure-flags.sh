#!/bin/bash

# ruleid: tar-insecure-flags
tar -xvf --absolute-paths archive.tar

# ruleid: tar-insecure-flags
tar -xvf -P archive.tar

# ok: tar-insecure-flags
tar -xvf --Psomeotherflag archive.tar

# ok: tar-insecure-flags
tar -xvf archive.tar
