package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getTerminalNode(t *testing.T) {
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType             = "image/png"
		functionName         = "getTerminalNode"
		testTeamName         = "testTeam1"
		expectedFirstName    = "Paul"
		expectedLastName     = "Allen"
		expectedEmail        = "paul.allen@example.com"
		expectedPhone        = "737.444.0988"
		expectedDescription  = "Test description"
		pRead                = "read"
		expectedWorkflowName = "test workflow"
	)
	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var workflowId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		sqldbtest.CheckError(t, cleanUpObject(db, "workflows", workflowId))
		sqldbtest.CheckError(t, cleanUpObject(db, "teams", teamId))
		sqldbtest.CheckError(t, cleanUpObject(db, "users", ownerId))
		sqldbtest.CheckError(t, cleanUpObject(db, "icons", iconId))
		sqldbtest.CheckError(t, cleanUpObject(db, "avatars", avatarId))
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{thisWorkflow},"+
			"pt:{uuid},"+
			"rt:uuid", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
	workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId,
		pRead, pRead, pRead, expectedDescription)

	var stopNodeId uuid.UUID

	t.Run("getTerminalNode", func(t *testing.T) {
		rows, err := db.Query("select getTerminalNode('%s');", workflowId)
		if err != nil {
			t.Fatalf("Fail: query error: %v", err)
		}
		if !rows.Next() {
			t.Fatalf("Fail: no rows returned")
		}
		if err = rows.Scan(&stopNodeId); err != nil {
			t.Fatalf("Fail: scan error: %v", err)
		}
	})
	t.Run("verify node", func(t *testing.T) {
		rows, err := db.Query("select id from workflowSteps where workflowId='%s' and name ='terminate'",
			workflowId)
		if err != nil {
			t.Fatalf("Fail: query error: %v", err)
		}
		if !rows.Next() {
			t.Fatalf("Fail: no rows returned")
		}
		var actualNodeId uuid.UUID
		if err = rows.Scan(&actualNodeId); err != nil {
			t.Fatalf("Fail: scan error: %v", err)
		}
		if actualNodeId != stopNodeId {
			t.Fatalf("Fail: node mismatch")
		}
	})
}
