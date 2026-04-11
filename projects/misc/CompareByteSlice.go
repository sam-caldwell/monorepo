package misc

const (
	Unknown = iota
	BytesEqual
	BytesLess
	BytesGreater
)

// CompareByteSlice - Compare two byte slices and return an integer result (1 equal, 2 LT, 3 GT)
func CompareByteSlice(lhs, rhs []byte) (result int) {
	maxIndex := PickSmallestInt(len(lhs), len(rhs))
	for i := 0; i < maxIndex; i++ {
		if lhs[i] == rhs[i] {
			result = BytesEqual
		} else if lhs[i] <= rhs[i] {
			result = BytesLess
		} else {
			result = BytesGreater
		}
	}
	return result
}
