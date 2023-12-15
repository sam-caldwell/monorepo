package psqlTrackerDb

import (
	"testing"
)

func TestSqlDbFunc_getTeamByOwnerId(t *testing.T) {
	//const (
	//	avatarHash          = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
	//	iconUrl             = "http://localhost/myfakeicon.ico"
	//	functionName        = "getTeamByOwnerId"
	//	expectedFirstName   = "Marie"
	//	expectedLastName    = "Curie"
	//	expectedEmail       = "glowing.physicist@example.com"
	//	expectedPhone       = "314.153.9356"
	//	expectedDescription = "Test description"
	//	expectedTeamName    = "Amazing Discoverer"
	//)
	//var avatarId uuid.UUID
	//var iconId uuid.UUID
	//var teamId uuid.UUID
	//var userId uuid.UUID
	//
	//db := sqldbtest.InitializeTestDbConn(t)
	//
	//t.Cleanup(func() {
	//	// Note: we only clean up the avatar we expect to have created.
	//	//       this should safeguard against an accidental run on prod.
	//	_, _ = db.Query("delete from %s where email='%s'", "users", expectedEmail)
	//	_, _ = db.Query("delete from %s where name='%s'", "teams", expectedTeamName)
	//	_, _ = db.Query("delete from %s where url='%s'", "avatars", avatarUrl)
	//	_, _ = db.Query("delete from %s where url='%s'", "icons", iconUrl)
	//	err := db.Close()
	//	sqldbtest.CheckError(t, err)
	//})
	//
	//t.Run("verify the function structure (params, return)", func(t *testing.T) {
	//	sqldbtest.VerifyFunctionStructure(t, db,
	//		strings.ToLower(functionName),
	//		fmt.Sprintf("fn:%s,"+
	//			"pn:{teamOwnerId},"+
	//			"pt:{uuid},"+
	//			"rt:jsonb", strings.ToLower(functionName)))
	//})
	//
	//t.Run("call createAvatar()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select createAvatar('%s'::mimeType,'%s');", avatarHash, avatarType)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
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
	//t.Run("call createUser()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
	//		expectedFirstName, expectedLastName, avatarId, expectedEmail, expectedPhone, expectedDescription)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	err = rows.Scan(&raw)
	//	if userId, err = uuid.Parse(raw); err != nil {
	//		t.Fatal(err)
	//	}
	//	if userId.String() == "00000000-0000-0000-0000-000000000000" {
	//		t.Fatal("illegal zero uuid")
	//	}
	//})
	//
	//t.Run("call createTeam()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select createTeam('%s','%s','%s','%s','%s','%s','%s');",
	//		expectedTeamName, iconId, userId, "read", "read", "read", expectedDescription)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	if err = rows.Scan(&raw); err != nil {
	//		t.Fatal(err)
	//	}
	//	teamId, err = uuid.Parse(raw)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//})
	//
	//t.Run("getTeamByOwnerId()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	if rows, err = db.Query("select getTeamByOwnerId('%s');", userId); err != nil {
	//		t.Fatal(err)
	//	}
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	var team TrackerTeam
	//	if err = rows.Scan(&raw); err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	if err = json.Unmarshal([]byte(raw), &team); err != nil {
	//		t.Fatal(err)
	//	}
	//	if team.Id != teamId {
	//		t.Fatalf("teamId mismatch")
	//	}
	//	if team.Name != expectedTeamName {
	//		t.Fatalf("Name mismatch")
	//	}
	//	if team.IconId != iconId {
	//		t.Fatalf("IconId mismatch")
	//	}
	//	if team.OwnerId != userId {
	//		t.Fatalf("OwnerId mismatch")
	//	}
	//	if team.Owner != "read" {
	//		t.Fatalf("Owner mismatch. Got %v", team.Owner)
	//	}
	//	if team.Team != "read" {
	//		t.Fatalf("Team mismatch Got %v", team.Team)
	//	}
	//	if team.Everyone != "read" {
	//		t.Fatalf("Everyone mismatch Got %v", team.Everyone)
	//	}
	//})
}
