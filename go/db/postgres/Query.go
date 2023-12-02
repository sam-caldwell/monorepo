package Postgres

import (
	"database/sql"
)

func (db *Db) Query(sql string) (*sql.Rows, error) {
	if db.conn == nil {
		panic("uninitialized db connection")
	}
	return db.conn.Query(sql)
}
