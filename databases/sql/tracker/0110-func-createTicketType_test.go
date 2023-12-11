package psqlTrackerDb

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/db/sqldbtest"
	"strings"
	"testing"
)

func TestSqlDbFunc_createTicketType(t *testing.T) {
	t.Skip("disabled until we can finish tests for the dependencies.")
	const (
		functionName     = "createTicketType"
		tableName        = "ticketTypes"
		testTicketType   = "testTicketType"
		testWorkflowName = "testWorkflowName"
	)
	//var typeId string
	var avatarId uuid.UUID
	var iconId uuid.UUID
	var workflowId uuid.UUID
	var ownerId uuid.UUID
	var teamId uuid.UUID
	//var typeName string

	db := sqldbtest.InitializeTestDbConn(t)

	t.Cleanup(func() {
		// Note: we only clean up the avatar we expect to have created.
		//       this should safeguard against an accidental run on prod.
		_, _ = db.Query("delete from %s where name='%s' cascade;", tableName, testTicketType)

		err := db.Close()
		sqldbtest.CheckError(t, err)
	})

	t.Run("verify the function structure (params, return)", func(t *testing.T) {
		sqldbtest.VerifyFunctionStructure(t, db,
			strings.ToLower(functionName),
			fmt.Sprintf("fn:%s,"+
				"pn:{name,iconid,workflowid,description},"+
				"pt:{text,varchar,uuid},"+
				"rt:uuid", strings.ToLower(functionName)))
	})

	t.Run("create test data", func(t *testing.T) {
		var err error
		t.Run("create avatarId", func(t *testing.T) {
			/*
			 * We need to create an avatar to create a user (ownerId)
			 */
			var rows *sql.Rows
			rows, err = db.Query("select createIcons('http://localhost/myfakeavatar.jpeg');")
			if err != nil {
				t.Fatal(err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var raw string
			err = rows.Scan(&raw)
			if avatarId, err = uuid.Parse(raw); err != nil {
				t.Fatal(err)
			}
			if avatarId.String() == "00000000-0000-0000-0000-000000000000" {
				t.Fatal("illegal zero uuid")
			}
		})
		t.Run("create iconId", func(t *testing.T) {
			/*
			 * We need to create an icon for use creating a Workflow and Team
			 */
			var rows *sql.Rows
			rows, err = db.Query("select createIcons('http://localhost/myfakeicon.jpeg');")
			if err != nil {
				t.Fatal(err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var raw string
			err = rows.Scan(&raw)
			if iconId, err = uuid.Parse(raw); err != nil {
				t.Fatal(err)
			}
			if iconId.String() == "00000000-0000-0000-0000-000000000000" {
				t.Fatal("illegal zero uuid")
			}
		})
		t.Run("create createUser (ownerId)", func(t *testing.T) {
			/*
			 * We need to create a user (ownerId) to create a workflow
			 */
			var rows *sql.Rows
			rows, err = db.Query("select createUser('wendell','fertig','%s',"+
				"'example@example.com','512-123-4567','test user');",
				avatarId)
			if err != nil {
				t.Fatal(err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var raw string
			err = rows.Scan(&raw)
			if ownerId, err = uuid.Parse(raw); err != nil {
				t.Fatal(err)
			}
			if ownerId.String() == "00000000-0000-0000-0000-000000000000" {
				t.Fatal("illegal zero uuid")
			}
		})
		t.Run("create createTeam (teamId)", func(t *testing.T) {
			/*
			 * We need to create a user (ownerId) to create a workflow
			 */
			var rows *sql.Rows
			rows, err = db.Query("select createTeam('testTeam','%s','%s',"+
				"'read','read','read','test team');",
				iconId, ownerId)
			if err != nil {
				t.Fatalf("createTeam() failed %v\n"+
					"iconId:  %v\n"+
					"ownerId: %v", err, iconId, ownerId)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var raw string
			err = rows.Scan(&raw)
			if teamId, err = uuid.Parse(raw); err != nil {
				t.Fatal(err)
			}
			if teamId.String() == "00000000-0000-0000-0000-000000000000" {
				t.Fatal("illegal zero uuid")
			}
		})
		t.Run("create workflowId", func(t *testing.T) {
			/*
			 * We have to create a workflow to create a ticketType.
			 */
			var rows *sql.Rows
			rows, err = db.Query(""+
				"select createWorkflow('%s','%s','%s','%s','read','read','read','testDescription');",
				testWorkflowName, iconId, ownerId, teamId)
			if err != nil {
				t.Fatal(err)
			}
			if !rows.Next() {
				t.Fatal("no row returned")
			}
			var rawIconId string
			err = rows.Scan(&rawIconId)
			if workflowId, err = uuid.Parse(rawIconId); err != nil {
				t.Fatal(err)
			}
			if workflowId.String() == "00000000-0000-0000-0000-000000000000" {
				t.Fatal("illegal zero uuid")
			}
		})
	})
}
