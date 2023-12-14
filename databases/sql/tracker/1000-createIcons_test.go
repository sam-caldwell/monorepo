package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createIcons(t *testing.T) {
	const (
		functionName = "createIcons"
		tableName    = "icons"
		testHash     = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
		testType     = "image/png"
	)
	var iconId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		t.Log("cleanup...")
		_, _ = db.Query("delete from %s where hash='%s'", tableName, testHash)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{iconHash,iconType},"+
				"pt:{varchar,entitytype},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("test the create operation", func(t *testing.T) {

		var rows *sql.Rows
		var err error

		var actualId uuid.UUID
		var actualHash string
		var actualMimeType string

		t.Run("create iconId", func(t *testing.T) {
			if rows, err = db.Query("select createIcons('%s','%s');", testHash, testType); err != nil {
				t.Fatal(err)
			}
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			err = rows.Scan(&iconId)
			sqldbtest.CheckError(t, err)
		})

		t.Run("Verify the entityId", func(t *testing.T) {
			rows, err = db.Query("select id from entity where id='icon';")
			sqldbtest.CheckError(t, err)
			defer func() { _ = rows.Close() }()
			if err = rows.Scan(&actualId); err != nil {
				t.Fatal(err)
			}
			if actualId != iconId {
				t.Fatal("actualId mismatch")
			}
		})

		t.Run("verify the icon record", func(t *testing.T) {
			var thisContext string
			rows, err = db.Query("select id, hash, mimetype from icons where id='%s';", iconId)
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			if err = rows.Scan(&actualId, &actualHash, &actualMimeType); err != nil {
				t.Fatal(err)
			}
			if thisContext != "icon" {
				t.Fatal("context mismatch")
			}
		})
	})
}
