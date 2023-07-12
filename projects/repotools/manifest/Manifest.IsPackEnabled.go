package projectmanifest

func (manifest *Manifest) IsPackEnabled() bool {
	return manifest.Options.PackEnabled
}
