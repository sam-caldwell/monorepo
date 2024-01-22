virtualization/vmDarwin
=======================

## Description
Create a CGO interface between our virtualization project and Apple's virtualization framework

## Helpful Resources
* [Code-Hex/vz/wiki](https://github.com/Code-Hex/vz/wiki)

## WARNING: You will need this.
1. There is a `vz.entitlements.xml` file (granting both access to the virtualization api and network api.)

2. You must sign any binaries compiled with this project using this xml file as follows:
```shell
  codesign --entitlements vz.entitlements -s - <YOUR BINARY PATH>
```
