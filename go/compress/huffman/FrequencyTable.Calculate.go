package huffman

// Calculate takes a symbol slice and returns a map with the frequency of each symbol.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (freq *FrequencyTable) Calculate(text []byte) {
	*freq = make(FrequencyTable)
	for _, c := range text {
		(*freq)[c]++
	}
}
