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

func TestSqlDbFunc_getTeamByOwnerId(t *testing.T) {
	const (
		avatarHash      = "7470786da2026111d353918be67c48f136706d64ed176ae02237daa274ec1faf"
		avatarType      = "image/png"
		iconHash        = "6a7d7cbaa5472de7d1376829ccaf334bd8837084fab1218214e0107950585745"
		iconType        = "image/png"
		functionName    = "getTeamByOwnerId"
		userFirstName   = "Marie"
		userLastName    = "Curie"
		userEmail       = "glowing.physicist@example.com"
		userPhone       = "314.153.9356"
		userDescription = "Test description"
		teamName        = "Amazing Discoverer"
		teamDescription = "this is the team"
		pRead           = "read"
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
			"pn:{teamOwnerId},"+
			"pt:{uuid},"+
			"rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)
	teamId = createTeam(t, db, teamName, iconId, ownerId, pRead, pRead, pRead, teamDescription)

	t.Run("getTeamByOwnerId()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		if rows, err = db.Query("select getTeamByOwnerId('%s');", ownerId); err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		var team TrackerTeam
		if err = rows.Scan(&raw); err != nil {
			t.Fatal(err)
		}

		if err = json.Unmarshal([]byte(raw), &team); err != nil {
			t.Fatal(err)
		}
		if team.Id != teamId {
			t.Fatalf("teamId mismatch")
		}
		if team.Name != teamName {
			t.Fatalf("Name mismatch")
		}
		if team.IconId != iconId {
			t.Fatalf("IconId mismatch")
		}
		if team.OwnerId != ownerId {
			t.Fatalf("OwnerId mismatch")
		}
		if team.Owner != "read" {
			t.Fatalf("Owner mismatch. Got %v", team.Owner)
		}
		if team.Team != "read" {
			t.Fatalf("Team mismatch Got %v", team.Team)
		}
		if team.Everyone != "read" {
			t.Fatalf("Everyone mismatch Got %v", team.Everyone)
		}
	})
}
