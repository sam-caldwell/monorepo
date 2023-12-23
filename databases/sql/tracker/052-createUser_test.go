package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createUser(t *testing.T) {
	const (
		functionName    = "createUser"
		avatarHash      = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
		avatarType      = "image/png"
		userFirstName   = "Wendell"
		userLastName    = "Fertig"
		userEmail       = "wendell.fertig@example.com"
		userPhone       = "512.123.4567"
		userDescription = "Test description"
	)
	var avatarId uuid.UUID
	var userId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_ = cleanUpObject(db, "users", userId)
		_ = cleanUpObject(db, "avatars", avatarId)
		sqldbtest.CheckError(t, db.Close())
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{newFirstName,newLastName,newAvatarId,newEmailAddress,newPhoneNumber,newDescription},"+
				"pt:{text,varchar,uuid},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	userId = createUser(t, db, userFirstName, userLastName, avatarId, userEmail, userPhone, userDescription)

	t.Run("verify the user record", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		var actualId, actualAvatarId uuid.UUID
		var actualFirstName, actualLastName, actualEmail, actualPhone, actualDescription sql.NullString

		rows, err = db.Query("select id,firstName,lastName,avatarId,email,phoneNumber,description "+
			"from users where id='%s'", userId)
		if err != nil {
			t.Fatal(err)
		}

		defer func() { _ = rows.Close() }()

		if !rows.Next() {
			t.Fatal("no row returned")
		}

		err = rows.Scan(&actualId, &actualFirstName, &actualLastName, &actualAvatarId,
			&actualEmail, &actualPhone, &actualDescription)
		if err != nil {
			t.Fatalf("Failed to read result: %v", err)
		}

		if actualId != userId {
			t.Fatal("UserId mismatch")
		}

		if actualFirstName.String != userFirstName {
			t.Fatalf("FirstName mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualFirstName.String, userFirstName)
		}

		if actualLastName.String != userLastName {
			t.Fatalf("LastName mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualLastName.String, userLastName)
		}

		if actualAvatarId != avatarId {
			t.Fatalf("AvatarId mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualAvatarId, avatarId)
		}

		if actualEmail.String != userEmail {
			t.Fatalf("Email mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualEmail.String, userEmail)
		}

		if actualPhone.String != userPhone {
			t.Fatalf("PhoneNumber mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualPhone.String, userPhone)
		}

		if actualDescription.String != userDescription {
			t.Fatalf("Description mismatch\n"+
				"actual:   '%s'\n"+
				"expected: '%s'", actualDescription.String, userDescription)
		}
	})
}
