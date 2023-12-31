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
		sqldbtest.CheckError(t, db.Close())
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

		for _, expectedContext := range []string{"", "test context"} {
			for _, expectedType := range []string{"user", "project", "ticket"} {

				entityId = createEntity(t, db, expectedType, expectedContext)

				testName := fmt.Sprintf("Verify the entityId (context %s, type: %v)",
					expectedContext, expectedType)
				t.Run(testName, func(t *testing.T) {
					rows, err = db.Query("select id, type, context from entity where id='%s';", entityId)
					if err != nil {
						t.Fatalf("Fail: (query): %v", err)
					}
					defer func() { _ = rows.Close() }()
					if !rows.Next() {
						t.Fatalf("Fail: no row returned")
					}
					if err = rows.Scan(&actualId, &actualType, &actualContext); err != nil {
						t.Fatalf("Fail: scan error: %v", err)
					}
					if actualId != entityId {
						t.Fatalf("Fail: actualId mismatch")
					}
					if actualType != expectedType {
						t.Fatalf("Fail: actualType mismatch")
					}
					if actualContext != expectedContext {
						t.Fatalf("Fail: context mismatch\n"+
							"actual:   '%v'\n"+
							"expected: '%v'", actualContext, expectedContext)
					}
				})
			}
		}
	})
}
