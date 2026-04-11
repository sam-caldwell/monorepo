package Postgres

import (
	"database/sql"
)

func (db *Db) Exec(sql string) (sql.Result, error) {
	if db.conn == nil {
		panic("uninitialized db connection")
	}
	return db.conn.Exec(sql)
}
