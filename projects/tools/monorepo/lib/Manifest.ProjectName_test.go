package monorepo

import "testing"

func TestManifest_ProjectName(t *testing.T) {
	var m Manifest

	t.Run("test 1", func(t *testing.T) {
		m.FileName = "/Users/samcaldwell/git/monorepo/go/trees/AsciiTree/Manifest.yml"
		expected := "trees/AsciiTree"
		if project := m.ProjectName("/Users/samcaldwell/git/monorepo"); project != expected {
			t.Fatalf("\nproject did not parse properly.\n"+
				"Actual: %s\n"+
				"Expect: %s\n", project, expected)
		}
	})
	t.Run("test 2", func(t *testing.T) {
		m.FileName = "/Users/samcaldwell/git/monorepo/go/project1/subprojectA/component/Manifest.yml"
		expected := "project1/subprojectA/component"
		if project := m.ProjectName("/Users/samcaldwell/git/monorepo"); project != expected {
			t.Fatalf("\nproject did not parse properly.\n"+
				"Actual: %s\n"+
				"Expect: %s\n", project, expected)
		}
	})
	t.Run("test 2", func(t *testing.T) {
		m.FileName = "/Users/samcaldwell/git/monorepo/go/project1/subprojectA/component/child/Manifest.yml"
		expected := "project1/subprojectA/component/child"
		if project := m.ProjectName("/Users/samcaldwell/git/monorepo"); project != expected {
			t.Fatalf("\nproject did not parse properly.\n"+
				"Actual: %s\n"+
				"Expect: %s\n", project, expected)
		}
	})

}
