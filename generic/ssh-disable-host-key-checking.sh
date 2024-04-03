#!/bin/bash

# ruleid: ssh-disable-host-key-checking
ssh -o StrictHostKeyChecking=no user@hostname

# ok: ssh-disable-host-key-checking
ssh user@hostname
