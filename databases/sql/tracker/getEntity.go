package psqlTrackerDb

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// getEntity - Get a record, unmarshall it and return the result (passed by reference)
func getEntity(t *testing.T, db *Postgres.Db, entityId uuid.UUID, result *TrackerEntity) {
	var err error
	var rows *sql.Rows
	t.Run("call getEntity()", func(t *testing.T) {
		rows, err = db.Query("select getEntity('%v');", entityId)
		if err != nil {
			t.Fatalf("query failed %v id:  %v", err, entityId)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var raw sql.NullString
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
