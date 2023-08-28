package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/convert"
	"github.com/sam-caldwell/monorepo/v2/projects/go/wrappers/os"
	"strconv"
)

// GetIntp - Return the int value of a given environment variable
func GetIntp(name *string) (result int, err error) {
	var value int64
	value, err = strconv.ParseInt(os.Getenv(*name), intBase, intVarSize)
	if err != nil {
		return defaultIntValue, fmt.Errorf(errReadingValue, err)
	}
	return convert.Int64ToIntSafe(value)
}
