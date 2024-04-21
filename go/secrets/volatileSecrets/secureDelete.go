package volatileSecrets

// secureDelete - overwrite RAM with 0x00 then free the memory
func secureDelete(b *[]byte) {
	if *b != nil {
		for i := range *b {
			(*b)[i] = 0x00
		}
		*b = nil
	}
}
