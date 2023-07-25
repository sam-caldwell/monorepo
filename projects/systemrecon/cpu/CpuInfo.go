package systemrecon

import (
	"fmt"
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
)

// CpuInfo - Returns a map of CPU characteristics
func CpuInfo() (info keyvalue.KeyValue, err error) {
	//ToDo: This needs functionality
	return info, fmt.Errorf("not implemented")
}
