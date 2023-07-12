package projectmanifest

func (manifest *Manifest) DisableSign() *Manifest {
	manifest.Options.SignEnabled = false
	return manifest
}
