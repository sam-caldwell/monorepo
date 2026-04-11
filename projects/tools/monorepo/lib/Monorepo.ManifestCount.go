package monorepo

// ManifestCount return the number of records in manifestList
func (m *Monorepo) ManifestCount() int {
	return len(m.manifestList)
}
