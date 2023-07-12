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
	monorepo "github.com/sam-caldwell/go/v2/projects/__system__"
	"strings"
)

func main() {
	fmt.Printf("%s\n", strings.Join(monorepo.GetSupportedLanguages(), ", "))
}
