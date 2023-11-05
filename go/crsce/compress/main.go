package main

/*
 * CRSCE Compress Tool
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Proof of concept for CRSCE compression.
 */
import (
	"github.com/sam-caldwell/monorepo/go/crsce/BitFile"
)

const commandUsage = `

crsce-compress --in <input_file> --out <output_file> [--bs <blocksize>]

`

func main() {
	var args Arguments
	args.GetArgs()

	var source bitfile.BitFile
	if err := source.Open(args.In); err != nil {
		panic(err)
	}
	defer source.Close()

	var target bitfile.BitFile
	if err := target.Open(args.Out); err != nil {
		panic(err)
	}
	defer target.Close()

	var fileSize int64
	if fileSize, err := source.Size(); err != nil {
		panic(err)
	}

}
