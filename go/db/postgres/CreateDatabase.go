package Postgres

import (
	"fmt"
)

func (db *Db) CreateDatabase(dbName string, owner string) error {
	if !SanitizeDBName(&dbName) {
		return fmt.Errorf("invalid database name (%s)", dbName)
	}
	if _, err := db.conn.Exec(fmt.Sprintf("create database %s", dbName)); err != nil {
		return fmt.Errorf("error creating the database %s: %v\n", dbName, err)
	}
	_, err := db.conn.Exec(fmt.Sprintf("grant all privileges on database %s to %s", dbName, owner))
	if err != nil {
		return fmt.Errorf("failed to grant privileges on the database %s to %s.  error: %v\n",
			dbName, owner, err)
	}

	return nil
}
