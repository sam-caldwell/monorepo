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

func TestSqlDbFunc_getUserByEmail(t *testing.T) {
	const (
		avatarHash          = "153398fef5f49ace88d82f0bb2f19ffa05f794904adeff6e54441326e2a7714c"
		avatarType          = "image/png"
		functionName        = "getUserByEmail"
		expectedFirstName   = "Alan"
		expectedLastName    = "Turing"
		expectedEmail       = "Alan.Turing@example.com"
		expectedPhone       = "713.123.4567"
		expectedDescription = "Test description"
	)
	var avatarId uuid.UUID
	var userId uuid.UUID
	var actualUser TrackerUser

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from avatar where hash='%s';", avatarHash)
		_, _ = db.Query("delete from user where id='%s'", userId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{emailaddress},"+
			"pt:{varchar},"+
			"rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	userId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

	t.Run("call getUserByEmail(emailAddress)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select getUserByEmail('%s');", expectedEmail)
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
