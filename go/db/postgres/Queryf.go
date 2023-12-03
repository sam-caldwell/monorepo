package Postgres

import (
    "database/sql"
)

func (db *Db) Queryf(sql string, values ...any) (*sql.Rows, error) {
	if db.conn == nil {
		panic("uninitialized db connection")
	}
	return db.conn.Query(sql, values...)
}
