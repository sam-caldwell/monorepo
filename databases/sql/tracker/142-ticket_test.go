package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_Ticket(t *testing.T) {
	const tableName = "ticket"

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
			"ColumnName:projectId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:authorUserId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:assignedUserId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:ticketTypeId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:owner,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'delete'::permissions",
			"ColumnName:team,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:everyone,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:summary,datatype:character varying,size:2048,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:workflowStepId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:description,datatype:text,size:-1,IsNullable:yes,ColumnDefault:<<null>>",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "projects", "projectId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "users", "authorUserId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "users", "assignedUserId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "ticketTypes", "ticketTypeId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "workflowSteps", "workflowStepId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "entity", "id", "id")
	})
	t.Run("verify indexes: summary", func(t *testing.T) {
		columnNames := []string{
			"summary",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
	t.Run("verify indexes: created", func(t *testing.T) {
		columnNames := []string{
			"created",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
	t.Run("verify indexes: owner", func(t *testing.T) {
		columnNames := []string{
			"owner",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
	t.Run("verify indexes: team", func(t *testing.T) {
		columnNames := []string{
			"team",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
	t.Run("verify indexes: everyone", func(t *testing.T) {
		columnNames := []string{
			"everyone",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
}
