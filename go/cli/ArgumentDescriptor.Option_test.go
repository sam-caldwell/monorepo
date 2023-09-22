package cli

/*
 * Copyright © 2023 Sam Caldwell <mail@samcaldwell.net>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/stretchr/testify/assert"
	"go/types"
	"testing"
)

func TestArgumentDescriptor_Option_WithString(t *testing.T) {
	const (
		programName        = "testProgram"
		programDescription = "this is a test"
		shortOption        = "-t"
		longOption         = "--test"
		testDescription    = "this is a test option"
		testType           = types.String
		testValue          = "testValue"
	)
	arg := NewArgParse(programName, programDescription).
		Argument("validate container runtime config", words.Config, words.Cri, words.Check).
		NoValue().
		Option([]string{shortOption, longOption}, testDescription, testType, Required, testValue).
		Done(NoAction)

	assert.NotNil(t, arg.command[arg.lastCommand], "arg.command should not be nil")
	assert.NotNil(t, arg.command[arg.lastCommand].value, "arg.command.value should not be nil")

}
