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
	source.Open(args.In)
	defer source.Close()
	var target bitfile.BitFile
	target.Open(args.Out)
	defer target.Close()

}
