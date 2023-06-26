//go:build darwin
// +build darwin

package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
)

// This will simulate the actual commandline call to sysctl -n <arg>
func mockTestFunc(m *runcommand.MockCommandExecutor,
	name string, arg ...string) (output []byte, err error) {

	m.Validate(2, name, arg...)
	output = []byte(m.Outputs[arg[1]])
	err = m.Errors[arg[1]]
	return output, err
}

var testCases = map[string]runcommand.TestData{
	"Happy path": {
		ExpectedResult: "64:256:1024",
		ExpectedError:  nil,
		Executor: runcommand.MockCommandExecutor{
			ProcessFunc: mockTestFunc,
			Outputs: map[string]string{
				"hw.l1icachesize": "65536",   //64K
				"hw.l2cachesize":  "262144",  //256K
				"hw.l3cachesize":  "1048576", //1024K
			},
			Errors: map[string]error{
				"hw.l1icachesize": nil,
				"hw.l2cachesize":  nil,
				"hw.l3cachesize":  nil,
			},
		},
	},
	"Sad path": {
		ExpectedResult: "64:-1:",
		ExpectedError:  fmt.Errorf("command Execution Error"),
		Executor: runcommand.MockCommandExecutor{
			ProcessFunc: mockTestFunc,
			Outputs: map[string]string{
				"hw.l1icachesize": "65536",   //64K
				"hw.l2cachesize":  "262144",  //256K
				"hw.l3cachesize":  "1048576", //1024K
			},
			Errors: map[string]error{
				"hw.l1icachesize": nil,
				"hw.l2cachesize":  fmt.Errorf("command Execution Error"),
				"hw.l3cachesize":  nil,
			},
		},
	},
}
