#!/bin/bash

# ruleid: container-privileged
docker run --privileged hello-world

# ruleid: container-privileged
podman run --privileged hello-world

# ok: container-privileged
docker run hello-world

# ok: container-privileged
podman run hello-world
