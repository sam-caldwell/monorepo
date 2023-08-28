package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/keyvalue"
)

// CpuInfo - Returns a map of CPU characteristics
func CpuInfo() (info keyvalue.KeyValue, err error) {
	//ToDo: This needs functionality
	return info, fmt.Errorf("not implemented")
}
