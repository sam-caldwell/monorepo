package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// deleteAvatar - delete a new avatar object
func deleteAvatar(t *testing.T, db *Postgres.Db, entityId uuid.UUID) int {
	var err error
	var count int
	var rows *sql.Rows
	t.Run("call deleteAvatar();", func(t *testing.T) {
		rows, err = db.Query("select deleteAvatar('%s');", entityId)
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
