package database

import (
    "fmt"
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

    //t.Run("happy path:run the function", func(t *testing.T) {
    //	var avatarId uuid.UUID
    //	var avatarUrl string
    //	t.Run("createAvatar", func(t *testing.T) {
    //		rows, err := db.Query("select createAvatar('%s');", testUrl)
    //		database.CheckError(t, err)
    //		rows.Next()
    //		err = rows.Scan(&avatarId)
    //		database.CheckError(t, err)
    //	})
    //	t.Run("Query for expected avatarId", func(t *testing.T) {
    //		rows, err := db.Query("select id,url from avatars where url='%s';", testUrl)
    //	})
    //})

}
