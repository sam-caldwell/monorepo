package psqlTrackerDb

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getAvatarById(t *testing.T) {
	const (
		tableName    = "avatar"
		functionName = "getAvatarById"
		testUrl      = "http://testUrlButNeverReal.tld/thisShouldNeverExistInTheDb.png"
	)
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where url='%s';", tableName, testUrl)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{id},"+
				"pt:{uuid},"+
				"rt:jsonb", strings.ToLower(functionName)))
	})

	t.Run("happy path:run the function", func(t *testing.T) {
		t.Skip("not implemented")
	})

}
