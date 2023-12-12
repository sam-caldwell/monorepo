package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deleteIcons(t *testing.T) {
	const (
		functionName = "deleteIcons"
		tableName    = "icons"
		testUrl      = "http://localhost/myfakeicon.jpeg"
	)
	var err error
	var iconId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where url='%s'", tableName, testUrl)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{iconid},"+
				"pt:{uuid},"+
				"rt:int4", strings.ToLower(functionName)))
	})

	t.Run("create iconId", func(t *testing.T) {
		var rows *sql.Rows

		if rows, err = db.Query("select createIcons('%s');", testUrl); err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if iconId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("verify the icon record", func(t *testing.T) {
		var rows *sql.Rows

		if rows, err = db.Query("select url from icons where id='%s';", iconId); err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var url string
		err = rows.Scan(&url)
		if err = rows.Scan(&url); err != nil {
			t.Fatal(err)
		}
		if url != testUrl {
			t.Fatalf("url mismatch\n"+
				"Got:      %s\n"+
				"Expected: %s", url, testUrl)
		}
	})

	t.Run("delete the icon", func(t *testing.T) {
		var rows *sql.Rows

		if rows, err = db.Query("select deleteIcons('%s');", iconId); err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		err = rows.Scan(&count)
		if err = rows.Scan(&count); err != nil {
			t.Fatal(err)
		}
		if count != 1 {
			t.Fatalf("deleteIcons() should return 1 but returned %d", count)
		}
	})

	t.Run("count the number of matching icons (expect zero)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select count(id) from icons where id=('%s');", iconId)
		if err != nil {
			t.Fatalf("count query failed %v\n"+
				"teamId:  %v", err, iconId)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		err = rows.Scan(&count)
		if count != 0 {
			t.Fatalf("expected count 0 but got %d", count)
		}
	})
}
