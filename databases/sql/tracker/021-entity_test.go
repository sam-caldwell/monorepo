package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_Entity(t *testing.T) {
	const tableName = "entity"

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		sqldbtest.CheckError(t, db.Close())
	})

	t.Run("verify we can query the table", func(t *testing.T) {
		rows, err := db.Query("select id, context, created from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:context,DataType:character varying,size:2048,IsNullable:NO,ColumnDefault:''::character varying",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
			"ColumnName:type,DataType:user-defined,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
		})
	})

	t.Run("verify indexes: context", func(t *testing.T) {
		columnNames := []string{
			"context",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})

	t.Run("verify indexes: created", func(t *testing.T) {
		columnNames := []string{
			"created",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})

}
