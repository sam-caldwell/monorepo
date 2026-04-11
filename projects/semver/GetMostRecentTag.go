package semver

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"regexp"
	"sort"
	"strings"
)

// GetMostRecentTag - Query the local git repo for the most recent tag
func (ver *SemanticVersion) GetMostRecentTag() (err error) {
	var tag string
	//Get the most recent tag (string)
	if err = func() error {
		var repo *git.Repository
		var tagNames []string
		// Open the repository
		if repo, err = ver.openRepo(defaultRepo); err != nil {
			return err
		}
		// Get the list of tags
		var tagRefs storer.ReferenceIter
		if tagRefs, err = repo.Tags(); err != nil {
			return err
		} else {
			// Iterate over the tags and get their names
			if err = tagRefs.ForEach(func(ref *plumbing.Reference) error {
				thisTag := strings.TrimSpace(ref.Name().Short())
				if regexp.MustCompile(tagPattern).Match([]byte(thisTag)) {
					tagNames = append(tagNames, thisTag)
				}
				return nil
			}); err != nil {
				return err
			}
		}
		// Sort the tag names in descending order
		sort.Sort(sort.Reverse(sort.StringSlice(tagNames)))
		// Get the most recent tag
		if len(tagNames) > 0 {
			tag = tagNames[0]
			return err
		}
		// there is no tag
		tag = defaultTag
		return err
	}(); err != nil {
		return err
	}
	//Parse string tag into the semantic version
	return ver.ParseP(&tag)
}
