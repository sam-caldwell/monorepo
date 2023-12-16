package psqlTrackerDb

import (
	"testing"
)

func TestSqlDbFunc_removeUserFromTeam(t *testing.T) {
	t.Skip("disabled for debugging")
	//const (
	//	avatarHash          = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
	//	iconUrl             = "http://localhost/myfakeicon.ico"
	//	functionName        = "removeUserFromTeam"
	//	expectedFirstName   = "Augustus"
	//	expectedLastName    = "Caesar"
	//	expectedEmail       = "augustus.caesar%d@example.com"
	//	expectedPhone       = "222.315.999%d"
	//	expectedDescription = "Test description"
	//	expectedTeamName    = "Coliseum"
	//)
	//var avatarId uuid.UUID
	//var iconId uuid.UUID
	//var users []uuid.UUID
	//var teamId uuid.UUID
	//var ownerId uuid.UUID
	//
	//db := sqldbtest.InitializeTestDbConn(t)
	//
	//t.Cleanup(func() {
	//	// Note: we only clean up the avatar we expect to have created.
	//	//       this should safeguard against an accidental run on prod.
	//	_, _ = db.Query("delete from users where email='%s'", expectedEmail)
	//	_, _ = db.Query("delete from teams where name='%s'", expectedTeamName)
	//	_, _ = db.Query("delete from avatars where url='%s'", avatarUrl)
	//	_, _ = db.Query("delete from icons where url='%s'", iconUrl)
	//	for _, id := range users {
	//		_, _ = db.Query("delete from teamMembership where userId='%s'", id)
	//	}
	//	err := db.Close()
	//	sqldbtest.CheckError(t, err)
	//})
	//
	//t.Run("verify the function structure (params, return)", func(t *testing.T) {
	//	sqldbtest.VerifyFunctionStructure(t, db,
	//		strings.ToLower(functionName),
	//		fmt.Sprintf("fn:%s,"+
	//			"pn:{uid,tid},"+
	//			"pt:{uuid},"+
	//			"rt:int4", strings.ToLower(functionName)))
	//})
	//
	//t.Run("call createAvatar()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select createAvatar('%s'::mimeType,'%s');", avatarHash, avatarType)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	err = rows.Scan(&raw)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if avatarId, err = uuid.Parse(raw); err != nil {
	//		t.Fatal(err)
	//	}
	//	if avatarId.String() == "00000000-0000-0000-0000-000000000000" {
	//		t.Fatal("illegal zero uuid")
	//	}
	//})
	//
	//t.Run("call createIcons()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select createIcons('%s');", iconUrl)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	err = rows.Scan(&raw)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if iconId, err = uuid.Parse(raw); err != nil {
	//		t.Fatal(err)
	//	}
	//})
	//
	//for i := 0; i < 3; i++ {
	//	t.Run("call createUser() to create 3 users", func(t *testing.T) {
	//		var rows *sql.Rows
	//		var userId uuid.UUID
	//		var err error
	//		email := fmt.Sprintf(expectedEmail, i)
	//		phone := fmt.Sprintf(expectedPhone, i)
	//		rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
	//			expectedFirstName, expectedLastName, avatarId, email, phone, expectedDescription)
	//		if err != nil {
	//			t.Fatal(err)
	//		}
	//		defer func() { _ = rows.Close() }()
	//		if !rows.Next() {
	//			t.Fatal("no row returned")
	//		}
	//		var raw string
	//		err = rows.Scan(&raw)
	//		if userId, err = uuid.Parse(raw); err != nil {
	//			t.Fatal(err)
	//		}
	//		if userId.String() == "00000000-0000-0000-0000-000000000000" {
	//			t.Fatal("illegal zero uuid")
	//		}
	//		users = append(users, userId)
	//	})
	//}
	//
	//ownerId = users[0]
	//
	//t.Run("call createTeam()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select createTeam('%s','%s','%s','%s','%s','%s','%s');",
	//		expectedTeamName, iconId, ownerId, "read", "read", "read", expectedDescription)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//
	//	err = rows.Scan(&raw)
	//	if teamId, err = uuid.Parse(raw); err != nil {
	//		t.Fatal(err)
	//	}
	//})
	//
	//for _, userId := range users {
	//	t.Run("call addUserToTeam()", func(t *testing.T) {
	//		var err error
	//		var rows *sql.Rows
	//		rows, err = db.Query("select addUserToTeam('%s','%s');", userId, teamId)
	//		if err != nil {
	//			t.Fatal(err)
	//		}
	//		defer func() { _ = rows.Close() }()
	//		if !rows.Next() {
	//			t.Fatal("no row returned")
	//		}
	//		var count int
	//		if err = rows.Scan(&count); err != nil {
	//			t.Fatal(err)
	//		}
	//		if count != 1 {
	//			t.Fatalf("expected count 1 but got %d", count)
	//		}
	//	})
	//}
	//
	//t.Run("Remove users[0] from the team", func(t *testing.T) {
	//	var err error
	//	var rows *sql.Rows
	//	rows, err = db.Query("select removeUserFromTeam('%s','%s');", users[0], teamId)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var count int
	//	if err = rows.Scan(&count); err != nil {
	//		t.Fatal(err)
	//	}
	//	if count != 1 {
	//		t.Fatalf("expected count 1 but got %d", count)
	//	}
	//})
	//
	//t.Run("confirm count after delete", func(t *testing.T) {
	//	var err error
	//	var rows *sql.Rows
	//	rows, err = db.Query(""+
	//		"select count(*) "+
	//		"from teamMembership "+
	//		"where userId='%s' "+
	//		"and teamId='%s'",
	//		users[0], teamId)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var count int
	//	if err = rows.Scan(&count); err != nil {
	//		t.Fatal(err)
	//	}
	//	if count != 0 {
	//		t.Fatalf("expected count 0 but got %d", count)
	//	}
	//})
}
