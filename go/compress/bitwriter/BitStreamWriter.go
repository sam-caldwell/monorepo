package bitwriter

import "io"

// BitStreamWriter - A writer capable of writing one bit at a time.
//
//	(c) 2023 Sam Caldwell.  MIT License
type BitStreamWriter struct {
	writer io.Writer
	buffer byte
	count  int
}
