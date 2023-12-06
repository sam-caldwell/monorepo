package monorepo

// Build - Run project-specific Build operations (as defined in the project manifest).
func (m *Monorepo) Build() (err error) {

	for _, manifest := range m.manifestList {
		if err = manifest.Run("build", m.Debug); err != nil {
			return err
		}
	}
	return err
}
