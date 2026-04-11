package Postgres

import (
	"database/sql"
	"fmt"
)

func (db *Db) Connect(dbHost string, dbPort uint, dbName string, dbUser string, dbPass string, useTls bool) error {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ",
		dbHost, dbPort, dbUser, dbPass, dbName)

	if useTls {
		connectionString += "sslmode=enable"
	} else {
		connectionString += "sslmode=disable"
	}

	dbConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to the database %s: %v", dbName, err)
	}

	db.conn = dbConn

	return nil
}
