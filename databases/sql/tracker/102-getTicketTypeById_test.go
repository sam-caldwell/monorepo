package psqlTrackerDb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getTicketTypeById(t *testing.T) {
	const (
		avatarHash           = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType           = "image/png"
		iconHash             = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType             = "image/png"
		functionName         = "getTicketTypeById"
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
	var actual TrackerTicketType

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		sqldbtest.CheckError(t, cleanUpObject(db, "ticketTypes", ticketTypeId))
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
			"pn:{typeid},"+
			"pt:{uuid},"+
			"rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)
	workflowId = createWorkflow(t, db, expectedWorkflowName, iconId, ownerId, teamId,
		pRead, pRead, pRead, expectedDescription)
	ticketTypeId = createTicketType(t, db, expectedTicketType, iconId, workflowId, expectedDescription)

	t.Run("get the ticket", func(t *testing.T) {
		var err error
		var raw sql.NullString
		var rows *sql.Rows

		rows, err = db.Query("select getTicketTypeById('%v');", ticketTypeId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}

		defer func() { _ = rows.Close() }()

		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}

		if err = rows.Scan(&raw); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}

		if !raw.Valid {
			t.Fatalf("Fail: no record or invalid record")
		}

		if err = json.Unmarshal([]byte(raw.String), &actual); err != nil {
			t.Fatalf("Fail: parse error %v", err)
		}

		if actual.Id != ticketTypeId {
			t.Fatalf("ticketTypeId mismatch")
		}
		if actual.Name != expectedTicketType {
			t.Fatalf("TicketType Name mismatch")
		}
		if actual.IconId != iconId {
			t.Fatalf("iconId mismatch")
		}
		if actual.WorkflowId != workflowId {
			t.Fatalf("WorkflowId mismatch")
		}
		if actual.Description != expectedDescription {
			t.Fatalf("Description mismatch")
		}
	})
}
