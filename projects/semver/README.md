Semantic Version Tools
======================

## Description

A git-aware semantic versioning tool built in golang, which 
includes a re-usable GitHub Action. This thing is intended to 
automate more so that I can tag less and drink more coffee.

## Usage

> Each semantic version (`SemanticVersion`) object consists of
> a major, minor and patch version, where each part of the
> larger object is an unsigned 8-bit integer (0-255) and where
> the `SemanticVersion` bump methods will carry-over the overflow
> of patch to minor and minor to major. However, if major goes
> above 255, an error will be thrown.
>
> The version number parts are tracked using the `VersionNumber`
> datatype and this could be changed from a size 8 to size 16 or
> or size 32 relatively quickly with only two changes (i.e. the
> VersionNumber type to uint16 or uint32, etc.) and the constant
> `versionNumberSize` to the appropriate bit size. Tests and code
> should handle the change transparently.

### Installation

### In your go project?

`go get "github.com/sam-caldwell/go/v2/projects/semver"`

```go
var version SemanticVersion
if err:=version.BumpMajor(); err!=nil{
    panic(err)
}
if err:=version.BumpMinor(); err!=nil{
    panic(err)
}
if err:=version.BumpPatch(); err!=nil{
    panic(err)
}
if err:=version.UpdateTag(); err!= nil {
    panic(err)
}
fmt.Println(version.String())
```
