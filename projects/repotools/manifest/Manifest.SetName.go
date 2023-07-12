package projectmanifest

func (manifest *Manifest) SetName(name string) *Manifest {
	if manifest.err != nil {
		return manifest
	}
	manifest.Name = name
	return manifest
}
