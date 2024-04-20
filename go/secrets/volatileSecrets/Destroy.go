package volatileSecrets

// Destroy - Overwrite memory and free it.
func (p *Password) Destroy() {
	secureDelete(&p.data)
}
