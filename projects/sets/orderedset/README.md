orderedSet
==========

## Description

A simple golang ordered set because who wants to reinvent wheels all the time
when you really just want to solve problems?

## Status

[![Go Tests](https://github.com/sam-caldwell/orderedset/actions/workflows/go-tests.yaml/badge.svg)](https://github.com/sam-caldwell/simpleSet/actions/workflows/go-tests.yaml)

## Supported Versions

- 1.18
- 1.19
- 1.20

## Methods

### `.Add(item any) error`

> Add a new item to the set and throw an error if it already exists.

### `.Count() int`

> Return a count of records

### `.Delete(pos int) error`

> Delete an existing

### `.List() []any`

> Return a list of elements `[]any` from the given set

### `.Pop() any`

> Pop element 0 from the set and return its value

### `.Push(item any)`

> Push an element to the top of the set.

## Build / test

> For local builds, there is a `Makefile` but there is also a GitHub action for this
> project.

`make test` - run tests 

