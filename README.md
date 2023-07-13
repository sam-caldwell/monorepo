Go Monorepo
===========

## Description

This is a monorepo for my various projects because reinventing wheels for different projects
gets real old. This repository is a go mod project but the plan is to support
multiple languages from Assembly and C to Go and Typescript with a common tool
chain to start projects fast and get things done.

---

## Status

![Version](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/VERSION.svg)

|          |                                                                                               |
|----------|-----------------------------------------------------------------------------------------------|
| Linter   | ![Linter](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/LINT.svg?branch=main) |
| Security | ![Snyk](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/SNYK.svg?branch=main)   |
| Test     | ![Test](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/TEST.svg?branch=main)   |
| Build    | ![Build](https://raw.githubusercontent.com/sam-caldwell/go/main/badges/BUILD.svg?branch=main) |

---

## Structure

There are two major project types in this repository:

### Command Projects

A command project exists in `cmd/` and represents a project which can be built to
produce a binary executable or other artifact.

### Reusable Code Projects

A project which simply creates reusable source code libraries (such as code in `go`
you can import from different remote projects) would be located in `projects/`. This
is code that is intended to be reusable, mission specific and testable.

### Documentation

Most documentation should be close to the code relevant to the information. But certain
general documents can be found in [./docs](./docs/README.md)

---

## Projects

## Command Projects

This monorepo creates several tools. Some tools are for maintaining and interacting with
the monorepo. Others are tools built for other purposes. These projects can be found in
the [cmd/](./cmd/) directory.

### General Commands

| category             | command                                                                  | description                                 |
|----------------------|--------------------------------------------------------------------------|---------------------------------------------|
| ***General***        |                                                                          |                                             | 
|                      | [color](./cmd/ansi/color/README.md)                                      | print color to the terminal without effort. |
|                      | [hasExecutable](./cmd/tools/hasExecutable/README.md)                     | Determine if an executable is on the system |
|                      | [locks](./cmd/tools/locks/README.md)                                     | Create/Manage locks on the file system      |
|                      | [what](./cmd/tools/what/README.md)                                       | Collect intelligence on a target system     |
|                      |                                                                          |                                             |
| ***Networking***     |                                                                          |                                             |
|                      | [calculateSubnets](./cmd/subnetting/calculateSubnets/README.md)          | easily calculate subnets                    |
|                      |                                                                          |                                             |
| ***Monorepo Tools*** |                                                                          |                                             |
|                      | [badge-maker](./cmd/tools/badge-maker/README.md)                         | Create SVG status badges                    |
|                      | [bump-version](./cmd/tools/bump-version/README.md)                       | Bump the repo version                       |
|                      | [create-project](./cmd/tools/create-project/README.md)                   | Create a new project skeleton               |
|                      | [get-supported-languages](./cmd/tools/get-supported-languages/README.md) | Print the list of supported languages       |
|                      | [get-supported-platforms](./cmd/tools/get-supported-platforms/README.md) | Print the list of supported platforms       |
|                      | [list-projects](./cmd/tools/list-projects/README.md)                     | List the projects in the monorepo           |
|                      | [set-version](./cmd/tools/set-version/README.md)                         | Set the current repo version                |

## Reusable-Code Projects

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
