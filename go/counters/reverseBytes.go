package counters

/*
 * reverseBytes
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * reverse byte order.
 */

func reverseBytes(data []byte) []byte {
	sz := len(data)
	out := make([]byte, sz)
	for i := 0; i < sz; i++ {
		out[i] = data[(sz-1)-i]
	}
	return out
}
