package database

import (
	"github.com/sam-caldwell/monorepo/databases/tools"
	"testing"
)

func TestSqlDbTable_PropertyKeys(t *testing.T) {
	const tableName = "propertykeys"
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database.database.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		_, err := db.Query("select id, name from %s limit 1;", tableName)
		database.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		database.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:gen_random_uuid()",
			"ColumnName:name,DataType:character varying,size:64,IsNullable:NO,ColumnDefault:<<null>>",
		})
	})
}
