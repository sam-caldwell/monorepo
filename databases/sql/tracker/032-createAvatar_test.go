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
			rows, err = db.Query("select id, type from entity where id='%s';", entityId)
			if err != nil {
				t.Fatalf("Fail: (query): %v", err)
			}
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("Fail: no row returned")
			}

			var entityType string
			if err = rows.Scan(&entityId, &entityType); err != nil {
				t.Fatal(err)
			}
			if entityId != entityId {
				t.Fatal("Fail: actualId mismatch")
			}
			if entityType != "avatar" {
				t.Fatal("Fail: context mismatch")
			}
		})

		//t.Run("Verify the avatar record", func(t *testing.T) {
		//  var actualHash string
		//  var actualMimeType string
		//	rows, err = db.Query(""+
		//		"select id, hash, mimeType"+
		//		"from avatars "+
		//		"where hash='%s';", testHash)
		//	sqldbtest.CheckError(t, err)
		//	defer func() { _ = rows.Close() }()
		//	if !rows.Next() {
		//		t.Fatal("Fail: no rows returned")
		//	}
		//	sqldbtest.CheckError(t, rows.Scan(&actualId, &actualHash, &actualMimeType))
		//})
	})
}
