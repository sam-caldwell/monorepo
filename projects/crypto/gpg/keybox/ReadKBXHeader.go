package keybox

import (
	"encoding/binary"
	"os"
)

// ReadKBXHeader  - read the KBX file header
//
//	(c) 2023 Sam Caldwell.  MIT License
func ReadKBXHeader(file *os.File) (KBXHeader, error) {
	var header KBXHeader
	err := binary.Read(file, binary.BigEndian, &header)
	if err != nil {
		return KBXHeader{}, err
	}
	return header, nil
}
