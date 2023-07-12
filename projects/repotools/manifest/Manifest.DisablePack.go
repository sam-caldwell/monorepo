package projectmanifest

func (manifest *Manifest) DisablePack() *Manifest {
	manifest.Options.PackEnabled = false
	return manifest
}
