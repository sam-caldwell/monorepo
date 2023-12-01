package Postgres

func (db *Db) Error() error {
	if db.err != nil {
		return db.err
	}
	return nil
}
