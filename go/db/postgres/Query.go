package Postgres

import (
	"database/sql"
	"fmt"
)

func (db *Db) Query(sql string, values ...any) (*sql.Rows, error) {
	if db.conn == nil {
		panic("uninitialized db connection")
	}
	return db.conn.Query(fmt.Sprintf(sql, values...))
}
