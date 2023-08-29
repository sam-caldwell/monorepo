Command: bump-version
=====================

## Description

A git-aware semantic versioning command which uses the semver
project to bump the version tag (in git) for a repository.

## Usage
```
bump-version [-major] [-minor] [-patch] [-updateTag]

-The tool will query the local (.) git repo for the most recent tag

-If there are no tags on the local repo at (.), v0.0.0 is assumed
 as the starting value.

  -major
        Bump major version
  -minor
        Bump minor version
  -patch
        Bump patch version
  -updateTag
        Update the git tag
  -v    verbose output
  -verbose
        verbose output
```
