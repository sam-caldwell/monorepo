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
	t.Skip("debugging")
	const (
		tableName    = "avatar"
		functionName = "createAvatar"
		testHash     = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
		testType     = "image/png"
		badTestUrl   = "this is bad"
		zeroUuid     = "00000000-0000-0000-0000-000000000000"
	)
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
				"pn:{avatarType,avatarHash},"+
				"pt:{varchar},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("test the create operation", func(t *testing.T) {

		var rows *sql.Rows
		var err error

		var actualId uuid.UUID
		var actualHash string
		var actualMimeType string

		t.Run("call createAvatar();", func(t *testing.T) {
			rows, err = db.Query("select createAvatar('%s','%s');", testHash, testType)
			sqldbtest.CheckError(t, err)
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			err = rows.Scan(&actualId)
			sqldbtest.CheckError(t, err)
		})

		t.Run("Verify the entityId", func(t *testing.T) {
			rows, err = db.Query("select id, context from entity where context='avatar';")
			sqldbtest.CheckError(t, err)
			defer func() { _ = rows.Close() }()
			var thisId uuid.UUID
			var thisContext string
			if err = rows.Scan(&thisId, &thisContext); err != nil {
				t.Fatal(err)
			}
			if actualId != thisId {
				t.Fatal("actualId mismatch")
			}
			if thisContext != "avatar" {
				t.Fatal("context mismatch")
			}
		})

		t.Run("Verify the avatar record", func(t *testing.T) {
			rows, err = db.Query(""+
				"select id, hash,mimeType"+
				"from avatars "+
				"where hash='%s';", testHash)
			sqldbtest.CheckError(t, err)
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("no rows returned")
			}
			sqldbtest.CheckError(t, rows.Scan(&actualId, &actualHash, &actualMimeType))
		})
	})
}
