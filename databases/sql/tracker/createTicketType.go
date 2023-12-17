package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// createTicketType - create a new ticketType object
func createTicketType(t *testing.T, db *Postgres.Db, name string, iconId, workflowId uuid.UUID,
	description string) (entityId uuid.UUID) {

	var err error
	var rows *sql.Rows

	t.Run("call createTicketType();", func(t *testing.T) {

		rows, err = db.Query("select createTicketType('%v','%v','%v','%v');",
			name, iconId, workflowId, description)

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
