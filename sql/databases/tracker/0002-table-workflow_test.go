package database

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestSqlDbTable_workflow(t *testing.T) {
	const tableName = "workflow"
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
			"ColumnName:id,DataType:uuid,IsNullable:NO,ColumnDefault:gen_random_uuid()",
			"ColumnName:description,DataType:text,IsNullable:yes,ColumnDefault:<<null>>",
			"ColumnName:name,DataType:character varying,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:iconId,datatype:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:ownerId,DataType:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:teamId,DataType:uuid,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:owner,datatype:user-defined,IsNullable:no,ColumnDefault:'delete'::permissions",
			"ColumnName:team,datatype:user-defined,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:everyone,datatype:user-defined,IsNullable:no,ColumnDefault:'read'::permissions",
		}
		actualColumns := database.GetTableColumns(t, db, tableName)
		database.CompareTwoStringLists(t, actualColumns, expectedColumns)
	})
}
