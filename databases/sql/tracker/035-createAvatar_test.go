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
		tableName    = "avatars"
		functionName = "createAvatar"
		testHash     = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
		testType     = "image/png"
	)
	var entityId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, tableName, entityId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{t,h},"+
			"pt:{varchar,mimeType},"+
			"rt:uuid", strings.ToLower(functionName)))

	t.Run("test the create operation", func(t *testing.T) {

		entityId = createAvatar(t, db, testType, testHash)

		t.Run("verify entity", func(t *testing.T) {
			var actualEntity TrackerEntity
			getEntity(t, db, entityId, &actualEntity)
			if actualEntity.Id != entityId {
				t.Fatalf("entityId mismatch. got: %v", actualEntity.Id)
			}
			if actualEntity.Type != "avatar" {
				t.Fatalf("entity Type mismatch. got: %v", actualEntity.Type)
			}
			if actualEntity.Context != "" {
				t.Fatalf("entity Context mismatch. got: %v", actualEntity.Type)
			}
		})

		t.Run("Verify the avatar", func(t *testing.T) {
			var rows *sql.Rows
			var err error
			var actualId uuid.UUID
			var actualHash string
			var actualMimeType string
			rows, err = db.Query("select id, hash, mimetype from %s where id='%s';", tableName, entityId)
			if err != nil {
				t.Fatalf("Fail: (query): %v", err)
			}
			defer func() { _ = rows.Close() }()
			if !rows.Next() {
				t.Fatal("Fail: no row returned")
			}
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
