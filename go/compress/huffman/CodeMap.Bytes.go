package huffman

import "bytes"

// Bytes method to convert CodeMap to a single []byte
//
//	(c) 2023 Sam Caldwell.  MIT License
func (cm *CodeMap) Bytes() []byte {
	var buffer bytes.Buffer

	for key, value := range *cm {
		buffer.WriteByte(key)              // Write the key byte
		buffer.WriteByte(byte(len(value))) // Write the length of the value
		buffer.Write(value)                // Write the value bytes
	}

	return buffer.Bytes()
}
