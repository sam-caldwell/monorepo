package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createWorkflowStep(t *testing.T) {
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74e6b8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dafd0b4ec2b0192e28c3b73336b71ea7b4"
		iconType             = "image/png"
		functionName         = "createWorkflowStep"
		expectedFirstName    = "Billy"
		expectedLastName     = "Gibbons"
		expectedEmail        = "billy.gibbons@example.com"
		expectedPhone        = "235.235.0248"
		expectedDescription  = "Test description"
		testTeamName         = "ZZ Top"
		expectedProject      = "Blues"
		expectedWorkflowName = "Haw haw haw"
		expectedStepName     = "sing"
		pRead                = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var projectId uuid.UUID
	var workflowId uuid.UUID
	var ticketTypeId uuid.UUID
	var prevStepId uuid.UUID //fake object (entity uuid only)
	var nextStepId uuid.UUID //fake object (entity uuid only)
	var stepId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "workflowSteps", stepId)
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
		fmt.Sprintf("fn:%s,"+
			"pn:{stepName,thisWorkflowId,pStepId,nStepId,stepDescription},"+
			"pt:{text,varchar,uuid},rt:uuid", strings.ToLower(functionName)))

	t.Run("setup", func(t *testing.T) {
		avatarId = createAvatar(t, db, avatarType, avatarHash)
		iconId = createIcon(t, db, iconType, iconHash)
		ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId,
			expectedEmail, expectedPhone, expectedDescription)
		teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
		projectId = createProject(t, db, expectedProject, iconId, ownerId, teamId, pRead, pRead, pRead, expectedDescription)
		workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId, pRead, pRead, pRead,
			expectedDescription)
		stepId = createWorkflowStep(t, db, workflowId, expectedStepName, prevStepId, nextStepId, expectedDescription)
	})

	t.Run("verify", func(t *testing.T) {
		var err error
		var rows *sql.Rows

		rows, err = db.Query(""+
			"select id, workflowId, name, prevStepId, nextStepId, description "+
			"from workflowSteps "+
			"where id=%s", stepId)

		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}

		defer func() { _ = rows.Close() }()

		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var actualStepId uuid.UUID
		var actualWorkflowId uuid.UUID
		var actualName string
		var actualPrevStepId uuid.UUID
		var actualNextStepId uuid.UUID
		var actualDescription string

		err = rows.Scan(&actualStepId, &actualWorkflowId, &actualName, &actualPrevStepId, &actualNextStepId, &actualDescription)
		if err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
		if actualStepId != stepId {
			t.Fatalf("stepId mismatch")
		}
		if actualWorkflowId != workflowId {
			t.Fatalf("workflowId mismatch")
		}
		if actualName != expectedStepName {
			t.Fatalf("expectedStepName mismatch")
		}
		if actualPrevStepId != prevStepId {
			t.Fatalf("prevStepId mismatch")
		}
		if actualNextStepId != nextStepId {
			t.Fatalf("nextStepId mismatch")
		}
		if actualDescription != expectedDescription {
			t.Fatalf("expectedDescription mismatch")
		}
	})
}
