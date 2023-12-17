package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// removeUserFromTeam - remove a user from a given team
func removeUserFromTeam(t *testing.T, db *Postgres.Db, userId, teamId uuid.UUID) int {
	var rows *sql.Rows
	var count int
	var err error

	t.Run("removeUserFromTeam()", func(t *testing.T) {

		rows, err = db.Query("select removeUserFromTeam('%s','%s');", userId, teamId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()

		if !rows.Next() {
			t.Fatal("Fail: expect at least one row")
		}
		if err = rows.Scan(&count); err != nil {
			t.Fatal(err)
		}
	})
	return count
}
