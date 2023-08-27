Version
=======

## Description
This project is really just a place to store the current project for the entire monorepo.
The constant (`version`) created in the file `version.go` is updated by the `cmd/tools/set-version`
tool which uses the `projects/semver` project to determine and bump the new version as part of our
release automation.

> YOU SHOULD NOT BE HERE.  THIS ENTIRE PROJECT IS FOR AUTOMATION.
