package database

import (
	"github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestSqlDbType_logPriority(t *testing.T) {
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database.CheckError(t, err)
	})

	t.Run("verify the type exists", func(t *testing.T) {
		_, err := db.Query("SELECT 1 FROM pg_type WHERE typname = 'logPriority';")
		database.CheckError(t, err)
	})
	t.Run("verify the type is enum", func(t *testing.T) {
		_, err := db.Query("SELECT 1 FROM pg_type WHERE typname = 'logpriority' AND typtype='e';")
		database.CheckError(t, err)
	})
	t.Run("verify the enumerated type values", func(t *testing.T) {
		actual := database.GetEnumValues(t, db, "logpriority")
		expected := []string{"critical", "warn", "info", "debug"}
		database.CompareTwoStringLists(t, actual, expected)
	})

}
