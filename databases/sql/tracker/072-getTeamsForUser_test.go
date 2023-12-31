package psqlTrackerDb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getTeamsForUser(t *testing.T) {
	const (
		avatarHash       = "a2601e31f65f266a1a94f08ad46918c8d0f9f09f995aa7fbdbfa113ad6911ba6"
		avatarType       = "image/png"
		iconHash         = "69357df9edaa759985b300c4d0341cd906bff5519ff55035a04b58c0af5237c3"
		iconType         = "image/png"
		functionName     = "getTeamsForUser"
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
		numberTeams      = 5
	)
	var avatarId uuid.UUID
	var iconId uuid.UUID
	var userId uuid.UUID
	var teamSet []uuid.UUID
	var ownerId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		rows, _ := db.Query("delete from teammemberships where userId='%s'", userId)
		defer func() { _ = rows.Close() }()
		for _, teamId := range teamSet {
			sqldbtest.CheckError(t, cleanUpObject(db, "teams", teamId))
		}
		sqldbtest.CheckError(t, cleanUpObject(db, "icons", iconId))
		sqldbtest.CheckError(t, cleanUpObject(db, "users", userId))
		sqldbtest.CheckError(t, cleanUpObject(db, "users", ownerId))
		sqldbtest.CheckError(t, cleanUpObject(db, "avatars", avatarId))
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{thisuserid},"+
			"pt:{uuid},"+
			"rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, ownerFirstName, ownerLastName, avatarId, ownerEmail, ownerPhone, ownerDescription)
	userId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)

	for i := 0; i < numberTeams; i++ {
		thisName := fmt.Sprintf("%v%d", teamName, i)
		teamId := createTeam(t, db, thisName, iconId, ownerId, pRead, pRead, pRead, userDescription)
		teamSet = append(teamSet, teamId)
		if count := addUserToTeam(t, db, userId, teamId); count != 1 {
			t.Fatalf("count expects 1 but got %d", count)
		}
	}

	type CollectionOfTeams struct {
		TeamID string `json:"teamId"`
	}

	t.Run("getTeamsForUser() and verify", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select getTeamsForUser('%s');", userId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()

		var actualTeams []CollectionOfTeams
		for rows.Next() {
			var raw string
			if err = rows.Scan(&raw); err != nil {
				t.Fatal(err)
			}
			if err = json.Unmarshal([]byte(raw), &actualTeams); err != nil {
				t.Fatalf("error: failed to parse: %v\n"+
					"raw: %v", err, raw)
			}
		}

		if len(actualTeams) != len(teamSet) {
			t.Fatalf("Error: expected count %d but got %d", len(teamSet), len(actualTeams))
		}
	})
}
