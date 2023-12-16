package psqlTrackerDb

import (
	"database/sql"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// deleteByField - delete records in a given table (tableName) with matching field and value
func deleteByField(t *testing.T, db *Postgres.Db, tableName, fieldName, value any) int {
	var err error
	var count int
	var rows *sql.Rows
	t.Run("delete by field and value", func(t *testing.T) {
		rows, err = db.Query("delete from %s where %s=('%s');", tableName, fieldName, value)
		if err != nil {
			t.Fatalf("delete query failed %v field (%s) value (%v)", err, fieldName, value)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		err = rows.Scan(&count)
	})
	return count
}
