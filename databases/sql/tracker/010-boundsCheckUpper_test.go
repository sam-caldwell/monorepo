package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbFunc_boundsCheckUpper(t *testing.T) {

	db := sqldbtest.InitializeTestDbConn(t)

	t.Run("Happy path", func(t *testing.T) {
		_, err := db.Query("select boundsCheckUpper(0,1);")
		if err != nil {
			t.Fatalf("Query error: %v", err)
		}
	})
	t.Run("Sad path (upper)", func(t *testing.T) {
		_, err := db.Query("select boundsCheckUpper(1,0);")
		if err == nil {
			t.Fatalf("Query error: %v", err)
		}
		if msg := err.Error(); msg != "pq: upper bounds check error" {
			t.Fatalf("error message mismatch: %s", msg)
		}
	})
}
