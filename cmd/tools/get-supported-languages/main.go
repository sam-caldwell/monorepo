package main

/*
 * get-supported-languages
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program lists the programming languages currently
 * supported by the monorepo tools
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/__system__"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"strings"
)

func main() {
	exit.IfVersionRequested()
	fmt.Printf("%s\n", strings.Join(monorepo.GetSupportedLanguages(), ", "))
}
