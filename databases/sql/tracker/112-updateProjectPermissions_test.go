package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateProjectPermissions(t *testing.T) {
	const (
		avatarHash          = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType          = "image/png"
		iconHash            = "182e31fa48267c22d598dfcddb66e2dafd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType            = "image/png"
		functionName        = "updateProjectPermissions"
		expectedFirstName   = "Sam"
		expectedLastName    = "Caldwell"
		expectedEmail       = "sam.caldwell@example.com"
		expectedPhone       = "444.333.2222"
		expectedDescription = "monorepo coder"
		testTeamName        = "Code Monkey"
		expectedProject     = "TrackerDb"
		pRead               = "read"
		pNone               = "none"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var originalOwnerId uuid.UUID
	var newOwnerId uuid.UUID
	var projectId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "projects", projectId)
		_ = cleanUpObject(db, "teams", teamId)
		_ = cleanUpObject(db, "users", originalOwnerId)
		_ = cleanUpObject(db, "users", newOwnerId)
		_ = cleanUpObject(db, "icons", iconId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,pn:{projectId,powner,pteam,peveryone},pt:{uuid,permissions},rt:int4",
			strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)

	iconId = createIcon(t, db, iconType, iconHash)

	originalOwnerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	newOwnerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, "newOwner@ex.co",
		"111.212.4435", expectedDescription)

	teamId = createTeam(t, db, testTeamName, iconId, originalOwnerId, pRead, pRead, pRead, expectedDescription)

	projectId = createProject(t, db, expectedProject, iconId, originalOwnerId, teamId,
		pRead, pRead, pRead, expectedDescription)

	t.Run("get project", func(t *testing.T) {
		var count int
		var err error
		var rows *sql.Rows

		rows, err = db.Query("select updateProjectPermissions('%v','%v'::permissions,'%v'::permissions,"+
			"'%v'::permissions);", projectId, pNone, pNone, pNone)
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
			t.Fatalf("count expect 1 but got %d", count)
		}
	})
	t.Run("verify", func(t *testing.T) {
		var err error
		var rows *sql.Rows
		var owner, team, everyone string
		rows, err = db.Query("select owner,team,everyone from projects where id='%s';", projectId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		if err = rows.Scan(&owner, &team, &everyone); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
		if owner != pNone {
			t.Fatalf("owner permission mismatch.  Got %s", owner)
		}
		if team != pNone {
			t.Fatalf("team permission mismatch.  Got %s", team)
		}
		if everyone != pNone {
			t.Fatalf("everyone permission mismatch.  Got %s", everyone)
		}
	})
}
