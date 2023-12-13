package sqldbtest

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/convert"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"strings"
	"testing"
)

// ValidateIndex - Verify that a given tableName has an expected index
//
//	Follow the indexName convention: ndx${tableName}${columnNames}
func ValidateIndex(t *testing.T, db *Postgres.Db, tableName string, expectedColumns []string) {

	convert.CapitalizeStringList(&expectedColumns)

	c := strings.Join(expectedColumns, "")

	indexName := fmt.Sprintf("ndx%s%s", convert.Capitalize(tableName), c)
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
	if !rows.Next() {
		t.Fatalf("no expected rows returned\n"+
			"table: %s\n"+
			"index: %s\n"+
			"err:   %v", tableName, indexName, err)
	}
	var actualIndexName, actualIndexDef string
	if err := rows.Scan(&actualIndexName, &actualIndexDef); err != nil {
		t.Fatalf("Failed to scan results\n"+
			"table: %s\n"+
			"index: %s\n"+
			"err:   %v", tableName, indexName, err)
	}
}
