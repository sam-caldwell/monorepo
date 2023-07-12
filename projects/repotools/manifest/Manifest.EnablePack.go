package projectmanifest

func (manifest *Manifest) EnablePack() *Manifest {
	manifest.Options.PackEnabled = true
	return manifest
}
