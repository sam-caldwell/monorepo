package bitreader

import (
    "bufio"
    "io"
)

// NewBitReader - create/configure a bitReader for reading from an io.Reader.
//
//	(c) 2023 Sam Caldwell.  MIT License
func NewBitReader(reader io.Reader) *BitReader {
    var (
        ok         bool
        byteReader io.ByteReader
    )

    if byteReader, ok = reader.(io.ByteReader); !ok {
        byteReader = bufio.NewReader(reader)
    }

    return &(BitReader{reader: byteReader})
}
