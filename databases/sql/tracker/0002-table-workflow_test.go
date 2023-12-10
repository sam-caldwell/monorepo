package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_workflow(t *testing.T) {
	const tableName = "workflow"

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("query the table", func(t *testing.T) {
		_, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
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
		sqldbtest.ValidateForeignKey(t, db, "workflow", "users", "ownerId", "id")
		sqldbtest.ValidateForeignKey(t, db, "workflow", "teams", "teamId", "id")
		sqldbtest.ValidateForeignKey(t, db, "workflow", "icons", "iconId", "id")
	})
}
