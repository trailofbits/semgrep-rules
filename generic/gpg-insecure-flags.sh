#!/bin/bash

# ruleid: gpg-insecure-flags
gpg --skip-verify --output doc --decrypt doc.gpg

# ok: gpg-insecure-flags
gpg --output doc --decrypt doc.gpg
