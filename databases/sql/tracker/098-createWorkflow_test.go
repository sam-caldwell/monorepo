package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createWorkflow(t *testing.T) {
	t.Skip("disabled for debugging")
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType             = "image/png"
		functionName         = "createWorkflow"
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
			"pn:{workflowName,workflowIconId,workflowOwnerId,workflowTeamId,"+
			"workflowPermissionOwner,workflowPermissionTeam,workflowPermissionEveryone,"+
			"workflowPermissionDescription},"+
			"pt:{text,varchar,uuid,permissions},"+
			"rt:uuid", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
	workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId,
		pRead, pRead, pRead, expectedDescription)

	//t.Run("verify workflow", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	var actualId uuid.UUID
	//	var actualName string
	//	var actualIconId uuid.UUID
	//	var actualOwnerId uuid.UUID
	//	var actualTeamId uuid.UUID
	//	var actualOwner string
	//	var actualTeam string
	//	var actualEveryone string
	//	var actualDescription string
	//
	//	rows, err = db.Query(""+
	//		"select id,name,iconId,ownerId,teamId,owner,team,everyone,description "+
	//		"from workflows where id = '%s';", workflowId)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//
	//	err = rows.Scan(&actualId, &actualName, &actualIconId, &actualOwnerId, &actualTeamId,
	//		&actualOwner, &actualTeam, &actualEveryone, &actualDescription)
	//
	//	if actualId != workflowId {
	//		t.Fatalf("workflowId mismatch")
	//	}
	//	if actualName != expectedWorkflowName {
	//		t.Fatalf("workflowName mismatch")
	//	}
	//	if actualIconId != iconId {
	//		t.Fatalf("iconId mismatch")
	//	}
	//	if actualOwnerId != ownerId {
	//		t.Fatalf("ownerId mismatch")
	//	}
	//	if actualTeamId != teamId {
	//		t.Fatalf("teamId mismatch")
	//	}
	//	if actualOwner != pRead {
	//		t.Fatalf("owner mismatch")
	//	}
	//	if actualTeam != pRead {
	//		t.Fatalf("team mismatch")
	//	}
	//	if actualEveryone != pRead {
	//		t.Fatalf("everyone mismatch")
	//	}
	//	if actualDescription != expectedDescription {
	//		t.Fatalf("description mismatch")
	//	}
	//})
}
