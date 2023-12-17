package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_PropertyKeys(t *testing.T) {
	const tableName = "propertykeys"
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		sqldbtest.CheckError(t, db.Close())
	})

	t.Run("query the table", func(t *testing.T) {
		rows, err := db.Query("select id, name from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:name,DataType:character varying,size:64,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
		})
	})
	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "entity", "id", "id")
	})
	t.Run("verify indexes: name", func(t *testing.T) {
		columnNames := []string{
			"name",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
	t.Run("verify indexes: created", func(t *testing.T) {
		columnNames := []string{
			"created",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
}
