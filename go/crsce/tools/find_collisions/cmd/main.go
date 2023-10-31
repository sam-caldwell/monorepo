package main

/*
 * FindCollisions Command
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a command to scan the keyspace of a given size
 * across a given number of concurrent workers.
 */

import (
	"flag"
	"fmt"
	hashScanner "github.com/sam-caldwell/monorepo/go/crsce/tools/find_collisions/lib"
	"os"
	"runtime"
)

const commandUsage = `
	FindCollisions <keySpaceSize>
	FindCollisions -h | --help
`

func main() {
	keySpaceSize := flag.Uint("keySpaceSize", 1024,
		"Number of bytes in the key space to scan")
	flag.Parse()

	NumberOfWorkers := runtime.NumCPU()
	workers := make([]hashScanner.Worker, NumberOfWorkers)
	for i := uint(0); i < uint(len(workers)); i++ {
		if err := workers[i].Star=-t(i, *keySpaceSize); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	}

}
