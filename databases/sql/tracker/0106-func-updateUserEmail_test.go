package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateUserEmail(t *testing.T) {
	const (
		originalAvatarUrl   = "http://localhost/myfakeavatar.jpeg"
		functionName        = "updateUserEmail"
		tableName           = "user"
		expectedFirstName   = "Ken"
		expectedLastName    = "Thompson"
		originalEmail       = "ken.thompson@example.com"
		newEmail            = "new.email@example.com"
		expectedPhone       = "555.123.4567"
		expectedDescription = "original description"
	)
	var avatarId uuid.UUID
	var userId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where email='%s'", tableName, newEmail)
		_, _ = db.Query("delete from %s where email='%s'", tableName, originalEmail)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{userId,newEmail},"+
				"pt:{varchar,uuid},"+
				"rt:int4", strings.ToLower(functionName)))
	})

	t.Run("call createAvatar() to create the original AvatarId", func(t *testing.T) {
		/*
		 * We need to create an avatar to create a user
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createAvatar('%s');", originalAvatarUrl)
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

	t.Run("call createUser()", func(t *testing.T) {
		/*
		 * create the user...
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
			expectedFirstName, expectedLastName, avatarId, originalEmail, expectedPhone, expectedDescription)
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

	t.Run("inspect and verify user", func(t *testing.T) {
		/*
		 * verify the user.
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select id,firstName,lastName,avatarId,email,phoneNumber,description "+
			"from users where id='%s'", userId)
		if err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}

		var actualId, actualAvatarId uuid.UUID
		var actualFirstName, actualLastName, actualEmail, actualPhone, actualDescription string

		err = rows.Scan(&actualId, &actualFirstName, &actualLastName, &actualAvatarId, &actualEmail, &actualPhone,
			&actualDescription)
		if err != nil {
			t.Fatalf("Failed to read result: %v", err)
		}
		if actualId != userId {
			t.Fatal("UserId mismatch")
		}
		if actualFirstName != expectedFirstName {
			t.Fatalf("FirstName mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualFirstName, expectedFirstName)
		}
		if actualLastName != expectedLastName {
			t.Fatalf("LastName mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualLastName, expectedLastName)
		}
		if actualAvatarId != avatarId {
			t.Fatalf("AvatarId mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualAvatarId, avatarId)
		}
		if actualEmail != originalEmail {
			t.Fatalf("Email mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualEmail, originalEmail)
		}
		if actualPhone != expectedPhone {
			t.Fatalf("PhoneNumber mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualPhone, expectedPhone)
		}
		if actualDescription != expectedDescription {
			t.Fatalf("Description mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualDescription, expectedDescription)
		}
	})

	t.Run("call updateUserEmail(userId,newEmail)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select updateUserEmail('%s','%s');", userId, newEmail)
		if err != nil {
			t.Fatalf("updateUserEmail() failed\n"+
				"userId:        %v\n"+
				"originalEmail: %v\n"+
				"newEmail:      %v\n"+
				"err: %v",
				userId, originalEmail, newEmail, err)
		}
		if !rows.Next() {
			t.Fatal("no row returned")
		}
		var count int
		if err = rows.Scan(&count); err != nil {
			t.Fatal(err)
		}
		if count != 1 {
			t.Fatalf("expected updated count 1 but got %d", count)
		}
	})
}
