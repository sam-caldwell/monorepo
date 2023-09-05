package environment

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/convert"
	"github.com/sam-caldwell/monorepo/go/projects/wrappers/os"
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
