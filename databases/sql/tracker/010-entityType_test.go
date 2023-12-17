package psqlTrackerDb

import (
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"testing"
)

func TestSqlDbType_entityType(t *testing.T) {
	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		err := db.Close()
		sqldbtest.CheckError(t, err)
	})
	t.Run("verify the enumerated type values", func(t *testing.T) {
		actual := sqldbtest.GetEnumValues(t, db, "entityType")
		expected := []string{"avatar", "icon", "user", "team", "teamAssociation", "workflow", "workflow_step",
			"workflow_action", "ticket_type", "project", "projectTypes", "ticket", "attachment", "comment",
			"property", "other"}
		sqldbtest.CompareTwoStringLists(t, actual, expected)
	})

}
