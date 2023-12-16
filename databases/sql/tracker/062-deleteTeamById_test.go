package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deleteTeamById(t *testing.T) {
	const (
		avatarHash          = "47ff27e37c131957af3d0ae9c402fe0ba206f6ad22048e2c4ede7566b58d3f91"
		avatarType          = "image/png"
		iconHash            = "182e31fa48267c22d598dfcddb66e2dffd0b4ec2b0192e28c3b73336b71ea8b4"
		iconType            = "image/png"
		functionName        = "deleteTeamById"
		testTeamName        = "testTeam4"
		expectedFirstName   = "Peter"
		expectedLastName    = "Norton"
		expectedEmail       = "someuser@example.com"
		expectedPhone       = "717.444.0988"
		expectedDescription = "Test description"
		pRead               = "read"
	)

	var avatarId uuid.UUID
	var iconId uuid.UUID
	var ownerId uuid.UUID
	var teamId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from teams where id='%s'", teamId)
		_, _ = db.Query("delete from users where id='%s'", ownerId)
		_, _ = db.Query("delete from icons where id='%s'", iconId)
		_, _ = db.Query("delete from avatars where id='%s'", avatarId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{teamid},"+
			"pt:{uuid},"+
			"rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)

	t.Run("call deleteTeamById(teamId)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select deleteTeamById('%s');", teamId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		err = rows.Scan(&count)
		if err != nil {
			t.Fatal(err)
		}
		if count != 1 {
			t.Fatalf("count expected 1 but got %d", count)
		}
	})

	if count := countById(t, db, "teams", teamId); count != 0 {
		t.Fatalf("expected count 0 but got %d", count)
	}
}
