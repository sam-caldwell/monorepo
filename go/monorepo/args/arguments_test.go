package args

import (
	"github.com/sam-caldwell/monorepo/go/testing/assert"
	"testing"
)

func TestArguments_Struct(t *testing.T) {
	var arg Arguments
	assert.False(t, arg.option.debug, "Expect debug false initially")
	assert.False(t, arg.option.color, "Expect color false initially")
	assert.False(t, arg.option.noop, "Expect noop false initially")

	assert.Equal(t, arg.group, "", "expect arg.group to be empty string")
	assert.Equal(t, arg.command, "", "expect arg.command to be empty string")
	if arg.parameters != nil {
		t.Fatal("arg.parameters should be nil")
	}
	arg.parameters = make(map[string]string)
	if arg.parameters == nil {
		t.Fatal("arg.parameters should initialize as map[string]string")
	}
}
