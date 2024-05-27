package compress

import (
	"compress/gzip"
	"io"
)

// StreamDecompress reads gzip-compressed data from the 'in' io.Reader,
// decompresses it, and writes the decompressed data to the 'out' io.Writer.
//
//	(c) 2023 Sam Caldwell.  MIT License
func StreamDecompress(in io.Reader, out io.Writer) error {

	gzipReader, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer func() { _ = gzipReader.Close() }()

	// Copy the decompressed data to the output writer
	_, err = io.Copy(out, gzipReader)
	if err != nil {
		return err
	}

	return nil
}
