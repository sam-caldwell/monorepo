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
	t.Skip("disabled for debugging")
	const (
		avatarHash          = "54742f98d09fdde6dbbb77d912c9e02448cd979590c19ecbf26039715db2e603"
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

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from avatar where hash='%s';", avatarHash)
		_, _ = db.Query("delete from %s where id='%s'", tableName, userId)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{userId},"+
			"pt:{uuid},"+
			"rt:jsonb", strings.ToLower(functionName)))

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
		var actualUser TrackerUser
		if err = json.Unmarshal([]byte(raw), &actualUser); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
	})
}
