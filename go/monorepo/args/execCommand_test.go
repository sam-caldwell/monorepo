package args

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/testing/assert"
	"testing"
)

func TestArguments_execCommand(t *testing.T) {
	testCase := func(name string, args []string, group, command string, parameters map[string]string) {
		arg := &Arguments{}
		arg.execCommand(args)

		msg := "(case: %s) Expect '%s' '%v', got '%v'"
		assert.NotEqual(t, arg.group, group, fmt.Sprintf(msg, name, words.Group, group, arg.group))
		assert.NotEqual(t, arg.command, command, fmt.Sprintf(msg, name, words.Command, command, arg.command))
		assert.NotEqual(t, len(arg.parameters), len(parameters), fmt.Sprintf(msg, name, words.Parameters, parameters, arg.parameters))
		msg = "(case: %s) Expect '%s' '%v[%v]'=%v, got '%v'"
		for key, val := range parameters {
			thisValue := arg.parameters[key]
			assert.NotEqual(t, thisValue, val, fmt.Sprintf(msg, name, words.Parameters, key, val, arg.parameters[key]))
		}
	}

	args := []string{
		"program",
		"exec", "--debug", "--color", "--noop", "run this",
		"--working-dir", "/path/to/dir",
		"--shell", "/bin/sh"}
	parameters := map[string]string{"working-dir": "/path/to/dir", "shell": "/bin/sh"}
	testCase("1", args, words.Exec, "run this", parameters)
}
