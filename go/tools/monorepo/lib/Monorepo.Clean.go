package monorepo

func (m *Monorepo) Clean() (err error) {

	for _, manifest := range m.manifestList {
		if err = manifest.Run("clean", m.Debug); err != nil {
			return err
		}
	}
	return err
}
