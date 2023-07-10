package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/convert"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"github.com/sam-caldwell/go/v2/projects/testing/hijack"
	"reflect"
)

func main() {
	const (
		initialValue = "original"
		alteredValue = "altered"
	)

	var actualValue = words.EmptyString
	targetFunc := func() {
		actualValue = initialValue
		fmt.Println("---targetFunc---")
	}
	imposterFunc := func() {
		actualValue = alteredValue
		fmt.Println("---imposterFunc---")
	}
	ansi.Blue().Println("setup").Reset()

	targetFuncPtr := reflect.ValueOf(targetFunc).Pointer()
	ansi.Yellow().Printf("targetFuncPtr:  %00x", targetFuncPtr).LF().Reset()

	imposterFuncPtr := reflect.ValueOf(imposterFunc).Pointer()

	ansi.Blue().Println("ready to hijack").
		Yellow().Printf("> targetFuncPtr:  %00x", targetFuncPtr).LF().
		Yellow().Printf("> imposterFuncPtr:%00x", imposterFuncPtr).LF().Reset()

	originalMemory, err := hijack.TestHijack(targetFuncPtr, imposterFuncPtr)
	ansi.Blue().Println("hijack completed").
		Yellow().Printf("> targetFuncPtr:  %00x", targetFuncPtr).LF().
		Yellow().Printf("> imposterFuncPtr:%00x", imposterFuncPtr).LF().
		Yellow().Printf("> originalMemory: %s", convert.ByteToHexString(originalMemory)).
		Reset()
	if err != nil {
		ansi.Red().
			Printf("hijackFunction returned an error: %v\n", err).
			Fatal(1)
	}
	if actualValue != words.EmptyString {
		ansi.Red().Println("expected actualValue to be empty").Fatal(1)
	}
	if targetFunc(); actualValue != alteredValue {
		ansi.Red().
			Printf("targetFunc should have altered the value. ActualValue:%s\n", actualValue).
			Fatal(1)

	}
	ansi.Green().Println("target is altered").Reset()
	if imposterFunc(); actualValue != alteredValue {
		ansi.Red().
			Printf("imposterFunc should have altered the value. ActualValue:%s\n", actualValue).
			Fatal(1)
	}
}
