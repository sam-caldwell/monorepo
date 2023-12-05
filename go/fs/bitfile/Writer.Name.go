package bitfile

/*
 * Writer.Name()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Return the file name of the given file
 */

import "github.com/sam-caldwell/monorepo/go/misc/words"

// Name - Return the file name of any currently open bitfile.
func (o *Writer) Name() string {

	if o.file != nil {
		return o.file.Name()
	}

	return words.EmptyString

}
