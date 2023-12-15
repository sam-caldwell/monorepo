package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// createUser - create a new user object
func createUser(t *testing.T, db *Postgres.Db, firstName, lastName string,
	avatarId uuid.UUID, email, phoneNumber, description string) (entityId uuid.UUID) {

	var err error

	var rows *sql.Rows

	t.Run("call createUser();", func(t *testing.T) {

		rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
			firstName, lastName, avatarId, email, phoneNumber, description)

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
