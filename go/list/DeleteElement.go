package list

import "github.com/sam-caldwell/monorepo/go/exit/errors"

// DeleteElement - Remove a given element from the input list
func DeleteElement[T any](slice []T, s int) []T {

	if s < 0 || s >= len(slice) {
		panic(errors.IndexOutOfRange)
	}

	return append(slice[:s], slice[s+1:]...)
}
