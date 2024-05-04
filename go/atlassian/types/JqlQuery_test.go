package AtlassianTypes

import "testing"

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
}
