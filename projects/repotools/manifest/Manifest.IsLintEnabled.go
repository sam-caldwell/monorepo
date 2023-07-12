package projectmanifest

func (manifest *Manifest) IsLintEnabled() bool {
	return manifest.Options.LintEnabled
}
