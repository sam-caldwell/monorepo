package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// updateTeamName - update the team name fields
func updatePermissions(t *testing.T, db *Postgres.Db, teamId uuid.UUID, po, pt, pe string) int {

	var err error
	var rows *sql.Rows
	var count int

	t.Run("update record", func(t *testing.T) {
		rows, err = db.Query("select updatePermissions('%s','%s','%s','%s')", teamId, po, pt, pe)
		if err != nil {
			t.Fatalf("Fail: Query error: %v", err)
		}
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		if err = rows.Scan(&count); err != nil {
			t.Fatal(err)
		}
	})
	return count
}
