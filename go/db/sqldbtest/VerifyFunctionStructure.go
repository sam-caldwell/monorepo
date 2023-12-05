package sqldbtest

import (
	"fmt"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"strings"
	"testing"
)

func VerifyFunctionStructure(t *testing.T, db *Postgres.Db, functionName string, expected string) {

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

	rows, err := db.Query(query)
	CheckError(t, err)
	defer func() {
		CheckError(t, rows.Close())
	}()
	var fn, pn, pt, rt string
	_ = rows.Next()
	CheckError(t, rows.Scan(&fn, &pn, &pt, &rt))
	actual := fmt.Sprintf("fn:%s,pn:%s,pt:%s,rt:%s", fn, pn, pt, rt)
	if actual != strings.ToLower(expected) {
		t.Fatalf("function structure mismatch\n"+
			"actual:   %s\n"+
			"expected: %s",
			actual,
			strings.ToLower(expected))
	}
}
