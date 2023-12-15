package psqlTrackerDb

import (
	"testing"
)

func TestSqlDbFunc_deleteWorkflowById(t *testing.T) {
	//const (
	//	avatarHash           = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
	//	avatarType           = "image/png"
	//	iconUrl              = "http://localhost/myfakeicon.ico"
	//	functionName         = "deleteWorkflowById"
	//	expectedFirstName    = "Wilbur"
	//	expectedLastName     = "Wright"
	//	expectedEmail        = "wilbur.wright@example.com"
	//	expectedPhone        = "002.002.0002"
	//	expectedDescription  = "Test description"
	//	expectedTeamName     = "WriteBrothersFlying"
	//	expectedWorkflowName = "KittyHawk"
	//)
	//var avatarId uuid.UUID
	//var iconId uuid.UUID
	//var teamId uuid.UUID
	//var ownerId uuid.UUID
	//var workflowId uuid.UUID
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
	//	_, _ = db.Query("delete from teamMembership where teamId='%s'", teamId)
	//	err := db.Close()
	//	sqldbtest.CheckError(t, err)
	//})
	//
	//t.Run("verify the function structure (params, return)", func(t *testing.T) {
	//	sqldbtest.VerifyFunctionStructure(t, db,
	//		strings.ToLower(functionName),
	//		fmt.Sprintf("fn:%s,"+
	//			"pn:{workflowId},"+
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
	//t.Run("call createUser()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select createUser('%s','%s','%s','%s','%s','%s');",
	//		expectedFirstName, expectedLastName, avatarId, expectedEmail, expectedPhone, expectedDescription)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	err = rows.Scan(&raw)
	//	if ownerId, err = uuid.Parse(raw); err != nil {
	//		t.Fatal(err)
	//	}
	//})
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
	//	err = rows.Scan(&raw)
	//	if teamId, err = uuid.Parse(raw); err != nil {
	//		t.Fatal(err)
	//	}
	//	t.Run("call addUserToTeam()", func(t *testing.T) {
	//		var rows *sql.Rows
	//		var err error
	//		rows, err = db.Query("select addUserToTeam('%s','%s');", ownerId, teamId)
	//		if err != nil {
	//			t.Fatal(err)
	//		}
	//		defer func() { _ = rows.Close() }()
	//		if !rows.Next() {
	//			t.Fatal("no row returned")
	//		}
	//		var count int
	//		err = rows.Scan(&count)
	//		if count != 1 {
	//			t.Fatalf("expected count 1 but got %d", count)
	//		}
	//	})
	//	t.Run("verify user membership", func(t *testing.T) {
	//		var rows *sql.Rows
	//		var err error
	//		rows, err = db.Query(""+
	//			"select count(userId) "+
	//			"from teamMembership "+
	//			"where userId='%s' "+
	//			"and teamId='%s';", ownerId, teamId)
	//		if err != nil {
	//			t.Fatal(err)
	//		}
	//		defer func() { _ = rows.Close() }()
	//		if !rows.Next() {
	//			t.Fatal("no row returned")
	//		}
	//		var count int
	//		err = rows.Scan(&count)
	//		if count != 1 {
	//			t.Fatalf("expected count 1 but got %d", count)
	//		}
	//	})
	//})
	//
	//t.Run("createWorkflow()", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select createWorkflow('%s','%s','%s','%s','%s','%s','%s','%s');",
	//		expectedWorkflowName, iconId, ownerId, teamId, "read", "read", "read", expectedDescription)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var raw string
	//	err = rows.Scan(&raw)
	//	if workflowId, err = uuid.Parse(raw); err != nil {
	//		t.Fatal(err)
	//	}
	//})
	//
	//t.Run("verify workflow", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query(""+
	//		"select id,name,iconId,ownerId,teamId,owner,team,everyone,description "+
	//		"from workflow where id = '%s';", workflowId)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var actualId uuid.UUID
	//	var actualName string
	//	var actualIconId uuid.UUID
	//	var actualOwnerId uuid.UUID
	//	var actualTeamId uuid.UUID
	//	var actualOwner string
	//	var actualTeam string
	//	var actualEveryone string
	//	var actualDescription string
	//	err = rows.Scan(&actualId, &actualName, &actualIconId, &actualOwnerId, &actualTeamId,
	//		&actualOwner, &actualTeam, &actualEveryone, &actualDescription)
	//
	//	if actualId != workflowId {
	//		t.Fatalf("workflowId mismatch")
	//	}
	//	if actualName != expectedWorkflowName {
	//		t.Fatalf("workflowName mismatch")
	//	}
	//	if actualIconId != iconId {
	//		t.Fatalf("iconId mismatch")
	//	}
	//	if actualOwnerId != ownerId {
	//		t.Fatalf("ownerId mismatch")
	//	}
	//	if actualTeamId != teamId {
	//		t.Fatalf("teamId mismatch")
	//	}
	//	if actualOwner != "read" {
	//		t.Fatalf("owner mismatch")
	//	}
	//	if actualTeam != "read" {
	//		t.Fatalf("team mismatch")
	//	}
	//	if actualEveryone != "read" {
	//		t.Fatalf("everyone mismatch")
	//	}
	//	if actualDescription != expectedDescription {
	//		t.Fatalf("description mismatch")
	//	}
	//})
	//
	//t.Run("delete workflow", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query(""+
	//		"select deleteWorkflowById('%s');", workflowId)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var count int
	//	err = rows.Scan(&count)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if count != 1 {
	//		t.Fatalf("expected count 1 but got %d", count)
	//	}
	//})
	//
	//t.Run("verify delete", func(t *testing.T) {
	//	var rows *sql.Rows
	//	var err error
	//	rows, err = db.Query("select count(*) from workflow where id='%s';", workflowId)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("no row returned")
	//	}
	//	var count int
	//	err = rows.Scan(&count)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if count != 0 {
	//		t.Fatalf("expected count 0 but got %d", count)
	//	}
	//})
}
