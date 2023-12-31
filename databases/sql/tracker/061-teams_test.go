package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_teams(t *testing.T) {
	const tableName = "teams"

	db := sqldbtest.InitializeTestDbConn(t)

    t.Cleanup(func() {
        sqldbtest.CheckError(t, db.Close())
    })

	t.Run("query the table", func(t *testing.T) {
		rows, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:description,DataType:text,size:-1,IsNullable:yes,ColumnDefault:<<null>>",
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:name,DataType:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:iconId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:ownerId,DataType:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:owner,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'delete'::permissions",
			"ColumnName:team,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:everyone,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'read'::permissions",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, "teams", "users", "ownerId", "id")
		sqldbtest.ValidateForeignKey(t, db, "teams", "icons", "iconId", "id")
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
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
	t.Run("verify indexes: ownerId", func(t *testing.T) {
		columnNames := []string{
			"ownerId",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
	t.Run("verify indexes: iconId", func(t *testing.T) {
		columnNames := []string{
			"iconId",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
	t.Run("verify indexes: owner", func(t *testing.T) {
		columnNames := []string{
			"owner",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
	t.Run("verify indexes: team", func(t *testing.T) {
		columnNames := []string{
			"team",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
	t.Run("verify indexes: everyone", func(t *testing.T) {
		columnNames := []string{
			"everyone",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})
}
