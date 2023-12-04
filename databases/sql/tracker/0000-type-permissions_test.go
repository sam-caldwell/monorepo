package database

import (
	"github.com/sam-caldwell/monorepo/databases/tools"
	"testing"
)

func TestSqlDbType_permissions(t *testing.T) {
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		database.CheckError(t, err)
	})
	t.Run("verify the enumerated type values", func(t *testing.T) {
		actual := database.GetEnumValues(t, db, "permissions")
		expected := []string{"none", "read", "create", "update", "delete"}
		database.CompareTwoStringLists(t, actual, expected)
	})

}
