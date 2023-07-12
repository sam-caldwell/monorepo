package projectmanifest

func (manifest *Manifest) EnableSign() *Manifest {
	manifest.Options.SignEnabled = true
	return manifest
}
