package sqldbtest

import (
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

func ValidateTable(t *testing.T, db *Postgres.Db, tableName string, expectedColumns []string) {
	actualColumns := GetTableColumns(t, db, tableName)
	CompareTwoStringLists(t, actualColumns, expectedColumns)
}
