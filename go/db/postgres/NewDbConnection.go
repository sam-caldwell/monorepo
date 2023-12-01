package Postgres

func NewDbConnection(dbHost string, dbPort uint, dbName string, dbUser string, dbPass string, useTls bool) *Db {
	var db Db
	return db.Connect(dbHost, dbPort, dbName, dbUser, dbPass, useTls)
}
