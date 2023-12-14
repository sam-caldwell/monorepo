package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createIcon(t *testing.T) {
	const (
		functionName = "createIcon"
		testHash     = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
		testType     = "image/png"
	)

	var rows *sql.Rows
	var err error
	var entityId uuid.UUID
	var iconId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		t.Log("cleanup...")
		_ = db.Close()
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{t,h},"+
				"pt:{varchar,mimeType},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("test the create operation", func(t *testing.T) {

		t.Run("call function", func(t *testing.T) {
			rows, err = db.Query("select createIcon('%s'::mimeType,'%s');", testType, testHash)
			if err != nil {
				t.Fatalf("Fail: (query): %v", err)
			}
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("Fail: no row returned")
			}
			err = rows.Scan(&entityId)
			if err != nil {
				t.Fatalf("row scan failed: %v", err)
			}
		})

		t.Run("Verify the entityId", func(t *testing.T) {
			rows, err = db.Query("select id, type from entity where id='%s';", entityId)
			if err != nil {
				t.Fatalf("Fail: (query): %v", err)
			}
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("Fail: no row returned")
			}
			var entityType string
			if err = rows.Scan(&iconId, &entityType); err != nil {
				t.Fatal(err)
			}
			if entityId != iconId {
				t.Fatal("Fail: iconId mismatch")
			}
			if entityType != "icon" {
				t.Fatal("Fail: context mismatch")
			}
		})

		t.Run("verify the icon record", func(t *testing.T) {
			rows, err = db.Query("select id, hash, mimetype from icons where id='%s';", iconId)
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var actualId uuid.UUID
			var actualHash string
			var actualMimeType string
			if err = rows.Scan(&actualId, &actualHash, &actualMimeType); err != nil {
				t.Fatal(err)
			}
			if actualId != iconId {
				t.Fatal("entityId mismatch")
			}
			if actualHash != testHash {
				t.Fatal("actualHash mismatch")
			}
			if actualMimeType != testType {
				t.Fatal("actualMimeType mismatch")
			}
		})
	})
}
