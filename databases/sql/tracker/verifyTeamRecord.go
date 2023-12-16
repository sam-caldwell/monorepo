package psqlTrackerDb

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// verifyTeamRecord - query for and verify the content of the record
func verifyTeamRecord(t *testing.T, db *Postgres.Db, teamId uuid.UUID, teamName string, iconId uuid.UUID,
	ownerId uuid.UUID, owner, team, everyone, description string) {

	var err error
	var rows *sql.Rows
	var actualTeam TrackerTeam

	t.Run("verify outcome", func(t *testing.T) {
		rows, err = db.Query("select getTeamById('%s')", teamId)
		if err != nil {
			t.Fatalf("Fail: Query error: %v", err)
		}
		if !rows.Next() {
			t.Fatal("Fail: no row returned")
		}
		var raw sql.NullString
		err = rows.Scan(&raw)
		if !raw.Valid {
			t.Fatal("Fail: invalid sql result (null) but expected at least one row")
		}
		if err := json.Unmarshal([]byte(raw.String), &actualTeam); err != nil {
			t.Fatalf("Failed to read result: %v", err)
		}
		if teamId != actualTeam.Id {
			t.Fatal("Fail: teamId mismatch")
		}
		if teamName != actualTeam.Name {
			t.Fatal("Fail: team name mismatch")
		}
		if iconId != actualTeam.IconId {
			t.Fatalf("Fail: iconId mismatch")
		}
		if ownerId != actualTeam.OwnerId {
			t.Fatalf("Fail: ownerId mismatch")
		}
		if actualTeam.Owner != owner {
			t.Fatalf("Fail: owner permission mismatch")
		}
		if actualTeam.Team != team {
			t.Fatalf("Fail: team permission mismatch")
		}
		if actualTeam.Everyone != everyone {
			t.Fatalf("Fail: everyone permission mismatch")
		}
		if actualTeam.Description != description {
			t.Fatalf("actualDescription mismatch\n"+
				"actual: %s\n"+
				"expected: %s", actualTeam.Description, description)
		}
	})
}
