package repoScanner

// Scan - Run all security scans for the repo or designated project
func Scan(noop bool, projectPath string) error {
	if noop {
		return nil
	}
	return nil
}
