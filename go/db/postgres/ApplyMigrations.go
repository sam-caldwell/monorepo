package Postgres

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func (db *Db) ApplyMigration(locations string) *Db {

	files, err := filepath.Glob(filepath.Join(locations, fmt.Sprintf("*.sql")))
	if err != nil {
		db.err = err
		return db
	}

	slices.Sort(files)

	for _, sqlFile := range files {
		content, err := os.ReadFile(sqlFile)
		if err != nil {
			db.err = err
			break
		}
		if _, err = db.conn.Exec(string(content)); err != nil {
			db.err = fmt.Errorf("error executing query from file %s: %v", sqlFile, err)
			break
		}
	}
	return db
}
