package threatQL

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/types"
	"path/filepath"
)

// NewQueryFile - Create a new file and write the given query to the same.
// Return the QueryFile (in-memory) object or error.  If error is returned
// test first because QueryFile may be nil.
//
//	Query files are written to <path>/<client>/<timestamp>.tql
func NewQueryFile(path directory.Path, clientId types.ClientId, query Query) (queryFile *QueryFile, err error) {
	fullPath := filepath.Join(path.String(), clientId.String())
	fileName, err := file.GenerateUniqueFileName(fullPath, queryFileExtension, maxRetries)
	if err != nil {
		return nil, err
	}
	qry, err := query.Bytes()
	if err != nil {
		return nil, err
	}
	queryFileName := file.File(fileName)
	return &(QueryFile{
			query:    query,
			fileName: &queryFileName,
		}),
		file.Create(fileName, qry, 0600)
}
