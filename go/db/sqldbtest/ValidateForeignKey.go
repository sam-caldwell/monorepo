package database

import (
	"fmt"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"strings"
	"testing"
)

func ValidateForeignKey(t *testing.T, db *Postgres.Db, table, referencedTable, column, referencedColumn string) {
	query := fmt.Sprintf(`
       SELECT
	       count(*)
	   FROM
	       pg_constraint
	           JOIN pg_attribute AS lhs
	               ON lhs.attnum = ANY(pg_constraint.conkey) AND lhs.attrelid = pg_constraint.conrelid
	           JOIN pg_attribute AS rhs
	               ON rhs.attnum = ANY(pg_constraint.confkey) AND rhs.attrelid = pg_constraint.confrelid
	   WHERE
	           conrelid = '%s'::regclass
	       AND confrelid = '%s'::regclass
           AND lhs.attname = '%s'
           AND rhs.attname = '%s';`,
		strings.ToLower(table),
		strings.ToLower(referencedTable),
		strings.ToLower(column),
		strings.ToLower(referencedColumn))
	rows, err := db.Query(query)
	CheckError(t, err)
	defer func() {
		CheckError(t, rows.Close())
	}()
	var count int
	_ = rows.Next()
	CheckError(t, rows.Scan(&count))
	if count != 1 {
		t.Fatalf("missing foriegn key\n"+
			"from %s::%s\n"+
			"to %s::%s", table, column, referencedTable, referencedColumn)
	}
	CheckError(t, rows.Close())
}
