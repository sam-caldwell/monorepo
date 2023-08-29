package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/wrappers/os"
	"strings"
)

// RequireStringp - Return the string value of a given environment variable (return error if not set)
func RequireStringp(name *string) (r string, e error) {
	value := os.Getenv(*name)
	if strings.TrimSpace(value) == "" {
		return defaultStringValue, fmt.Errorf(errEnvVarNotFound)
	}
	return value, nil
}
