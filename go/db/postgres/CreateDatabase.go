package Postgres

import (
	"fmt"
)

func (db *Db) CreateDatabase(dbName string, owner string) *Db {
	var sql string
	if !SanitizeDBName(&dbName) {
		db.err = fmt.Errorf("invalid database name (%s)", dbName)
	}
	sql = fmt.Sprintf("create database '%s';", dbName)
	if _, err := db.conn.Exec(sql); err != nil {
		db.err = fmt.Errorf("error creating the database %s: %v", dbName, err)
		return db
	}

	sql = fmt.Sprintf("grant all privileges on database %s to %s", dbName, owner)
	if _, err := db.conn.Exec(sql); err != nil {
		db.err = fmt.Errorf("error granting privileges on the database %s to %s: %v", dbName, owner, err)
		return db
	}
	return db
}
