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
		rows, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:ticketId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:authorId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:author,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'delete'::permissions",
			"ColumnName:team,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:everyone,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "ticket", "ticketId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "users", "authorId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "entity", "id", "id")
	})
	t.Run("verify indexes: author", func(t *testing.T) {
		columnNames := []string{
			"author",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames)
	})
	t.Run("verify indexes: team", func(t *testing.T) {
		columnNames := []string{
			"team",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames)
	})
	t.Run("verify indexes: everyone", func(t *testing.T) {
		columnNames := []string{
			"everyone",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames)
	})
}
