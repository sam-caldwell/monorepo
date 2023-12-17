package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// deleteTicketTypeById - delete a new ticketType object by Id
func deleteTicketTypeById(t *testing.T, db *Postgres.Db, ticketTypeId uuid.UUID) (count int) {

	var err error
	var rows *sql.Rows

	t.Run("call deleteTicketTypeById();", func(t *testing.T) {

		rows, err = db.Query("select deleteTicketTypeById('%v');", ticketTypeId)

		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}

		defer func() { _ = rows.Close() }()

		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}

		if err = rows.Scan(&count); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
	})
	return count
}
