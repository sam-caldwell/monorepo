package filters

/*
 * projects/repotools/common/Filter.go
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines a filtering mechanism for repotools.ListProjects()
 *
 * See README.md
 */

// Filter - A project filter's state
type Filter struct {
	Commands     bool //Indicates a projects in cmd/<project>/<program> should be shown...only
	BuildEnabled bool
	LintEnabled  bool
	PackEnabled  bool
	ScanEnabled  bool
	SignEnabled  bool
	Language     string
	os           string
	arch         string
}
