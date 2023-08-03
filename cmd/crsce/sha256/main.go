package main

import (
	"fmt"
	crypto "github.com/sam-caldwell/go/v2/projects/crypto/sha256stream"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"os"
)

const (
	programName  = "sha256"
	commandUsage = `
sha256 -h|-help
	show this screen

sha256 string
	Hash the contents of the args[]
`
)

func main() {
	exit.OnCondition(len(os.Args) <= 1, exit.MissingArg, "Missing inputs", commandUsage)
	hashStream := crypto.NewSha256Stream()

	input := os.Args[1]
	for _, c := range input {
		for i := 7; i >= 0; i-- {
			bit := ((c >> i) & 1) == 1
			hashStream.AddBit(bit)
		}
	}
	fmt.Println(hashStream.String())
}
