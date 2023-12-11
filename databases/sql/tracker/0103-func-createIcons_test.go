package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createIcons(t *testing.T) {
	const (
		functionName = "createIcons"
		tableName    = "icons"
		testUrl      = "http://localhost/myfakeicon.jpeg"
	)
	var iconId uuid.UUID
	var ownerId uuid.UUID
	var teamId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	//t.Cleanup(func() {
	//	// Note: we only clean up the avatar we expect to have created.
	//	//       this should safeguard against an accidental run on prod.
	//	_, _ = db.Query("delete from %s where name='%s'", tableName, testTeamName)
	//
	//	err := db.Close()
	//	sqldbtest.CheckError(t, err)
	//})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{name,iconid,ownerId,owner,team,everyone,description},"+
				"pt:{text,varchar,uuid,permissions},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("create iconId", func(t *testing.T) {
		/*
		 * We need to create an icon for use creating a Workflow and Team
		 */
		var rows *sql.Rows
		var err error
		if rows, err = db.Query("select createIcons('%s');", testUrl); err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if iconId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
	})
}
