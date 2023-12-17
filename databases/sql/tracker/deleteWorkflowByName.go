package psqlTrackerDb

import (
	"database/sql"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// deleteWorkflowByName - delete a workflow object by name
func deleteWorkflowByName(t *testing.T, db *Postgres.Db, workflowName string) (count int) {

	t.Run("delete workflow", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query(""+
			"select deleteWorkflowByName('%s');", workflowName)
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
