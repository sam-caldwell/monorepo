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
		badTestUrl1  = "this is bad"
		badTestUrl2  = "http://169.254.169.256/foo.png"
		badTestUrl3  = "ftp://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png"
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
	//t.Run("sad path: test the bad URLs", func(t *testing.T) {
	//	urls := []string{
	//		badTestUrl1, badTestUrl2, badTestUrl3,
	//	}
	//	for _, url := range urls {
	//		rows, err := db.Query("select isValidUrl('%s');", url)
	//		sqldbtest.CheckError(t, err)
	//		t.Run("isValidUrl() should return a row", func(t *testing.T) {
	//			if !rows.Next() {
	//				t.Fatal("no row returned")
	//			}
	//		})
	//		t.Run("isValidUrl() should be sad", func(t *testing.T) {
	//			var result bool
	//			err = rows.Scan(&result)
	//			sqldbtest.CheckError(t, err)
	//			if result {
	//				t.Fatalf("Expected false, got true result (%v)", result)
	//			}
	//		})
	//	}
	//})
}
