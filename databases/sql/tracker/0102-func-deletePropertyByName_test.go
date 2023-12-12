package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deletePropertyByName(t *testing.T) {
	const (
		functionName     = "createStringProperty"
		rootTable        = "propertyKeys"
		stringTableName  = "stringProperties"
		numericTableName = "numericProperties"
		testPropertyName = "testProperty"
		testStringValue  = "testStringValue"
		testIntegerValue = 1337
	)

	var propertyId string
	var propertyString string
	var propertyInteger int

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

	t.Run("Test the string property", func(t *testing.T) {
		t.Run("Create the string property", func(t *testing.T) {
			rows, err := db.Query("select createStringProperty('%s','%s');", testPropertyName, testStringValue)
			if err != nil {
				t.Fatalf("Failed when calling createStringProperty(): %v", err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			err = rows.Scan(&propertyId)
			sqldbtest.CheckError(t, err)
			if propertyId == "" {
				t.Fatal("empty propertyId returned")
			}
			raw, err := uuid.Parse(propertyId)
			if err != nil {
				t.Fatalf("Failed: returned value is not uuid\n"+
					"Expected: %v\n"+
					"Actual:   %v", propertyId, raw)
			}
		})
		t.Run("Verify string property", func(t *testing.T) {
			rows, err := db.Query("select value from %s where id='%s'", stringTableName, propertyId)
			if err != nil {
				t.Fatalf("failed to select expected value: %v", err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			if err = rows.Scan(&propertyString); err != nil {
				t.Fatalf("Error scanning row value: %v", err)
			}
			if propertyString != testStringValue {
				t.Fatalf("testPropertyValue mismatch")
			}
		})
		t.Run("delete string property", func(t *testing.T) {
			rows, err := db.Query("select deletePropertyByName('%s');", testPropertyName)
			if err != nil {
				t.Fatalf("Failed when calling deletePropertyByName(): %v", err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var count int
			if err = rows.Scan(&count); err != nil {
				t.Fatal(err)
			}
			if count != 1 {
				t.Fatalf("expected count 1 but got %d", count)
			}
		})
	})

	t.Run("Test the numeric property", func(t *testing.T) {
		t.Run("Create the numeric property", func(t *testing.T) {
			rows, err := db.Query("select createIntegerProperty('%s',%d);", testPropertyName, testIntegerValue)
			if err != nil {
				t.Fatalf("Failed when calling createIntegerProperty(): %v", err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			err = rows.Scan(&propertyId)
			sqldbtest.CheckError(t, err)
			if propertyId == "" {
				t.Fatal("empty propertyId returned")
			}
			raw, err := uuid.Parse(propertyId)
			if err != nil {
				t.Fatalf("Failed: returned value is not uuid\n"+
					"Expected: %v\n"+
					"Actual:   %v", propertyId, raw)
			}
		})
		t.Run("Verify numeric property", func(t *testing.T) {
			rows, err := db.Query("select value from %s where id='%s'", numericTableName, propertyId)
			if err != nil {
				t.Fatalf("failed to select expected value: %v", err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			if err = rows.Scan(&propertyInteger); err != nil {
				t.Fatalf("Error scanning row value: %v", err)
			}
			if propertyInteger != testIntegerValue {
				t.Fatalf("testIntegerValue mismatch")
			}
		})

		t.Run("delete numeric property", func(t *testing.T) {
			rows, err := db.Query("select deletePropertyByName('%s');", testPropertyName)
			if err != nil {
				t.Fatalf("Failed when calling deletePropertyByName(): %v", err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var count int
			if err = rows.Scan(&count); err != nil {
				t.Fatal(err)
			}
			if count != 1 {
				t.Fatalf("expected count 1 but got %d", count)
			}
		})
	})

}
