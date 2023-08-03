package crypto

/*
 * projects/crypto/NewSha256Stream
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides NewSha256Stream() which will
 * create and initialize a new sha256 stream object.
 */

// NewSha256Stream - Create and initialize a sha256 stream object.
func NewSha256Stream() *Sha256Stream {
	return &Sha256Stream{
		h: [8]uint32{
			uint32(ivA),
			uint32(ivB),
			uint32(ivC),
			uint32(ivD),
			uint32(ivE),
			uint32(ivF),
			uint32(ivG),
			uint32(ivH),
		},
		buffer:  [64]byte{},
		bitNdx:  0,
		byteNdx: 0,
		size:    uint32(0),
	}
}
