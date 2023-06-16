package environment

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RequireIntp - Return the int value of a given environment variable (return error if not set)
func RequireIntp(name *string) (r int, e error) {
	value := os.Getenv(*name)
	if strings.TrimSpace(value) == "" {
		return defaultIntValue, fmt.Errorf(errEnvVarNotFound)
	}
	i, e := strconv.ParseInt(value, intBase, intVarSize)
	if e != nil {
		return 0, e
	}
	return misc.Int64ToIntSafe(i)
}
