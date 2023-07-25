package projectmanifest

/*
 * projects/repotools/manifest/CreateManifest.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * CreateManifest will initialize a Manifest object and return
 * a point thereto with the manifest filename stored therein.
 */

// CreateManifest - Create a manifest object and store the corresponding manifest file to the internal state.
func CreateManifest(filename string) *Manifest {
	m := Manifest{}
	m.fileName = filename
	return &m
}
