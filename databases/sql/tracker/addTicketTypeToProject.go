package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// addTicketTypeToProject - associate a ticket type from a project
func addTicketTypeToProject(t *testing.T, db *Postgres.Db, projectId, ticketTypeId uuid.UUID) (entityId uuid.UUID) {

	var err error
	var rows *sql.Rows

	t.Run("call createTicketType();", func(t *testing.T) {
		rows, err = db.Query("select addTicketTypeToProject('%v','%v');", projectId, ticketTypeId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		if err = rows.Scan(&entityId); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
	})
	return entityId
}
