package projectmanifest

func (manifest *Manifest) DisableBuild() *Manifest {
	manifest.Options.BuildEnabled = false
	return manifest
}
