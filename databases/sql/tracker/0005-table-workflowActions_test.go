package pgsqlTracker

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_WorkflowActions(t *testing.T) {
	const tableName = "workflowactions"

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("query the table", func(t *testing.T) {
		_, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:gen_random_uuid()",
			"ColumnName:workflowStepId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:name,datatype:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:topic,datatype:character varying,size:2048,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:message,datatype:character varying,size:2048,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:description,datatype:text,size:-1,IsNullable:yes,ColumnDefault:<<null>>",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "workflowSteps", "workflowStepId", "id")
	})
}
