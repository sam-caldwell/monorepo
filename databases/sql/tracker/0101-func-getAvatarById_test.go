package psqlTrackerDb

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getAvatarById(t *testing.T) {
	type TrackerAvatar struct {
		Id  string `json:"id"`
		Url string `json:"url"`
	}

	const (
		tableName    = "avatar"
		functionName = "getAvatarById"
		testUrl      = "http://example.com/itMayExists.png"
	)
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where url='%s';", tableName, testUrl)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{avatarid},"+
				"pt:{uuid},"+
				"rt:jsonb", strings.ToLower(functionName)))
	})

	var avatarId uuid.UUID
	t.Run("create an avatar record", func(t *testing.T) {
		rows, err := db.Query("select createAvatar('%s');", testUrl)
		if err != nil {
			t.Fatalf("Failed to create record: %v", err)
		}

		t.Run("createAvatar() should return a row", func(t *testing.T) {
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			if err := rows.Scan(&avatarId); err != nil {
				t.Fatalf("err: %v", err)
			}
		})
	})

	t.Run("call getAvatarById() and inspect results", func(t *testing.T) {
		rows, err := db.Query("SELECT getAvatarById('%s')", avatarId)
		if err != nil {
			t.Fatalf("Failed to get record: %v", err)
		}
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var raw string
		if err := rows.Scan(&raw); err != nil {
			t.Fatalf("Failed to read result: %v", err)
		}
		if strings.TrimSpace(raw) == "" {
			t.Fatalf("Fail: unexpected empty rawResult")
		}
		var decodedResult TrackerAvatar
		if err := json.Unmarshal([]byte(raw), &decodedResult); err != nil {
			t.Fatalf("Failed to decode expected JSON: %v\n"+
				"expected Id: %v\n"+
				"got: %s",
				err, avatarId, raw)
		}

		if decodedResult.Id != avatarId.String() {
			t.Fatalf("Fail: avatarId not as expected.\n"+
				"Wanted: %s\n"+
				"got:    %s", avatarId, decodedResult.Id)
		}
		if strings.TrimSpace(decodedResult.Url) == "" {
			t.Fatalf("Fail: avatar url not as expected.")
		}
	})

}
