package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// createProject - create a new project object
func createProject(t *testing.T, db *Postgres.Db, name string, iconId, ownerId, teamId uuid.UUID,
	pOwner, pTeam, pEveryone, description string) (entityId uuid.UUID) {

	var err error
	var rows *sql.Rows

	t.Run("call createProject();", func(t *testing.T) {

		rows, err = db.Query("select createProject('%v','%v','%v','%v','%v','%v','%v','%v');",
			name, iconId, ownerId, teamId, pOwner, pTeam, pEveryone, description)

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
