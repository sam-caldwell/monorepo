package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// getAvatarById - Get a record, unmarshall it and return the result (passed by reference)
func getAvatarById(t *testing.T, db *Postgres.Db, entityId uuid.UUID, result *TrackerIcon) int {
	var err error
	var count int
	var rows *sql.Rows
	t.Run("call getAvatarById()", func(t *testing.T) {
		rows, err = db.Query("select getAvatarById('%s');", entityId)
		if err != nil {
			t.Fatalf("query failed %v teamId:  %v", err, entityId)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		err = rows.Scan(result)
	})
	return count
}
