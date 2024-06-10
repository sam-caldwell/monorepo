package huffman

import "bytes"

// FromBytes method to decode the []byte slice into the CodeMap
//
//	(c) 2023 Sam Caldwell.  MIT License
func (cm *CodeMap) FromBytes(b []byte) {
	*cm = make(CodeMap)
	buffer := bytes.NewBuffer(b)

	for buffer.Len() > 0 {
		key, _ := buffer.ReadByte()    // Read the key byte
		length, _ := buffer.ReadByte() // Read the length of the value
		value := make([]byte, length)
		_, err := buffer.Read(value) // Read the value bytes
		if err != nil {
			panic(err)
		}
		(*cm)[key] = value
	}
}
