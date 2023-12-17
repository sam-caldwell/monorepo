package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateTeamIcon(t *testing.T) {
	const (
		avatarHash      = "ce9e67fb9ac29bcc831ee93a5bb9431063c4652704fb3a715826b2beef57c6e2"
		avatarType      = "image/png"
		oIconHash       = "880f67ded3d5c953348039c0bd575919455fc3e0a5cbcf1ac408c493cf53cdd4"
		oIconType       = "image/png"
		nIconHash       = "19ed4500c67de78533c3b5b66d8279da6d7b119901b68326c9af6e72daad4bfd"
		nIconType       = "image/png"
		functionName    = "updateTeamIcon"
		userFirstName   = "Klaus"
		userLastName    = "Fuchs"
		userEmail       = "treason.weasel@example.com"
		userPhone       = "314.152.9246"
		userDescription = "Traitor"
		teamName        = "Manhattan Project"
		teamDescription = "this is the team"
		pRead           = "read"
	)
	var avatarId uuid.UUID
	var originalIcon uuid.UUID
	var newIcon uuid.UUID
	var ownerId uuid.UUID
	var teamId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from teams where id='%s'", teamId)
		_, _ = db.Query("delete from icons where id='%s'", originalIcon)
		_, _ = db.Query("delete from icons where id='%s'", newIcon)
		_, _ = db.Query("delete from users where id='%s'", ownerId)
		_, _ = db.Query("delete from avatars where id='%s'", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{teamId,newIconId},"+
			"pt:{uuid},"+
			"rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	originalIcon = createIcon(t, db, oIconType, oIconHash)
	newIcon = createIcon(t, db, nIconType, nIconHash)
	ownerId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)
	teamId = createTeam(t, db, teamName, originalIcon, ownerId, pRead, pRead, pRead, teamDescription)

	if count := updateTeamIcon(t, db, teamId, newIcon); count != 1 {
		t.Fatalf("expect 1 but got %d", count)
	}
	verifyTeamRecord(t, db, teamId, teamName, newIcon, ownerId, pRead, pRead, pRead, teamDescription)
}
