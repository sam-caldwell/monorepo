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
		avatarUrl           = "http://localhost/myfakeavatar.jpeg"
		iconUrl             = "http://localhost/myfakeicon.ico"
		functionName        = "getTeamById"
		expectedFirstName   = "Klaus"
		expectedLastName    = "Fuchs"
		expectedEmail       = "treason.weasel@example.com"
		expectedPhone       = "314.152.9246"
		expectedDescription = "Test description"
		expectedTeamName    = "Manhattan Project"
	)
	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teamId uuid.UUID
	var userId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where email='%s'", "users", expectedEmail)
		_, _ = db.Query("delete from %s where name='%s'", "teams", expectedTeamName)
		_, _ = db.Query("delete from %s where url='%s'", "avatars", avatarUrl)
		_, _ = db.Query("delete from %s where url='%s'", "icons", iconUrl)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{teamId},"+
				"pt:{uuid},"+
				"rt:jsonb", strings.ToLower(functionName)))
	})

	t.Run("call createAvatar()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createAvatar('%s');", avatarUrl)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if err != nil {
			t.Fatal(err)
		}
		if avatarId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
		if avatarId.String() == "00000000-0000-0000-0000-000000000000" {
			t.Fatal("illegal zero uuid")
		}
	})

	t.Run("call createIcons()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createIcons('%s');", iconUrl)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if err != nil {
			t.Fatal(err)
		}
		if iconId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("call createUser()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
			expectedFirstName, expectedLastName, avatarId, expectedEmail, expectedPhone, expectedDescription)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if userId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
		if userId.String() == "00000000-0000-0000-0000-000000000000" {
			t.Fatal("illegal zero uuid")
		}
	})

	t.Run("call createTeam()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createTeam('%s','%s','%s','%s','%s','%s','%s');",
			expectedTeamName, iconId, userId, "read", "read", "read", expectedDescription)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		err = rows.Scan(&raw)
		if teamId, err = uuid.Parse(raw); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("getTeamById()", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		if rows, err = db.Query("select getTeamById('%s');", teamId); err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
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
		if team.Name != expectedTeamName {
			t.Fatalf("Name mismatch")
		}
		if team.IconId != iconId {
			t.Fatalf("IconId mismatch")
		}
		if team.OwnerId != userId {
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
