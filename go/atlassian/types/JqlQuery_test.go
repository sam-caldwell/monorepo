package AtlassianTypes

import (
	"encoding/json"
	"testing"
)

func TestJqlQuery_Struct(t *testing.T) {
	t.Run("Test defaults", func(t *testing.T) {
		q := JqlQuery{}
		if len(q.Expand) != 0 {
			t.Fatal("expand should be empty")
		}
		if len(q.Fields) != 0 {
			t.Fatal("fields should be empty")
		}
		if q.FieldsByKeys {
			t.Fatalf("fieldsByKeys should be false")
		}
		if q.Jql != "" {
			t.Fatal("jql should be empty")
		}
		if q.MaxResults != 0 {
			t.Fatal("maxResults should be 0")
		}
		if q.StartAt != 0 {
			t.Fatal("startAt should be 0")
		}
		if q.ValidateQuery != "" {
			t.Fatal("validateQuery should be empty")
		}
	})
	t.Run("json unmarshall/marshall", func(t *testing.T) {
		var queryText = `{
          "expand": [
            "names",
            "schema",
            "operations"
          ],
          "fields": [
            "summary",
            "status",
            "assignee"
          ],
          "fieldsByKeys": true,
          "jql": "project = HSP",
          "maxResults": 15,
          "startAt": 0
        }`
		var queryObject JqlQuery
		if err := json.Unmarshal([]byte(queryText), &queryObject); err != nil {
			t.Fatal(err)
		}
		if queryObject.Expand[0] != "names" {
			t.Fatal("expand should be 'names'")
		}
		if queryObject.Expand[1] != "schema" {
			t.Fatal("expand should be 'schema'")
		}
		if queryObject.Expand[2] != "operations" {
			t.Fatal("expand should be 'operations'")
		}
		if queryObject.Fields[0] != "summary" {
			t.Fatal("fields should be 'summary'")
		}
		if queryObject.Fields[1] != "status" {
			t.Fatal("fields should be 'status'")
		}
		if queryObject.Fields[2] != "assignee" {
			t.Fatal("fields should be 'assignee'")
		}
		if !queryObject.FieldsByKeys {
			t.Fatal("fieldsByKeys should be true")
		}
		if queryObject.Jql != "project = HSP" {
			t.Fatal("jql should be 'project = HSP'")
		}
		if queryObject.MaxResults != 15 {
			t.Fatal("maxResults should be 15")
		}
		if queryObject.StartAt != 0 {
			t.Fatal("startAt should be 0")
		}
		if actual, err := json.Marshal(queryObject); err != nil {
			t.Fatal("error should be nil")
		} else {

			expected := "{\"expand\":[\"names\",\"schema\",\"operations\"]," +
				"\"fields\":[\"summary\",\"status\",\"assignee\"]," +
				"\"fieldsByKeys\":true,\"jql\":\"project = HSP\",\"maxResults\":15}"
			//Note: startAt is omitted because 0 is the default (omitted) value.

			if string(actual) != expected {
				t.Fatalf("text should be equal:\n"+
					"actual:\n%v\n---\nexpected:\n%v\n", string(actual), expected)
			}
		}
	})
}
