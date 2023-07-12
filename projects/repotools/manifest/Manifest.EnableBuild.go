package projectmanifest

func (manifest *Manifest) EnableBuild() *Manifest {
	manifest.Options.BuildEnabled = true
	return manifest
}
