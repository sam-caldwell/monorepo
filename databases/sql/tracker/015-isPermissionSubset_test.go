package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbFunc_isPermissionSubset(t *testing.T) {
	db := sqldbtest.InitializeTestDbConn(t)
	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})
	testFunc := func(lhs string, rhs string, expected bool) {
		t.Run("isPermissionSubset()", func(t *testing.T) {
			var actual bool
			rows, err := db.Query("select isPermissionSubset('%s','%s');", lhs, rhs)
			if err != nil {
				t.Fatalf("Fail: query error. %v", err)
			}
			if !rows.Next() {
				t.Fatalf("Fail: expected rows to be returned but got none")
			}
			if err = rows.Scan(&actual); err != nil {
				t.Fatalf("Fail: scan error: %v", err)
			}
			if expected != actual {
				t.Fatalf("outcome expected %v but got %v", expected, actual)
			}
		})
	}

	/*
	   - lhs is the offered permission.
	   - rhs is the expected permission.
	   - The result is whether the expected permission (lhs) is in the offered permission (rhs)
	*/
	testFunc("none", "none", false)
	testFunc("none", "read", false)
	testFunc("none", "create", false)
	testFunc("none", "update", false)
	testFunc("none", "delete", false)

	testFunc("read", "none", false)  // None is never going to pass
	testFunc("read", "read", true)   // Read is Read.
	testFunc("read", "create", true) // A Creator can read
	testFunc("read", "update", true) // An Updator can read
	testFunc("read", "delete", true) // A Deleter can read

	testFunc("create", "none", false)  // None is never going to pass
	testFunc("create", "read", false)  // Read cannot create
	testFunc("create", "create", true) // create can create
	testFunc("create", "update", true) // updaters can create
	testFunc("create", "delete", true) // a deleter can create

	testFunc("update", "none", false)   // None is never going to pass
	testFunc("update", "read", false)   // read cannot update
	testFunc("update", "create", false) // read cannot create
	testFunc("update", "update", true)  // updaters can update
	testFunc("update", "delete", true)  // a deleter can update (delete is the ultimate update)

	testFunc("delete", "none", false)   // None is never going to pass
	testFunc("delete", "read", false)   // read cannot delete
	testFunc("delete", "create", false) // create cannot delete
	testFunc("delete", "update", false) // update cannot delete
	testFunc("delete", "delete", true)  // delete can delete

}
