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
		avatarHash          = "f5fc32a488d160483076487deaf907b89fd46d76a8055819a0a62dc2fd55b111"
		avatarType          = "image/png"
		functionName        = "updateUserEmail"
		tableName           = "user"
		expectedFirstName   = "Ken"
		expectedLastName    = "Thompson"
		expectedEmail       = "ken.thompson@example.com"
		newEmail            = "new.email@example.com"
		expectedPhone       = "555.123.4567"
		expectedDescription = "original description"
	)
	var avatarId uuid.UUID
	var userId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from avatar where hash='%s';", avatarHash)
		_, _ = db.Query("delete from %s where id='%s'", tableName, userId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,pn:{userId,newEmail},pt:{varchar,uuid},rt:int4", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	userId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

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
		defer func() { _ = rows.Close() }()
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
		if actualEmail != expectedEmail {
			t.Fatalf("Email mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualEmail, expectedEmail)
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
				userId, expectedEmail, newEmail, err)
		}
		defer func() { _ = rows.Close() }()
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

	t.Run("verify user update", func(t *testing.T) {
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
		defer func() { _ = rows.Close() }()
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
		if actualEmail != newEmail {
			t.Fatalf("Email mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualEmail, newEmail)
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
}
