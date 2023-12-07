package monorepo

// Clean - Run project-specific clean-up operations (as defined in the project manifest).
func (m *Monorepo) Clean() (err error) {
	pass := 0
	fail := 0
	for _, manifest := range m.manifestList {
		if err = manifest.Run("clean", &m.Root, m.Debug); err != nil {
			fail++
			return err
		} else {
			pass++
		}
	}
	showStats(pass, fail)
	return err
}
