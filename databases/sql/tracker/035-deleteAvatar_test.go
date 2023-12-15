package psqlTrackerDb

import (
	"database/sql"
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
	var rows *sql.Rows
	var err error
	var entityId uuid.UUID

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
				"pn:{targetId},"+
				"pt:{uuid},"+
				"rt:int4", strings.ToLower(functionName)))
	})

	t.Run("create an avatar record", func(t *testing.T) {
		entityId = createAvatar(t, db)
	})
	t.Run("delete the avatar", func(t *testing.T) {
		deleteAvatar(t, db, entityId)
		rows, err = db.Query("select deleteAvatar('%s');", entityId)
		if err != nil {
			t.Fatalf("failed on call to deleteAvatar(): %v", err)
		}
		defer func() { _ = rows.Close() }()
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
		rows, err = db.Query("select deleteAvatar('%s');", entityId)
		if err != nil {
			t.Fatalf("failed on call to deleteAvatar(): %v", err)
		}
		defer func() { _ = rows.Close() }()
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

	t.Run("count the number of matching avatars (expect zero)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select count(id) from avatars where id=('%s');", entityId)
		if err != nil {
			t.Fatalf("count query failed %v\n"+
				"teamId:  %v", err, entityId)
		}
		defer func() { _ = rows.Close() }()
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
