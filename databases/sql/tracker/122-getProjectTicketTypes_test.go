package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getProjectTicketTypes(t *testing.T) {
	t.Skip("disabled for debugging")
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74e6b8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d5988fcddb66e2dafd0b4ec2b0192e28c3b73336b71ea7b4"
		iconType             = "image/png"
		functionName         = "getProjectTicketTypes"
		expectedFirstName    = "Neils1"
		expectedLastName     = "Bohr1"
		expectedEmail        = "probability.reality1@example.com"
		expectedPhone        = "235.235.0237"
		expectedDescription  = "Test description"
		testTeamName         = "PlayingDiceWithTheUniverse1"
		expProject           = "Quantum Insanity1"
		expectedWorkflowName = "WavesToParticles1"
		expectedTicketType   = "type_"
		pRead                = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var projects []uuid.UUID
	var workflowId uuid.UUID
	var ticketTypeId []uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		for _, tid := range ticketTypeId {
			_ = cleanUpObject(db, "ticketTypes", tid)
		}
		_ = cleanUpObject(db, "workflows", workflowId)
		for _, pid := range projects {
			_ = cleanUpObject(db, "projects", pid)
		}
		_ = cleanUpObject(db, "teams", teamId)
		_ = cleanUpObject(db, "users", ownerId)
		_ = cleanUpObject(db, "icons", iconId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,pn:{thisPid,maxRecords,startAt},pt:{int4,uuid},rt:jsonb",
			strings.ToLower(functionName)))

	t.Run("setup", func(t *testing.T) {
		/*
		 * This is all just trivial stuff we need bit which we can reuse.
		 */
		avatarId = createAvatar(t, db, avatarType, avatarHash)
		iconId = createIcon(t, db, iconType, iconHash)
		ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId,
			expectedEmail, expectedPhone, expectedDescription)
		teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
		workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId, pRead, pRead, pRead,
			expectedDescription)
		/*
		 * Let's create a project which will have one ticket type
		 */
		projects = append(projects, createProject(t, db, expProject+"0", iconId, ownerId, teamId,
			pRead, pRead, pRead, expectedDescription))
		/*
		 * Let's create a second project which will have several ticket types
		 */
		projects = append(projects, createProject(t, db, expProject+"2", iconId, ownerId, teamId,
			pRead, pRead, pRead, expectedDescription))
		/*
		 * Now let's create a set of 10 ticket types...
		 */
		for i := 0; i < 10; i++ {
			typ := fmt.Sprintf("%v%v", expectedTicketType, i)
			d := fmt.Sprintf("%v%d", expectedDescription, i)
			ticketTypeId = append(ticketTypeId, createTicketType(t, db, typ, iconId, workflowId, d))
		}
		/*
		 * We associate project0 with ticket type 0
		 */
		_ = addTicketTypeToProject(t, db, projects[0], ticketTypeId[0])
		/*
		 * We will then associate the other ticket types with project[1]
		 */
		for i := 1; i < 10; i++ {
			_ = addTicketTypeToProject(t, db, projects[1], ticketTypeId[i])
		}
	})
	t.Run("Inspect the outcome of our simple project[0] example", func(t *testing.T) {
		/*
		 * We can now query the ticket types for a given project (id) and inspect it.
		 */
		var supportedTicketTypes []TrackerTicketType
		supportedTicketTypes = getProjectTicketTypes(t, db, projects[0])
		if len(supportedTicketTypes) != 1 {
			t.Fatal("Fail: project0 expected only one associated ticket Type")
		}
		if supportedTicketTypes[0].Id != ticketTypeId[0] {
			t.Fatalf("Fail: project0 expected id mismatch.  Got %s", supportedTicketTypes[0].Id)
		}
		if supportedTicketTypes[0].Name != expectedTicketType+"0" {
			t.Fatalf("Fail: project0 expected name mismatch.  Got %s", supportedTicketTypes[0].Name)
		}
	})
	t.Run("Inspect the outcome of our complex project[1] example", func(t *testing.T) {
		/*
		 * We can now query the ticket types for a given project (id) and inspect it.
		 */
		var supportedTicketTypes []TrackerTicketType
		supportedTicketTypes = getProjectTicketTypes(t, db, projects[1])
		if len(supportedTicketTypes) != 9 {
			t.Fatalf("Fail: project1 (%v)\n"+
				"Expect:9 associated ticket Types\n"+
				"Got: %d\n"+
				"stt: %v", projects[1], len(supportedTicketTypes), supportedTicketTypes)
		}
		for i := 1; i < 10; i++ {
			if supportedTicketTypes[i-1].Id != ticketTypeId[i] {
				t.Fatalf("Fail: project1 expected id mismatch.\n"+
					"Got    %s\n"+
					"Wanted %s", ticketTypeId[1], supportedTicketTypes[i-1].Id)
			}
			if supportedTicketTypes[i-1].Name != fmt.Sprintf("%v%d", expectedTicketType, i) {
				t.Fatalf("Fail: project1 expected name mismatch.  Got %s", supportedTicketTypes[i-1].Name)
			}
		}
	})
}
