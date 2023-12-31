package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_Avatars(t *testing.T) {
	const tableName = "avatars"
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
        sqldbtest.CheckError(t, db.Close())
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		rows, err := db.Query("select id,hash,mimeType,created from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:hash,DataType:character varying,size:1024,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:mimeType,DataType:character varying,size:1024,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
		})
	})
	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "entity", "id", "id")
	})
	t.Run("verify indexes: created", func(t *testing.T) {
		columnNames := []string{
			"created",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
	t.Run("verify indexes: hash", func(t *testing.T) {
		columnNames := []string{
			"hash",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
	t.Run("verify indexes: mimeType", func(t *testing.T) {
		columnNames := []string{
			"mimeType",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
}
