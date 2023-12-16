package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_addUserToTeam(t *testing.T) {
	const (
		avatarHash       = "a2601e31f65f266a1a94f08ad46918c8d0f9f09f995aa7fbdbfa113ad6911ba6"
		avatarType       = "image/png"
		iconHash         = "69357df9edaa759985b300c4d0341cd906bff5519ff55035a04b58c0af5237c3"
		iconType         = "image/png"
		functionName     = "addUserToTeam"
		ownerFirstName   = "William"
		ownerLastName    = "Shakespeare"
		ownerEmail       = "will.shakespeare@example.com"
		ownerPhone       = "332.152.9247"
		ownerDescription = "Test description"
		userFirstName    = "Isaac"
		userLastName     = "Newton"
		userEmail        = "isaac.newton@example.com"
		userPhone        = "332.152.9246"
		userDescription  = "Test description"
		teamName         = "Gravity Research"
		pRead            = "read"
	)
	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var userId uuid.UUID
	var ownerId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from teammemberships where teamId='%s'", teamId)
		_, _ = db.Query("delete from teams where id='%s'", teamId)
		_, _ = db.Query("delete from users where id='%s'", userId)
		_, _ = db.Query("delete from users where id='%s'", ownerId)
		_, _ = db.Query("delete from icons where id='%s'", iconId)
		_, _ = db.Query("delete from avatars where id='%s'", avatarId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{userId,teamId},"+
			"pt:{uuid},"+
			"rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, ownerFirstName, ownerLastName, avatarId, ownerEmail, ownerPhone, ownerDescription)
	teamId = createTeam(t, db, teamName, iconId, ownerId, pRead, pRead, pRead, userDescription)
	userId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)

	if count := addUserToTeam(t, db, userId, teamId); count != 1 {
		t.Fatalf("count expects 1 but got %d", count)
	}

}
