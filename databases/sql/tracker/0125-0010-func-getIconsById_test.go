package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getIconsById(t *testing.T) {
	const (
		functionName = "getIconsById"
		tableName    = "icons"
		testIconUrl  = "http://localhost/myfakeicon.jpeg"
	)
	var err error
	var iconId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where url='%s'", tableName, testIconUrl)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{iconid},"+
				"pt:{uuid},"+
				"rt:text", strings.ToLower(functionName)))
	})

	t.Run("create iconId", func(t *testing.T) {
		var rows *sql.Rows

		if rows, err = db.Query("select createIcons('%s');", testIconUrl); err != nil {
			t.Fatal(err)
		}
        defer func() { _ = rows.Close() }()
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
        defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var url string
		err = rows.Scan(&url)
		if err = rows.Scan(&url); err != nil {
			t.Fatal(err)
		}
		if url != testIconUrl {
			t.Fatalf("url mismatch\n"+
				"Got:      %s\n"+
				"Expected: %s", url, testIconUrl)
		}
	})

	t.Run("call getIconsById()", func(t *testing.T) {
		var rows *sql.Rows

		if rows, err = db.Query("select getIconsById('%s');", iconId); err != nil {
			t.Fatal(err)
		}
        defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var iconUrl string
		err = rows.Scan(&iconUrl)
		if err = rows.Scan(&iconUrl); err != nil {
			t.Fatal(err)
		}
		if iconUrl != testIconUrl {
			t.Fatalf("url mismatch\n"+
				"Got:      %s\n"+
				"Expected: %s", iconUrl, testIconUrl)
		}
	})

}
