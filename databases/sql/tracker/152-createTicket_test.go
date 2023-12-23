package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createTicket(t *testing.T) {
	const (
		avatarHash          = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType          = "image/png"
		iconHash            = "182e31fa48267c22d598dfcddb66e9dafd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType            = "image/png"
		functionName        = "createTicket"
		expectedFirstName   = "Robert"
		expectedLastName    = "Oppenheimer"
		expectedEmail       = "Robert.Oppenheimer@example.com"
		expectedPhone       = "235.235.0238"
		expectedDescription = "Test description"
		testTeamName        = "Los Alamos Nerds"
		expectedProject     = "Manhattan Project"
		pRead               = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var projectId uuid.UUID
	var ticketId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "tickets", ticketId)
		_ = cleanUpObject(db, "projects", projectId)
		_ = cleanUpObject(db, "teams", teamId)
		_ = cleanUpObject(db, "users", ownerId)
		_ = cleanUpObject(db, "icons", iconId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{project,author,assignee,type,powner,pteam,peveryone,ticketsummary,ticketdescription},"+
			"pt:{text,varchar,uuid,permissions},"+
			"rt:uuid",
			strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)

	iconId = createIcon(t, db, iconType, iconHash)

	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId,
		expectedEmail, expectedPhone, expectedDescription)

	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)

	projectId = createProject(t, db, expectedProject, iconId, ownerId, teamId, pRead, pRead, pRead, expectedDescription)

	t.Run("verify record", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query(""+
			"select id, name, iconId, ownerId, teamId, owner, team, everyone, description "+
			"from projects "+
			"where id='%s'", projectId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var actualId uuid.UUID
		var actualName string
		var actualIconId uuid.UUID
		var actualOwnerId uuid.UUID
		var actualTeamId uuid.UUID
		var actualOwner string
		var actualTeam string
		var actualEveryone string
		var actualDescription string
		err = rows.Scan(&actualId, &actualName, &actualIconId, &actualOwnerId, &actualTeamId,
			&actualOwner, &actualTeam, &actualEveryone, &actualDescription)

		if actualId != projectId {
			t.Fatalf("actualId mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualId, projectId)
		}
		if actualName != expectedProject {
			t.Fatalf("expectedProject mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualName, expectedProject)
		}
		if actualIconId != iconId {
			t.Fatalf("actualIconId mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualIconId, iconId)
		}
		if actualOwnerId != ownerId {
			t.Fatalf("actualOwnerId mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualOwnerId, ownerId)
		}
		if actualTeamId != teamId {
			t.Fatalf("actualTeamId mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualTeamId, teamId)
		}
		if actualOwner != pRead {
			t.Fatalf("actualOwner mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualOwner, pRead)
		}
		if actualTeam != pRead {
			t.Fatalf("actualTeam mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualTeam, pRead)
		}
		if actualEveryone != pRead {
			t.Fatalf("actualEveryone mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualEveryone, pRead)
		}
		if actualDescription != expectedDescription {
			t.Fatalf("actualDescription mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualDescription, expectedDescription)
		}
	})
}
