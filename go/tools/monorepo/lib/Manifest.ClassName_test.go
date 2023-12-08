package monorepo

import "testing"

func TestManifest_ClassName(t *testing.T) {
	var m Manifest

	t.Run("test 1", func(t *testing.T) {
		m.FileName = "/Users/samcaldwell/git/monorepo/go/trees/AsciiTree/Manifest.yml"
		expected := "go"
		if class := m.ClassName("/Users/samcaldwell/git/monorepo"); class != expected {
			t.Fatalf("\nproject did not parse properly.\n"+
				"Actual: %s\n"+
				"Expect: %s\n", class, expected)
		}
	})
	t.Run("test 2", func(t *testing.T) {
		m.FileName = "/Users/samcaldwell/git/monorepo/java/project1/subprojectA/component/Manifest.yml"
		expected := "java"
		if class := m.ClassName("/Users/samcaldwell/git/monorepo"); class != expected {
			t.Fatalf("\nproject did not parse properly.\n"+
				"Actual: %s\n"+
				"Expect: %s\n", class, expected)
		}
	})
	t.Run("test 2", func(t *testing.T) {
		m.FileName = "/Users/samcaldwell/git/monorepo/go/project1/subprojectA/component/child/Manifest.yml"
		expected := "go"
		if class := m.ClassName("/Users/samcaldwell/git/monorepo"); class != expected {
			t.Fatalf("\nproject did not parse properly.\n"+
				"Actual: %s\n"+
				"Expect: %s\n", class, expected)
		}
	})

}
