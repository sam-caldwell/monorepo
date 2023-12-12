package sqldbtest

import (
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// GetEnumValues - Return a list of enumerated values.
func GetEnumValues(t *testing.T, db *Postgres.Db, typeName string) (values []string) {
	rows, err := db.Query("SELECT 1 as a FROM pg_type WHERE typname = '%s'AND typtype='e';", typeName)
	CheckError(t, err)
	defer func() { _ = rows.Close() }()
	rows, err = db.Query("SELECT unnest(enum_range(NULL::%s)) AS enum_value;", typeName)
	CheckError(t, err)
	defer func() {
		CheckError(t, rows.Close())
	}()
	return SqlRowsToColumns(t, rows)
}
