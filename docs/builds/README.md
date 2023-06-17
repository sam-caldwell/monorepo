Build Automation
================

## Description

This document details the build automation for our repo.

## A quick summary

* We only build projects in `cmd/<project>/<program>` which have `main.go`.
  These programs are the only binaries we produce in this repo though you
  may create other stuff by importing this repo as a Go module and have fun.

* You can **enable/disable/list builds for specific projects** See "Build Enable/Disable"
  (below).

## Controls and Tools

These are things you can do on your laptop...

### List the enabled projects

It's easy to find out what projects are in the `cmd` directory that are **ENABLED**:

```bash
make list/build/projects
```

This will exclude disabled projects, and the output will look like this:

```bash
(venv) samcaldwell@mbp go % make list/build/projects

current binary projects (enabled):
 - cmd/subnetting/calculateSubnets

(venv) samcaldwell@mbp go % 
```

### Disable a Build Project

You can disable builds for a project using the following:

```bash 
make disable/build/program/${YOUR_PROJECT_NAME}
```

* This will create `cmd/<project>/build.disabled` in your local repo.
* You **can** check in a `build.disabled` file and GitHub actions will respect it.
* The `build.disabled` contains the user, host and timestamp of the operation.
  For example:
  ```text
    samcaldwell@mbp.local : Fri Jun 16 21:56:57 CDT 2023
  ```

### Enable a Build Project

Build projects can be enabled (if disabled using the above process) by running
this command:

```bash
make enable/build/program/${YOUR_PROJECT_NAME}
```

(This simply deletes the `build.disabled` file.)
