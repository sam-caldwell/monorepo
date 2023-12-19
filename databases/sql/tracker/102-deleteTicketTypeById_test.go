package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deleteTicketTypeById(t *testing.T) {
    t.Skip("disabled for debugging")
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType             = "image/png"
		functionName         = "deleteTicketTypeById"
		expectedFirstName    = "Jack"
		expectedLastName     = "Cook"
		expectedEmail        = "jack.cook@example.com"
		expectedPhone        = "321.321.6543"
		expectedDescription  = "Test description"
		testTeamName         = "OceanExplorers"
		expectedWorkflowName = "NavigationProcess"
		expectedTicketType   = "epic"
		pRead                = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var workflowId uuid.UUID
	var ticketTypeId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "ticketTypes", ticketTypeId)
		_ = cleanUpObject(db, "workflows", workflowId)
		_ = cleanUpObject(db, "teams", teamId)
		_ = cleanUpObject(db, "users", ownerId)
		_ = cleanUpObject(db, "icons", iconId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{typeid},"+
			"pt:{uuid},"+
			"rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
	workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId,
		pRead, pRead, pRead, expectedDescription)
	ticketTypeId = createTicketType(t, db, expectedTicketType, iconId, workflowId, expectedDescription)
	t.Logf("ticketTypeId: %v", ticketTypeId)
	if count := deleteTicketTypeById(t, db, ticketTypeId); count != 1 {
		t.Fatalf("count 1 expected but got %d", count)
	}

}
