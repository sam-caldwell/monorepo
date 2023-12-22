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
		expected := []string{
			"attachment",
			"avatar",
			"comment",
			"icon",
			"project",
			"projectTypes",
			"property",
			"team",
			"teamAssociation",
			"ticket",
			"ticketType",
			"ticketTypeAssociation",
			"user",
			"workflow",
			"workflowAction",
			"workflowStep",
		}
		sqldbtest.CompareTwoStringLists(t, actual, expected)
	})
}
