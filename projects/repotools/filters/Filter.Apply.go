package filters

/*
 * projects/repotools/filters/Apply.go
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Given a project manifest object and the internal state of the filter
 * determine if the given project (by its manifest) matches the criteria
 * of the filter then return true so that the caller's operation can be
 * applied to the project.
 *
 * See README.md
 */

import (
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
)

// Apply - Apply the internal filter state to a project manifest and return a boolean answer
func (filter *Filter) Apply(manifest projectmanifest.Manifest) (filtered bool) {

	//Note: -commands (filter.Commands) is not used in this filter stage.
	//      it will be applied to the path of the file at a higher level.
	return (manifest.IsLintEnabled() == filter.LintEnabled) ||
		(manifest.IsScanEnabled() == filter.ScanEnabled) ||
		(manifest.IsBuildEnabled() == filter.BuildEnabled) ||
		(manifest.IsSignEnabled() == filter.SignEnabled) ||
		(manifest.IsPackEnabled() == filter.PackEnabled) ||
		manifest.SupportsOpsys(filter.os) ||
		manifest.SupportsArch(filter.os)
}
