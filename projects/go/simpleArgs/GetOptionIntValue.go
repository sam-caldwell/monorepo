package simpleArgs

import (
	"strconv"
)

func GetOptionIntValue(name string, required bool) (value int, err error) {
	var tmp string
	if tmp, err = GetOptionValue(name); err != nil {
		if required || (err.Error() != "option not found") {
			return 0, err
		}
		return 0, nil
	}
	value, err = strconv.Atoi(tmp)
	return value, err
}
