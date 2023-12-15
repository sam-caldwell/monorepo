package psqlTrackerDb

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_deleteIcon(t *testing.T) {
	const (
		tableName    = "icons"
		functionName = "deleteIcon"
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
			"pn:{targetId},"+
			"pt:{uuid},"+
			"rt:int4", strings.ToLower(functionName)))

	t.Run("create an record", func(t *testing.T) {
		entityId = createIcon(t, db, testType, testHash)
	})
	t.Run("delete the record", func(t *testing.T) {
		if count := deleteIcon(t, db, entityId); count != 1 {
			t.Fatalf("Failed to delete the record. count: %d", count)
		}
	})

	t.Run("verify outcome", func(t *testing.T) {
		if count := countById(t, db, "avatars", entityId); count != 0 {
			t.Fatalf("expected count 0 but got %d", count)
		}
	})
}
