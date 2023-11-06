package bitfile

/*
 * Tests for Block.BitRead()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines unit tests for Block.BitRead()
 */

import (
	"errors"
	"testing"
)

func TestBlock_ReadBit(t *testing.T) {
	blk := Block{
		buffer: []byte{0x00, 0xFF, 0x02, 0x04, 0x08, 0x10, 0x11, 0x12, 0x13, 0x14},
	}
	bit2int := func(b bool) uint {
		if b {
			return 1
		}
		return 0
	}

	testByte := func(pos uint, value byte) {
		if blk.buffer[pos] != value {
			t.Fatalf("Byte value mismatch (pos: %d, value: %02x)", pos, value)
		}
		//t.Logf("--Byte: %02x--", value)
	}

	testBit := func(pos, expBit uint, expErr error) {
		bit, err := blk.ReadBit(pos)
		if (err != nil) && (err.Error() != expErr.Error()) {
			t.Fatalf("Expected error not found.\n"+
				"expected:'%v'\n"+
				"actual:  '%v'", expErr, err)
		}
		if bit2int(bit) != expBit {
			t.Fatalf("expected bit mismatch (pos: %d, value: %02x)", pos, bit2int(bit))
		}
	}

	//Test bounds check
	testBit(uint(8*len(blk.buffer)), 0, errors.New("index out of range"))

	//Test bits 0x00
	testByte(0, 0x00)
	testBit(0, 0, nil)
	testBit(1, 0, nil)
	testBit(2, 0, nil)
	testBit(3, 0, nil)

	testBit(4, 0, nil)
	testBit(5, 0, nil)
	testBit(6, 0, nil)
	testBit(7, 0, nil)

	//Test bits 0xFF
	testByte(1, 0xFF)
	testBit(8, 1, nil)
	testBit(9, 1, nil)
	testBit(10, 1, nil)
	testBit(11, 1, nil)

	testBit(12, 1, nil)
	testBit(13, 1, nil)
	testBit(14, 1, nil)
	testBit(15, 1, nil)

	//Test bits 0x02
	testByte(2, 0x02)
	testBit(16, 0, nil)
	testBit(17, 1, nil)
	testBit(18, 0, nil)
	testBit(19, 0, nil)

	testBit(20, 0, nil)
	testBit(21, 0, nil)
	testBit(22, 0, nil)
	testBit(23, 0, nil)

	//Test bits 0x04
	testByte(3, 0x04)
	testBit(24, 0, nil)
	testBit(25, 0, nil)
	testBit(26, 1, nil)
	testBit(27, 0, nil)

	testBit(28, 0, nil)
	testBit(29, 0, nil)
	testBit(30, 0, nil)
	testBit(31, 0, nil)

	//Test bits 0x08
	testByte(4, 0x08)
	testBit(32, 0, nil)
	testBit(33, 0, nil)
	testBit(34, 0, nil)
	testBit(35, 1, nil)

	testBit(36, 0, nil)
	testBit(37, 0, nil)
	testBit(38, 0, nil)
	testBit(39, 0, nil)

	//Test bits 0x10
	testByte(5, 0x10)
	testBit(40, 0, nil)
	testBit(41, 0, nil)
	testBit(42, 0, nil)
	testBit(43, 0, nil)

	testBit(44, 1, nil)
	testBit(45, 0, nil)
	testBit(46, 0, nil)
	testBit(47, 0, nil)

	//Test bits 0x11
	testByte(6, 0x11)
	testBit(48, 1, nil)
	testBit(49, 0, nil)
	testBit(50, 0, nil)
	testBit(51, 0, nil)

	testBit(52, 1, nil)
	testBit(53, 0, nil)
	testBit(54, 0, nil)
	testBit(55, 0, nil)

	//Test bits 0x12
	testByte(7, 0x12)
	testBit(56, 0, nil)
	testBit(57, 1, nil)
	testBit(58, 0, nil)
	testBit(59, 0, nil)

	testBit(60, 1, nil)
	testBit(61, 0, nil)
	testBit(62, 0, nil)
	testBit(63, 0, nil)

	//Test bits 0x13
	testByte(8, 0x13)
	testBit(64, 1, nil)
	testBit(65, 1, nil)
	testBit(66, 0, nil)
	testBit(67, 0, nil)

	testBit(68, 1, nil)
	testBit(69, 0, nil)
	testBit(70, 0, nil)
	testBit(71, 0, nil)

	//Test bits 0x14
	testByte(9, 0x14)
	testBit(72, 0, nil)
	testBit(73, 0, nil)
	testBit(74, 1, nil)
	testBit(75, 0, nil)

	testBit(76, 1, nil)
	testBit(77, 0, nil)
	testBit(78, 0, nil)
	testBit(78, 0, nil)
}
