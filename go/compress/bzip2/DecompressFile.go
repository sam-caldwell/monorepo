package bzip2

import (
	"os"
)

// DecompressFile - given an input and output file/path name, decompress the input content
//
//	(c) 2023 Sam Caldwell.  MIT License
func DecompressFile(inputPath, outputPath string) (err error) {
	var (
		inputFile  *os.File
		outputFile *os.File
	)

	if inputFile, err = os.Open(inputPath); err != nil {
		return err
	}
	defer func() { _ = inputFile.Close() }()

	if outputFile, err = os.Create(outputPath); err != nil {
		return err
	}
	defer func() { _ = outputFile.Close() }()
	return Decompress(inputFile, outputFile)
}
