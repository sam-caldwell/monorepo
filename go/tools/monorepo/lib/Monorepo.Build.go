package monorepo

// Build - Run project-specific Build operations (as defined in the project manifest).
func (m *Monorepo) Build() (err error) {
	pass := 0
	fail := 0
	//minikube.Start()
	//defer minikube.Stop()
	for _, manifest := range m.manifestList {
		if err = manifest.Run("build", &m.Root, m.Debug); err != nil {
			fail++
			return err
		} else {
			pass++
		}
	}
	showStats(pass, fail)
	return err
}
