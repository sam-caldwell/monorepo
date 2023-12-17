package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_removeUserFromTeam(t *testing.T) {

	const (
		avatarHash       = "a2601e31f65f266a1a94f08ad46918c8d0f9f09f995aa7fbdbfa113ad6911ba6"
		avatarType       = "image/png"
		iconHash         = "69357df9edaa759985b300c4d0341cd906bff5819ff55035a04b58c0af5237c3"
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
		rows, _ := db.Query("delete from teammemberships where teamId='%s'", teamId)
		defer func() { _ = rows.Close() }()
		_ = cleanUpObject(db, "teams", teamId)
		_ = cleanUpObject(db, "icons", iconId)
		_ = cleanUpObject(db, "users", ownerId)
		_ = cleanUpObject(db, "users", userId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{userId,teamId},"+
			"pt:{uuid},"+
			"rt:int4", strings.ToLower(functionName)))

	t.Run("setup", func(t *testing.T) {
		avatarId = createAvatar(t, db, avatarType, avatarHash)
		iconId = createIcon(t, db, iconType, iconHash)
		ownerId = createUser(t, db, ownerFirstName, ownerLastName, avatarId, ownerEmail, ownerPhone, ownerDescription)
		teamId = createTeam(t, db, teamName, iconId, ownerId, pRead, pRead, pRead, userDescription)
		userId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)
		if count := addUserToTeam(t, db, userId, teamId); count != 1 {
			t.Fatalf("count expects 1 but got %d", count)
		}
	})
	t.Run("verify", func(t *testing.T) {
		users := getUsersInTeam(t, db, teamId)
		if len(users) != 1 {
			t.Fatal("expected 1 record")
		}
		if users[0] != userId {
			t.Fatalf("userId mismatch %v", users[0])
		}
	})
	t.Run("delete user from team", func(t *testing.T) {
		if count := removeUserFromTeam(t, db, userId, teamId); count != 1 {
			t.Fatalf("count expects 1 but got %d", count)
		}
	})

}
