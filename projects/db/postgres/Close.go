package Postgres

func (db *Db) Close() error {
	return db.conn.Close()
}
