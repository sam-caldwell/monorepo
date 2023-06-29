Go Monorepo
===========

## Description

This is a single monorepo for my various golang projects because reinventing wheels for different projects
gets real old, real quick whether here or when I just want to reuse some of my code in a different personal
or work project.

## Status
![Version](https://raw.githubusercontent.com/sam-caldwell/go/main/VERSION.svg)          

|          |                                                                         |
|----------|-------------------------------------------------------------------------|
| Linter   | ![Linter](https://github.com/sam-caldwell/badges/LINT.svg?branch=main)  |
| Security | ![Linter](https://github.com/sam-caldwell/badges/SNYK.svg?branch=main)  |
| Test     | ![Linter](https://github.com/sam-caldwell/badges/TEST.svg?branch=main)  |
| Build    | ![Linter](https://github.com/sam-caldwell/badges/BUILD.svg?branch=main) |





## Structure

1. If we create a binary, it's `main.go` goes [here in cmd/](./cmd). (`Makefile` will automatically build it).
2. If we create reusable code, it goes [here in projects/](./projects)
3. Makefile automation is [here in Makefile.d/](./Makefile.d)
4. Binaries we produce are [here in build/](./build)
5. Documentation is [here in docs/](./docs) (this includes example code for our projects)

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