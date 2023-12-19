package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// createWorkflowStep - create a new workflow step object in the database for testing
func createWorkflowStep(t *testing.T, db *Postgres.Db, workflowId uuid.UUID, stepName string,
	prevStepId, nextStepId uuid.UUID, description string) (entityId uuid.UUID) {

	t.Run("createWorkflow()", func(t *testing.T) {
		var err error
		var rows *sql.Rows
		rows, err = db.Query("select createWorkflowStep("+
			"'%s'::varchar,"+
			"'%s'::uuid,"+
			"'%s'::uuid,"+
			"'%s'::uuid,"+
			"'%s'::text"+
			");",
			stepName, workflowId, prevStepId, nextStepId, description)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		if err = rows.Scan(&entityId); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
	})

	return entityId

}
