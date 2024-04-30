package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"strings"
)

// FromString - Given a string, store the equivalent object state or return an error.
func (objects *Objects) FromString(c string) (err error) {
	switch strings.ToLower(c) {
	case "ticket":
		*objects = Ticket
	case "project":
		*objects = Project
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
