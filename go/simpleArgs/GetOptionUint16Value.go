package simpleArgs

import "github.com/sam-caldwell/monorepo/go/strconv"

func GetOptionUint16Value(name string, required bool) (value uint16, err error) {
	var tmp string
	if tmp, err = GetOptionValue(name); err != nil {
		if required || (err.Error() != "option not found") {
			return 0, err
		}
		return 0, nil
	}
	return strconv.ParseUInt16(tmp)
}
