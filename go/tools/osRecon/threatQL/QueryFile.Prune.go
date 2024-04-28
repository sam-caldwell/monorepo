package threatQL

import "github.com/sam-caldwell/monorepo/go/fs/file"

// Prune - If queryFile has a defined filename, delete it
func (q *QueryFile) Prune() error {
	if q.fileName != nil {
		//A query was loaded, delete this file.
		if err := file.Delete(string(*q.fileName)); err != nil {
			return err
		}
		q.fileName = nil
	}
	return nil
}
