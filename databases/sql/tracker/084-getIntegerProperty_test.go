package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getIntegerProperty(t *testing.T) {
	const (
		functionName      = "getIntegerProperty"
		rootTable         = "propertyKeys"
		testPropertyName  = "testIntegerProperty2"
		testPropertyValue = 1337
	)

	var propertyId string
	var propertyValue int

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from %s where name='%s';", rootTable, testPropertyName)
        sqldbtest.CheckError(t, db.Close())
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{propertyname},"+
				"pt:{varchar},"+
				"rt:int4", strings.ToLower(functionName)))
	})

	t.Run("create a numeric property", func(t *testing.T) {
		rows, err := db.Query("select createIntegerProperty('%s',%d);", testPropertyName, testPropertyValue)
		if err != nil {
			t.Fatalf("Failed to create integer property. %v", err)
		}
		defer func() { _ = rows.Close() }()
		t.Run("createIntegerProperty() should return a row", func(t *testing.T) {
			if !rows.Next() {
				t.Fatal("no row returned")
			}
		})
		t.Run("expect uuid result", func(t *testing.T) {
			err = rows.Scan(&propertyId)
			sqldbtest.CheckError(t, err)
			if propertyId == "" {
				t.Fatal("empty propertyId returned")
			}
			u, err := uuid.Parse(propertyId)
			if err != nil {
				t.Fatalf("Failed: returned value is not uuid\n"+
					"Expected: %v\n"+
					"Actual:   %v", propertyId, u)
			}
		})
	})
	t.Run("get the property value using property name", func(t *testing.T) {
		rows, err := db.Query("select %s('%s');", functionName, testPropertyName)
		if err != nil {
			t.Fatalf("Failed to fetch integer property. %v", err)
		}
		defer func() { _ = rows.Close() }()
		t.Run("function should return a row", func(t *testing.T) {
			if !rows.Next() {
				t.Fatal("no row returned")
			}
		})
		t.Run("function should return a value", func(t *testing.T) {
			err = rows.Scan(&propertyValue)
			if err != nil {
				t.Fatalf("Error scanning row value: %v", err)
			}
			if propertyValue != testPropertyValue {
				t.Fatalf("property value mismatch")
			}
		})
	})
}
