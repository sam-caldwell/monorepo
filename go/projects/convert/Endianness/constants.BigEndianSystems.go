//go:build ppc || ppc64 || s390 || s390x
// +build ppc ppc64 s390 s390x

package endianness

/*
 * projects/convert/Endianness/constants.BigEndianness.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides two simple constants to help us determine system endianness at
 * runtime without being unsafe or engaging in the pedantic dumbfuckery of a conversation
 * with the goLang-nuts forum of opinionated holier-than-though jack-wagons.  If this
 * offends you, consider the fact that so many of us have to create workarounds because
 * you live in perfect, pretty worlds while the rest of us write code in reality where
 * nasty stuff still exists.
 */
const (
	BigEndian    = true
	LittleEndian = false
)
