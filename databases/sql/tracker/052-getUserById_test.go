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

func TestSqlDbFunc_getUserById(t *testing.T) {
	const (
		avatarHash          = "1262190e91301435e25eb58cc5897a9833b90fef85a1d4f5a1965fb6c070e8f2"
		avatarType          = "image/png"
		functionName        = "getUserById"
		tableName           = "user"
		expectedFirstName   = "John"
		expectedLastName    = "Von Neumann"
		expectedEmail       = "John.vonNeumann@example.com"
		expectedPhone       = "415.123.4567"
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
		fmt.Sprintf("fn:%s,pn:{userId},pt:{uuid},rt:jsonb", strings.ToLower(functionName)))

	avatarId = createAvatar(t, db, avatarType, avatarHash)
	userId = createUser(t, db, expectedFirstName, expectedLastName, avatarId, expectedEmail,
		expectedPhone, expectedDescription)

	t.Run("call getUserById(userId)", func(t *testing.T) {
		var rows *sql.Rows
		var err error
		rows, err = db.Query("select getUserById('%s');", userId)
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
