#!/bin/bash

# ruleid: wget-no-check-certificate
wget --no-check-certificate https://google.com

# ruleid: wget-no-check-certificate
wget --no-hsts https://google.com

# ok: wget-no-check-certificate
wget https://google.com
