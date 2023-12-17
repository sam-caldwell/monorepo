package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_stringProperties(t *testing.T) {
	const tableName = "stringproperties"
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		sqldbtest.CheckError(t, db.Close())
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		rows, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:value,DataType:text,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "propertyKeys", "id", "id")
	})
}
