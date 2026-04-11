package volatileSecrets

// Get - Retrieve the secret in memory (yeah...this will put it in cleartext in memory for a time).
// Tip: keep the returned variable local and remove it as soon as possible.
func (p *Password) Get() string {
	return string(p.GetBytes())
}
