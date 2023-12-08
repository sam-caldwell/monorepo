package sqldbtest

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

const (
	valueSerial = "nextval('logs_id_seq'::regclass)"
)

func TestSqlDbTable_Logs(t *testing.T) {
	const tableName = "logs"

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("query the table (verifies permissions of user and existence of table)", func(t *testing.T) {
		_, err := db.Query("select 1 from %s limit 1;", tableName)
		sqldbtest.CheckError(t, err)
	})

	t.Run("check table schema", func(t *testing.T) {
		sqldbtest.ValidateTable(t, db, tableName, []string{
			"ColumnName:id,datatype:integer,size:-1,IsNullable:no,ColumnDefault:" + valueSerial,
			"ColumnName:priority,datatype:user-defined,size:-1,IsNullable:no,ColumnDefault:'info'::logpriority",
			"ColumnName:userId,datatype:uuid,size:-1,IsNullable:no,ColumnDefault:<<null>>",
			"ColumnName:message,datatype:character varying,size:2048,IsNullable:no,ColumnDefault:<<null>>",
		})
	})

	t.Run("check foreign keys", func(t *testing.T) {
		sqldbtest.ValidateForeignKey(t, db, "workflow", "users", "ownerId", "id")
		sqldbtest.ValidateForeignKey(t, db, "workflow", "teams", "teamId", "id")
		sqldbtest.ValidateForeignKey(t, db, "workflow", "icons", "iconId", "id")
	})
}
