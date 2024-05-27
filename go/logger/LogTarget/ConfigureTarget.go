package LogTarget

import (
	"fmt"
)

// ConfigureTarget - Provide an interface for Target configuration
type ConfigureTarget map[string]string

func (cfg *ConfigureTarget) ExpectOrPanic(name string) string {
	record, ok := (*cfg)[name]
	if !ok {
		panic(fmt.Sprintf("missing %s", name))
	}
	return record
}
