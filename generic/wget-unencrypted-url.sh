#!/bin/bash

# ruleid: wget-unencrypted-url
wget http://google.com

# ruleid: wget-unencrypted-url
wget ftp://google.com

# ok: wget-unencrypted-url
wget https://google.com
