#!/bin/bash

# ruleid: container-user-root
docker run -u root hello-world

# ruleid: container-user-root
podman run --user root hello-world

# ok: container-user-root
docker run hello-world

# ok: container-user-root
podman run hello-world
