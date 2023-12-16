package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateTeamPermissions(t *testing.T) {
	const (
		avatarHash      = "ce9e67fb9ac29bcc831ee93a5bb9431063c4652704fb3a715826b2beef57c6e2"
		avatarType      = "image/png"
		oIconHash       = "880f67ded3d5c953348039c0bd575919455fc3e0a5cbcf1ac408c493cf53cdd4"
		oIconType       = "image/png"
		functionName    = "updatePermissions"
		userFirstName   = "Amy"
		userLastName    = "Budwit"
		userEmail       = "amy.budwitt1@example.com"
		userPhone       = "123.991.3544"
		userDescription = "A good mentor...missed"
		teamName        = "Engineering Professionals"
		teamDescription = "this is the team"
		pRead           = "read"
		pNone           = "none"
	)
	var avatarId uuid.UUID
	var IconId uuid.UUID
	var ownerId uuid.UUID
	var teamId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from teams where id='%s'", teamId)
		_, _ = db.Query("delete from users where id='%s'", ownerId)
		_, _ = db.Query("delete from icons where id='%s'", IconId)
		_, _ = db.Query("delete from avatars where id='%s'", avatarId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{teamId,o,t,e},"+
			"pt:{uuid,permissions},"+
			"rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	IconId = createIcon(t, db, oIconType, oIconHash)
	ownerId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)
	teamId = createTeam(t, db, teamName, IconId, ownerId, pRead, pRead, pRead, teamDescription)

	if count := updatePermissions(t, db, teamId, pNone, pNone, pNone); count != 1 {
		t.Fatalf("expect 1 but got %d", count)
	}
	verifyTeamRecord(t, db, teamId, teamName, IconId, ownerId, pNone, pNone, pNone, teamDescription)
}
