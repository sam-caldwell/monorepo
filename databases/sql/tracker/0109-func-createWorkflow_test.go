package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createWorkflow(t *testing.T) {
	const (
		avatarUrl            = "http://localhost/myfakeavatar.jpeg"
		iconUrl              = "http://localhost/myfakeicon.ico"
		functionName         = "createWorkflow"
		expectedFirstName    = "Alexander"
		expectedLastName     = "Bell"
		expectedEmail        = "abell@example.com"
		expectedPhone        = "001.001.0001"
		expectedDescription  = "Test description"
		expectedTeamName     = "PhoneGuy"
		expectedWorkflowName = "ETCallHome"
	)
	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from users where email='%s'", expectedEmail)
		_, _ = db.Query("delete from teams where name='%s'", expectedTeamName)
		_, _ = db.Query("delete from avatars where url='%s'", avatarUrl)
		_, _ = db.Query("delete from icons where url='%s'", iconUrl)
		_, _ = db.Query("delete from teamMembership where teamId='%s'", teamId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{name,iconid,ownerid,teamid,owner,team,everyone,description},"+
				"pt:{text,varchar,uuid,permissions},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("call createAvatar()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createAvatar('%s');", avatarUrl)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if err != nil {
			t.Fatal(err)
		}
		if avatarId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
		if avatarId.String() == "00000000-0000-0000-0000-000000000000" {
			t.Fatal("illegal zero uuid")
		}
	})

	t.Run("call createIcons()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createIcons('%s');", iconUrl)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if err != nil {
			t.Fatal(err)
		}
		if iconId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("call createUser()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
			expectedFirstName, expectedLastName, avatarId, expectedEmail, expectedPhone, expectedDescription)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if ownerId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("call createTeam()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createTeam('%s','%s','%s','%s','%s','%s','%s');",
			expectedTeamName, iconId, ownerId, "read", "read", "read", expectedDescription)
		if err != nil {
			t.Fatal(err)
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
		t.Run("call addUserToTeam()", func(t *testing.T) {
			var rows *sql.Rows
			var err error
			rows, err = db.Query("select addUserToTeam('%s','%s');", ownerId, teamId)
			if err != nil {
				t.Fatal(err)
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
		t.Run("verify user membership", func(t *testing.T) {
			var rows *sql.Rows
			var err error
			rows, err = db.Query(""+
				"select count(userId) "+
				"from teamMembership "+
				"where userId='%s' "+
				"and teamId='%s';", ownerId, teamId)
			if err != nil {
				t.Fatal(err)
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
	})

	t.Run("createWorkflow() and verify", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createWorkflow('%s','%s','%s','%s','%s','%s','%s','%s');",
			expectedWorkflowName, iconId, ownerId, teamId, "read", "read", "read", expectedDescription)
		if err != nil {
			t.Fatal(err)
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

	})
}
