package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_WorkFlowSteps(t *testing.T) {
	const tableName = "workflowSteps"

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("query the table", func(t *testing.T) {
		rows, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:workflowId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:name,datatype:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:iconId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:prevStepId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:nextStepId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:description,datatype:text,size:-1,IsNullable:yes,ColumnDefault:<<null>>",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "workflow", "workflowId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "workflowSteps", "prevStepId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "workflowSteps", "nextStepId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "icons", "iconId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "entity", "id", "id")
	})
	t.Run("verify indexes: name", func(t *testing.T) {
		columnNames := []string{
			"name",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames)
	})
	t.Run("verify indexes: iconId", func(t *testing.T) {
		columnNames := []string{
			"iconId",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames)
	})
	t.Run("verify indexes: prevStepId", func(t *testing.T) {
		columnNames := []string{
			"prevStepId",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames)
	})
	t.Run("verify indexes: nextStepId", func(t *testing.T) {
		columnNames := []string{
			"nextStepId",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames)
	})
	t.Run("verify indexes: created", func(t *testing.T) {
		columnNames := []string{
			"created",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames)
	})
}
