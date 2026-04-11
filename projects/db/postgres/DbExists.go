package Postgres

import (
	_ "github.com/lib/pq"
)

func (db *Db) DbExists(dbName string) bool {
	var exists bool
	_ = db.conn.QueryRow("SELECT 1 FROM pg_database WHERE datname = %s", dbName).Scan(&exists)
	return exists
}
