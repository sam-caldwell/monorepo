package environment

import (
	"github.com/sam-caldwell/go/v2/projects/go/wrappers/os"
	"strconv"
)

// GetBoolp - Return the boolean value of an environment variable.
func GetBoolp(name *string) (r bool, e error) {
	return strconv.ParseBool(os.Getenv(*name))
}
