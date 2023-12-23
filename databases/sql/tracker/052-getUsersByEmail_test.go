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

func TestSqlDbFunc_getUsersByEmail(t *testing.T) {
	const (
		avatarHash          = "153398fef5f49ace88d82f0bb2f19ffa05f794904adeff6e54441326e2a7714c"
		avatarType          = "image/png"
		functionName        = "getUsersByEmail"
		expectedFirstName   = "Alan"
		expectedLastName    = "Turing"
		expectedEmail       = "Alan.Turing@example.com"
		expectedPhone       = "713.123.4567"
		expectedDescription = "Test description"
	)
	var avatarId uuid.UUID
	var userId uuid.UUID
	var actualUser []TrackerUser

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "users", userId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{emailaddress,pageLimit,pageOffset},"+
			"pt:{int4,varchar},"+
			"rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	userId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

	t.Run("call getUsersByEmail(emailAddress)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select getUsersByEmail('%s',1000,0);", expectedEmail)
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
		if actualUser[0].Id != userId {
			t.Fatalf("userId mismatch")
		}
		if actualUser[0].FirstName != expectedFirstName {
			t.Fatalf("firstname mismatch")
		}
		if actualUser[0].LastName != expectedLastName {
			t.Fatalf("lastname mismatch")
		}
		if actualUser[0].Email != expectedEmail {
			t.Fatalf("email mismatch")
		}
		if actualUser[0].PhoneNumber != expectedPhone {
			t.Fatalf("phone mismatch")
		}
		if actualUser[0].Description != expectedDescription {
			t.Fatalf("description mismatch")
		}
	})
}
