package psqlTrackerDb

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// getUsersInTeam - get a list of users in a given team
func getUsersInTeam(t *testing.T, db *Postgres.Db, teamId uuid.UUID) (users []uuid.UUID) {
	var rows *sql.Rows
	var err error

	t.Run("getUsersInTeam() and verify", func(t *testing.T) {

		rows, err = db.Query("select getUsersInTeam('%s');", teamId)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = rows.Close() }()

		type CollectionOfUsers struct {
			UserId string `json:"userId"`
		}

		var actual []CollectionOfUsers
		for rows.Next() {
			var raw string
			if err = rows.Scan(&raw); err != nil {
				t.Fatal(err)
			}
			if err = json.Unmarshal([]byte(raw), &actual); err != nil {
				t.Fatalf("error: failed to parse: %v\n"+
					"raw: %v", err, raw)
			}
		}
		for _, user := range actual {
			userId, err := uuid.Parse(user.UserId)
			if err != nil {
				t.Fatal(err)
			}
			users = append(users, userId)
		}
	})
	return users
}
