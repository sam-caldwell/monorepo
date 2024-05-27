package convert

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// ErrorToString - Safely turn an error into a string
//
//	(c) 2023 Sam Caldwell.  MIT License
func ErrorToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return words.EmptyString
}
