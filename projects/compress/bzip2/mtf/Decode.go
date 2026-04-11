package mtf

// Decode - Perform Move-To-Front Decode Operation
//
//	(c) 2023 Sam Caldwell.  MIT License
func Decode(input []byte) []byte {
	symbols := make([]byte, 256)
	for i := range symbols {
		symbols[i] = byte(i)
	}

	output := make([]byte, len(input))
	for i, val := range input {
		symbol := symbols[val]
		output[i] = symbol
		copy(symbols[1:val+1], symbols[:val])
		symbols[0] = symbol
	}
	return output
}
