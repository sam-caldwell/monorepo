package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"strings"
)

// FromString - Given a string, store the equivalent command state or return an error.
func (command *Commands) FromString(c string) (err error) {
	switch strings.ToLower(c) {
	case "create":
		*command = Create
	case "read":
		*command = Read
	case "update":
		*command = Update
	case "delete":
		*command = Delete
	case "list":
		*command = List
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
