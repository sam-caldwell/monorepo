package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_TicketTypes(t *testing.T) {
	const tableName = "tickettypes"

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		rows, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:name,datatype:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:iconId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:workflowId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:description,datatype:text,size:-1,IsNullable:yes,ColumnDefault:<<null>>",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "icons", "iconId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "workflow", "workflowId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "entity", "id", "id")
	})
	t.Run("verify indexes: name", func(t *testing.T) {
		columnNames := []string{
			"name",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
	t.Run("verify indexes: created", func(t *testing.T) {
		columnNames := []string{
			"created",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
}
