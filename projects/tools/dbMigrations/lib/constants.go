package dbMigrations

const (
	CliRoot   = "root_dir"
	CliDbHost = "db_host"
	CliDbPort = "db_port"
	CliDbUser = "db_user"
	CliDbPass = "db_pass"
	CliDbTls  = "use_tls"

	CliUsageRootDir  = "root directory for SQL projects"
	CliUsageDbHost   = "database host"
	CliUsageDbPort   = "database port"
	CliUsageDbUser   = "database user"
	CliUsageDbPass   = "database password"
	CliUsageDbUseTls = "Use TLS"

	DefaultPgHost = "localhost"
	DefaultPgPort = 5432
	DefaultPgUser = "postgres"
	DefaultPgPass = "password"

	EnvDbHost = "db_host"
	EnvDbPort = "db_port"
	EnvDbUser = "db_user"
	EnvDbPass = "db_pass"

	DefaultDb = "postgres"

	SqlDirectoryPath = "databases/sql"
)
