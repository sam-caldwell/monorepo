package counters

/*
 * LargeCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * LargeCounter will create an array of n uin64 elements, where the array represents
 * some significant large number.  This is more efficient than ByteCounter because
 * the counter does not perform as many carry operations since each element is a
 * large 8-byte number rather than a single byte.
 */

// LargeCounter - A struct which embodies the large number in memory
type LargeCounter struct {
	sz uint     //Number of bits in the large number
	v  []uint64 // uint64 value array
}
