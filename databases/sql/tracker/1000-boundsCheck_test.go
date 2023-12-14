package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbFunc_boundsCheck(t *testing.T) {

	db := sqldbtest.InitializeTestDbConn(t)

	t.Run("Happy path", func(t *testing.T) {
		_, err := db.Query("select boundsCheck(10,0,100);")
		if err != nil {
			t.Fatalf("Query error: %v", err)
		}
	})
	t.Run("Sad path (lower)", func(t *testing.T) {
		_, err := db.Query("select boundsCheck(-1,0,100);")
		if err == nil {
			t.Fatalf("Query error: %v", err)
		}
	})
	t.Run("Sad path (upper)", func(t *testing.T) {
		_, err := db.Query("select boundsCheck(1000,0,100);")
		if err == nil {
			t.Fatalf("Query error: %v", err)
		}
	})
}
