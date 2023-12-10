package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_numericProperties(t *testing.T) {
	const tableName = "numericproperties"
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
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:value,DataType:integer,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
		})
	})
	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "propertyKeys", "id", "id")
	})
}
