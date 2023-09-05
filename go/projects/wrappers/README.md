Wrappers
========

## Description

This path creates wrapper functionality around some of Golang's parts (and other external systems) to give this
repo abstraction for testing.

## Discussion

Many of the opinionated Golang patterns come from Google and are dictated by Google's needs. Unfortunately, the rest
of us live in a different reality and our needs are a bit less constrained than they are in the halls of the Mountain
View lair where Golang decisions are made. One significant problem is in testing. We can't all create interfaces
for all the things to mock up all the things as we work the way Googlers apparently can, given their opinions on the
subject.

This repo attempts to isolate itself further from the opinion-cult and create some abstraction to allow better testing
of code and to allow a pivot point when the language or external dependencies make decisions that would otherwise cost
in terms of productivity.

## Usage

A "wrapper" in the context of this project is a simple package level function pointer variable. For example, in the
repo's wrapper of `os.Exit()` (found in `projects/wrappers/os`) we see this--

```golang
package os

import (
	"github.com/sam-caldwell/monorepo//projects/wrappers/os"
)

var Exit = os.Exit
var Stdout = os.Stdout
```

This allows the monorepo to begin using the `wrappers/os` package instead of golang's native `os` package, creating
an abstraction layer for purposes of testing. Now, our software can use `os.Exit()` from `wrappers` and its tests can
use `init()` to override `os.Exit` with some mock function that contains test functionality. For example:

```golang
package foo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo//projects/wrappers/os"
	"testing"
)

func init() {
	os.Exit = func(n int) {
		fmt.Print("I'm sorry, Dave.  But I cannot let you do that.")
	}
}

func TestExit(t *testing.T) {
	os.Exit(1)           //wrappers will use our current os.Exit()
	os.ResetOsExitWrapper() //Reset our function pointers
	os.Exit(42)          // This will exit with an exit code 42
}
```
For each wrapper