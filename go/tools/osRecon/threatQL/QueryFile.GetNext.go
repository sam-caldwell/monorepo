package threatQL

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/types"
	"io/fs"
	"strings"
)

// GetNext - Find the next file in the client query queue and set the filename.
// Note this does not load the file, it just sets the internal state.  Do not call
// Load() again until the file has been read into memory, or it will be deleted.
func (q *QueryFile) getNext(path directory.Path, clientId types.ClientId) (err error) {
	//
	// Make sure we don't run a file twice by not cleaning up after ourselves
	// by only allowing getNext() if we have disposed of and reset q.fileName
	//
	if q.fileName != nil {
		return fmt.Errorf("cannot get next file until last file is cleared")
	}
	//
	// Glob the given path for any files containing the query file extension.
	// Anything in a subdirectory will be ignored.
	//
	fileList, err := file.Glob(path.String(), fmt.Sprintf("*.%s", queryFileExtension))
	//
	// Iterate over the globed file list until we find a file that meets our requirements.
	// This will eliminate--
	//      - directories
	//      - files that do not have .tql extensions
	//      - files that exceed our maximum size.
	//
	for _, targetFile := range fileList {
		var info fs.FileInfo
		if info, err = file.Stat(targetFile); err != nil {
			// Bail, we had an error getting the file info.
			return err
		} else {
			if info.IsDir() || !(strings.HasSuffix(info.Name(), queryFileExtension)) || (info.Size() > maxQueryFileSz) {
				// Skip file that fails filter rules
				continue
			} else {
				// We found the first file, capture it and terminate the loop
				f := file.File(targetFile)
				q.fileName = &f
				break
			}
		}
	}
	return err
}
