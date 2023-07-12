package projectmanifest

func (manifest *Manifest) DisableLint() *Manifest {
	manifest.Options.LintEnabled = false
	return manifest
}
