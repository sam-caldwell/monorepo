RepoTool/Manifest
=================

## Description

Each project in the monorepo must have a `MANIFEST.yaml` descriptor file
to inform the `monorepo` command and its dependency `repotool` how to treat
the project. This package (`projectmanifest`) facilitates the read and write
of that `MANIFEST.yaml` file.

## Example `MANIFEST.yaml` file

```yaml
---
#
# This is an example Manifest.yaml file
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
toolchain:
  language: go

  lint: enabled
  scan: enabled
  build: disabled
  pack: disabled
  sign: disabled

supported:
  - opsys: darwin
    arch: amd64
  - opsys: darwin
    arch: arm64
  - opsys: linux
    arch: amd64
  - opsys: linux
    arch: arm64
  - opsys: windows
    arch: amd64
```

## Usage: Creating a Manifest

The `projectmanifest` project can be implemented as a chain of methods, as follows:

```golang
err := projectmanifest.
CreateManifest(manifestFileName).
SetName(projectName).
SetAuthor(*author).
SetLanguage("go").
EnableLint().
DisableBuild().
DisablePack().
DisableScan().
DisableSign().
AddPlatform(runtime.GOOS, runtime.GOARCH).
AddPlatform("linux", "amd64").
AddPlatform("windows", "amd64").
AddDependency("ansi/color").
AddDependency("exit").
AddDependency("misc").
WriteFile().
Error()
```

In the above example, we see that `CreateManifest(<filename>)` is called first to
Create the `Manifest` object and record the intended filename. Any state (including)
error state will be passed through this chain of methods until it is handled by `.Error()`
in the end.

Each method in the above example chains together to add/update/manipulate the state of
the object store. Only `.WriteFile()` will persist this state to disk as the `MANIFEST.yaml`
file defined when `CreateManifest()` was first called.

At the end of the chain we call `.Error()` to return our error state to the `err` variable
provided by the caller for error handling. One should note that `.Error()` must come at the
end as it does not allow chaining otherwise.

## Usage: Loading and Accessing a Manifest

We can load a manifest using...

```golang
var manifest projectmanifest.Manifest

err := manifest.LoadFile(<filename>)
```

This command will load the manifest (yaml) content into the `Manifest` struct.
From there, data can be accessed using the following commands:

* `GetName() string` - return the project name
* `GetAuthor() string` - return the project author
* `GetLanguage() string` - return the project language
* `IsBuildEnabled() bool` - return if Build is enabled
* `IsLintEnabled() bool` - return if Lint is enabled
* `IsPackEnabled() bool` - return if Pack is enabled
* `IsScanEnabled() bool` - return if Scan is enabled
* `IsSignEnabled() bool` - return if Sign is enabled
