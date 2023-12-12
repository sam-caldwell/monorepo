package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbTable_users(t *testing.T) {
	const tableName = "users"
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		rows, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
		defer func() { _ = rows.Close() }()
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,DataType:uuid,size:-1,IsNullable:NO,ColumnDefault:gen_random_uuid()",
			"ColumnName:description,DataType:text,size:-1,IsNullable:yes,ColumnDefault:<<null>>",
			"ColumnName:email,DataType:character varying,size:256,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:firstName,DataType:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:lastName,DataType:character varying,size:64,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:avatarId,DataType:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:phonenumber,DataType:character varying,size:20,IsNullable:YES,ColumnDefault:<<null>>",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, "users", "avatars", "avatarId", "id")
	})
}
