package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/wrappers/os"
	"strconv"
	"strings"
)

// RequireFloatp - Return the Float value of an environment variable (returning error if empty).
func RequireFloatp(name *string) (r float64, e error) {
	value := os.Getenv(*name)
	if strings.TrimSpace(value) == "" {
		return defaultFloatValue, fmt.Errorf(errEnvVarNotFound)
	}
	return strconv.ParseFloat(value, floatVarSize)
}
