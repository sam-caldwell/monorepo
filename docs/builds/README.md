Build Automation
================

## Description

This document details the build automation for our repo.

## A Quick Summary
This repository has two types of projects:
* `projects/` - reusable source code that can be imported into other projects using our go module.
* `cmd/` - binaries produced from this repo as finished products/releases.

The build automation discussed here, covers that second set of `cmd` projects and how they are 
all built.


## Controls and Tools

These are things you can do on your laptop...
 
* [Build Local Binaries](build-the-projects.md)

* [List the Enabled Projects](list-enabled-build-projects.md)

* [Disable a Build project](disable-a-build-project.md)

* [Enable a Build Project](enable-a-build-project.md)

* [Show the Current Repo Version](show-current-repo-version.md)

There are some things that can only happen in GitHub Actions, such as--

* Bumping and setting the version.  `make version` will not run locally.