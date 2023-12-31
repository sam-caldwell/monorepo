package sqldbtest

import (
	"database/sql"
	"fmt"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"strings"
	"testing"
)

func VerifyFunctionStructure(t *testing.T, db *Postgres.Db, functionName string, expected string) {
	var err error
	var rows *sql.Rows
	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		query := fmt.Sprintf("SELECT "+
			"proname AS fn, "+
			"proargnames AS pn, "+
			"("+
			"SELECT array_agg(typname) "+
			"FROM pg_type "+
			"WHERE oid = ANY(proargtypes)"+
			") AS pt, "+
			"("+
			"SELECT typname "+
			"FROM pg_type "+
			"WHERE oid = prorettype "+
			") AS rt "+
			"FROM pg_proc "+
			"WHERE proname = lower('%s')", functionName)

		rows, err = db.Query(query)
		CheckError(t, err)

		if !rows.Next() {
			t.Fatalf("Fail: expected at least one row returned.")
		}
		defer func() { _ = rows.Close() }()
		var actualName, actualPNames, actualPTypes, actualReturnTypes sql.NullString

		if err = rows.Scan(&actualName, &actualPNames, &actualPTypes, &actualReturnTypes); err != nil {
			t.Fatalf("failed scanning row: %v", err)
		}
		if !actualName.Valid {
			t.Fatalf("function not found: %v", functionName)
		}

		actual := fmt.Sprintf("fn:%s,pn:%s,pt:%s,rt:%s",
			actualName.String, actualPNames.String, actualPTypes.String, actualReturnTypes.String)

		if actual != strings.ToLower(expected) {
			t.Fatalf("fail: function structure mismatch\n"+
				"actual:   %s\n"+
				"expected: %s",
				actual,
				strings.ToLower(expected))
		}
	})
}
