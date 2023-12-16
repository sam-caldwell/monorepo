package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateTeamDescription(t *testing.T) {
	const (
		avatarHash         = "6a8434431319cfc231f4313e4663657f63ef99890635c6c0981316c591bc8a19"
		avatarType         = "image/png"
		iconHash           = "b14e141d72fc7b31ef2fa19e034b2d4c615dcbf26f1ad41a12365615dbcfa2e1"
		iconType           = "image/png"
		functionName       = "updateTeamDescription"
		userFirstName      = "Klaus"
		userLastName       = "Fuchs"
		userEmail          = "treason.weasel@example.com"
		userPhone          = "314.152.9246"
		userDescription    = "Test description"
		teamName           = "Manhattan Project"
		teamDescription    = "this is the team"
		newTeamDescription = "new team description"
		pRead              = "read"
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
			"pn:{teamId,d},"+
			"pt:{text,uuid},"+
			"rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)
	teamId = createTeam(t, db, teamName, iconId, ownerId, pRead, pRead, pRead, teamDescription)

	if count := updateTeamDescription(t, db, teamId, newTeamDescription); count != 1 {
		t.Fatalf("expect 1 but got %d", count)
	}
	verifyTeamRecord(t, db, teamId, teamName, iconId, ownerId, pRead, pRead, pRead, newTeamDescription)
}
