package monorepo

func (m *Monorepo) Build() (err error) {

	for _, manifest := range m.manifestList {
		if err = manifest.Run("build", m.Debug); err != nil {
			return err
		}
	}
	return err
}
