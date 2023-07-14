package projectmanifest

func (manifest *Manifest) DependencyCount() int {
	return len(manifest.Dependencies)
}
