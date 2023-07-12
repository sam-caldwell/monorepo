package projectmanifest

func (manifest *Manifest) IsScanEnabled() bool {
	return manifest.Options.ScanEnabled
}
