package pgsqlTracker

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_TeamMembership(t *testing.T) {
	const tableName = "teammembership"

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		_, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:userId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:teamId,DataType:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, "workflow", "users", "ownerId", "id")
		sqldbtest.ValidateForeignKey(t, db, "workflow", "teams", "teamId", "id")
		sqldbtest.ValidateForeignKey(t, db, "workflow", "icons", "iconId", "id")
	})
}
