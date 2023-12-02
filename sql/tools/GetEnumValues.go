package database

import (
	"fmt"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// GetEnumValues - Return a list of enumerated values.
func GetEnumValues(t *testing.T, db *Postgres.Db, typeName string) (values []string) {

	query := fmt.Sprintf("SELECT unnest(enum_range(NULL::%s)) AS enum_value;", typeName)

	rows, err := db.Query(query)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			t.Fatal(err)
		}
		values = append(values, v)

	}
	return values
}
