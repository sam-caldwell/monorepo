package main

/*
 * CRSCE Compress Tool
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Proof of concept for CRSCE compression.
 */
/*
import (
	"fmt"
	crsce "github.com/sam-caldwell/monorepo/go/crsce/compress/lib"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/file"
	file2 "github.com/sam-caldwell/monorepo/go/fs/file"
)
*/
/*
func main() {
	//
	// evaluate command-line arguments
	//
	var args crsce.Arguments
	args.GetArgs()
	//
	// open the input (source) file
	//
	var source file2.BitFile
	exit.CheckError(source.OpenRead(&args.In))
	defer source.Close()
	//
	// open the output (target) file
	//
	var target file2.BitFile
	exit.CheckError(target.OpenRead(&args.Out))
	defer target.Close()
	//
	// Determine the file size
	//
	fileSize, err := source.Size()
	exit.CheckError(err)

	fmt.Printf("compressing file with %d bytes", fileSize)

	exit.CheckError(writeFileHeader(&target, fileSize, args.BlockSize))
	//
	// Read the entire source (input) file block by block
	// For each block, read in the bits for that block to crate
	// a cross-sums pair (LSM, VSM) and write these blocks to disk
	//
	var endOfFile bool
	for blockId := 0; !endOfFile; blockId++ {
		block, err := source.ReadBytes(blockSize)
		endOfFile, err = file.IsEndOfFile(err)
		exit.CheckErrorf(err)
		exit.CheckError(crsce.Compress(block, target))
	}
}
*/
