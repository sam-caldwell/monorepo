package environment

import (
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
)

// GetStringp - Return the int value of a given environment variable
func GetStringp(name *string) (r string, e error) {
	return os.Getenv(*name), nil
}
