package psqlTrackerDb

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createPropertyKey(t *testing.T) {
	const (
		functionName     = "createPropertyKey"
		tableName        = "propertyKeys"
		testPropertyName = "testPropertyName"
	)

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where name='%s' cascade;", tableName, testPropertyName)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{propertyname},"+
				"pt:{varchar},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	var propertyId string
	t.Run("create a property using the function", func(t *testing.T) {
		rows, err := db.Query("select createPropertyKey('%s');", testPropertyName)
		if err != nil {
			t.Fatalf("Failed when calling createPropertyKey(): %v", err)
		}
		t.Run("createPropertyKey() should return a row", func(t *testing.T) {
			if !rows.Next() {
				t.Fatal("no row returned")
			}
		})

		t.Run("expect result is a uuid", func(t *testing.T) {
			err = rows.Scan(&propertyId)
			sqldbtest.CheckError(t, err)
		})
		if propertyId == "" {
			t.Fatal("empty propertyId returned")
		}
	})

	t.Run("select and verify the propertyId and propertyNAme", func(t *testing.T) {
		rows, err := db.Query("select name from propertyKeys where id='%s'", propertyId)
		if err != nil {
			t.Fatalf("select query failed: %v", err)
		}
		t.Run("createPropertyKey() should return a row", func(t *testing.T) {
			if !rows.Next() {
				t.Fatal("no row returned")
			}
		})
		var propertyName string
		t.Run("expect the propertyName can be scanned", func(t *testing.T) {
			err = rows.Scan(&propertyName)
			if err != nil {
				t.Fatalf("Error scanning row value: %v", err)
			}
			if propertyName != testPropertyName {
				t.Fatalf("propertyName mismatch")
			}
		})
	})
}
