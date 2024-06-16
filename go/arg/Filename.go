package arg

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

type ExistsOrNot bool

const (
	Exists    ExistsOrNot = true
	NotExists ExistsOrNot = false
)

// Filename - get a filename from the command-line and verify if it exists if requireExists is true
//
//	(c) 2023 Sam Caldwell.  MIT License
func Filename(name, defaultValue, usage string, requireExists ExistsOrNot) (*string, error) {
	value := flag.String(name, defaultValue, usage)
	if bool(requireExists) && !file.Exists(*value) {
		return nil, fmt.Errorf(errors.NotFound+errors.Details, name+":"+*value)
	}
	return value, nil
}
