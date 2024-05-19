package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestSortManifestsByDependencies(t *testing.T) {
	// Create a Monorepo instance with unsorted manifests
	repo := &Monorepo{
		StartTime: time.Now(),
		Root:      "/tmp/TestSortManifestsByDependencies",
		Debug:     true,
		manifestList: []Manifest{
			{
				FileName: "class2/A.yaml",
				config: Config{
					Header: HeaderSection{
						Dependencies: []string{"class1/G.yaml"},
					},
				},
			},
			{
				FileName: "class3/B.yaml",
				config: Config{
					Header: HeaderSection{
						Dependencies: []string{"class2/A.yaml"},
					},
				},
			},
			{
				FileName: "class3/C.yaml",
				config: Config{
					Header: HeaderSection{
						Dependencies: []string{},
					},
				},
			},
			{
				FileName: "class4/F.yaml",
				config: Config{
					Header: HeaderSection{
						Dependencies: []string{},
					},
				},
			},
			{
				FileName: "class9/D.yaml",
				config: Config{
					Header: HeaderSection{
						Dependencies: []string{"class4/E.yaml", "class4/F.yaml"},
					},
				},
			},
			{
				FileName: "class4/E.yaml",
				config: Config{
					Header: HeaderSection{
						Dependencies: []string{"class1/G.yaml"},
					},
				},
			},
			{
				FileName: "class1/G.yaml",
				config: Config{
					Header: HeaderSection{
						Dependencies: []string{},
					},
				},
			},
		},
	}
	t.Cleanup(func() {
		for _, m := range repo.manifestList {
			d := filepath.Dir(filepath.Join(repo.Root, m.FileName))
			if directory.Exists(d) {
				if err := os.RemoveAll(d); err != nil {
					t.Fatal(err)
				}
			}
		}
	})
	for _, m := range repo.manifestList {
		d := filepath.Dir(filepath.Join(repo.Root, m.FileName))
		if !directory.Exists(d) {
			if err := os.MkdirAll(d, 0744); err != nil {
				t.Fatal(err)
			}
		}
		if err := os.WriteFile(filepath.Join(repo.Root, m.FileName), []byte("ok"), 0744); err != nil {
			t.Fatal(err)
		}
	}

	// Expected order after sorting
	expectedOrder := []string{
		"class1/G.yaml", "class4/E.yaml", "class4/F.yaml",
		"class3/C.yaml", "class3/B.yaml", "class9/D.yaml",
		"class2/A.yaml",
	}

	// Call the method to be tested
	repo.SortManifestsByDependencies(false)

	// Verify the order after sorting
	actualOrder := make([]string, len(repo.manifestList))
	for i, manifest := range repo.manifestList {
		actualOrder[i] = manifest.FileName
	}

	if !reflect.DeepEqual(actualOrder, expectedOrder) {
		t.Fatalf("Unexpected order after sorting. Expected: %v, Actual: %v", expectedOrder, actualOrder)
	}
}
