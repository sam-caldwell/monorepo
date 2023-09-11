package args

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/testing/assert"
	"testing"
)

func TestArguments_execCommand(t *testing.T) {
	const command string = "run script"

	testCase := func(name string, args []string, group, thisCommand string, parameters map[string]string) {
		arg := &Arguments{}
		arg.parameters = make(map[string]string)
		arg.execCommand(args)

		msg := "(case: %s) Expect '%s' '%v', got '%v'"
		assert.Equal(t, arg.group, group, fmt.Sprintf(msg, name, words.Group, group, arg.group))
		assert.Equal(t, arg.command, thisCommand, fmt.Sprintf(msg, name, words.Command, thisCommand, arg.command))
		assert.Equal(t, len(arg.parameters), len(parameters), fmt.Sprintf(msg, name, words.Parameters, parameters, arg.parameters))
		msg = "(case: %s) Expect '%s' '%v[%v]'==%v, got '%v'"
		for key, val := range parameters {
			thisValue := arg.parameters[key]
			assert.Equal(t, thisValue, val, fmt.Sprintf(msg, name, thisValue, words.Parameters, key, val, arg.parameters[key]))
		}
	}

	var caseData = []struct {
		name       string
		args       []string
		parameters map[string]string
	}{
		{
			"case 1: no arguments",
			[]string{"program", words.Exec, command},
			map[string]string{},
		},
		{
			"case 2: no parameters, global arguments",
			[]string{"program", words.Exec, "--debug", "--color", "--noop", command, "--working-dir", "/path/to/dir", "--shell", "/bin/sh"},
			map[string]string{"working-dir": "/path/to/dir", "shell": "/bin/sh"},
		},
		{
			"case 2: includes all possible parameters, global arguments",
			[]string{"program", words.Exec, "--debug", "--color", "--noop", command, "--working-dir", "/path/to/dir", "--shell", "/bin/sh"},
			map[string]string{"working-dir": "/path/to/dir", "shell": "/bin/sh"},
		},
		{
			"case 2: includes all possible parameters, global arguments",
			[]string{"program", words.Exec, "--color", "--debug", "--noop", command, "--working-dir", "/path/to/dir", "--shell", "/bin/sh"},
			map[string]string{"working-dir": "/path/to/dir", "shell": "/bin/sh"},
		},
		{
			"case 2: includes all possible parameters, global arguments",
			[]string{"program", words.Exec, "--color", "--debug", "--noop", command, "--shell", "/bin/sh", "--working-dir", "/path/to/dir"},
			map[string]string{"working-dir": "/path/to/dir", "shell": "/bin/sh"},
		},
	}
	for _, testData := range caseData {
		testCase(testData.name, testData.args, words.Exec, command, testData.parameters)
	}

}
