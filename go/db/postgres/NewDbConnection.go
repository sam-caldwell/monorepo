package Postgres

func NewDbConnection(dbHost string, dbPort uint, dbName string, dbUser string, dbPass string, useTls bool) (*Db, error) {
	var db Db
	err := db.Connect(dbHost, dbPort, dbName, dbUser, dbPass, useTls)
	return &db, err
}
