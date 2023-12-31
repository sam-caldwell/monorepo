package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getProjectById(t *testing.T) {
	const (
		avatarHash          = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType          = "image/png"
		iconHash            = "182e31fa48267c22d598dfcddb66e2dafd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType            = "image/png"
		functionName        = "getProjectById"
		expectedFirstName   = "Jack"
		expectedLastName    = "Rackham"
		expectedEmail       = "jack.rackham@example.com"
		expectedPhone       = "321.321.6593"
		expectedDescription = "Test description"
		testTeamName        = "OceanExplorers1"
		expectedProject     = "testProject1"
		pRead               = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var projectId uuid.UUID
	//var actual TrackerProject

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "projects", projectId)
		_ = cleanUpObject(db, "teams", teamId)
		_ = cleanUpObject(db, "users", ownerId)
		_ = cleanUpObject(db, "icons", iconId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,pn:{projectid},pt:{uuid},rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)

	iconId = createIcon(t, db, iconType, iconHash)

	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)

	projectId = createProject(t, db, expectedProject, iconId, ownerId, teamId, pRead, pRead, pRead, expectedDescription)
	//
	//t.Run("get project", func(t *testing.T) {
	//	var raw string
	//	var err error
	//	var rows *sql.Rows
	//
	//	rows, err = db.Query("select getProjectById('%v');", projectId)
	//	if err != nil {
	//		t.Fatalf("Fail: function call failed: %v", err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("Fail: no row returned")
	//	}
	//	if err = rows.Scan(&raw); err != nil {
	//		t.Fatalf("Failed to scan rows. %v", err)
	//	}
	//	if err = json.Unmarshal([]byte(raw), &actual); err != nil {
	//		t.Fatalf("failed to unmarshal object. %v", err)
	//	}
	//})
	//t.Run("verify", func(t *testing.T) {
	//	if actual.Id != projectId {
	//		t.Fatalf("projectId mismatch")
	//	}
	//	if actual.Name != expectedProject {
	//		t.Fatalf("project name mismatch")
	//	}
	//	if actual.IconId != iconId {
	//		t.Fatalf("iconId mismatch")
	//	}
	//	if actual.OwnerId != ownerId {
	//		t.Fatalf("ownerId mismatch")
	//	}
	//	if actual.TeamId != teamId {
	//		t.Fatalf("teamId mismatch")
	//	}
	//	if actual.Description != expectedDescription {
	//		t.Fatalf("description mismatch")
	//	}
	//})
}
