package volatileSecrets

// secureDelete - overwrite RAM
func secureDelete(b *[]byte) {
	if *b != nil {
		for i := range *b {
			(*b)[i] = 0x00
		}
		*b = nil
	}
}
