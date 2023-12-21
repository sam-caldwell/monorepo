package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getWorkflowsByTeamId(t *testing.T) {
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType             = "image/png"
		functionName         = "getWorkflowsByTeamId"
		expectedFirstName    = "Jack"
		expectedLastName     = "Cook"
		expectedEmail        = "jack.cook@example.com"
		expectedPhone        = "321.321.6543"
		expectedDescription  = "Test description"
		testTeamName         = "OceanExplorers"
		expectedWorkflowName = "NavigationProcess"
		pRead                = "read"
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
			"pn:{workflowTeamId,pageLimit,pageOffset},"+
			"pt:{int4,uuid},"+
			"rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
	workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId,
		pRead, pRead, pRead, expectedDescription)

	t.Run("call getWorkflowByTeamId (bad bounds)", func(t *testing.T) {
		var err error
		if _, err = db.Query("select getWorkflowByTeamId('%s',0,10);", ownerId); err == nil {
			t.Fatal(err)
		}
	})
	t.Run("call getWorkflowByTeamId (bad bounds)", func(t *testing.T) {
		//var rows *sql.Rows
		var err error
		if _, err = db.Query("select getWorkflowByTeamId('%s',1001,0);", teamId); err == nil {
			t.Fatal(err)
		}
	})
	//t.Run("call getWorkflowByTeamId", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select getWorkflowsByOwnerId('%s',1,0);", ownerId)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	if err = rows.Scan(&raw); err != nil {
	//		t.Fatal(err)
	//	}
	//	var actualWorkflow []TrackerWorkflow
	//	if err = json.Unmarshal([]byte(raw), &actualWorkflow); err != nil {
	//		t.Fatal(err)
	//	}
	//	for _, wf := range actualWorkflow {
	//		if wf.Id != workflowId {
	//			t.Fatal("Error: workflowId mismatch")
	//		}
	//
	//		if wf.Name != expectedWorkflowName {
	//			t.Fatal("Error: Name mismatch")
	//		}
	//
	//		if wf.IconId != iconId {
	//			t.Fatal("Error: IconId mismatch")
	//		}
	//
	//		if wf.OwnerId != ownerId {
	//			t.Fatal("Error: OwnerId mismatch")
	//		}
	//
	//		if wf.TeamId != teamId {
	//			t.Fatal("Error: TeamId mismatch")
	//		}
	//
	//		if wf.Owner != "read" {
	//			t.Fatal("Error: Owner mismatch")
	//		}
	//
	//		if wf.Team != "read" {
	//			t.Fatal("Error: Team mismatch")
	//		}
	//
	//		if wf.Everyone != "read" {
	//			t.Fatal("Error: Everyone mismatch")
	//		}
	//
	//		if wf.Description != expectedDescription {
	//			t.Fatal("Error: Description mismatch")
	//		}
	//	}
	//})
}
