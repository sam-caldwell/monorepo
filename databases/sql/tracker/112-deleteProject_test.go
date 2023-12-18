package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deleteProject(t *testing.T) {
	const (
		avatarHash          = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType          = "image/png"
		iconHash            = "182e31fa48267c22d598dfcddb66e2dafd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType            = "image/png"
		functionName        = "deleteProject"
		expectedFirstName   = "Aleksandr"
		expectedLastName    = "Solzhenitsyn"
		expectedEmail       = "Aleksandr.Solzhenitsyn@example.com"
		expectedPhone       = "812.612.0632"
		expectedDescription = "How to piss off your government and still die of old age"
		testTeamName        = "Zek"
		expectedProject     = "WritingBooks"
		pRead               = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var projectId uuid.UUID

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
		fmt.Sprintf("fn:%s,pn:{projectid},pt:{uuid},rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)

	iconId = createIcon(t, db, iconType, iconHash)

	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)

	projectId = createProject(t, db, expectedProject, iconId, ownerId, teamId, pRead, pRead, pRead, expectedDescription)

	t.Run("delete project", func(t *testing.T) {
		var count int
		var err error
		var rows *sql.Rows

		rows, err = db.Query("select deleteProject('%v');", projectId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		if err = rows.Scan(&count); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
		if count != 1 {
			t.Fatalf("expect count 1 but got %d", count)
		}
	})
	t.Run("verify", func(t *testing.T) {
		if count := countById(t, db, "projects", projectId); count != 0 {
			t.Fatalf("expect count 0 but got %d", count)
		}
	})
}
