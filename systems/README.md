Systems
=======

## Description
This top-level directory contains systems tooling (virtual machine builders, vagrant box builders, container 
descriptors, etc.).  The goal of this directory is to represent all systems in terms of code with zero manual
configuration for both local build environments and for remote deployment to servers.


## Technologies
* [Docker Containers](docker/README.md)
  > These are the repo base images

* [Packer Image Builders](packer/README.md)
  > These are the packer image builders which create our vagrant boxes for general-purpose use
  > Eventually this will be extended to create immutable machine images for other purposes. 

* [Vagrant Boxes](vagrant/README.md)
  > These are the mission-specific vagrant boxes we use for build/test, etc.
