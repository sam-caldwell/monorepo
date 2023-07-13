package repotools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentGitHash - Return the current git hash for the monorepo
func GetCurrentGitHash() (hash string, err error) {
	/*
	 * Get the repo root.
	 */
	var repoRoot string
	if repoRoot, err = FindRepoRoot(); err != nil {
		return hash, err
	}
	/*
	 * Read head from the GIT repository itself.
	 * (NOTE: no shell popped, no command executed.  Just raw git. Like the good Lord intended)
	 */
	refBytes, err := os.ReadFile(filepath.Join(repoRoot, ".git", "HEAD"))
	if err != nil {
		return "", err
	}
	/*
	 * Parse the reference to get the commit hash
	 */
	ref := strings.TrimSpace(string(refBytes))
	if !strings.HasPrefix(ref, "ref:") {
		return hash, fmt.Errorf(".git/HEAD should contain ref: refs/<path> but does not")
	}
	refBytes, err = os.ReadFile(
		filepath.Join(
			repoRoot,
			".git",
			strings.TrimSpace(
				strings.TrimPrefix(
					ref,
					"ref:"))))

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(refBytes)), err
}
