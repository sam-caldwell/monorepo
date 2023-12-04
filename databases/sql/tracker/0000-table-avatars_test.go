package database

import (
	database2 "github.com/sam-caldwell/monorepo/databases/tools"
	"testing"
)

func TestSqlDbTable_Avatars(t *testing.T) {
	const tableName = "avatars"
	db := database2.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database2.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		_, err := db.Query("select id, url from %s limit 1;", tableName)
		database2.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		database2.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:gen_random_uuid()",
			"ColumnName:url,DataType:text,size:-1,IsNullable:YES,ColumnDefault:<<null>>",
		})
	})
}
