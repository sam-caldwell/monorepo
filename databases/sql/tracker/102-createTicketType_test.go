package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createTicketType(t *testing.T) {
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType             = "image/png"
		functionName         = "createTicketType"
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
			"pn:{typeName,typeIconId,wid,typeDescription},"+
			"pt:{text,varchar,uuid},"+
			"rt:uuid", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
	workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId,
		pRead, pRead, pRead, expectedDescription)
	ticketTypeId = createTicketType(t, db, expectedTicketType, iconId, workflowId, expectedDescription)

	t.Run("verify record", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query(""+
			"select id, name, iconId, workflowId, description "+
			"from ticketTypes "+
			"where id='%s'", ticketTypeId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var actualId uuid.UUID
		var actualName string
		var actualIcon uuid.UUID
		var actualWorkflow uuid.UUID
		var actualDescription string
		err = rows.Scan(&actualId, &actualName, &actualIcon, &actualWorkflow, &actualDescription)

		if actualId != ticketTypeId {
			t.Fatalf("id mismatch\n"+
				"actual:   %v\n"+
				"expected: %v", actualId, ticketTypeId)
		}
		if actualName != expectedTicketType {
			t.Fatalf("name mismatch")
		}
		if actualIcon != iconId {
			t.Fatalf("iconId mismatch")
		}
		if actualWorkflow != workflowId {
			t.Fatalf("workflowId mismatch")
		}
		if actualDescription != expectedDescription {
			t.Fatalf("Description mismatch")
		}
	})
}
