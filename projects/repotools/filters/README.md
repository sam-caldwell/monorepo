Project:repotools.filters
=========================

## Description

This project creates a system of filters which will allow a user to filter
a list of projects given a filter and attributes from the project `MANIFEST.yaml`.

## Theory of Operation

The `projectfilters` package provides two (2) levels of functionality:

* A CLI args processor that will scan os.Args for the flags and return a `Filter` struct
  containing the filter settings derived from these CLI arguments.


* A filter-decision method `Filter.Apply(manifest)` which can take the filter state and
  compare it to the contents of a project's manifest file and returning a boolean
  filter/no-filter decision.

### Implementation:

* This means we need a `struct` to contain our filter state, called `Filter`
* We need a `FromCliArgs()` method to scan for flags and set our internal state
* We need an `Apply(manifest)`method to evaluate a project manifest against our filter state.

### Method: `Apply()`

This method takes a loaded project `Manifest` and its own internal filter state and evaluates
if each element of the filter and manifest agree that the record should be shown. If any of
the record-pairs return true, the record is shown.

### Method: `FromCliArgs()`

This method scans the `os.Args` list for any filter flags and when found, validates them and
stores their value in the `Filter` internal state for use by other features of the `Filter`.

#### filter flags:
```
  -commands
  -buildEnabled
  -lintEnabled
  -scanEnabled
  -signEnabled
  -packEnabled
  -os <opsys>
  -arch <arch>
```
