package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createStringProperty(t *testing.T) {
	const (
		functionName = "createStringProperty"
		rootTable    = "propertyKeys"
		//tableName        = "stringProperties"
		testPropertyName  = "testStringProperty"
		testPropertyValue = "testPropertyValue"
	)

	var propertyId string
	var propertyValue string

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where name='%s' cascade;", rootTable, testPropertyName)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{propertyname,propertyvalue},"+
				"pt:{text,varchar},"+
				"rt:uuid", strings.ToLower(functionName)))
	})
	t.Run("Create the string property", func(t *testing.T) {
		rows, err := db.Query("select createStringProperty('%s','%s');",
			testPropertyName, testPropertyValue)
		if err != nil {
			t.Fatalf("Failed when calling createStringProperty(): %v", err)
		}
		t.Run("createStringProperty() should return a row", func(t *testing.T) {
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
	t.Run("expect the propertyValue can be scanned", func(t *testing.T) {
		rows, err := db.Query("select value from stringProperties where id='%s'", propertyId)
		if err != nil {
			t.Fatalf("failed to select expected value: %v", err)
		}
		t.Run("createStringProperty() should return a row", func(t *testing.T) {
			if !rows.Next() {
				t.Fatal("no row returned")
			}
		})
		t.Run("createStringProperty() should contain a propertyValue", func(t *testing.T) {
			err = rows.Scan(&propertyValue)
			if err != nil {
				t.Fatalf("Error scanning row value: %v", err)
			}
			if propertyValue != testPropertyValue {
				t.Fatalf("testPropertyValue mismatch")
			}
		})
	})
}
