dockerImages
============

## Objective
* Develop a set of common docker images with CI/CD pipeline for maintenance.
* All images are built from the local `opsys/ubuntu:24.04` base — no third-party images pulled from the internet.
* Language containers use multi-stage builds to produce separate **builder** and **runtime** images.
* Runtime images use a minimal distroless base to minimize attack surface and image size.

## Architecture

```
opsys/
  ubuntu/24.04    -> opsys/ubuntu:24.04 (opsys/ubuntu:latest)
  distroless      -> opsys/distroless:latest (built from ubuntu:24.04, scratch-based)

language/
  java/21         -> language/java:21-builder  (Ubuntu + OpenJDK 21 JDK)
                  -> language/java:21-runtime  (distroless + minimal JRE via jlink)

  node            -> language/node:builder     (Ubuntu + nvm + Node 20, 21, latest + npm + yarn)
                  -> language/node:runtime     (distroless + Node)

  python          -> language/python:builder   (Ubuntu + Python 3.14 + latest from source + pip3 + venv)
                  -> language/python:runtime   (distroless + Python)

services/
  ceph            -> service containers built on top of base images
  dns-server
  postgresql
```

## Build order
Images must be built in dependency order:
1. `opsys/ubuntu:24.04` and `opsys/distroless:latest`
2. `language/*` (depends on both opsys images)
3. `services/*` (depends on language or opsys images)

## Build instructions
```text
make build/docker
```

## Clean images
```text
make clean/docker
```
