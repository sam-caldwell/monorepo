package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_getIconsById(t *testing.T) {
	const (
		tableName    = "icons"
		functionName = "getIconById"
		testHash     = "b5bb9d8014a0f9b1d61e21e796d78dccdf1351f23cd32812f4850b878ae4944c"
		testType     = "image/png"
	)
	var entityId uuid.UUID

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		_, _ = db.Query("delete from %s where hash='%s';", tableName, testHash)
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	sqldbtest.VerifyFunctionStructure(t, db,
		strings.ToLower(functionName),
		fmt.Sprintf("fn:%s,"+
			"pn:{entityId},"+
			"pt:{uuid},"+
			"rt:jsonb", strings.ToLower(functionName)))

	t.Run("create an record", func(t *testing.T) {
		entityId = createIcon(t, db, testType, testHash)
	})

	t.Run("get the record and verify", func(t *testing.T) {
		var record TrackerIcon
		getIconById(t, db, entityId, &record)

		if record.Id != entityId {
			t.Fatalf("Fail: id mismatch\ngot '%v'", record.Id)
		}
		if record.Hash != testHash {
			t.Fatalf("Fail: hash mismatch\ngot '%v'", record.Hash)
		}
		if record.MimeType != testType {
			t.Fatalf("Fail: mime-type mismatch\ngot '%v'", record.MimeType)
		}
	})
}
