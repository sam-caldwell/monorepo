package simpleArgs

import (
	"github.com/sam-caldwell/monorepo/go/strconv"
)

// GetOptionIntValue - Get a commandline option, expecting an integer value
//
//	(c) 2023 Sam Caldwell.  MIT License
func GetOptionIntValue(name string, required bool) (value int, err error) {
	var tmp string
	if tmp, err = GetOptionValue(name); err != nil {
		if required || (err.Error() != OptionNotFound) {
			return 0, err
		}
		return 0, nil
	}
	return strconv.ParseInt(tmp)
}
