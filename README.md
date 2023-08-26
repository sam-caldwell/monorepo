Go Monorepo
===========

## Description

This is a single monorepo for my various golang projects because reinventing wheels for different projects
gets real old, real quick whether here or when I just want to reuse some of my code in a different personal
or work project.

## Status
![Version](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/VERSION.svg)          

|          |                                                                                               |
|----------|-----------------------------------------------------------------------------------------------|
| Linter   | ![Linter](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/LINT.svg?branch=main) |
| Security | ![Snyk](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/SNYK.svg?branch=main)   |
| Test     | ![Test](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/TEST.svg?branch=main)   |
| Build    | ![Build](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/BUILD.svg?branch=main) |





## Structure

1. If we create a binary, it's `main.go` goes [here in cmd/](./cmd). (`Makefile` will automatically build it).
2. If we create reusable code, it goes [here in projects/](./projects)
4. Binaries we produce are [here in build/](./build)
5. Documentation is [here in docs/](./docs) (this includes example code for our projects)

## Commands
This section covers the commands in the go monorepo:
* [Command: Ansi/Color](./cmd/ansi/color/README.md)
* [Command: net/calculateSubnets](./cmd/net/calculate-subnets/README.md)
* [Command: Tools/badge-maker](./cmd/tools/badge-maker/README.md)
* [Command: What](./cmd/tools/what/README.md)

## Projects

* [Ansi Color Printing](projects/go/ansi/README.md)

* [Cli Argument Parser](projects/go/argparse/README.md)

* [Counter](projects/go/counters/README.md)

* [Environment Variable Tools](projects/go/environment/README.md)

* [Sets](projects/go/sets/README.md)
    * [Ordered Set](projects/go/sets/orderedset/README.md)
    * [Simple Set](projects/go/sets/simpleset/README.md)

* [Networking Stuff](projects/go/net/README.md)
    * [Subnet Tools](projects/go/net/subnetting/calculate-subnets/README.md)

## Other Documentation and Stuff

### Build, Test, etc...

* [Install Local Git Hooks](docs/git/hooks.md)
* [Build Automation](docs/builds/README.md)
* [Test Automation](docs/tests/README.md)

### Makefile Tooling

* [Makefiles](Makefile.d)
* [Determine Operating System Using Make](Makefile.d/check/check.os.mk)