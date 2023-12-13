package sqldbtest

import (
	"fmt"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"strings"
	"testing"
)

// ValidateIndex - Verify that a given tableName has an expected index
//
//	Follow the indexName convention: ndx${tableName}${columnNames}
func ValidateIndex(t *testing.T, db *Postgres.Db, tableName string, expectedColumns []string) {
	tableName = strings.ToLower(tableName)
	columnsForName := strings.ToLower(strings.Join(expectedColumns, ""))
	indexName := fmt.Sprintf("ndx%s%s", tableName, columnsForName)

	t.Run("verify index ("+indexName+")", func(t *testing.T) {

		rows, err := db.Query(""+
			"select indexname,indexdef "+
			"from pg_indexes "+
			"where tablename=lower('%s') "+
			"and indexname=lower('%s')", tableName, indexName)
		if err != nil {
			t.Fatalf("Failed to query index\n"+
				"table: %s\n"+
				"index: %s\n"+
				"err:   %v", tableName, indexName, err)
		}
		defer func() { _ = rows.Close() }()
		if !rows.Next() {
			t.Fatalf("no expected rows returned\n"+
				"table: %s\n"+
				"index: %s\n", tableName, indexName)
		}

		var actualIndexName, actualIndexDef string
		actualIndexName = strings.ToLower(actualIndexName)
		actualIndexDef = strings.ToLower(actualIndexDef)
		if err := rows.Scan(&actualIndexName, &actualIndexDef); err != nil {
			t.Fatalf("Failed to scan results\n"+
				"table: '%s'\n"+
				"index: '%s'\n"+
				"err:   %v", tableName, indexName, err)
		}
		if actualIndexName != indexName {
			t.Fatalf("indexName mismatch\n"+
				"table: '%s'\n"+
				"index: '%s'\n"+
				"actual:'%s'", tableName, indexName, actualIndexDef)
		}
		if actualIndexDef == "" {
			t.Fatalf("indexdef mismatch\n"+
				"table: '%s'\n"+
				"index: '%s'\n"+
				"actual:'%s'", tableName, indexName, actualIndexDef)
		}
	})
}
