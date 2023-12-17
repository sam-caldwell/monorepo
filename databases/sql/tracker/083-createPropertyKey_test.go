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
		testPropertyName = "testPropertyName"
	)

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from propertyKeys where name='%s';", testPropertyName)
        sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{propertyname},"+
			"pt:{varchar},"+
			"rt:uuid", strings.ToLower(functionName)))

	var propertyId string

	t.Run("create a property using the function", func(t *testing.T) {
		rows, err := db.Query("select createPropertyKey('%s');", testPropertyName)
		if err != nil {
			t.Fatalf("Failed when calling createPropertyKey(): %v", err)
		}
		defer func() { _ = rows.Close() }()
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
		defer func() { _ = rows.Close() }()
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
