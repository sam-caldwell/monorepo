package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deleteAvatar(t *testing.T) {
	const (
		tableName    = "avatar"
		functionName = "deleteAvatar"
		testUrl      = "http://example.com/itMayExists.png"
	)
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where url='%s';", tableName, testUrl)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{avatarid},"+
				"pt:{uuid},"+
				"rt:int4", strings.ToLower(functionName)))
	})

	var avatarId uuid.UUID
	t.Run("create an avatar record", func(t *testing.T) {
		rows, err := db.Query("select createAvatar('%s');", testUrl)
		if err != nil {
			t.Fatalf("Failed to create record: %v", err)
		}

		t.Run("createAvatar() should return a row", func(t *testing.T) {
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			if err := rows.Scan(&avatarId); err != nil {
				t.Fatalf("err: %v", err)
			}
		})
	})
	t.Run("delete the avatar", func(t *testing.T) {
		rows, err := db.Query("select deleteAvatar('%s');", avatarId)
		if err != nil {
			t.Fatalf("failed on call to deleteAvatar(): %v", err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		if err := rows.Scan(&count); err != nil {
			t.Fatalf("err: %v", err)
		}
		if count == 0 {
			t.Fatal("Failed to delete the record.")
		}
	})

	t.Run("delete the avatar...again and expect 0 results", func(t *testing.T) {
		rows, err := db.Query("select deleteAvatar('%s');", avatarId)
		if err != nil {
			t.Fatalf("failed on call to deleteAvatar(): %v", err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		if err := rows.Scan(&count); err != nil {
			t.Fatalf("err: %v", err)
		}
		if count != 0 {
			t.Fatal("We shouldn't be able to delete a record twice..")
		}
	})
}
