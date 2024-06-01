package simpleArgs

import "github.com/sam-caldwell/monorepo/go/strconv"

// GetOptionUint16Value - Get a commandline option, expecting an integer value
//
//	(c) 2023 Sam Caldwell.  MIT License
func GetOptionUint16Value(name string, required bool) (value uint16, err error) {
	var tmp string
	if tmp, err = GetOptionValue(name); err != nil {
		if required || (err.Error() != OptionNotFound) {
			return 0, err
		}
		return 0, nil
	}
	return strconv.ParseUInt16(tmp)
}
