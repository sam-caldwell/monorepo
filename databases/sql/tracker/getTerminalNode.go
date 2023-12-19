package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// getTerminalNode - get the terminal node of a workflow
func getTerminalNode(t *testing.T, db *Postgres.Db, workflowId uuid.UUID) (entityId uuid.UUID) {
	var rows *sql.Rows
	var err error

	t.Run("getTerminalNode() and verify", func(t *testing.T) {

		rows, err = db.Query("select getTerminalNode('%s');", workflowId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()

		if !rows.Next() {
			t.Fatalf("no rows returned")
		}
		if err = rows.Scan(&entityId); err != nil {
			t.Fatal(err)
		}
	})
	return entityId
}
