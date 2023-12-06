package monorepo

func (m *Monorepo) Test() (err error) {
	for _, manifest := range m.manifestList {
		if err = manifest.Run("test", m.Debug); err != nil {
			return err
		}
	}
	return err
}
