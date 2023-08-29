package semver

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing"
)

// UpdateTag - update the current local git repo (current directory) with new tag.
func (ver *SemanticVersion) UpdateTag() error {

	repo, err := ver.openRepo(defaultRepo)
	if err != nil {
		return err
	}

	tagRef, err := repo.Tags()
	if err != nil {
		return err
	}

	// Check if the tag already exists
	tagExists := false
	err = tagRef.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().Short() == ver.String() {
			tagExists = true
		}
		return nil
	})
	if err != nil {
		return err
	}

	// If tag already exists, return an error or handle it as per your requirement
	if tagExists {
		return fmt.Errorf(errTagAlreadyExists, ver.String())
	}

	// Get the HEAD reference
	headRef, err := repo.Head()
	if err != nil {
		return err
	}

	// Get the commit object for the HEAD reference
	commit, err := repo.CommitObject(headRef.Hash())
	if err != nil {
		return err
	}

	// Create the tag in the repository
	if _, err := repo.CreateTag(ver.String(), commit.Hash, nil); err != nil {
		return err
	}
	return nil
}
