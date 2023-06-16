package environment

import (
	"fmt"
	"os"
	"strconv"
)

// GetIntp - Return the int value of a given environment variable
func GetIntp(name *string) (result int, err error) {
	var value int64
	value, err = strconv.ParseInt(os.Getenv(*name), intBase, intVarSize)
	if err != nil {
		err = fmt.Errorf(errReadingValue, err)
	}
	return misc.Int64ToIntSafe(value)
}
