package huffman

// huffmanCode - Object representing a symbol, code and code length a symbol, its code and code length.
//
//	(c) 2023 Sam Caldwell.  MIT License
type huffmanCode[CodeType, SymbolType IndexTypes] struct {
	code    CodeType
	codeLen uint8
	symbol  SymbolType
}
