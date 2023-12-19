package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// createWorkflow - create a workflow object for testing
func createWorkflow(t *testing.T, db *Postgres.Db, name string, iconId, ownerId, teamId uuid.UUID,
	pOwner, pTeam, pEveryone, description string) (workflowId uuid.UUID) {

	t.Run("createWorkflow()", func(t *testing.T) {

		var err error
		var rows *sql.Rows

		rows, err = db.Query("select createWorkflow("+
			"'%s'::varchar,"+
			"'%s'::uuid,'%s'::uuid,'%s'::uuid,"+
			"'%s'::permissions,'%s'::permissions,'%s'::permissions,"+
			"'%s'::text);",
			name, iconId, ownerId, teamId, pOwner, pTeam, pEveryone, description)

		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}

		defer func() { _ = rows.Close() }()

		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}

		if err = rows.Scan(&workflowId); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}

	})

	return workflowId

}
