package database

import (
	database "github.com/sam-caldwell/monorepo/sql/tools"
	"testing"
)

func TestDbConnectionForTests(t *testing.T) {
	var err error
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err = db.Close()
		database.CheckError(t, err)
	})

	t.Run("verify db connection works", func(t *testing.T) {
		_, err = db.Query("select 1;")
		database.CheckError(t, err)
	})

}
