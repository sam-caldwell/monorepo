package runcommand

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
	"testing"
)

// Execute - Intercept a command execution and instead return predefined response content
func (m SimpleMockCommandExecutor) Execute(name string, arg ...string) ([]byte, error) {
	//output := m.Outputs[arg[2]]
	//err := m.Errors[arg[2]]
	return m.ProcessFunc(&m, name, arg...)
}

func (m SimpleMockCommandExecutor) Validate(expectedNumberArgs int, name string, arg ...string) {
	var t *testing.T
	if strings.TrimSpace(name) == words.EmptyString {
		t.Fatalf("empty command (name) argument passed to SimpleMockCommandExecutor.Executor()")
	}
	if len(arg) != expectedNumberArgs {
		t.Fatalf("expected number of arguments (%d) does not match actual (%d)", expectedNumberArgs, len(arg))
	}
	for i, a := range arg {
		if strings.TrimSpace(name) == words.EmptyString {
			t.Fatalf("argument cannot be empty: %d [%v]", i, a)
		}
	}
}
