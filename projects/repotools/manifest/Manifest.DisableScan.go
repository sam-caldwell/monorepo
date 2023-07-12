package projectmanifest

func (manifest *Manifest) DisableScan() *Manifest {
	manifest.Options.ScanEnabled = false
	return manifest
}
