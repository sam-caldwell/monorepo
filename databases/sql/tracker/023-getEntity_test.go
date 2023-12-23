package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getEntity(t *testing.T) {
	const (
		tableName    = "entity"
		functionName = "getEntity"
		entityType   = "user"
	)
	var entityId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{entityId},"+
			"pt:{uuid},"+
			"rt:jsonb", strings.ToLower(functionName)))

	t.Run("create an record", func(t *testing.T) {
		entityId = createEntity(t, db, entityType, "getEntity() context")
	})

	t.Run("Verify the entityId [sanity check before calling getEntity()]", func(t *testing.T) {
		var actualId uuid.UUID
		var actualType string
		var actualContext string
		var rows *sql.Rows
		var err error

		rows, err = db.Query("select id, type, context from entity where id='%s';", entityId)
		if err != nil {
			t.Fatalf("Fail: (query): %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		if err = rows.Scan(&actualId, &actualType, &actualContext); err != nil {
			t.Fatal(err)
		}
		if actualId != entityId {
			t.Fatal("Fail: actualId mismatch")
		}
		if actualType != "user" {
			t.Fatal("Fail: actualType mismatch")
		}
		if actualContext != "getEntity() context" {
			t.Fatal("Fail: context mismatch")
		}
	})

	t.Run("get the record and verify", func(t *testing.T) {
		var record TrackerEntity
		getEntity(t, db, entityId, &record)

		if record.Id != entityId {
			t.Fatalf("Fail: id mismatch\n"+
				"got:   '%v'\n"+
				"expect:'%v'", record.Id, entityId)
		}
	})
}
