#!/bin/bash

# ruleid: postgres-insecure-sslmode
psql postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=disable

# ok: postgres-insecure-sslmode
psql postgresql://myapplicationuser:mypass@myhost:1234/applicationdb?sslmode=require
