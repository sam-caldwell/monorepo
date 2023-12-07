package monorepo

// Test - Run project-specific Test operations (as defined in the project manifest).
func (m *Monorepo) Test() (err error) {
	for _, manifest := range m.manifestList {
		if err = manifest.Run("test", &m.Root, m.Debug); err != nil {
			return err
		}
	}
	return err
}
