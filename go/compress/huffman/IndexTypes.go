package huffman

// IndexTypes - the type used for indexing the huffmanNodes.
//
// For example:
//
//		when implementing bzip2, we use uint16 to handle MTF indexes
//		in the tree and magic values for run-length encoding and EOF
//		and 256 possible symbols.
//
//	 The total size of codeLen + symbol must be less than the 64-bit
//	 code (uint64).  See huffmanCode struct, a generic with a definable
//	 symbol size.
//
//		(c) 2023 Sam Caldwell.  MIT License
type IndexTypes interface {
	uint16 | uint32 | uint64
}
