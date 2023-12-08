package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"strconv"
	"strings"
)

// RequireBoolp - Return the boolean value of an environment variable (returning error if empty).
func RequireBoolp(name *string) (r bool, e error) {
	value := os.Getenv(*name)
	if strings.TrimSpace(value) == "" {
		return defaultBoolValue, fmt.Errorf(errEnvVarNotFound+errors.Details, *name)
	}
	return strconv.ParseBool(value)
}
