package compress

import (
	"compress/gzip"
	"io"
)

// StreamCompress reads raw data from the 'in' io.Reader,
// compresses it, and writes the compressed data to the 'out' io.Writer.
func StreamCompress(in io.Reader, out io.Writer) (err error) {

	gzipWriter := gzip.NewWriter(out)
	defer func() { _ = gzipWriter.Close() }()

	if _, err = io.Copy(gzipWriter, in); err != nil {
		return err
	}

	// Close the gzip writer to flush the compressed data
	if err = gzipWriter.Close(); err != nil {
		return err
	}

	return err
}
