package crsce

import "github.com/sam-caldwell/go/v2/projects/fs/file"

type CSM struct {
	buffer   file.Reader
	count    uint64 //number of bytes processed
	position uint64 //the bit position within the stream
}
