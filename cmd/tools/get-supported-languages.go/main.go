package main

import (
	"fmt"
	monorepo "github.com/sam-caldwell/go/v2/projects/__system__"
	"strings"
)

func main() {
	fmt.Printf("%s\n", strings.Join(monorepo.GetSupportedLanguages(), ", "))
}
