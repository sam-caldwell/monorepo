package persistence

import (
	"github.com/google/uuid"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// FetchQueryList - Get a list of up to count files for a given client (uuid).
func (qry *QueryCollector) FetchQueryList(clientId uuid.UUID, count uint16) (err error) {
	const (
		maxQueryFileSz     = 16384
		queryFileExtension = ".tql"
	)
	var (
		dir      *os.File
		dirEntry []os.DirEntry
	)
	defer func() {
		if dir != nil {
			_ = dir.Close()
		}
	}()
	//
	// Open the directory
	//
	directory := filepath.Join(qry.path.String(), clientId.String())
	if dir, err = os.Open(directory); err != nil {
		return err
	}
	//
	// Read the contents
	//
	if dirEntry, err = dir.ReadDir(int(count)); err != nil {
		return err
	}
	//
	// filter out any directories or files that do not meet our size or extension requirements
	//
	for _, file := range dirEntry {
		var info fs.FileInfo
		if info, err = file.Info(); err != nil {
			return err
		}
		if file.IsDir() || !(strings.HasSuffix(info.Name(), queryFileExtension)) || (info.Size() > maxQueryFileSz) {
			//Skip file that fails filter rules
			continue
		}
		qry.queue[clientId] = append(qry.queue[clientId], filepath.Join(directory, info.Name()))
	}
	return err
}
