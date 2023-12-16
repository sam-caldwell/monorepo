package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// updateTeamName - update the team name fields
func updateTeamName(t *testing.T, db *Postgres.Db, teamId uuid.UUID, teamName string) int {

	var err error
	var rows *sql.Rows
	var count int

	t.Run("update record", func(t *testing.T) {
		rows, err = db.Query("select updateTeamName('%s','%s')", teamId, teamName)
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
