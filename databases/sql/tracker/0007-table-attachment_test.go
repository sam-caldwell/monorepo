package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_Attachment(t *testing.T) {
	const tableName = "attachment"

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
			"ColumnName:id,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:gen_random_uuid()",
			"ColumnName:ticketId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:authorId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:author,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'delete'::permissions",
			"ColumnName:team,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:everyone,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "ticket", "ticketId", "id")
	})
}
