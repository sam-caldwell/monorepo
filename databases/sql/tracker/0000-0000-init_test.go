package pgsqlTracker

import (
	"database/sql"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestDbConnectionForTests(t *testing.T) {
	var err error
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("Error when cleaning up\nerr:%v", err)
		}
	})

	t.Run("verify db connection works", func(t *testing.T) {
		var rows *sql.Rows
		rows, err = db.Query("select 1;")
		if err != nil {
			t.Fatalf("unexpected response\nerr:%v", err)
		}
		defer func() { _ = rows.Close() }()
		rows.Next()
	})

}
