package bzip2

import (
	"github.com/sam-caldwell/monorepo/go/compress/bitwriter"
	"github.com/sam-caldwell/monorepo/go/compress/bzip2/bwt"
	"github.com/sam-caldwell/monorepo/go/compress/bzip2/mtf"
	"github.com/sam-caldwell/monorepo/go/compress/huffman"
	"github.com/sam-caldwell/monorepo/go/convert"
	"io"
)

// Compress - compress the given input Bytes
//
//	(c) 2023 Sam Caldwell.  MIT License
func Compress(input io.Reader, output io.Writer) (err error) {

	var inputBytes, indexBytes []byte

	if inputBytes, err = io.ReadAll(input); err != nil {
		return err
	}

	bwtOutput, index := bwt.Encode(inputBytes)

	mtfOutput := mtf.Encode(bwtOutput)

	freqs := make(map[byte]int)
	for _, b := range mtfOutput {
		freqs[b]++
	}

	codes := make(map[byte][]bool)

	huffman.Encode(freqs, codes)

	bitWriter := bitwriter.NewBitStreamWriter(output)
	if indexBytes, err = convert.Int32ToBytes(index); err != nil {
		return err
	}
	if _, err = output.Write(indexBytes); err != nil {
		return err
	}
	for _, b := range mtfOutput {
		code := codes[b]
		for _, bit := range code {
			if err := bitWriter.WriteBit(bit); err != nil {
				return err
			}
		}
	}
	if err = bitWriter.Flush(); err != nil {
		return err
	}
	return nil
}
