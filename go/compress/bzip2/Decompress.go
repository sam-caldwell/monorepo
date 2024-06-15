package bzip2

import (
	"encoding/binary"
	"github.com/sam-caldwell/monorepo/go/compress/bitreader"
	"github.com/sam-caldwell/monorepo/go/compress/bzip2/bwt"
	"github.com/sam-caldwell/monorepo/go/compress/bzip2/mtf"
	"github.com/sam-caldwell/monorepo/go/compress/huffman"
	"github.com/sam-caldwell/monorepo/go/convert"
	"io"
)

// Decompress - Decompress the given input and write to the given output
//
//	(c) 2023 Sam Caldwell.  MIT License
func Decompress(input io.Reader, output io.Writer) (err error) {
	var (
		encodedBits []bool
		freqs       map[byte]int
		index       int32
		indexBytes  []byte
	)
	indexBytes = make([]byte, 4)
	if _, err = input.Read(indexBytes); err != nil {
		return err
	}
	if index, err = convert.BytesToInt32(indexBytes); err != nil {
		return err
	}
	if freqs, err = deserializeFreqs(input); err != nil {
		return err
	}
	bitReader := bitreader.NewBitStreamReader(input)
	for {
		bit, err := bitReader.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		encodedBits = append(encodedBits, bit)
	}
	bwtOutput := bwt.Decode(
		mtf.Decode(huffman.Decode(freqs, encodedBits)),
		int(index))

	if _, err := output.Write(bwtOutput); err != nil {
		return err
	}
	return nil
}

// deserializeFreqs - deserialize the frequency table
func deserializeFreqs(reader io.Reader) (map[byte]int, error) {
	freqs := make(map[byte]int)
	for {
		b := make([]byte, 1)
		if _, err := reader.Read(b); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		freqBytes := make([]byte, 4)
		if _, err := reader.Read(freqBytes); err != nil {
			return nil, err
		}
		freq := int(binary.BigEndian.Uint32(freqBytes))
		freqs[b[0]] = freq
	}
	return freqs, nil
}
