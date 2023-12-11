package psqlTrackerDb

import (
	"fmt"
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
		sqldbtest.CheckError(t, err)
		t.Run("createAvatar() should return a row", func(t *testing.T) {
			if !rows.Next() {
				t.Fatal("no row returned")
			}
		})
		var propertyId string
		t.Run("expect result is a uuid", func(t *testing.T) {
			err = rows.Scan(&propertyId)
			sqldbtest.CheckError(t, err)
		})
		if propertyId == "" {
			t.Fatal("empty propertyId returned")
		}
	})
}
