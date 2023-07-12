package projectmanifest

func (manifest *Manifest) SetAuthor(author string) *Manifest {
	if manifest.err != nil {
		return manifest
	}
	manifest.Author = author
	return manifest
}
