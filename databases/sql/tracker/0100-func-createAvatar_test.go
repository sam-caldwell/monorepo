package database

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/databases/tools"
	"strings"
	"testing"
)

func TestSqlDbFunc_createAvatar(t *testing.T) {

	const (
		tableName    = "avatar"
		functionName = "createAvatar"
		testUrl      = "http://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png"
		badTestUrl1  = "this is bad"
		badTestUrl2  = "http://169.254.169.256/foo.png"
		badTestUrl3  = "ftp://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png"
		zeroUuid     = "00000000-0000-0000-0000-000000000000"
	)
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		t.Log("cleanup...")
		_, _ = db.Query("delete from %s where url='%s';", tableName, testUrl)

		err := db.Close()
		database.database.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		database.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{avatarurl},"+
				"pt:{text},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("happy path:run the function", func(t *testing.T) {
		var avatarId uuid.UUID
		var avatarUrl string

		t.Run("execute createAvatar('"+testUrl+"');", func(t *testing.T) {
			rows, err := db.Query("select createAvatar('%s');", testUrl)
			database.CheckError(t, err)
			t.Run("createAvatar() should return a row", func(t *testing.T) {
				if !rows.Next() {
					t.Fatal("no row returned")
				}
			})
			t.Run("expect result is a uuid", func(t *testing.T) {
				err = rows.Scan(&avatarId)
				database.CheckError(t, err)
			})
			t.Run("Query for expected avatarId", func(t *testing.T) {
				rows, err := db.Query(""+
					"select id, url "+
					"from avatars "+
					"where url='%s';", testUrl)
				database.CheckError(t, err)
				rows.Next()
				database.CheckError(t, rows.Scan(&avatarId, &avatarUrl))
				if avatarUrl != testUrl {
					t.Fatalf("url mismatch\n"+
						"actual:   %s\n"+
						"expected: %s",
						avatarUrl, testUrl)
				}
			})
		})
	})

	t.Run("sad path:run the function with badUrl", func(t *testing.T) {
		var avatarId uuid.UUID
		for _, badUrl := range []string{badTestUrl1, badTestUrl2, badTestUrl3} {
			t.Run("execute createAvatar('"+badUrl+"');", func(t *testing.T) {
				t.Logf("createAvatar('%s');", badUrl)
				rows, err := db.Query("select createAvatar('%s');", badUrl)
				database.CheckError(t, err)
				t.Run("createAvatar() should return a row", func(t *testing.T) {
					if !rows.Next() {
						t.Fatal("no row returned")
					}
				})
				t.Run("expect zeroUuid Result", func(t *testing.T) {
					err = rows.Scan(&avatarId)
					database.CheckError(t, err)
					if avatarId.String() == zeroUuid {
						t.Fatalf("invalid AvatarId: %s", avatarId)
					}
				})
				t.Run("Query for expected avatarId (expect none)", func(t *testing.T) {
					t.Logf("query table for data: %s", badUrl)
					rows, err := db.Query(""+
						"select id, url "+
						"from avatars "+
						"where url='%s';", badUrl)
					t.Logf("query returned %s", badUrl)
					database.CheckError(t, err)
					t.Logf("query had no error")
					if rows.Next() {
						t.Fatal("We should not have created an avatar entry on invalid url")
					}
					t.Logf("result ok")
				})
			})
		}
	})
}
