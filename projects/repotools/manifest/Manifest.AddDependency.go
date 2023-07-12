package projectmanifest

func (manifest *Manifest) AddDependency(dep string) *Manifest {
	//ToDo: verify the dependency project exists (it must be in projects/ and have a MANIFEST.yaml file)
	manifest.Dependencies = append(manifest.Dependencies, dep)
	return manifest
}
