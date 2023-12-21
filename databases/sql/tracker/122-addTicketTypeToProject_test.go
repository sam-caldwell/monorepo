package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_addTicketTypeToProject(t *testing.T) {
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74e6b8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dafd0b4ec2b0192e28c3b73336b71ea7b4"
		iconType             = "image/png"
		functionName         = "addTicketTypeToProject"
		expectedFirstName    = "Robert"
		expectedLastName     = "Oppenheimer"
		expectedEmail        = "Robert.Oppenheimer@example.com"
		expectedPhone        = "235.235.0238"
		expectedDescription  = "Test description"
		testTeamName         = "Los Alamos Nerds"
		expectedProject      = "Manhattan Project"
		expectedWorkflowName = "NavigationProcess"
		expectedTicketType   = "epic"
		pRead                = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var projectId uuid.UUID
	var workflowId uuid.UUID
	var ticketTypeId uuid.UUID
	var assocId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "projectTicketTypes", assocId)
		_ = cleanUpObject(db, "ticketTypes", ticketTypeId)
		_ = cleanUpObject(db, "workflows", workflowId)
		_ = cleanUpObject(db, "projects", projectId)
		_ = cleanUpObject(db, "teams", teamId)
		_ = cleanUpObject(db, "users", ownerId)
		_ = cleanUpObject(db, "icons", iconId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,pn:{thisPid,thisTid},pt:{uuid},rt:uuid", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId,
		expectedEmail, expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
	projectId = createProject(t, db, expectedProject, iconId, ownerId, teamId, pRead, pRead, pRead, expectedDescription)

	workflowId = createWorkflow(t, db, expectedWorkflowName,
		iconId, ownerId, teamId,
		pRead, pRead, pRead,
		expectedDescription)

	ticketTypeId = createTicketType(t, db, expectedTicketType, iconId, workflowId, expectedDescription)

	assocId = addTicketTypeToProject(t, db, projectId, ticketTypeId)

	t.Run("verify record", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select projectId, ticketTypeId from projectTicketTypes where id='%s'", assocId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var actualProjectId uuid.UUID
		var actualTicketTypeId uuid.UUID

		err = rows.Scan(&actualProjectId, &actualTicketTypeId)

		if actualProjectId != projectId {
			t.Fatalf("actualId mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualProjectId, projectId)
		}
		if actualTicketTypeId != ticketTypeId {
			t.Fatalf("expectedProject mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualTicketTypeId, ticketTypeId)
		}
	})
}
