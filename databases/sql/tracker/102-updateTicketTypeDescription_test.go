package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateTicketTypeDescription(t *testing.T) {
	t.Skip("disabled for debugging")
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType             = "image/png"
		functionName         = "updateTicketTypeDescription"
		expectedFirstName    = "Jack"
		expectedLastName     = "Cook"
		expectedEmail        = "jack.cook@example.com"
		expectedPhone        = "321.321.6543"
		expectedDescription  = "Test description"
		newDescription       = "new description"
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
			"pn:{typeId,typeDescription},"+
			"pt:{text,uuid},"+
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

	t.Run("verify", func(t *testing.T) {
		var err error
		var rows *sql.Rows
		rows, err = db.Query("select description from ticketTypes where id='%s';", ticketTypeId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var actualDescription string
		if err = rows.Scan(&actualDescription); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
		if actualDescription != expectedDescription {
			t.Fatalf("description mismatch.  Got %s", actualDescription)
		}
	})
	t.Run("call createTicketType();", func(t *testing.T) {
		var err error
		var rows *sql.Rows
		rows, err = db.Query("select updateTicketTypeDescription('%v','%v');", ticketTypeId, newDescription)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var count int
		if err = rows.Scan(&count); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
		if count != 1 {
			t.Fatalf("expected count 1 but got %d", count)
		}
	})

	t.Run("verify", func(t *testing.T) {
		var err error
		var rows *sql.Rows
		rows, err = db.Query("select description from ticketTypes where id='%s';", ticketTypeId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var actualDescription string
		if err = rows.Scan(&actualDescription); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
		if actualDescription != newDescription {
			t.Fatalf("description mismatch.  Got %s", actualDescription)
		}
	})
}
