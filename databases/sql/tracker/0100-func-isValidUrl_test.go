package psqlTrackerDb

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_isValidUrl(t *testing.T) {
	const (
		functionName = "isValidUrl"
		zeroUuid     = "00000000-0000-0000-0000-000000000000"
	)
	db := sqldbtest.InitializeTestDbConn(t)

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{url},"+
				"pt:{text},"+
				"rt:bool", strings.ToLower(functionName)))
	})

	t.Run("happy path: test the good URLs", func(t *testing.T) {

		urls := []string{
			"http://google.com",
			"https://google.com",
			"http://simpleHostName",
			"http://happyUrl.tld/good.image.name.png",
			"http://happyUrl.tld/good.image.name.png",
			"http://happyUrl.tld/good.image.name.png/?query=param",
			"http://happyUrl.tld/good.image.name.png/?query1=param1&query1=param1",
			"https://happyUrl.tld/good.image.name.png",
			"https://happyUrl.tld/good.image.name.png/?query=param",
			"https://happyUrl.tld/good.image.name.png/?query1=param1&query1=param1",
		}

		for _, url := range urls {

			rows, err := db.Query("select isValidUrl('%s');", url)
			if err != nil {
				t.Fatalf("Fail: Query returned an error %s", err)
			}

			t.Run("isValidUrl() should return a row", func(t *testing.T) {
				if !rows.Next() {
					t.Fatal("Fail: no row returned")
				}
			})

			t.Run("isValidUrl() should be happy", func(t *testing.T) {
				var result bool
				err = rows.Scan(&result)
				if !result {
					t.Fatalf("Fail: Expected valid URL (%s). isValidUrl() returned %v", url, result)
				}
			})
		}
	})
	t.Run("sad path: test the bad URLs", func(t *testing.T) {
		urls := []string{
			"this is bad",
			"ftp://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png",
			"htp://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png",
			"smb://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png",
			"ssh://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png",
			"nfs://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png",
			"tftp://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png",
		}
		for _, url := range urls {
			rows, err := db.Query("select isValidUrl('%s');", url)
			sqldbtest.CheckError(t, err)
			t.Run("isValidUrl() should return a row", func(t *testing.T) {
				if !rows.Next() {
					t.Fatal("no row returned")
				}
			})
			t.Run("isValidUrl() should be sad", func(t *testing.T) {
				// returning true means a bad url made it through
				var result bool
				err = rows.Scan(&result)
				if result {
					t.Fatalf("Fail: Expected invalid URL (%s) to fail, returned %v", url, result)
				}
			})
		}
	})
}
