package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"strconv"
	"strings"
)

// RequireFloatp - Return the Float value of an environment variable (returning error if empty).
func RequireFloatp(name *string) (r float64, e error) {
	value := os.Getenv(*name)
	if strings.TrimSpace(value) == "" {
		return defaultFloatValue, fmt.Errorf(errEnvVarNotFound+errors.Details, *name)
	}
	return strconv.ParseFloat(value, floatVarSize)
}
