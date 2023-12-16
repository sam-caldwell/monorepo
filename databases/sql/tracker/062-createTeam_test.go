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
		avatarHash          = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType          = "image/png"
		iconHash            = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType            = "image/png"
		functionName        = "createTeam"
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
		_, _ = db.Query("delete from teams where id='%s'", teamId)
		_, _ = db.Query("delete from users where id='%s'", ownerId)
		_, _ = db.Query("delete from icons where id='%s'", iconId)
		_, _ = db.Query("delete from avatars where id='%s'", avatarId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{name,iconid,ownerId,owner,team,everyone,description},"+
			"pt:{text,varchar,uuid,permissions},"+
			"rt:uuid", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

	t.Run("createTeam (teamId)", func(t *testing.T) {
		teamId = createTeam(t, db, testTeamName, iconId, ownerId,
			"read", "read", "read", expectedDescription)
	})

	t.Run("inspect and verify team", func(t *testing.T) {
		var err error
		var rows *sql.Rows
		var actualTeamId uuid.UUID
		var actualTeamName string
		var actualIconId uuid.UUID
		var actualOwnerId uuid.UUID
		var actualPermissionOwner, actualPermissionTeam, actualPermissionEveryone string
		var actualDescription string

		rows, err = db.Query("select id,name,iconId,ownerId,owner,team,everyone,description "+
			"from teams where id='%s'", teamId)
		if err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}

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
