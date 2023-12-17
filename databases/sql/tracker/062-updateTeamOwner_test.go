package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateTeamOwner(t *testing.T) {
	const (
		avatarHash      = "ce9e67fb9ac29bcc831ee93a5bb9431063c4652704fb3a715826b2beef57c6e2"
		avatarType      = "image/png"
		oIconHash       = "880f67ded3d5c953348039c0bd575919455fc3e0a5cbcf1ac408c493cf53cdd4"
		oIconType       = "image/png"
		functionName    = "updateTeamOwner"
		userFirstName   = "Amy"
		userLastName    = "Budwit"
		userEmail       = "amy.budwitt1@example.com"
		userDescription = "A good mentor...missed"
		teamName        = "Engineering Professionals"
		teamDescription = "this is the team"
		pRead           = "read"
	)
	var avatarId uuid.UUID
	var IconId uuid.UUID
	var originalOwner uuid.UUID
	var newOwner uuid.UUID
	var teamId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from teams where id='%s'", teamId)
		_, _ = db.Query("delete from icons where id='%s'", IconId)
		_, _ = db.Query("delete from users where id='%s'", originalOwner)
		_, _ = db.Query("delete from users where id='%s'", newOwner)
		_, _ = db.Query("delete from avatars where id='%s'", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{teamId,newOwnerId},"+
			"pt:{uuid},"+
			"rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	IconId = createIcon(t, db, oIconType, oIconHash)
	originalOwner = createUser(t, db, "F", "L", avatarId, "e@a.co", "314.152.9346", userDescription)
	newOwner = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, "314.152.9347", userDescription)
	teamId = createTeam(t, db, teamName, IconId, originalOwner, pRead, pRead, pRead, teamDescription)

	if count := updateTeamOwner(t, db, teamId, newOwner); count != 1 {
		t.Fatalf("expect 1 but got %d", count)
	}
	verifyTeamRecord(t, db, teamId, teamName, IconId, newOwner, pRead, pRead, pRead, teamDescription)
}
