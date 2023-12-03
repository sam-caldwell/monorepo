package database

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestSqlDbTable_Membership(t *testing.T) {
	const tableName = "teammembership"

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
			"ColumnName:userId,datatype:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:teamId,DataType:uuid,IsNullable:no,ColumnDefault:<<null>>",
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
