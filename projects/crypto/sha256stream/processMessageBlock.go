package crypto

import "encoding/binary"

/*
 * projects/crypto/processMessageBlock.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file processes the hash.buffer message block
 * and resets the buffer and indexes afterward.
 */

// processMessageBlock -
func (hash *Sha256Stream) processMessageBlock() {

	// Constants (from SHA-256 specification)
	var K = []uint32{
		0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
		0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
		0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
		0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
		0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
		0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
		0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
		0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
	}
	var a = hash.h[_a]
	var b = hash.h[_b]
	var c = hash.h[_c]
	var d = hash.h[_d]
	var e = hash.h[_e]
	var f = hash.h[_f]
	var g = hash.h[_g]
	var h = hash.h[_h]

	// Step 3: Message Scheduling
	var words [64]uint32
	for i := 0; i < 16; i++ {
		words[i] = binary.BigEndian.Uint32(hash.buffer[i*4 : (i+1)*4])
	}

	for i := 16; i < 64; i++ {
		s0 := (words[i-15]>>7 | words[i-15]<<25) ^ (words[i-15]>>18 | words[i-15]<<14) ^ (words[i-15] >> 3)
		s1 := (words[i-2]>>17 | words[i-2]<<15) ^ (words[i-2]>>19 | words[i-2]<<13) ^ (words[i-2] >> 10)
		words[i] = words[i-16] + s0 + words[i-7] + s1
	}

	// Step 4: Hash Computation (Main Loop)
	for i := 0; i < 64; i++ {
		var s1 uint32 = (e>>6 | e<<26) ^ (e>>11 | e<<21) ^ (e>>25 | e<<7)
		var ch uint32 = (e & f) ^ ((^e) & g)
		k := K[i]
		w := words[i]
		temp1 := h + s1 + ch + k + w

		var s0 uint32 = (a>>2 | a<<30) ^ (a>>13 | a<<19) ^ (a>>22 | a<<10)
		var maj uint32 = (a & b) ^ (a & c) ^ (b & c)
		temp2 := s0 + maj

		h = g
		g = f
		f = e
		e = (d + temp1) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
		d = c
		c = b
		b = a
		a = (temp1 + temp2) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
	}

	// Step 5: Update Hash Values
	hash.h[_a] = (hash.h[_a] + a) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
	hash.h[_b] = (hash.h[_b] + b) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
	hash.h[_c] = (hash.h[_c] + c) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
	hash.h[_d] = (hash.h[_d] + d) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
	hash.h[_e] = (hash.h[_e] + e) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
	hash.h[_f] = (hash.h[_f] + f) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
	hash.h[_g] = (hash.h[_g] + g) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
	hash.h[_h] = (hash.h[_h] + h) & 0xffffffff // Use bitwise AND operation to keep within 32-bit range
}
