package database

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestSqlDbTable_numericProperties(t *testing.T) {
	const tableName = "numericproperties"
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
			"ColumnName:id,DataType:uuid,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:value,DataType:integer,IsNullable:NO,ColumnDefault:<<null>>",
		}
		actualColumns := database.GetTableColumns(t, db, tableName)
		database.CompareTwoStringLists(t, actualColumns, expectedColumns)
	})
	t.Run("check foreign keys", func(t *testing.T) {
		database.ValidateForeignKey(t, db, tableName, "propertyKeys", "id", "id")
	})
}
