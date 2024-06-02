package bitreader

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"io"
)

// ReadBits64 - read the specified number of bits from the BitReader and returns them as uint64.
//
// The BitReader structure maintains an internal buffer (br.buffer) and a count of how many bits are
// currently in the buffer (br.numberBitsInBuffer). This function extracts the requestedNumberOfBits
// from the buffer and returns them as uint64 (`bitString`).
//
// If there are not enough bits currently in the buffer, the function reads more bytes from the
// underlying (io.ByteReader) until there are enough bits to fulfill the request. The new bytes
// are added to the internal buffer (br.buffer), and the bit count is updated accordingly.
//
// The function works as follows:
//
//  1. While requestedNumberOfBits is greater than the br.numberBitsInBuffer, read another byte from the
//     underlying io.ByteReader:
//     - If reading a byte returns an io.EOF error, the method returns 0 and sets the error state internally.
//     - Otherwise, the byte is added to the buffer (`br.buffer`), and the bit count is incremented by 8.
//
//  2. Once there are enough bits in the buffer, the requested number of bits (requestedNumberOfBits) are
//     extracted:
//     - The bits are extracted by shifting the buffer right by the difference between the total bit count and
//     the requested bit count, then masking the result to get the lowest `requestedNumberOfBits` bits.
//
//  3. The internal bit count (`br.bits`) is decremented by the number of bits that were read, and the function
//     returns the extracted bits as uint64 (`bitString`).
//
//     (c) 2023 Sam Caldwell.  MIT License
func (br *BitReader) ReadBits64(requestedNumberOfBits uint8) (bitString uint64) {
	const (
		BitsInByte              = 8
		MaxNumberOfBitsInOutput = 64
	)
	if requestedNumberOfBits > MaxNumberOfBitsInOutput {
		br.err = fmt.Errorf(errors.BoundsCheckError)
	}
	for uint64(requestedNumberOfBits) > br.numberBitsInBuffer {
		b, err := br.reader.ReadByte()
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		if err != nil {
			br.err = err
			return 0
		}
		br.buffer <<= BitsInByte
		br.buffer |= uint64(b)
		br.numberBitsInBuffer += BitsInByte
	}
	shiftRequired := br.numberBitsInBuffer - uint64(requestedNumberOfBits)
	bitString = (br.buffer >> (shiftRequired)) & ((1 << requestedNumberOfBits) - 1)
	br.numberBitsInBuffer -= uint64(requestedNumberOfBits)
	return
}
