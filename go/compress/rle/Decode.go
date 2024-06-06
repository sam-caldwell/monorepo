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

	for i := 0; i < length; {
		if i+1 >= length {
			return nil, fmt.Errorf(errors.MalformedInput)
		}

		char := input[i]
		j := i + 1

		// Find the end of the numeric count
		for j < length && input[j] >= '0' && input[j] <= '9' {
			j++
		}
		count, err := strconv.Atoi(string(input[i+1 : j]))
		if err != nil {
			return nil, fmt.Errorf(errors.MalformedInput)
		}

		for k := 0; k < count; k++ {
			decoded = append(decoded, char)
		}
		i = j
	}

	return decoded, nil
}
