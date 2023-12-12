package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getTeamsForUser(t *testing.T) {
	const (
		avatarUrl           = "http://localhost/myfakeavatar.jpeg"
		iconUrl             = "http://localhost/myfakeicon.ico"
		functionName        = "getTeamsForUser"
		expectedFirstName   = "Benoit"
		expectedLastName    = "Mandelbrot"
		expectedEmail       = "benoit.mandelbrot@example.com"
		expectedPhone       = "432.192.9746"
		expectedDescription = "Test description"
		expectedTeamName    = "FractalMaker"
	)
	var avatarId uuid.UUID
	var iconId uuid.UUID
	var teams []uuid.UUID
	var userId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where email='%s'", "users", expectedEmail)
		_, _ = db.Query("delete from %s where name='%s'", "teams", expectedTeamName)
		_, _ = db.Query("delete from %s where url='%s'", "avatars", avatarUrl)
		_, _ = db.Query("delete from %s where url='%s'", "icons", iconUrl)
		for _, teamId := range teams {
			_, _ = db.Query("delete from %s where teamId='%s'", "teamMembership", teamId)
		}
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{userId},"+
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

	for i := 0; i < 3; i++ {
		t.Run("call createTeam()", func(t *testing.T) {
			var rows *sql.Rows
			var err error
			teamName := fmt.Sprintf("%s %d", expectedTeamName, i)

			rows, err = db.Query("select createTeam('%s','%s','%s','%s','%s','%s','%s');",
				teamName, iconId, userId, "read", "read", "read", expectedDescription)
			if err != nil {
				t.Fatal(err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var raw string
			var teamId uuid.UUID
			err = rows.Scan(&raw)
			if teamId, err = uuid.Parse(raw); err != nil {
				t.Fatal(err)
			}
			t.Run("call addUserToTeam()", func(t *testing.T) {
				var rows *sql.Rows
				var err error
				rows, err = db.Query("select addUserToTeam('%s','%s');", userId, teamId)
				if err != nil {
					t.Fatal(err)
				}
				if !rows.Next() {
					t.Fatal("no row returned")
				}
				var count int
				err = rows.Scan(&count)
				if count != 1 {
					t.Fatalf("expected count 1 but got %d", count)
				}
			})
			t.Run("verify user membership", func(t *testing.T) {
				var rows *sql.Rows
				var err error
				rows, err = db.Query(""+
					"select count(userId) "+
					"from teamMembership "+
					"where userId='%s' "+
					"and teamId='%s';", userId, teamId)
				if err != nil {
					t.Fatal(err)
				}
				if !rows.Next() {
					t.Fatal("no row returned")
				}
				var count int
				err = rows.Scan(&count)
				if count != 1 {
					t.Fatalf("expected count 1 but got %d", count)
				}
			})
			teams = append(teams, teamId)
		})
	}
	t.Run("getTeamsForUser() and verify", func(t *testing.T) {
		var rows *sql.Rows
		var err error

		rows, err = db.Query("select getTeamsForUser('%s');", userId)
		if err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		err = rows.Scan(&count)
		if count != len(teams) {
			t.Fatalf("expected count %d but got %d", len(teams), count)
		}
	})
}
