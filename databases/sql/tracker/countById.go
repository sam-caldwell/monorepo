package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// countById - count records in a given table (tableName) with matching entityId
func countById(t *testing.T, db *Postgres.Db, tableName string, entityId uuid.UUID) int {
	var err error
	var count int
	var rows *sql.Rows
	t.Run("count entities by id in table '"+tableName+"';", func(t *testing.T) {
		rows, err = db.Query("select count(id) from %s where id=('%s');", tableName, entityId)
		if err != nil {
			t.Fatalf("Fail: count query failed %v teamId:  %v", err, entityId)
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
