package projectmanifest

func CreateManifest(filename string) *Manifest {
	m := Manifest{}
	m.fileName = filename
	return &m
}
