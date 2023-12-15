package psqlTrackerDb

import (
	"testing"
)

func TestSqlDbFunc_getAvatarById(t *testing.T) {
	//const (
	//	tableName     = "avatar"
	//	functionName  = "getAvatarById"
	//	testAvatarUrl = "http://example.com/itMayExists.png"
	//)
	//db := sqldbtest.InitializeTestDbConn(t)
	//
	//t.Cleanup(func() {
	//	// Note: we only clean up the avatar we expect to have created.
	//	//       this should safeguard against an accidental run on prod.
	//	_, _ = db.Query("delete from %s where url='%s';", tableName, testAvatarUrl)
	//
	//	err := db.Close()
	//	sqldbtest.CheckError(t, err)
	//})
	//
	//t.Run("verify the function structure (params, return)", func(t *testing.T) {
	//	sqldbtest.VerifyFunctionStructure(t, db,
	//		strings.ToLower(functionName),
	//		fmt.Sprintf("fn:%s,"+
	//			"pn:{targetId},"+
	//			"pt:{uuid},"+
	//			"rt:text", strings.ToLower(functionName)))
	//})
	//
	//var avatarId uuid.UUID
	//t.Run("create an avatar record", func(t *testing.T) {
	//	rows, err := db.Query("select createAvatar('%s');", testAvatarUrl)
	//	if err != nil {
	//		t.Fatalf("Failed to create record: %v", err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//
	//	t.Run("createAvatar() should return a row", func(t *testing.T) {
	//		if !rows.Next() {
	//			t.Fatal("no row returned")
	//		}
	//		if err := rows.Scan(&avatarId); err != nil {
	//			t.Fatalf("err: %v", err)
	//		}
	//	})
	//})
	//
	//t.Run("call getAvatarById() and inspect results", func(t *testing.T) {
	//	rows, err := db.Query("SELECT getAvatarById('%s')", avatarId)
	//	if err != nil {
	//		t.Fatalf("Failed to get record: %v", err)
	//	}
	//	defer func() { _ = rows.Close() }()
	//	if !rows.Next() {
	//		t.Fatal("Fail: no row returned")
	//	}
	//	var avatarUrl string
	//	if err := rows.Scan(&avatarUrl); err != nil {
	//		t.Fatalf("Failed to read result: %v", err)
	//	}
	//	if avatarUrl != testAvatarUrl {
	//		t.Fatalf("Fail: avatarUrl not as expected.\n"+
	//			"Wanted: %s\n"+
	//			"got:    %s", avatarUrl, testAvatarUrl)
	//	}
	//})
}
