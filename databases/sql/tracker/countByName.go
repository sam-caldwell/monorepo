package psqlTrackerDb

import (
	"database/sql"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// countByName - count records in a given table (tableName) with matching Name
func countByName(t *testing.T, db *Postgres.Db, tableName, value string) int {
	var err error
	var count int
	var rows *sql.Rows
	t.Run("count entities by id in table '"+tableName+"';", func(t *testing.T) {
		rows, err = db.Query("select count(id) from %s where name='%v';", tableName, value)
		if err != nil {
			t.Fatalf("Fail: count query failed %v\n\tvalue:  %v", err, value)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var count int
		err = rows.Scan(&count)
	})
	return count
}
