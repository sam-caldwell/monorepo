package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// createTeam - create a new team object
func createTeam(t *testing.T, db *Postgres.Db, teamName string, iconId, ownerId uuid.UUID,
	description string) (entityId uuid.UUID) {

	var err error
	var rows *sql.Rows

	t.Run("call createUser();", func(t *testing.T) {

		rows, err = db.Query("select createTeam('%s','%s','%s','read','read','read','%s');",
			teamName, iconId, ownerId, description)

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
