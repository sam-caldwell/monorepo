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

func TestSqlDbFunc_getUserByPhone(t *testing.T) {
	const (
		avatarHash          = "e8172424cebb838aed7e9278793d3aa46d99a97b5caf16d562ee35f7fb8fdf2c"
		avatarType          = "image/png"
		functionName        = "getUserByPhone"
		expectedFirstName   = "Ada"
		expectedLastName    = "Lovelace"
		expectedEmail       = "ada.lovelace@example.com"
		expectedPhone       = "915.123.4567"
		expectedDescription = "Test description"
	)
	var avatarId uuid.UUID
	var userId uuid.UUID
	var actualUser TrackerUser

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from users where id='%s'", userId)
		_, _ = db.Query("delete from avatars where id='%s'", avatarId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,pn:{phone},pt:{varchar},rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	userId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

	t.Run("call getUserByPhone(expectedPhone)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select getUserByPhone('%s');", expectedPhone)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var raw string
		if err = rows.Scan(&raw); err != nil {
			t.Fatal(err)
		}
		if err = json.Unmarshal([]byte(raw), &actualUser); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
	})
	t.Run("verify", func(t *testing.T) {
		if actualUser.Id != userId {
			t.Fatalf("userId mismatch")
		}
		if actualUser.FirstName != expectedFirstName {
			t.Fatalf("firstname mismatch")
		}
		if actualUser.LastName != expectedLastName {
			t.Fatalf("lastname mismatch")
		}
		if actualUser.Email != expectedEmail {
			t.Fatalf("email mismatch")
		}
		if actualUser.PhoneNumber != expectedPhone {
			t.Fatalf("phone mismatch")
		}
		if actualUser.Description != expectedDescription {
			t.Fatalf("description mismatch")
		}
	})
}
