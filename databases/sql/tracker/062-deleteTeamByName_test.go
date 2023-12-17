package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deleteTeamByName(t *testing.T) {
	const (
		avatarHash          = "7de85b7ceec1b76bb42e6efef98c7ba24418cc3ef8c69c8609118e7c587d2e4c"
		avatarType          = "image/png"
		iconHash            = "b00061aac177dd79b667351dde07edf223313499112635d0c3ffc08bdbadcd8b"
		iconType            = "image/png"
		functionName        = "deleteTeamByName"
		testTeamName        = "testTeam3"
		expectedFirstName   = "Peter"
		expectedLastName    = "Norton"
		expectedEmail       = "anotherUser@example.com"
		expectedPhone       = "727.444.0988"
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
		_, _ = db.Query("delete from icons where id='%s'", iconId)
		_, _ = db.Query("delete from users where id='%s'", ownerId)
		_, _ = db.Query("delete from avatars where id='%s'", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{teamName},"+
			"pt:{varchar},"+
			"rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)
	teamId = createTeam(t, db, testTeamName, iconId, ownerId, pRead, pRead, pRead, expectedDescription)

	t.Run("call deleteTeamByName(teamName)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select deleteTeamByName('%s');", testTeamName)
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

	if count := countByName(t, db, "teams", testTeamName); count != 0 {
		t.Fatalf("expected count 0 but got %d", count)
	}
}
