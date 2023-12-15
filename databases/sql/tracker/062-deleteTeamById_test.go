package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deleteTeamById(t *testing.T) {
	const (
		avatarHash          = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
		avatarType          = "image/png"
		iconUrl             = "http://localhost/myfakeicon.jpeg"
		functionName        = "deleteTeamById"
		tableName           = "teams"
		testTeamName        = "testTeam4"
		expectedFirstName   = "Peter"
		expectedLastName    = "Norton"
		expectedEmail       = "someuser@example.com"
		expectedPhone       = "717.444.0988"
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
				"pn:{teamId},"+
				"pt:{uuid},"+
				"rt:int4", strings.ToLower(functionName)))
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
		defer func() { _ = rows.Close() }()
		defer func() { _ = rows.Close() }()
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

	t.Run("createAvatars()", func(t *testing.T) {
		/*
		 * We need to create an avatar to create a user (ownerId)
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createAvatar('%s'::mimeType,'%s');", avatarHash, avatarType)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		defer func() { _ = rows.Close() }()
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
		defer func() { _ = rows.Close() }()
		defer func() { _ = rows.Close() }()
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
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createTeam('%s','%s','%s','read','read','read','test team');",
			testTeamName, iconId, ownerId)
		if err != nil {
			t.Fatalf("createTeam() failed %v\n"+
				"iconId:  %v\n"+
				"ownerId: %v",
				err, iconId, ownerId)
		}
		defer func() { _ = rows.Close() }()
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

	t.Run("deleteTeamById(teamId)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select deleteTeamById('%s');", teamId)
		if err != nil {
			t.Fatalf("deleteTeam() failed %v\n"+
				"teamId:  %v", err, teamId)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		err = rows.Scan(&count)
		if count != 1 {
			t.Fatalf("expected count 1 but got %d", count)
		}
	})

	t.Run("count the number of matching teams (expect zero)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select count(id) from teams where id=('%s');", teamId)
		if err != nil {
			t.Fatalf("count query failed %v\n"+
				"teamId:  %v", err, teamId)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		err = rows.Scan(&count)
		if count != 0 {
			t.Fatalf("expected count 0 but got %d", count)
		}
	})
}
