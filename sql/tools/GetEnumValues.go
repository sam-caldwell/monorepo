package database

import (
	"fmt"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// GetEnumValues - Return a list of enumerated values.
func GetEnumValues(t *testing.T, db *Postgres.Db, typeName string) (values []string) {
	_, err := db.Query(
		fmt.Sprintf(""+
			"SELECT 1 "+
			"FROM pg_type "+
			"WHERE typname = '%s'AND typtype='e';",
			typeName))
	CheckError(t, err)
	rows, err := db.Query(fmt.Sprintf(""+
		"SELECT unnest(enum_range(NULL::%s)) AS enum_value;", typeName))
	CheckError(t, err)
	defer func() {
		CheckError(t, rows.Close())
	}()
	return SqlRowsToColumns(t, rows)
}
