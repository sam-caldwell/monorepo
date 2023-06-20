package semver

import "github.com/go-git/go-git/v5"

// openRepo - open a repo
func (ver *SemanticVersion) openRepo(path string) (*git.Repository, error) {
	return git.PlainOpen(path)
}
