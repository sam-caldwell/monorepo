package rle

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"strconv"
)

// Decode performs run-length decoding on the input byte slice and returns the decoded byte slice.
// It returns an error if the input is malformed.
func Decode(input []byte) ([]byte, error) {
	var decoded []byte
	length := len(input)
	if length == 0 {
		return decoded, nil
	}
	if length%2 != 0 {
		return decoded, fmt.Errorf(errors.MalformedInput)
	}

	for i := 0; i < length; i += 2 {
		char := input[i]
		count, err := strconv.Atoi(string(input[i+1]))
		if err != nil {
			return []byte{}, fmt.Errorf(errors.MalformedInput)
		}
		for j := 0; j < count; j++ {
			decoded = append(decoded, char)
		}
	}

	return decoded, nil
}
