package crsce

import "github.com/sam-caldwell/go/v2/projects/fs/file"

// CSM - Virtual CSM object.
type CSM struct {
	buffer   file.Reader
	position uint64 //the bit position within the stream
	err      error  //the reader error state.
}
