CRSCE: Cross-Sums Compression/Expansion
=======================================

# Description
This is yet another attempt to implement Cross Sums Compression/Expansion now that the math seems pretty solid.

# Usage

```golang
package main

import "github.com/sam-caldwell/go/v2/projects/crsce"

func main() {
	const (
		mySourceFile = "test.txt"
		myTargetFile = "test.crsce"
		bufferSize   = 64  //Read buffer size
		blockSize    = 128 //Block size is the length/width of CSM or the length of the LSM/VSM cross sums.
	)

	//var c crsce.Crsce

	// Open a source file
	//if err := c.Open(mySourceFile, bufferSize); err != nil {
	//	panic(err)
	//}
	//
	Compress the file, naming the target file.
	//if err := c.Compress(myTargetFile, blockSize); err != nil {
    //    panic(err)
	//}
}

```