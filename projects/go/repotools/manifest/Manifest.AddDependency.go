package projectmanifest

/*
 * projects/repotool/manifest/Manifest.AddDependency.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file contains a function called AddDependency
 * which will append the manifest.Dependencies list with
 * the name of a project upon which the project described
 * by the Manifest itself.
 */

// AddDependency - Add a project dependency record
func (manifest *Manifest) AddDependency(dep string) *Manifest {
	//ToDo: verify the dependency project exists (it must be in projects/ and have a MANIFEST.yaml file)
	manifest.Dependencies = append(manifest.Dependencies, dep)
	return manifest
}
