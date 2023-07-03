Packer Builder
==============

## Description

This project contains hashicorp packer build manifests
which will generate vagrant boxes and other virtual
machine images from a uniform configuration-as-code source.

These vagrant boxes and virtual machines are intended as--

1. Base immutable machines
2. Test machines for automated testing
3. Cloud-deployable machines
4. Virtual hosts for deployment to local hosting

## Usage

### To Build an image

`make build/packer`

