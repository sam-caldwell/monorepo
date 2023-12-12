package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createTeam(t *testing.T) {
	const (
		avatarUrl           = "http://localhost/myfakeavatar.jpeg"
		iconUrl             = "http://localhost/myfakeicon.jpeg"
		functionName        = "createTeam"
		tableName           = "teams"
		testTeamName        = "testTeam1"
		expectedFirstName   = "Paul"
		expectedLastName    = "Allen"
		expectedEmail       = "paul.allen@example.com"
		expectedPhone       = "737.444.0988"
		expectedDescription = "Test description"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var ownerId uuid.UUID
	var teamId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where name='%s'", tableName, testTeamName)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{name,iconid,ownerId,owner,team,everyone,description},"+
				"pt:{text,varchar,uuid,permissions},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("createIcons()", func(t *testing.T) {
		/*
		 * We need to create an icon for use creating a Workflow and Team
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createIcons('%s');", iconUrl)
		if err != nil {
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
		t.Logf("iconId created: %v", iconId)
	})

	t.Run("avatarId", func(t *testing.T) {
		/*
		 * We need to create an avatar to create a user (ownerId)
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createAvatar('%s');", avatarUrl)
		if err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if avatarId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
		if avatarId.String() == "00000000-0000-0000-0000-000000000000" {
			t.Fatal("illegal zero uuid")
		}
	})

	t.Run("createUser (ownerId)", func(t *testing.T) {
		/*
		 * We need to create a user (ownerId) to create a workflow
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
			expectedFirstName, expectedLastName, avatarId, expectedEmail, expectedPhone, expectedDescription)
		if err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if ownerId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
		if ownerId.String() == "00000000-0000-0000-0000-000000000000" {
			t.Fatal("illegal zero uuid")
		}
	})

	t.Run("createTeam (teamId)", func(t *testing.T) {
		/*
		 * We need to create a user (ownerId) to create a workflow
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createTeam('%s','%s','%s','read','read','read','%s');",
			testTeamName, iconId, ownerId, expectedDescription)
		if err != nil {
			t.Fatalf("createTeam() failed %v\n"+
				"iconId:  %v\n"+
				"ownerId: %v",
				err, iconId, ownerId)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if teamId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
		if teamId.String() == "00000000-0000-0000-0000-000000000000" {
			t.Fatal("illegal zero uuid")
		}
	})

	t.Run("inspect and verify team", func(t *testing.T) {
		/*
		 * verify the user.
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select id,name,iconId,ownerId,owner,team,everyone,description "+
			"from teams where id='%s'", teamId)
		if err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}

		var actualTeamId uuid.UUID
		var actualTeamName string
		var actualIconId uuid.UUID
		var actualOwnerId uuid.UUID
		var actualPermissionOwner, actualPermissionTeam, actualPermissionEveryone string
		var actualDescription string

		err = rows.Scan(&actualTeamId, &actualTeamName, &actualIconId, &actualOwnerId,
			&actualPermissionOwner, &actualPermissionTeam, &actualPermissionEveryone,
			&actualDescription)

		if err != nil {
			t.Fatalf("Failed to read result: %v", err)
		}
		if teamId != actualTeamId {
			t.Fatal("teamId mismatch")
		}
		if testTeamName != actualTeamName {
			t.Fatal("team name mismatch")
		}
		if iconId != actualIconId {
			t.Fatalf("iconId mismatch")
		}
		if ownerId != actualOwnerId {
			t.Fatalf("ownerId mismatch")
		}
		if actualPermissionOwner != "read" {
			t.Fatalf("owner permission mismatch")
		}
		if actualPermissionTeam != "read" {
			t.Fatalf("team permission mismatch")
		}
		if actualPermissionEveryone != "read" {
			t.Fatalf("everyone permission mismatch")
		}
		if actualDescription != expectedDescription {
			t.Fatalf("actualDescription mismatch\n"+
				"actual: %s\n"+
				"expected: %s", actualDescription, expectedDescription)
		}
	})
}
