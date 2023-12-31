package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"strconv"
	"strings"
)

// RequireIntp - Return the int value of a given environment variable (return error if not set)
func RequireIntp(name *string) (r int, e error) {
	value := os.Getenv(*name)
	if strings.TrimSpace(value) == "" {
		return defaultIntValue, fmt.Errorf(errEnvVarNotFound+errors.Details, *name)
	}
	i, e := strconv.ParseInt(value, intBase, intVarSize)
	if e != nil {
		return 0, e
	}
	return convert.Int64ToIntSafe(i)
}
