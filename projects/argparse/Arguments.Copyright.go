package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/valid"
	"strings"
)

// Copyright - Append the Copyright string to the preamble.
func (arg *Arguments) Copyright(year int, author string, email string) *Arguments {

	var err error

	yearString, err := valid.IsValidYear(year)
	if err != nil {
		_ = arg.err.Push(err)
	}

	author = strings.TrimSpace(author)

	email = strings.TrimSpace(email)

	return arg.Postscript(
		fmt.Sprintf("(c) %s %s <%s>",
			yearString,
			author,
			email))
}
