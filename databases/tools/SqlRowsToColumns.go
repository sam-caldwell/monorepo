package database

import (
	"database/sql"
	"testing"
)

// SqlRowsToColumns - convert sql rows to a list of strings
func SqlRowsToColumns(t *testing.T, rows *sql.Rows) (values []string) {
	if rows == nil {
		t.Fatal("nil sql rows")
	}
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			t.Fatal(err)
		}
		values = append(values, v)

	}
	return values
}
