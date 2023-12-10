package main

import (
	"flag"
	_ "github.com/lib/pq"
	"github.com/sam-caldwell/monorepo/go/ansi"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	env "github.com/sam-caldwell/monorepo/go/environment"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"os"
	"path/filepath"
)

const (
	cliRoot   = "root_dir"
	cliDbHost = "db_host"
	cliDbPort = "db_port"
	cliDbUser = "db_user"
	cliDbPass = "db_pass"
	cliDbTls  = "use_tls"

	cliUsageRootDir  = "root directory for SQL projects"
	cliUsageDbHost   = "database host"
	cliUsageDbPort   = "database port"
	cliUsageDbUser   = "database user"
	cliUsageDbPass   = "database password"
	cliUsageDbUseTls = "Use TLS"

	defaultPgHost = "localhost"
	defaultPgPort = 5432
	defaultPgUser = "postgres"
	defaultPgPass = "password"

	envDbHost = "pg_db_host"
	envDbPort = "pg_db_port"
	envDbUser = "pg_db_user"
	envDbPass = "pg_db_pass"

	defaultDatabase = "postgres"

	sqlDirectoryPath = "databases/sql"
)

func main() {
	var databasesDirectory string
	/*
	 * Parse command-line arguments with env variables as backup and defaults thereafter.
	 */
	rootDirectory := flag.String(cliRoot, directory.GetCurrent(), cliUsageRootDir)
	pgDbHost := flag.String(cliDbHost, env.GetStringOrDefault(envDbHost, defaultPgHost), cliUsageDbHost)
	pgDbPort := flag.Uint(cliDbPort, env.GetUintOrDefault(envDbPort, defaultPgPort), cliUsageDbPort)
	pgDbUser := flag.String(cliDbUser, env.GetStringOrDefault(envDbUser, defaultPgUser), cliUsageDbUser)
	pgDbPass := flag.String(cliDbPass, env.GetStringOrDefault(envDbPass, defaultPgPass), cliUsageDbPass)
	pgDbUseTls := flag.Bool(cliDbTls, false, cliUsageDbUseTls)
	flag.Parse()

	ansi.Blue().
		Printf("%s : %s\n", cliRoot, *rootDirectory).
		Printf("%s : %v\n", cliDbHost, *pgDbHost).
		Printf("%s : %v\n", cliDbPort, *pgDbPort).
		Printf("%s : %v\n", cliDbUser, *pgDbUser).
		Printf("%s : %v\n", cliDbPass, *pgDbPass).
		Printf("%s : %v\n", cliDbTls, *pgDbUseTls).
		LF().
		Reset()

	/*
	 * Validate the database directory exists
	 */
	databasesDirectory = filepath.Join(*rootDirectory, sqlDirectoryPath)
	if !directory.Exists(databasesDirectory) {
		ansi.Red().
			Time().
			Printf("directory does not exist (%s)\n", databasesDirectory).
			Reset().
			Fatal(exit.NotFound)
	}
	ansi.Green().Time().Printf("Databases Directory confirmed (%s)\n", databasesDirectory).Reset()

	/*
	 * Establish a default database connection (postgres)
	 */
	postgresDb, err := Postgres.NewDbConnection(*pgDbHost, *pgDbPort, defaultDatabase, *pgDbUser, *pgDbPass, *pgDbUseTls)
	if err != nil {
		ansi.Red().
			Time().
			Printf("Error connecting to database (postgres): %v\n", err).
			Reset().
			Fatal(exit.ConnectionFailed)
	}
	ansi.Green().Time().Println("connection to postgres is established").Reset()

	defer func() {
		if err := postgresDb.Close(); err != nil {
			ansi.Red().
				Time().
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
			if err := postgresDb.CreateDatabase(dbName, *pgDbUser); err != nil {
				ansi.Red().
					Time().
					Printf("Error creating database: %v\n", err).
					Reset().
					Fatal(exit.ConnectionFailed)
			}
		}

		dbConn, err := Postgres.NewDbConnection(*pgDbHost, *pgDbPort, dbName, *pgDbUser, *pgDbPass, *pgDbUseTls)
		if err != nil {
			ansi.Red().
				Time().
				Printf("Error connecting to database (%s): %v\n", dbName, err).
				Reset().
				Fatal(exit.ConnectionFailed)
		}
		defer func() {
			if err := dbConn.Close(); err != nil {
				ansi.Red().
					Time().
					Printf("Error closing database connector: %v\n", err).
					Reset().
					Fatal(exit.ConnectionFailed)
			}
			ansi.Green().Time().Println("Terminating OK").Reset()
		}()
		if err := dbConn.ApplyMigration(filepath.Join(databasesDirectory, dbName)); err != nil {
			ansi.Red().
				Time().
				Printf("error applying migration to %s: %v\n", dbName, err).
				Reset().
				Fatal(exit.GeneralError)
		}
		return nil
	}); err != nil {
		ansi.Red().
			Time().
			Printf("Error: %v\n", err).
			Reset().
			Fatal(exit.GeneralError)
	}
}
