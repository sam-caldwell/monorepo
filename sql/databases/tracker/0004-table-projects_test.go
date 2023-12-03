package database

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestSqlDbTable_Projects(t *testing.T) {
	const tableName = "projects"

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
			"ColumnName:ownerId,datatype:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:teamId,datatype:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:owner,datatype:user-defined,IsNullable:no,ColumnDefault:'delete'::permissions",
			"ColumnName:team,datatype:user-defined,IsNullable:no,ColumnDefault:'none'::permissions",
			"ColumnName:everyone,datatype:user-defined,IsNullable:no,ColumnDefault:'none'::permissions",
			"ColumnName:defaultTicketType,datatype:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:description,datatype:text,IsNullable:yes,ColumnDefault:<<null>>",
		}
		actualColumns := database.GetTableColumns(t, db, tableName)
		database.CompareTwoStringLists(t, actualColumns, expectedColumns)
	})

	t.Run("check foreign keys", func(t *testing.T) {
		database.ValidateForeignKey(t, db, tableName, "users", "ownerId", "id")
		database.ValidateForeignKey(t, db, tableName, "teams", "teamId", "id")
		database.ValidateForeignKey(t, db, tableName, "icons", "iconId", "id")
		database.ValidateForeignKey(t, db, tableName, "ticketTypes", "defaultTicketType", "id")
	})
}
