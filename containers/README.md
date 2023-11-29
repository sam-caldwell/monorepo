dockerImages
============

## Objective
* Develop a set of common docker images with CI/CD pipeline for maintenance.

## Architecture
* The base layer is the operating system located in `./docker/opsys/<osname>`.
* The language layer builds from the base layer to add a specific language/platform.
* User-defined containers extend from the language layer most often but could 
  directly build from the language layer.

## Build instructions
```text
make build/docker
```

## Clean images
```text
make clean/docker
```
