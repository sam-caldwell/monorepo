package psqlTrackerDb

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// addTicketTypeToProject - associate a ticket type from a project
func getProjectTicketTypes(t *testing.T, db *Postgres.Db, projectId uuid.UUID) (ticketTypes []TrackerTicketType) {

	var err error
	var rows *sql.Rows

	t.Run("call getProjectTicketTypes();", func(t *testing.T) {
		rows, err = db.Query("select getProjectTicketTypes('%v');", projectId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatalf("no rows returned")
		}
		var data string
		if err = rows.Scan(&data); err != nil {
			t.Fatalf("Fail: error scanning record: %v", err)
		}
		if err = json.Unmarshal([]byte(data), &ticketTypes); err != nil {
			t.Fatalf("Fail: failed to unmarshall json.\n"+
				"error: %v\n"+
				"data:  %v", err, data)
		}
		//for rows.Next() {
		//	var data string
		//	if err = rows.Scan(&data); err != nil {
		//		t.Fatalf("Fail: error scanning record. Err: %v", err)
		//	}
		//	var tt TrackerTicketType
		//	if err = json.Unmarshal([]byte(data), &tt); err != nil {
		//		t.Fatalf("Fail: error unmarshaling record. Err: %v", err)
		//	}
		//	ticketTypes = append(ticketTypes, tt)
		//}
	})
	return ticketTypes
}
