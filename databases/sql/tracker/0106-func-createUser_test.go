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
	type TrackerUser struct {
		Id          uuid.UUID `yaml:"id"`
		FirstName   string    `yaml:"firstName"`
		LastName    string    `yaml:"lastName"`
		AvatarId    uuid.UUID `yaml:"avatarId"`
		Email       string    `yaml:"email"`
		PhoneNumber string    `yaml:"phoneNumber"`
		Description string    `yaml:"description"`
	}

	const (
		avatarUrl           = "http://localhost/myfakeavatar.jpeg"
		functionName        = "createUser"
		tableName           = "user"
		expectedFirstName   = "Wendell"
		expectedLastName    = "Fertig"
		expectedEmail       = "wendell.fertig@example.com"
		expectedPhone       = "512.712.3095"
		expectedDescription = "Test description"
	)
	var avatarId uuid.UUID
	var userId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where email='%s'", tableName, expectedEmail)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{firstname,lastname,avatarid,email,phonenumber,description},"+
				"pt:{text,varchar,uuid},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("create avatarId", func(t *testing.T) {
		/*
		 * We need to create an avatar to create a user
		 */
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

	t.Run("create createUser (userId)", func(t *testing.T) {
		/*
		 * create the user...
		 */
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

	//t.Run("inspect and verify user", func(t *testing.T) {
	//    /*
	//     * verify the user.
	//     */
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select id,firstName,lastName,avatarId,phoneNumber,description "+
	//		"from users where id='%s'", userId)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	if err := rows.Scan(&raw); err != nil {
	//		t.Fatalf("Failed to read result: %v", err)
	//	}
	//	if strings.TrimSpace(raw) == "" {
	//		t.Fatalf("Fail: unexpected empty rawResult")
	//	}
	//	var actualUser TrackerUser
	//	if err := json.Unmarshal([]byte(raw), &actualUser); err != nil {
	//		t.Fatalf("Failed to decode expected JSON: %v\n"+
	//			"expected Id: %v\n"+
	//			"got: %s",
	//			err, avatarId, raw)
	//	}
	//	if actualUser.Id != avatarId {
	//		t.Fatalf("Fail: avatarId not as expected.\n"+
	//			"Wanted: %s\n"+
	//			"got:    %s", avatarId, actualUser.Id)
	//	}
	//	if actualUser.FirstName == expectedFirstName {
	//		t.Fatal("FirstName mismatch")
	//	}
	//	if actualUser.LastName == expectedLastName {
	//		t.Fatal("LastName mismatch")
	//	}
	//	if actualUser.AvatarId == avatarId {
	//		t.Fatal("AvatarId mismatch")
	//	}
	//	if actualUser.Email == expectedEmail {
	//		t.Fatal("Email mismatch")
	//	}
	//	if actualUser.PhoneNumber == expectedPhone {
	//		t.Fatal("PhoneNumber mismatch")
	//	}
	//	if actualUser.Description == expectedDescription {
	//		t.Fatal("Description mismatch")
	//	}
	//})
}
