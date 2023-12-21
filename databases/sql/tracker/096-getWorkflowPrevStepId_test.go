package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getWorkflowPrevStepId(t *testing.T) {
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74e6b8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dafd0b4ec2b0192e28c3b73336b71ea7b4"
		iconType             = "image/png"
		functionName         = "getWorkflowPrevStepId"
		expectedFirstName    = "Billy"
		expectedLastName     = "Gibbons"
		expectedEmail        = "billy.gibbons@example.com"
		expectedPhone        = "235.235.0248"
		expectedDescription  = "Test description"
		testTeamName         = "ZZ Top"
		expectedWorkflowName = "Haw haw haw"
		expectedStepName     = "sing"
		pRead                = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var workflowId uuid.UUID
	var prevStepId uuid.UUID
	var nextStepId uuid.UUID
	var stepId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		sqldbtest.CheckError(t, cleanUpObject(db, "workflowSteps", stepId))
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
			"pn:{stepId},"+
			"pt:{uuid},rt:uuid", strings.ToLower(functionName)))

	t.Run("setup", func(t *testing.T) {
		avatarId = createAvatar(t, db, avatarType, avatarHash)
		iconId = createIcon(t, db, iconType, iconHash)
		ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId,
			expectedEmail, expectedPhone, expectedDescription)
		teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
		workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId, pRead, pRead, pRead,
			expectedDescription)

		prevStepId = getStartNode(t, db, workflowId)
		nextStepId = getTerminalNode(t, db, workflowId)
		stepId = createWorkflowStep(t, db, workflowId, expectedStepName, prevStepId, nextStepId, expectedDescription)
		t.Logf("workflowId: %v", workflowId)
		t.Logf("stepId: %v", stepId)
	})

	var err error
	var actualStepName string
	var actualPrevStepId uuid.UUID
	var targetNodeId uuid.UUID
	t.Run("fetch actual record", func(t *testing.T) {
		var rows *sql.Rows
		/*
		 * Get a prevStepId, which should be 'start' node.
		 */
		targetNodeId = nextStepId
		rows, err = db.Query("select name, prevStepId from workflowSteps where id='%s'", targetNodeId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		if err = rows.Scan(&actualStepName, &actualPrevStepId); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
	})
	t.Run("verify starting point", func(t *testing.T) {
		if actualStepName != "terminate" {
			t.Fatalf("Fail: step name mismatch (starting point)\n"+
				"workflow:   '%v'\n"+
				"actualName: '%s'\n"+
				"expected:   'start'", workflowId, actualStepName)
		}
	})
	t.Run("get the next node", func(t *testing.T) {
		var rows *sql.Rows
		targetNodeId = actualPrevStepId
		rows, err = db.Query("select id,name from workflowSteps where id='%s'", targetNodeId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		if err = rows.Scan(&actualPrevStepId, &actualStepName); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
	})
	t.Run("verify next node", func(t *testing.T) {
		if actualStepName != expectedStepName {
			t.Fatalf("Fail: step name mismatch (next point)\n"+
				"workflow:   '%v'\n"+
				"actualName: '%s'\n"+
				"expected:   '%s'", workflowId, actualStepName, expectedStepName)
		}
		if actualPrevStepId != stepId {
			t.Fatalf("Fail: stepId mismatch\n"+
				"workflow:       '%v'\n"+
				"ActualNextStep: '%v'\n"+
				"StepId:         '%v'", workflowId, actualPrevStepId, targetNodeId)
		}
	})
}
