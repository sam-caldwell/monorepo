package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbType_logPriority(t *testing.T) {
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})
	t.Run("verify the enumerated type values", func(t *testing.T) {
		actual := sqldbtest.GetEnumValues(t, db, "logpriority")
		expected := []string{"critical", "warn", "info", "debug"}
		sqldbtest.CompareTwoStringLists(t, actual, expected)
	})

}
