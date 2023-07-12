package projectmanifest

func (manifest *Manifest) IsSignEnabled() bool {
	return manifest.Options.SignEnabled
}
