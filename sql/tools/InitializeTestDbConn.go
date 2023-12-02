package database

import (
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	env "github.com/sam-caldwell/monorepo/go/environment"
	"testing"
)

// InitializeTestDbConn - For testing purposes, this establishes a connection to the database
func InitializeTestDbConn(t *testing.T) *Postgres.Db {
	var db Postgres.Db

	pgHost, err := env.RequireString("db_host")
	CheckError(t, err)
	pgPort, err := env.RequireInt("db_port")
	CheckError(t, err)
	pgUser, err := env.RequireString("db_user")
	CheckError(t, err)
	pgPass, err := env.RequireString("db_pass")
	CheckError(t, err)
	pgUseTls, err := env.RequireBool("use_tls")
	CheckError(t, err)
	dbName, err := env.RequireString("db_name")
	CheckError(t, err)

	if err = db.Connect(pgHost, uint(pgPort), dbName, pgUser, pgPass, pgUseTls); err != nil {
		t.Fatal(err)
	}
	return &db
}
