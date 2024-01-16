#!/bin/bash

# ruleid: installer-allow-untrusted
sudo installer -pkg /path/to/package.pkg -target / -allowUntrusted

# ok: installer-allow-untrusted
sudo installer -pkg /path/to/package.pkg -target /
