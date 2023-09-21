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
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestArgumentDescriptor_GlobalFlag(t *testing.T) {
	//t.Skip("Disabled until search() function can pass tests")
	const (
		testProgramName        = "testProgram"
		testProgramDescription = "test program description"
		testCommand            = "test"
		testDescription        = "test command description"
		initialValue           = "initial value"
		newValue               = "new value"
	)

	testValue := initialValue

	fn := func() {
		testValue = newValue
	}
	flags := []string{"-t", "--test"}

	for _, flag := range flags {
		os.Args = append(os.Args, flag)
		arg := NewArgParse(testProgramName, testProgramDescription).
			GlobalFlag(flags, "test global faction", fn).
			Argument(testDescription, testCommand)

		assert.Equal(t, arg.programName, testProgramName, "1.test program name mismatch")
		assert.Equal(t, arg.programDescription, testProgramDescription, "2.test program description mismatch")
		assert.NotNil(t, arg.command, "3.ArgumentDescriptor.command should not be nil")
		assert.NotNil(t, arg.helpText, "4.ArgumentDescriptor.helpText should not be nil")
		assert.Equal(t, arg.lastCommand, testCommand, "5.ArgumentDescriptor.lastCommand mismatch")
		assert.NotNil(t, arg.command[arg.lastCommand], "6.ArgumentDescriptor.command not nil")
		assert.Equal(t, arg.helpText[arg.lastCommand], testDescription, "7.ArgumentDescriptor.helpText mismatch")
		assert.Equal(t, testValue, newValue, "8.Expected global flag to transform the testValue state")
	}
}
