package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createEntity(t *testing.T) {
	const (
		functionName = "createEntity"
	)
	var rows *sql.Rows
	var err error
	var entityId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{type,context},"+
				"pt:{varchar,entityType},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("test the create operation", func(t *testing.T) {

		var actualId uuid.UUID
		var actualType string
		var actualContext string

		entityId = createEntity(t, db, "other")

		t.Run("Verify the entityId", func(t *testing.T) {
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
			if actualType != "other" {
				t.Fatal("Fail: actualType mismatch")
			}
			if actualContext != "" {
				t.Fatal("Fail: context mismatch")
			}
		})
	})
}
