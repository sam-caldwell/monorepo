package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createAvatar(t *testing.T) {
	const (
		functionName = "createAvatar"
		testHash     = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
		testType     = "image/png"
	)
	var rows *sql.Rows
	var err error
	var entityId uuid.UUID
	//var avatarId uuid.UUID

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
			rows, err = db.Query("select createAvatar('%s'::mimeType,'%s');", testType, testHash)
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
			rows, err = db.Query("select id, hash, mimetype from avatars where id='%s';", entityId)
			if err != nil {
				t.Fatalf("Fail: (query): %v", err)
			}
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("Fail: no row returned")
			}

			var actualId uuid.UUID
			var actualHash string
			var actualMimeType string
			if err = rows.Scan(&actualId, &actualHash, &actualMimeType); err != nil {
				t.Fatal(err)
			}
			if actualId != entityId {
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
