package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateUserAvatar(t *testing.T) {
	const (
		avatarHash          = "53cc4ca6f01c300005754f5c790f64bc8d99b2d4bdee6d67d5f03a76abc4a313"
		avatarType          = "image/png"
		newAvatarHash       = "53cc4ca6f01c300005754f5c790f64bc8d99b2d4bdee6d67d5f03a76abc4a313"
		functionName        = "updateUserAvatar"
		expectedFirstName   = "Grace"
		expectedLastName    = "Hopper"
		expectedEmail       = "grace.hopper@example.com"
		expectedPhone       = "214.123.4567"
		expectedDescription = "Test description"
	)
	var originalAvatarId uuid.UUID
	var userId uuid.UUID
	var newAvatarId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from avatars where hash='%s';", avatarHash)
		_, _ = db.Query("delete from avatars where hash='%s';", newAvatarHash)
		_, _ = db.Query("delete from users where id='%s'", userId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,pn:{userId,newAvatarId},pt:{uuid},rt:int4", strings.ToLower(functionName)))

	originalAvatarId = createAvatar(t, db, avatarType, avatarHash)
	newAvatarId = createAvatar(t, db, avatarType, newAvatarHash)
	userId = createUser(t, db, expectedFirstName, expectedLastName, originalAvatarId, expectedEmail,
		expectedPhone, expectedDescription)

	t.Run("call updateUserAvatar(userId,avatarId)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select updateUserAvatar('%s','%s');", userId, newAvatarId)
		if err != nil {
			t.Fatalf("updateUserAvatar() failed\n"+
				"userId:           %v\n"+
				"originalAvatarId: %v\n"+
				"newAvatarId:      %v\n"+
				"err: %v",
				userId, originalAvatarId, newAvatarId, err)
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
		if actualAvatarId != newAvatarId {
			t.Fatalf("AvatarId mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualAvatarId, newAvatarId)
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
}
