package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_updateUserName(t *testing.T) {
	const (
		avatarHash          = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
		functionName        = "updateUserName"
		tableName           = "user"
		originalFirstName   = "Anthony"
		newFirstName        = "Hannibal"
		originalLastName    = "HopKins"
		newLastName         = "Lecter"
		emailAddress        = "anthony.hopkins@example.com"
		expectedPhone       = "866.123.4567"
		expectedDescription = "original description"
	)
	var avatarId uuid.UUID
	var userId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where email='%s'", tableName, emailAddress)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{userId,newFirstName,newLastName},"+
				"pt:{varchar,uuid},"+
				"rt:int4", strings.ToLower(functionName)))
	})

	t.Run("call createAvatar() to create the original AvatarId", func(t *testing.T) {
		/*
		 * We need to create an avatar to create a user
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createAvatar('%s'::mimeType,'%s');", avatarHash, avatarType)
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

	t.Run("call createUser()", func(t *testing.T) {
		/*
		 * create the user...
		 */
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
			originalFirstName, originalLastName, avatarId, emailAddress, expectedPhone, expectedDescription)
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
		if actualFirstName != originalFirstName {
			t.Fatalf("FirstName mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualFirstName, originalFirstName)
		}
		if actualLastName != originalLastName {
			t.Fatalf("LastName mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualLastName, originalLastName)
		}
		if actualAvatarId != avatarId {
			t.Fatalf("AvatarId mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualAvatarId, avatarId)
		}
		if actualEmail != emailAddress {
			t.Fatalf("Email mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualEmail, emailAddress)
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

	t.Run("call updateUserName(userId,newFirstName,newLastName)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select updateUserName('%s','%s','%s');", userId, newFirstName, newLastName)
		if err != nil {
			t.Fatalf("updateUserEmail() failed\n"+
				"userId:        %v\n"+
				"newFirstName: %v\n"+
				"newLastName:      %v\n"+
				"err: %v",
				userId, newFirstName, newLastName, err)
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
		if actualFirstName != newFirstName {
			t.Fatalf("FirstName mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualFirstName, newFirstName)
		}
		if actualLastName != newLastName {
			t.Fatalf("LastName mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualLastName, newLastName)
		}
		if actualAvatarId != avatarId {
			t.Fatalf("AvatarId mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualAvatarId, avatarId)
		}
		if actualEmail != emailAddress {
			t.Fatalf("Email mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualEmail, emailAddress)
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
