package database

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestSqlDbTable_TicketTypes(t *testing.T) {
	const tableName = "tickettypes"

	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		_, err := db.Query(fmt.Sprintf("select 1 from %s limit 1;", tableName))
		database.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		expectedColumns := []string{
			"ColumnName:id,datatype:uuid,IsNullable:no,ColumnDefault:gen_random_uuid()",
			"ColumnName:name,datatype:character varying,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:iconId,datatype:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:workflowId,datatype:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:description,datatype:text,IsNullable:yes,ColumnDefault:<<null>>",
		}
		actualColumns := database.GetTableColumns(t, db, tableName)
		database.CompareTwoStringLists(t, actualColumns, expectedColumns)
	})

	t.Run("check foreign keys", func(t *testing.T) {
		database.ValidateForeignKey(t, db, "workflow", "users", "ownerId", "id")
		database.ValidateForeignKey(t, db, "workflow", "teams", "teamId", "id")
		database.ValidateForeignKey(t, db, "workflow", "icons", "iconId", "id")
	})
}
