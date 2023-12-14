package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbFunc_boundsCheckLower(t *testing.T) {

	db := sqldbtest.InitializeTestDbConn(t)

	t.Run("Happy path", func(t *testing.T) {
		_, err := db.Query("select boundsCheckLower(10,0);")
		if err != nil {
			t.Fatalf("Query error: %v", err)
		}
	})
	t.Run("Sad path (lower)", func(t *testing.T) {
		_, err := db.Query("select boundsCheckLower(-1,0);")
		if err == nil {
			t.Fatalf("Query error: %v", err)
		}
	})
}
