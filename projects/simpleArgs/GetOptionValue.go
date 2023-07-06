package simpleArgs

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os"
	"strings"
)

// GetOptionValue - Return the value (next arg) if the named option exists.
func GetOptionValue(name string) (value string, err error) {
	if len(os.Args) < 4 {
		return words.EmptyString, err
	}
	argList := os.Args[2:]
	for pos, option := range argList {
		if option == strings.TrimSpace(strings.ToLower(name)) {
			if pos >= len(argList)-2 {
				err = fmt.Errorf("insufficient arguments for %s", name)
				break
			}
			value = argList[pos+1]
			if strings.HasPrefix(value, "--") {
				err = fmt.Errorf("option has no value")
				break
			}
		}
	}
	return value, err
}
