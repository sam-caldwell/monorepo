package database

import (
	"github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestSqlDbTable_Icons(t *testing.T) {
	const tableName = "icons"
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database.CheckError(t, err)
	})

	t.Run("query the avatars table", func(t *testing.T) {
		_, err := db.Query("select id, url from avatars limit 1;")
		database.CheckError(t, err)
	})

	t.Run("check avatars table schema", func(t *testing.T) {
		expectedColumns := []database.Record{
			{"id", "uuid", "NO", "gen_random_uuid()"},
			{"url", "text", "YES", "<<null>>"},
		}
		actualColumns := database.GetTableColumns(t, db, tableName)
		database.CompareTableColumns(t, actualColumns, expectedColumns)
	})
}
