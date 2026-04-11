package bzip2

import (
	"os"
)

// CompressFile - given input and output path/filename, compress the contents and return any errors.
//
//	(c) 2023 Sam Caldwell.  MIT License
func CompressFile(inputPath, outputPath string) (err error) {

	var inputFile *os.File

	var outputFile *os.File

	if inputFile, err = os.Open(inputPath); err != nil {
		return err
	}

	defer func() { _ = inputFile.Close() }()

	if outputFile, err = os.Create(outputPath); err != nil {
		return err
	}

	defer func() { _ = outputFile.Close() }()

	return Compress(inputFile, outputFile)

}
