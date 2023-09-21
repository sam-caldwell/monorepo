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
	"strings"
	"testing"
)

func TestArgumentDescriptor_Argument(t *testing.T) {
	const (
		testProgramName        = "testProgram"
		TestProgramDescription = "test program description"
		testCommand            = "testCommand"
		testSubcommand         = "subTest"
		testAction             = "testAction"
		testDescription        = "testDescription"
	)
	arg := NewArgParse(testProgramName, TestProgramDescription).
		Argument(testDescription, testCommand, testSubcommand, testAction)
	/*
	 * Validate the state of the ArgumentDescriptor
	 */
	fullyQualifiedCommand := strings.Join([]string{testCommand, testSubcommand, testAction}, words.Space)

	assert.Equal(t, arg.programName, testProgramName, "test program name mismatch")
	assert.Equal(t, arg.programDescription, TestProgramDescription, "test program description mismatch")
	assert.NotNil(t, arg.command, "ArgumentDescriptor.command should not be nil")
	assert.NotNil(t, arg.helpText, "ArgumentDescriptor.helpText should not be nil")

	assert.Equal(t, arg.lastCommand, fullyQualifiedCommand, "ArgumentDescriptor.lastCommand mismatch")
	assert.NotNil(t, arg.command[arg.lastCommand], "ArgumentDescriptor.command not nil")
	assert.Equal(t, arg.helpText[arg.lastCommand], testDescription, "ArgumentDescriptor.helpText mismatch")

}
