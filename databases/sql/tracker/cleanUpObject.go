package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
)

// cleanUpObject - delete an object from a given table by id
func cleanUpObject(db *Postgres.Db, tableName string, entityId uuid.UUID) (err error) {
	if entityId.String() == "00000000-0000-0000-0000-000000000000" {
		return nil
	}
	fmt.Printf("cleanUpObject() start.  Entity (table: %v): %v\n", tableName, entityId)
	var rows *sql.Rows
	rows, err = db.Query("delete from %s where id='%s'", tableName, entityId)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	fmt.Printf("cleanUpObject() done.  Entity (table: %v): %v\n", tableName, entityId)
	return err
}
