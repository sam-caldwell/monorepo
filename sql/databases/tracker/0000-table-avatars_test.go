package database

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestSqlDbTable_Avatars(t *testing.T) {
	const tableName = "avatars"
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		_, err := db.Query(fmt.Sprintf("select id, url from %s limit 1;", tableName))
		database.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		database.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:gen_random_uuid()",
			"ColumnName:url,DataType:text,size:-1,IsNullable:YES,ColumnDefault:<<null>>",
		})
	})
}
