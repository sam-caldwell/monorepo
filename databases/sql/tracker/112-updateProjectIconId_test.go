package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateProjectIconId(t *testing.T) {
	t.Skip("disabled for debugging")
	const (
		avatarHash          = "4ab7b2cbfa7a2120025400e1d08ace0ec81b9a27a5411b00e1ec75e74edb8f51"
		avatarType          = "image/png"
		iconHash1           = "182e31fa48267c22d598dfcddb66e2dafe0b4ec2b0192e28c3b73336b71ea8b4"
		iconHash2           = "182e31fa48267c22d598dfcddb66e2daff0b4ec2b0192e28c3b73336b71ea8b4"
		iconType            = "image/png"
		functionName        = "updateProjectIconId"
		expectedFirstName   = "Sam"
		expectedLastName    = "Caldwell"
		expectedEmail       = "sam.caldwell@example.com"
		expectedPhone       = "444.333.2222"
		expectedDescription = "monorepo coder"
		testTeamName        = "Code Monkey"
		expectedProject     = "TrackerDb"
		pRead               = "read"
	)

	var avatarId uuid.UUID
	var originalIconId uuid.UUID
	var newIconId uuid.UUID
	var teamId uuid.UUID
	var ownerId uuid.UUID
	var projectId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "projects", projectId)
		_ = cleanUpObject(db, "teams", teamId)
		_ = cleanUpObject(db, "users", ownerId)
		_ = cleanUpObject(db, "icons", originalIconId)
		_ = cleanUpObject(db, "icons", newIconId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,pn:{projectId,projectIconId},pt:{uuid},rt:int4",
			strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)

	originalIconId = createIcon(t, db, iconType, iconHash1)
	newIconId = createIcon(t, db, iconType, iconHash2)

	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

	teamId = createTeam(t, db, testTeamName, originalIconId, ownerId, pRead, pRead, pRead, expectedDescription)

	projectId = createProject(t, db, expectedProject, originalIconId, ownerId, teamId, pRead, pRead, pRead,
		expectedDescription)

	t.Run("get project", func(t *testing.T) {
		var count int
		var err error
		var rows *sql.Rows

		rows, err = db.Query("select updateProjectIconId('%v','%v');", projectId, newIconId)
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
		rows, err = db.Query("select iconId from projects where id='%s';", projectId)
		if err != nil {
			t.Fatalf("Fail: function call failed: %v", err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var actual uuid.UUID
		if err = rows.Scan(&actual); err != nil {
			t.Fatalf("Failed to scan rows. %v", err)
		}
		if actual != newIconId {
			t.Fatalf("description mismatch.  Got %s", actual)
		}
	})
}
