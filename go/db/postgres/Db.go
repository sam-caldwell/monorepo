package Postgres

import "database/sql"

type Db struct {
	conn *sql.DB
}
