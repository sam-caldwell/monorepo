package projectmanifest

func (manifest *Manifest) EnableScan() *Manifest {
	manifest.Options.ScanEnabled = true
	return manifest
}
