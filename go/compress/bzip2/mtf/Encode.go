package mtf

// Encode - Perform Move-To-Front Encode Operation
//
//	(c) 2023 Sam Caldwell.  MIT License
func Encode(input []byte) []byte {
	symbols := make([]byte, 256)
	for i := range symbols {
		symbols[i] = byte(i)
	}

	output := make([]byte, len(input))
	for i, val := range input {
		index := 0
		for symbols[index] != val {
			index++
		}
		output[i] = byte(index)
		copy(symbols[1:index+1], symbols[:index])
		symbols[0] = val
	}
	return output
}
