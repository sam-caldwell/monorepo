package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
)

// SetAnyp - Set an environment variable of any type
func SetAnyp(name *string, value *any) error {
	return os.Setenv(*name, fmt.Sprintf("%v", *value))
}
