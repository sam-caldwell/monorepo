package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// getStartNode - get the start node of a workflow
func getStartNode(t *testing.T, db *Postgres.Db, workflowId uuid.UUID) (entityId uuid.UUID) {
	var rows *sql.Rows
	var err error

	t.Run("getStartNode() and verify", func(t *testing.T) {

		rows, err = db.Query("select getStartNode('%s');", workflowId)
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
