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

func TestSqlDbFunc_getTeamById(t *testing.T) {
	const (
		avatarHash      = "6a8434431319cfc231f4313e4663657f63ef99890635c6c0981316c591bc8a19"
		avatarType      = "image/png"
		iconHash        = "b14e141d72fc7b31ef2fa19e034b2d4c615dcbf26f1ad41a12365615dbcfa2e1"
		iconType        = "image/png"
		functionName    = "getTeamById"
		userFirstName   = "Klaus"
		userLastName    = "Fuchs"
		userEmail       = "treason.weasel@example.com"
		userPhone       = "314.152.9246"
		userDescription = "Test description"
		teamName        = "Manhattan Project"
		teamDescription = "this is the team"
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
			"pn:{teamId},"+
			"pt:{uuid},"+
			"rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	iconId = createIcon(t, db, iconType, iconHash)
	ownerId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)
	teamId = createTeam(t, db, teamName, iconId, ownerId,
		"read", "read", "read", teamDescription)

	t.Run("call getTeamById()", func(t *testing.T) {
		var err error
		var raw string
		var rows *sql.Rows
		var team TrackerTeam
		if rows, err = db.Query("select getTeamById('%s');", teamId); err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
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
			t.Fatalf("Name mismatch. got: '%v'", team.Name)
		}
		if team.IconId != iconId {
			t.Fatalf("IconId mismatch")
		}
		if team.OwnerId != ownerId {
			t.Fatalf("OwnerId mismatch")
		}
		if team.Owner != "read" {
			t.Fatalf("Owner mismatch. Got '%v'", team.Owner)
		}
		if team.Team != "read" {
			t.Fatalf("Team mismatch Got '%v'", team.Team)
		}
		if team.Everyone != "read" {
			t.Fatalf("Everyone mismatch Got '%v'", team.Everyone)
		}
	})
}
