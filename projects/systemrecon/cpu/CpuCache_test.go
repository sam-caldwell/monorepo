package systemrecon

import (
	runcommand "github.com/sam-caldwell/go/v2/projects/RunCommand"
	"testing"
)

func TestCpuCache(t *testing.T) {
	mockExecutor := runcommand.MockCommandExecutor{
		Output: "32768",
		Error:  nil,
	}

	expected := "32:32:32"
	got, gotErr := CpuCache(mockExecutor)
	if gotErr != nil {
		t.Errorf("CpuCache() returned error:%v", gotErr)
	}

	if got != expected {
		t.Errorf("CpuCache() = %v, want %v", got, expected)
	}
}
