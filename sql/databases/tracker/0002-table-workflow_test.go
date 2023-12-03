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
		database.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:gen_random_uuid()",
			"ColumnName:name,DataType:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:iconId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:ownerId,DataType:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:teamId,DataType:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:owner,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'delete'::permissions",
			"ColumnName:team,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:everyone,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:description,DataType:text,size:-1,IsNullable:yes,ColumnDefault:<<null>>",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		database.ValidateForeignKey(t, db, "workflow", "users", "ownerId", "id")
		database.ValidateForeignKey(t, db, "workflow", "teams", "teamId", "id")
		database.ValidateForeignKey(t, db, "workflow", "icons", "iconId", "id")
	})
}
