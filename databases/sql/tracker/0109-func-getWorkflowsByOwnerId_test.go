package psqlTrackerDb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getWorkflowsByOwnerId(t *testing.T) {
	const (
		avatarUrl            = "http://localhost/myfakeavatar.jpeg"
		iconUrl              = "http://localhost/myfakeicon.ico"
		functionName         = "getWorkflowsByOwnerId"
		expectedFirstName    = "William"
		expectedLastName     = "Kidd"
		expectedEmail        = "william.kidd@example.com"
		expectedPhone        = "321.321.6544"
		expectedDescription  = "Test description"
		expectedTeamName     = "OceanExplorers1701"
		expectedWorkflowName = "NavigationProcess1701"
	)
	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var workflowId uuid.UUID

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
				"pn:{workflowOwnerId,pageLimit,pageOffset},"+
				"pt:{int4,uuid},"+
				"rt:jsonb", strings.ToLower(functionName)))
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

	t.Run("createWorkflow()", func(t *testing.T) {
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
		if workflowId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("call getWorkflowByOwnerId (bad bounds)", func(t *testing.T) {
		//var rows *sql.Rows
		var err error
		if _, err = db.Query("select getWorkflowByOwnerId('%s',0,10);", ownerId); err == nil {
			t.Fatal(err)
		}
	})
	t.Run("call getWorkflowByOwnerId (bad bounds)", func(t *testing.T) {
		//var rows *sql.Rows
		var err error
		if _, err = db.Query("select getWorkflowByOwnerId('%s',1001,0);", ownerId); err == nil {
			t.Fatal(err)
		}
	})

	t.Run("call getWorkflowByOwnerId", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select getWorkflowsByOwnerId('%s',1,0);", ownerId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		if err = rows.Scan(&raw); err != nil {
			t.Fatal(err)
		}
		var actualWorkflow []TrackerWorkflow
		if err = json.Unmarshal([]byte(raw), &actualWorkflow); err != nil {
			t.Fatal(err)
		}
		for _, wf := range actualWorkflow {
			if wf.Id != workflowId {
				t.Fatal("Error: workflowId mismatch")
			}

			if wf.Name != expectedWorkflowName {
				t.Fatal("Error: Name mismatch")
			}

			if wf.IconId != iconId {
				t.Fatal("Error: IconId mismatch")
			}

			if wf.OwnerId != ownerId {
				t.Fatal("Error: OwnerId mismatch")
			}

			if wf.TeamId != teamId {
				t.Fatal("Error: TeamId mismatch")
			}

			if wf.Owner != "read" {
				t.Fatal("Error: Owner mismatch")
			}

			if wf.Team != "read" {
				t.Fatal("Error: Team mismatch")
			}

			if wf.Everyone != "read" {
				t.Fatal("Error: Everyone mismatch")
			}

			if wf.Description != expectedDescription {
				t.Fatal("Error: Description mismatch")
			}
		}
	})
}
