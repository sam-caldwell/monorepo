package psqlTrackerDb

import (
	"database/sql"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// deleteWorkflowById - delete a workflow object by id
func deleteWorkflowById(t *testing.T, db *Postgres.Db, workflowId uuid.UUID) (count int) {

	t.Run("delete workflow", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query(""+
			"select deleteWorkflowById('%s');", workflowId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		err = rows.Scan(&count)
		if err != nil {
			t.Fatal(err)
		}
	})

	return count
}
