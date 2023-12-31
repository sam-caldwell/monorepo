package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// deleteIcon - delete a new icon object
func deleteIcon(t *testing.T, db *Postgres.Db, entityId uuid.UUID) int {
	var err error
	var count int
	var rows *sql.Rows
	t.Run("call deleteIcon();", func(t *testing.T) {
		rows, err = db.Query("select deleteIcon('%s');", entityId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		err = rows.Scan(&count)
		if err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
		if count != 1 {
			t.Fatalf("Fail: count expects 1 but got %d", count)
		}
	})
	return count
}
