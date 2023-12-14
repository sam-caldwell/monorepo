package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbFuncPreventUpdate(t *testing.T) {
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("preventUpdate()", func(t *testing.T) {
		_, err := db.Query("select preventUpdate();")
		if err == nil {
			t.Fatal("expected error but encountered none")
		}
		if eStr := err.Error(); eStr != "pq: trigger functions can only be called as triggers" {
			t.Fatalf("unexpected error message: '%s'", eStr)
		}
	})
}
