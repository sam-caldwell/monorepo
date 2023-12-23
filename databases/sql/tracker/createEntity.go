package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// createEntity - create a new entity object
func createEntity(t *testing.T, db *Postgres.Db, entityType, context string) (entityId uuid.UUID) {
	var err error
	var rows *sql.Rows
	t.Run("call createEntity();", func(t *testing.T) {
		t.Logf("Create entity type:'%v' context:'%v'", entityType, context)
		rows, err = db.Query("select createEntity('%s'::entityType, '%s');", entityType, context)
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
		t.Logf("create entityId: %v", entityId)
	})
	return entityId
}
