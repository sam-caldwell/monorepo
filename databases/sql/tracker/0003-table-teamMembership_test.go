package database

import (
	"github.com/sam-caldwell/monorepo/databases/tools"
	"testing"
)

func TestSqlDbTable_TeamMembership(t *testing.T) {
	const tableName = "teammembership"

	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		_, err := db.Query("select 1 from %s limit 1;", tableName)
		database.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		database.ValidateTable(t, db, tableName, []string{
			"ColumnName:userId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:teamId,DataType:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		database.ValidateForeignKey(t, db, "workflow", "users", "ownerId", "id")
		database.ValidateForeignKey(t, db, "workflow", "teams", "teamId", "id")
		database.ValidateForeignKey(t, db, "workflow", "icons", "iconId", "id")
	})
}
