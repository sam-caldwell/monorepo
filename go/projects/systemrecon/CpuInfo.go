package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/keyvalue"
)

// CpuInfo - Returns a map of CPU characteristics
func CpuInfo() (info keyvalue.KeyValue, err error) {
	//ToDo: This needs functionality
	return info, fmt.Errorf("not implemented")
}
