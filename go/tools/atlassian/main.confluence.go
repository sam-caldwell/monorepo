package main

import "github.com/sam-caldwell/monorepo/go/ansi"

func main() {
	ansi.Red().Println("not implemented").Fatal(1).Reset()
}
