package main

import (
	"flag"
	_ "github.com/lib/pq"
	"github.com/sam-caldwell/monorepo/go/ansi"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	env "github.com/sam-caldwell/monorepo/go/environment"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	lib "github.com/sam-caldwell/monorepo/go/tools/dbMigrations/lib"
	"os"
	"path/filepath"
)

func main() {
	var databasesDirectory string
	/*
	 * Parse command-line arguments with env variables as backup and defaults thereafter.
	 */
	rootDirectory := flag.String(lib.CliRoot, directory.GetCurrent(), lib.CliUsageRootDir)
	pgDbHost := flag.String(lib.CliDbHost, env.GetStringOrDefault(lib.EnvDbHost, lib.DefaultPgHost), lib.CliUsageDbHost)
	pgDbPort := flag.Uint(lib.CliDbPort, env.GetUintOrDefault(lib.EnvDbPort, lib.DefaultPgPort), lib.CliUsageDbPort)
	pgDbUser := flag.String(lib.CliDbUser, env.GetStringOrDefault(lib.EnvDbUser, lib.DefaultPgUser), lib.CliUsageDbUser)
	pgDbPass := flag.String(lib.CliDbPass, env.GetStringOrDefault(lib.EnvDbPass, lib.DefaultPgPass), lib.CliUsageDbPass)
	pgDbUseTls := flag.Bool(lib.CliDbTls, false, lib.CliUsageDbUseTls)
	flag.Parse()

	ansi.Blue().
		Printf("%s : %s\n", lib.CliRoot, *rootDirectory).
		Printf("%s : %v\n", lib.CliDbHost, *pgDbHost).
		Printf("%s : %v\n", lib.CliDbPort, *pgDbPort).
		Printf("%s : %v\n", lib.CliDbUser, *pgDbUser).
		Printf("%s : <redacted: %d chars>\n", lib.CliDbPass, len(*pgDbPass)).
		Printf("%s : %v\n", lib.CliDbTls, *pgDbUseTls).
		LF().
		Reset()

	/*
	 * Validate the database directory exists
	 */
	databasesDirectory = filepath.Join(*rootDirectory, lib.SqlDirectoryPath)
	if !directory.Exists(databasesDirectory) {
		ansi.Red().
			Printf("directory does not exist (%s)\n", databasesDirectory).
			Reset().
			Fatal(exit.NotFound)
	}
	ansi.Green().Printf("Databases Directory confirmed (%s)\n", databasesDirectory).Reset()

	/*
	 * Establish a default database connection (postgres)
	 */
	postgresDb, err := Postgres.NewDbConnection(*pgDbHost, *pgDbPort, lib.DefaultDb, *pgDbUser, *pgDbPass, *pgDbUseTls)
	if err != nil {
		ansi.Red().
			Printf("Error connecting to database (postgres): %v\n", err).
			Reset().
			Fatal(exit.ConnectionFailed)
	}
	ansi.Green().Println("connection to postgres is established").Reset()

	defer func() {
		if err := postgresDb.Close(); err != nil {
			ansi.Red().
				Printf("Error closing database connector: %v\n", err).
				Reset().
				Fatal(exit.ConnectionFailed)
		}
	}()

	ansi.Blue().Println("Walk through database directories and process each database discovered").Reset()
	ansi.Blue().Printf("databasesDirectory:%s\n", databasesDirectory).Reset()
	if err := filepath.Walk(databasesDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		dbName := info.Name()
		if dbName == "databases" || dbName == "sql" || dbName == "Makefile.d" {
			return nil //Skip the "databases" directory
		}
		if postgresDb.DbExists(dbName) {
			ansi.Yellow().Printf("database exists: %s\n", dbName).Reset()
		} else {
			ansi.Blue().Printf("database not found (%s) ...create one\n", dbName).Reset()
			if err := postgresDb.CreateDatabase(dbName, *pgDbUser, false); err != nil {
				ansi.Red().
					Printf("Error creating database: %v\n", err).
					Reset().
					Fatal(exit.ConnectionFailed)
			}
		}

		dbConn, err := Postgres.NewDbConnection(*pgDbHost, *pgDbPort, dbName, *pgDbUser, *pgDbPass, *pgDbUseTls)
		if err != nil {
			ansi.Red().
				Printf("Error connecting to database (%s): %v\n", dbName, err).
				Reset().
				Fatal(exit.ConnectionFailed)
		}
		defer func() {
			if err := dbConn.Close(); err != nil {
				ansi.Red().
					Printf("Error closing database connector: %v\n", err).
					Reset().
					Fatal(exit.ConnectionFailed)
			}
			ansi.Green().Println("Terminating OK").Reset()
		}()
		if err := dbConn.ApplyMigration(filepath.Join(databasesDirectory, dbName)); err != nil {
			ansi.Red().
				Printf("error applying migration to %s: %v\n", dbName, err).
				Reset().
				Fatal(exit.GeneralError)
		}
		return nil
	}); err != nil {
		ansi.Red().
			Printf("Error: %v\n", err).
			Reset().
			Fatal(exit.GeneralError)
	}
}
