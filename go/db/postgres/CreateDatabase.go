package Postgres

import (
	"fmt"
)

// CreateDatabase - Create database if not exists (with option to drop the database first)
func (db *Db) CreateDatabase(dbName string, owner string, dropIfExists bool) error {
	if !SanitizeDBName(&dbName) {
		return fmt.Errorf("invalid database name (%s)", dbName)
	}
	databaseExists := db.DbExists(dbName)
	if dropIfExists && databaseExists {
		if _, err := db.conn.Exec(fmt.Sprintf("drop database %s", dbName)); err != nil {
			return fmt.Errorf("error dropping database %s (before create): %v\n", dbName, err)
		}
	}
	if !databaseExists {
		if _, err := db.conn.Exec(fmt.Sprintf("create database %s", dbName)); err != nil {
			return fmt.Errorf("error creating the database %s: %v\n", dbName, err)
		}

		_, err := db.conn.Exec(fmt.Sprintf("grant all privileges on database %s to %s", dbName, owner))
		if err != nil {
			return fmt.Errorf("failed to grant privileges on the database %s to %s.  error: %v\n",
				dbName, owner, err)
		}
	}
	return nil
}
