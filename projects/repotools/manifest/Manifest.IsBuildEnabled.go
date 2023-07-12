package projectmanifest

func (manifest *Manifest) IsBuildEnabled() bool {
	return manifest.Options.BuildEnabled
}
