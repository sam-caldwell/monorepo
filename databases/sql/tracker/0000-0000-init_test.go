package sqldbtest

import (
    "github.com/sam-caldwell/monorepo/go/db/sqldbtest"
    "testing"
)

func TestDbConnectionForTests(t *testing.T) {
	var err error
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err = db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify db connection works", func(t *testing.T) {
		_, err = db.Query("select 1;")
		sqldbtest.CheckError(t, err)
	})

}
