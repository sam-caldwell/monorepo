package pair

import (
	"bytes"
	"fmt"
)

// KeyToBytes converts a key of KeyType to a []byte.
//
//	(c) 2023 Sam Caldwell.  MIT License
func KeyToBytes[KeyType comparable](key KeyType) ([]byte, error) {
	const conversionError = "error converting key to bytes: %w"
	var buf bytes.Buffer

	if err := writeValueToBytes(&buf, key); err != nil {
		return nil, fmt.Errorf(conversionError, err)
	}

	return buf.Bytes(), nil
}
