package huffman

func Encode(input []byte) (output []byte) {

	encode := func(input []byte, codes map[byte]string) []byte {
		var encodedStr string
		for _, b := range input {
			encodedStr += codes[b]
		}
		return []byte(encodedStr)
	}

	if len(input) == 0 {
		return nil
	}

	// Build the Huffman tree
	root := buildHuffmanTree(input)

	// Generate Huffman codes
	codes := make(map[byte]string)
	generateHuffmanCodes(root, "", codes)

	// Encode the input
	encodedData := encode(input, codes)

	return encodedData
}
