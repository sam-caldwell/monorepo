package psqlTrackerDb

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// getAvatarById - Get a record, unmarshall it and return the result (passed by reference)
func getAvatarById(t *testing.T, db *Postgres.Db, entityId uuid.UUID, result *TrackerAvatar) {
	var err error
	var rows *sql.Rows
	var raw sql.NullString
	t.Run("call getAvatarById()", func(t *testing.T) {
		t.Logf("calling getAvatarById('%s');", entityId)
		rows, err = db.Query("select getAvatarById('%s');", entityId)
		if err != nil {
			t.Fatalf("query failed %v id:  %v", err, entityId)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		err = rows.Scan(&raw)
		if err != nil {
			t.Fatalf("Fail: error scanning result. %v", err)
		}
		if raw.Valid {
			if err = json.Unmarshal([]byte(raw.String), result); err != nil {
				t.Fatalf("Fail: could not unmarshal result. %v", err)
			}
		} else {
			t.Fatalf("Fail: expected record but got NULL")
		}
	})
}
