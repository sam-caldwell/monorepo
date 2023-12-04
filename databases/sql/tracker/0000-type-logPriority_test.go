package database

import (
	"github.com/sam-caldwell/monorepo/databases/sql/tools"
	"testing"
)

func TestSqlDbType_logPriority(t *testing.T) {
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database.CheckError(t, err)
	})
	t.Run("verify the enumerated type values", func(t *testing.T) {
		actual := database.GetEnumValues(t, db, "logpriority")
		expected := []string{"critical", "warn", "info", "debug"}
		database.CompareTwoStringLists(t, actual, expected)
	})

}
