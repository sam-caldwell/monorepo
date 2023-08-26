package projectmanifest

/*
 * projects/repotool/manifest/Manifest.EnablePack.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the EnablePack() method
 * which will set the PackEnabled option to true.
 */

// EnablePack - Set Options.PackEnabled to true
func (manifest *Manifest) EnablePack() *Manifest {
	manifest.Options.PackEnabled = true
	return manifest
}
