package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_users(t *testing.T) {
	const tableName = "users"
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		sqldbtest.CheckError(t, db.Close())
	})

	t.Run("query the table", func(t *testing.T) {
		rows, err := db.Query("select id,firstName,lastName,avatarId,created,email,phoneNumber,description "+
			"from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:<<null>>",
			"ColumnName:description,DataType:text,size:-1,IsNullable:yes,ColumnDefault:<<null>>",
			"ColumnName:email,DataType:character varying,size:256,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:firstName,DataType:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:lastName,DataType:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:avatarId,DataType:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:phonenumber,DataType:character varying,size:20,IsNullable:YES,ColumnDefault:<<null>>",
			"ColumnName:created,DataType:timestamp without time zone,size:-1,IsNullable:NO,ColumnDefault:now()",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, tableName, "avatars", "avatarId", "id")
		sqldbtest.ValidateForeignKey(t, db, tableName, "entity", "id", "id")
	})

	t.Run("verify indexes: created", func(t *testing.T) {
		columnNames := []string{
			"created",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})

	t.Run("verify indexes: email", func(t *testing.T) {
		columnNames := []string{
			"email",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})

	t.Run("verify indexes: firstName", func(t *testing.T) {
		columnNames := []string{
			"firstName",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})

	t.Run("verify indexes: lastName", func(t *testing.T) {
		columnNames := []string{
			"lastName",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, false)
	})

	t.Run("verify indexes: phoneNumber", func(t *testing.T) {
		columnNames := []string{
			"phoneNumber",
		}
		sqldbtest.ValidateIndex(t, db, tableName, columnNames, true)
	})
}
