package Postgres

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
	"path/filepath"
	"slices"
)

func (db *Db) ApplyMigration(locations string) error {

	ansi.Blue().Time().Printf("Applying Migrations to database\n").Reset()

	files, err := filepath.Glob(filepath.Join(locations, fmt.Sprintf("*.sql")))
	if err != nil {
		return fmt.Errorf("SQL file glob failed. error: %v\n", err)
	}

	slices.Sort(files)

	for _, sqlFile := range files {

		ansi.Blue().Time().Printf(">>SQL File (applying): %s ", sqlFile).Reset()

		content, err := os.ReadFile(sqlFile)
		if err != nil {
			ansi.LF()
			return fmt.Errorf(err.Error())
		}

		if _, err = db.conn.Exec(string(content)); err != nil {
			ansi.LF()
			return fmt.Errorf("error executing query from file %s: %v", sqlFile, err)
		}
		ansi.Green().Print("...done\n").Reset()
	}
	ansi.Green().Time().Println("Migration applied successfully").Reset()
	return nil

}
