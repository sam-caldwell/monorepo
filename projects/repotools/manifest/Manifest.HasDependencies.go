package projectmanifest

func (manifest *Manifest) HasDependencies() bool {
	return len(manifest.Dependencies) > 0
}
