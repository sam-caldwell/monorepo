package database

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/sql/tools"
	"strings"
	"testing"
)

func TestSqlDbFunc_createAvatar(t *testing.T) {

	const (
		tableName    = "avatar"
		functionName = "createAvatar"
		testUrl      = "http://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png"
	)
	db := database.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where url='%s';", tableName, testUrl)

		err := db.Close()
		database.CheckError(t, err)
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

		t.Run("createAvatar", func(t *testing.T) {
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
				t.Log("here we are 2")
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
}
