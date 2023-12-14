package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbType_permissions(t *testing.T) {
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})
	t.Run("verify the enumerated type values", func(t *testing.T) {
		actual := sqldbtest.GetEnumValues(t, db, "permissions")
		expected := []string{"none", "read", "create", "update", "delete"}
		sqldbtest.CompareTwoStringLists(t, actual, expected)
	})

}
