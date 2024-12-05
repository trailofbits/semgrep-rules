#!/bin/bash

# ruleid: mysql-insecure-sslmode
mysql mysql://myapplicationuser:mypass@myhost:1234/applicationdb?SslMode=Disabled

# ruleid: mysql-insecure-sslmode
mysql mysql://myapplicationuser:mypass@myhost:1234/applicationdb?useSSL=false

# ok: mysql-insecure-sslmode
mysql mysql://myapplicationuser:mypass@myhost:1234/applicationdb?SslMode=Required
