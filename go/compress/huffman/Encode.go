package huffman

import "fmt"

// Encode encodes a given symbol using the Huffman tree.
func (t *huffmanTree[CodeType, SymbolType]) Encode(symbol SymbolType,
	lookupTable map[SymbolType]huffmanCode[CodeType, SymbolType]) (CodeType, uint8, error) {

	if code, ok := lookupTable[symbol]; ok {
		return code.code, code.codeLen, nil
	}

	return 0, 0, fmt.Errorf("symbol %d not found in Huffman tree", symbol)

}
