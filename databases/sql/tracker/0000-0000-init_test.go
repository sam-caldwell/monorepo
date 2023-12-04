package database

import (
	database2 "github.com/sam-caldwell/monorepo/databases/tools"
	"testing"
)

func TestDbConnectionForTests(t *testing.T) {
	var err error
	db := database2.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err = db.Close()
		database2.CheckError(t, err)
	})

	t.Run("verify db connection works", func(t *testing.T) {
		_, err = db.Query("select 1;")
		database2.CheckError(t, err)
	})

}
