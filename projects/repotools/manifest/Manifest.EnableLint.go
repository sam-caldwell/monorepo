package projectmanifest

func (manifest *Manifest) EnableLint() *Manifest {
	manifest.Options.LintEnabled = true
	return manifest
}
