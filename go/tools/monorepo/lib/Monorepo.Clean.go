package monorepo

func (m *Monorepo) Clean() (err error) {

	for _, path := range m.manifestList {
		var manifest Manifest
		if err = manifest.Load(path); err != nil {
			return err
		}
		if err = manifest.Run("clean", m.Debug); err != nil {
			return err
		}
	}
	return err
}
