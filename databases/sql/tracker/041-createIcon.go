package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// createIcon - create a new icon object
func createIcon(t *testing.T, db *Postgres.Db, mimeType, hash string) (entityId uuid.UUID) {
	var err error
	var rows *sql.Rows
	t.Run("call createIcon();", func(t *testing.T) {
		rows, err = db.Query("select createIcon('%s','%s');", mimeType, hash)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		err = rows.Scan(&entityId)
		if err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
	})
	return entityId
}
