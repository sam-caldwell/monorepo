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
3. Makefile automation is [here in Makefile.d/](./Makefile.d)
4. Binaries we produce are [here in build/](./build)
5. Documentation is [here in docs/](./docs) (this includes example code for our projects)

## Commands
This section covers the commands in the go monorepo:
* [Command: Ansi/Color](./cmd/ansi/color/README.md)
* [Command: subnetting/calculateSubnets](./cmd/subnetting/calculateSubnets/README.md)
* [Command: Tools/badge-maker](./cmd/tools/badge-maker/README.md)
* [Command: Tools/findBuildProjects](./cmd/tools/findBuildProjects/README.md)
* [Command: Tools/hasExecutable](cmd/tools/has-executable/README.md)
* [Command: Tools/Locks](./cmd/tools/locks/README.md)
* [Command: What](./cmd/tools/what/README.md)

## Projects

* [Ansi Color Printing](./projects/ansi/README.md)

* [Cli Argument Parser](./projects/argparse/README.md)

* [Counter](./projects/counter/README.md)

* [Environment Variable Tools](./projects/environment/README.md)

* [Sets](./projects/sets/README.md)
    * [Ordered Set](./projects/sets/orderedset/README.md)
    * [Simple Set](./projects/sets/simpleset/README.md)

* [Networking Stuff](./projects/net/README.md)
    * [Subnetting Tools](./projects/net/subnetting/README.md)

## Other Documentation and Stuff

### Build, Test, etc...

* [Install Local Git Hooks](docs/git/hooks.md)
* [Build Automation](docs/builds/README.md)
* [Test Automation](docs/tests/README.md)

### Makefile Tooling

* [Makefiles](Makefile.d)
* [Determine Operating System Using Make](Makefile.d/check/check.os.mk)