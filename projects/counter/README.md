Counter
=======

## Description

A set of counters which helps avoid reinvented wheels.

## Status

[![Go Tests](https://github.com/sam-caldwell/counters/actions/workflows/go-tests.yaml/badge.svg)](https://github.com/sam-caldwell/simpleSet/actions/workflows/go-tests.yaml)

## Usage

### Installation

`go get "github.com/sam-caldwell/counter/v2"`

### Tests

`make test`

### Features

#### `simpleCounter()`

> Just a simple counter with an increment and decrement capability.
> Returns current value then increments the internal state.

```golang
package main

import counters "github.com/sam-caldwell/counter/v2"

var count counters.SimpleCounter
count.Increment() //output: 0
count.Increment() //output: 1
count.Increment() //output: 2
count.Value(0) //output: 2

var countIf counters.ConditionalCounter
countif.Increment(true) //output: 0
countif.Value()          //output: 1
countif.Increment(false) //output: 1
countif.Decrement(true) //output:1

```

`countOk(ok bool)` - returns incremented count if input is `true` or `-1` if false
