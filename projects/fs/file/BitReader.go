package file

/*
 * projects/fs/file/BitReader.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * A bit-for-bit file reader.  This reader will allow
 * an asynchronous reading of a file at the bit level
 * to a channel buffer.
 */

type BitReader struct {
	buffer     chan bool
	count      int64
	readBuffer []byte
	done       bool
}
